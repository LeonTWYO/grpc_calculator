package main

import (
	"context"
	"log"
	"net"

	"github.com/LeonTWYO/grpc_calculator/gpb"
	"google.golang.org/grpc"
)

// Server implements the Calculator service
type Server struct{}

func (s *Server) Add(ctx context.Context, req *gpb.AddRequest) (*gpb.AddResponse, error) {
	result := req.GetOperand1() + req.GetOperand2()
	return &gpb.AddResponse{Result: result}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	gpb.RegisterCalculatorServer(grpcServer, &Server{})

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Starting gRPC server on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
