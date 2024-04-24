package main

import (
	"context"
	"log"
	"net"

	"github.com/LeonTWYO/grpc_calculator/protof" // Import your generated gRPC package

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Add(ctx context.Context, req *protof.AddRequest) (*protof.AddResponse, error) {
	result := req.Operand1 + req.Operand2
	return &protof.AddResponse{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, req *protof.MultiplyRequest) (*protof.MultiplyResponse, error) {
	result := req.Factor1 * req.Factor2
	return &protof.MultiplyResponse{Result: result}, nil
}

func (s *server) Divide(ctx context.Context, req *protof.DivideRequest) (*protof.DivideResponse, error) {
	if req.Divisor == 0 {
		return nil, grpc.Errorf(grpc.CodeInvalidArgument, "Cannot divide by zero")
	}
	quotient := req.Dividend / req.Divisor
	remainder := req.Dividend % req.Divisor
	return &protof.DivideResponse{Quotient: quotient, Remainder: remainder}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protof.RegisterCalculatorServer(s, &server{})
	log.Println("Starting gRPC server on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
