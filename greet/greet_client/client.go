package main

import (
	"context"
	"fmt"
	"log"

	greetpb "github.com/sagarsahu27/myGrpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello World !!")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Cannot connect to server %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//	fmt.Println("Client created %s", c)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sagar",
			LastName:  "Sahu",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling Greet RPC: ", err)
	}

	log.Println("Response from Greeting: ", res.Result)
}
