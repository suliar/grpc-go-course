package main

import (
	"context"
	"fmt"
	"github.com/suliar/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
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
	//doServerStreaming(c)
	//doClientStreaming(c)
	doBiDiStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("starting to do unary rpc...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Suli",
			LastName:  "Arubi",
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
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("starting to do a client streaming rpc")

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Suli",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Ayo",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Suli",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling long greet: %v", err)
	}

	// we iterate over our slice and send it message individually
	for _, req := range requests {
		fmt.Printf("Sending request\n: %v", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err:= stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiveing long greet: %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)

}


func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("starting to do a bidirectional streaming rpc")

	waitC := make(chan struct{})

	requests := []*greetpb.GreetEveryoneRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Suli",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Ayo",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Suli",
			},
		},
	}

	// we create the stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	// we send a bunch of messages to the client (go routines)
	go func() {
		// function to send bunch of messages

		for _, req := range requests {
			fmt.Printf("sending message %v", req)
			err = stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
			if err != nil {
				log.Fatalf("Cannot send request: %v", err)
			}
		}
		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("Error closing stream: %v", err)
		}
		close(waitC)
	}()

	// we receive a bunch of messages from the client (go routine)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving from stream: %v", err)
				break
			}

			fmt.Printf("Received\n: %v", res.GetResult())
		}
		close(waitC)

	}()

	// block until everything is done
	<-waitC

}