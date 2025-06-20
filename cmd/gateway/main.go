package main

import (
	"context"
	"log"
	"net/http"

	mainpb "grpcchatapp/proto/gen"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()            //->base context
	ctx, cancel := context.WithCancel(ctx) //->now can cancel anytime
	defer cancel()                         //->ensures cancelation at the end of main function

	//REST Multiplexer
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := mainpb.RegisterChatServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("Failed to register gateway handler: %v", err)
	}

	log.Println("REST gateway started on port 8080")
	handler := cors.AllowAll().Handler(mux)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
