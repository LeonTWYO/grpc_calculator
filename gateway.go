package main

import (
	"context"
	"log"
	"net/http"

	"github.com/LeonTWYO/grpc_calculator/protobuf"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func runGRPCGateway() {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	// Dial gRPC server
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := "localhost:50051"
	if err := protobuf.RegisterCalculatorHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.Fatalf("Failed to register gRPC gateway handler: %v", err)
	}

	// Start HTTP server
	log.Println("Starting gRPC gateway server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to serve gRPC gateway: %v", err)
	}
}
