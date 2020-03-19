package main

import (
	"fmt"
	"google.golang.org/grpc"
	"github.com/suliar/grpc-go-course/greet/greetpb"
	"log"
	"net"
)

func main() {
	fmt.Println("Hello World")
	lis, err :=  net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s:= grpc.NewServer()
}
