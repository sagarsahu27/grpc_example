package main

import (
	"context"
	"fmt"
	"log"

	calculatorpb "github.com/sagarsahu27/myGrpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello World !!")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Cannot connect to server %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	//	fmt.Println("Client created %s", c)
	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  23,
		SecondNumber: 22,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling Greet RPC: ", err)
	}

	log.Println("Response from Greeting: ", res.SumResult)
}
