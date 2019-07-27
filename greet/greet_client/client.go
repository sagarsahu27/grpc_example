package main

import (
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
	fmt.Println("Client created %s", c)
}
