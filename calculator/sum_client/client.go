package main

import (
	"context"
	"fmt"
	"github.com/suliar/grpc-go-course/calculator/sumpb"
	"google.golang.org/grpc"
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
	doSum(c)

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