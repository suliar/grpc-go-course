package main

import (
	"context"
	"fmt"
	"github.com/suliar/grpc-go-course/calculator/sumpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Hello I am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := sumpb.NewSumApiClient(cc)
	//doSum(c)
	doStreamSum(c)

}

func doSum(c sumpb.SumApiClient){
	req := &sumpb.DoSumRequest{
		Sum:                  &sumpb.Sum{
			FirstNumber:          3,
			SecondNumber:         10,
		},
	}

	res, err := c.DoSum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while callling sum rpc: %v", err)
	}
	log.Printf("response from Sum: %v", res.Result)

}

func doStreamSum(c sumpb.SumApiClient) {
	req := &sumpb.PrimeNumberRequest{
		PrimeNo: 12,
	}

	resStream, err := c.PrimeNumberDecom(context.Background(), req)
	if err != nil {
		log.Fatalf("error while callling sum rpc: %v", err)
	}
	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something wrong happened: %v", err)
		}

		fmt.Println(res.GetResult())
	}

}