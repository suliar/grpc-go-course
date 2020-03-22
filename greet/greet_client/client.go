package main

import (
	"context"
	"fmt"
	"github.com/suliar/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello, I am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("starting to do unary rpc...")
	req := &greetpb.GreetRequest{
		Greeting:             &greetpb.Greeting{
			FirstName:            "Suli",
			LastName:             "Arubi",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while called Greet rpc: %v", err)
	}
	log.Printf("response from Greet: %v", res.Result)
}
