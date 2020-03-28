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

func (s *server) PrimeNumberDecom(req *sumpb.PrimeNumberRequest, stream sumpb.SumApi_PrimeNumberDecomServer) error {
	fmt.Printf("PrimeNumberDecom function was invoked: %v", req)
	num := req.GetPrimeNo()
	divisor := int64(2)

	// Get the number of 2s that divide num
	for num > 1 {
		if num% 2 == 0 {
			stream.Send(&sumpb.PrimeNumberResponse{
				Result: divisor,
			})
			num = num / divisor
		} else {
			divisor++
			fmt.Printf("divisor has increased: %v", divisor)
		}

	}
	return nil
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