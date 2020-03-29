package main

import (
	"context"
	"fmt"
	"github.com/suliar/grpc-go-course/blog/blogpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello, I am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	doUnary(c)
}

func doUnary(c blogpb.BlogServiceClient) {
	fmt.Println("starting to do unary rpc...")
	req := &blogpb.CreateBlogRequest{
		Blog:                 &blogpb.Blog{
			AuthorId:             "Ange",
			Title:                "Second blog content",
			Content:              "Currently playing weird songs",
		},
	}

	fmt.Println("Creating Blog")
	res, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to created blog: %v", err)
	}
	fmt.Printf("Blog has been created: %v", res)
}
