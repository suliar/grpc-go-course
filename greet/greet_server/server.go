package main

import (
	"context"
	"fmt"
	"github.com/suliar/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

}

func main() {
	fmt.Println("Hello World")
	lis, err :=  net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s:= grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
