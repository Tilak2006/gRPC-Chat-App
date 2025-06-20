package main

import (
	mainpb "grpcchatapp/proto/gen"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
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

func main() {
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
