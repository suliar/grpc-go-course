package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/suliar/grpc-go-course/calculator/sumpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}

func (s *server) DoSum(ctx context.Context, req *sumpb.DoSumRequest) (*sumpb.DoSumResponse, error) {
	sum := req.Sum
	if sum.FirstNumber < 0 {
		return nil, errors.New("wrong number")
	}

	if sum.SecondNumber < 0 {
		return nil, errors.New("wrong number")
	}

	res := sum.GetFirstNumber() + sum.GetSecondNumber()

	rest := &sumpb.DoSumResponse{
		Result: res,
	}
	return rest, nil
}

func main() {
	fmt.Println("Hello World")
	lis, err :=  net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s:= grpc.NewServer()
	sumpb.RegisterSumApiServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}