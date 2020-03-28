package main

import (
	"context"
	"fmt"
	"github.com/suliar/grpc-go-course/calculator/sumpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
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
	//doStreamSum(c)
	doClientStreamAverage(c)

}

func doSum(c sumpb.SumApiClient) {
	req := &sumpb.DoSumRequest{
		Sum: &sumpb.Sum{
			FirstNumber:  3,
			SecondNumber: 10,
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

func doClientStreamAverage(c sumpb.SumApiClient) {
	fmt.Println("starting to do a client streaming rpc")

	numbers := []int64{1, 1, 1, 1, 1}

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("error while calling Compute Average: %v", err)
	}

	for _, number := range numbers {
		time.Sleep(1000 * time.Millisecond)
		err = stream.Send(&sumpb.ComputeAverageRequest{
			Number: number,
		})
		if err != nil {
			log.Fatalf("fail to send request: %v\n error is :%v", numbers, err)
		}

	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving Compute Average: %v", err)
	}

	fmt.Printf("Compute Average Response: %v", res)

}
