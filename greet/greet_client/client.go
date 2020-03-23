package main

import (
	"context"
	"fmt"
	"github.com/suliar/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"io"
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

	//doUnary(c)
	doServerStreaming(c)
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


func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("starting to do stream rpc...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Suli",
			LastName:  "Arubi",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err:= resStream.Recv()
		if err == io.EOF {
			break
	}
	if err != nil {
		log.Fatalf("error while reading stream: %v", err)
	}
	log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
}
}