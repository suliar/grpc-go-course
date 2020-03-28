package main

import (
	"context"
	"fmt"
	"github.com/suliar/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked: %v", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName

	return &greetpb.GreetResponse{
		Result: result,
	}, nil
}

func (s *server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked: %v", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		err := stream.Send(res)
		if err != nil {
			return err
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return nil

}

func (s *server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet function was invoked with a streaming request\n")

	result := "Hello "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// basically means we have finished reading the client stream
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})

		}
		if err != nil {
			log.Fatalf("Error whilst reading client stream: %v", err)
		}

		firstName := req.Greeting.GetFirstName()

		result += firstName + "! "

	}
	return nil
}

func main() {
	fmt.Println("Hello World")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
