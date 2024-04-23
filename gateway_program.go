package main

import (
	"context"
	"log"
	"net/http"

	"github.com/LeonTWYO/grpc_calculator/gpb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// StartHTTPServer starts the HTTP server for gRPC-Gateway
func StartHTTPServer() {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gpb.RegisterCalculatorHandlerFromEndpoint(ctx, mux, ":50051", opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC gateway: %v", err)
	}

	log.Println("Starting HTTP server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
