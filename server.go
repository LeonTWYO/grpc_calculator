package main

import (
	"context"
	"log"
	"net"

	"github.com/LeonTWYO/grpc_calculator/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	protobuf.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, req *protobuf.AddRequest) (*protobuf.AddResponse, error) {
	result := req.Operand1 + req.Operand2
	return &protobuf.AddResponse{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, req *protobuf.MultiplyRequest) (*protobuf.MultiplyResponse, error) {
	result := req.Factor1 * req.Factor2
	return &protobuf.MultiplyResponse{Result: result}, nil
}

func (s *server) Divide(ctx context.Context, req *protobuf.DivideRequest) (*protobuf.DivideResponse, error) {
	if req.Divisor == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "cannot divide by zero")
	}
	quotient := req.Dividend / req.Divisor
	remainder := req.Dividend % req.Divisor
	return &protobuf.DivideResponse{Quotient: quotient, Remainder: remainder}, nil
}

func runGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protobuf.RegisterCalculatorServer(s, &server{})
	log.Println("Starting gRPC server on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
