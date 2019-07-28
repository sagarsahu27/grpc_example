package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sagarsahu27/myGrpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Println("Received request from ", req)
	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()
	result := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Welcome to calculator Service ...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Couldn't list to given port %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to server %v", err)
	}
}
