package main

import (
	"context"
	mainpb "grpcchatapp/proto/gen"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatStream struct {
	stream mainpb.ChatService_ChatServer
	userID string
	ch     chan *mainpb.ChatMessage
}

type ChatServer struct {
	mainpb.UnimplementedChatServiceServer
	mu      sync.Mutex
	clients map[string]ChatStream
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		clients: make(map[string]ChatStream),
	}
}

func (s *ChatServer) Chat(stream mainpb.ChatService_ChatServer) error {
	mssg, err := stream.Recv()
	if err != nil {
		return err
	}
	userID := mssg.SenderId
	userName := mssg.SenderName

	mssgCh := make(chan *mainpb.ChatMessage, 100)
	client := ChatStream{
		stream: stream,
		userID: userID,
		ch:     mssgCh,
	}
	s.mu.Lock()
	s.clients[userID] = client
	s.mu.Unlock()

	log.Printf("%s joined the Chat\n", userName)
	defer func() {
		s.mu.Lock()
		delete(s.clients, userID)
		close(mssgCh)
		s.mu.Unlock()

		log.Printf("%s left the chat\n", userName)
	}()
	go func() {
		for mssg := range mssgCh {
			err := stream.Send(mssg)
			if err != nil {
				log.Printf("Error sending the message to %s: %v", userID, err)
				break
			}
		}
	}()

	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("%s disconnected: %v", userID, err)
			break
		}
		msg.Timestamp = time.Now().Unix()

		//this is to broadcast to all other members cnencted now
		s.mu.Lock()
		for id, cl := range s.clients {
			if id != userID {
				cl.ch <- msg
			}
		}
		s.mu.Unlock()
	}
	return nil
}

var messageHistory []*mainpb.ChatMessage //->empty slice pointing to struct(ChatMessage), used to hold all chat messages so far.

func (s *ChatServer) SendMessage(ctx context.Context, msg *mainpb.ChatMessage) (*emptypb.Empty, error) {
	msg.Timestamp = time.Now().Unix()

	s.mu.Lock()
	messageHistory = append(messageHistory, msg)

	for id, client := range s.clients {
		client.ch <- msg
		log.Printf("Sent message from [%s] to [%s]", msg.SenderName, id)
	}
	s.mu.Unlock()

	log.Printf("Message from %s: %s", msg.SenderName, msg.Message)

	return &emptypb.Empty{}, nil

}

func (s *ChatServer) GetMessages(ctx context.Context, _ *emptypb.Empty) (*mainpb.ChatMessages, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return &mainpb.ChatMessages{Messages: messageHistory}, nil
}

func main() {
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open Log file: %v", err)
	}
	defer logFile.Close()

	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("There was an error listening to the Server", err)
	}
	grpcServer := grpc.NewServer()
	mainpb.RegisterChatServiceServer(grpcServer, NewChatServer())
	log.Println("gRPC server started on port", port)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

}
