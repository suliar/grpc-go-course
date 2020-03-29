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

	//doCreateBlog(c)
	//doReadBlog(c)
	//doUpdateBlog(c)
	doDeleteBlog(c)
}

func doCreateBlog(c blogpb.BlogServiceClient) {
	fmt.Println("starting to do unary rpc...")
	req := &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			AuthorId: "Ange",
			Title:    "Second blog content",
			Content:  "Currently playing weird songs",
		},
	}

	fmt.Println("Creating Blog")
	res, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to created blog: %v", err)
	}
	fmt.Printf("Blog has been created: %v", res)
}

func doReadBlog(c blogpb.BlogServiceClient) {
	fmt.Println("starting to do unary rpc...")
	req := &blogpb.ReadBlogRequest{
		BlogId: "5e80ab03dc07445c58faf713",
	}

	fmt.Println("Reading Block")
	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to read blog: %v", err)
	}
	fmt.Printf("Blog Found: %v", res.Blog)
}

func doUpdateBlog(c blogpb.BlogServiceClient) {
	fmt.Println("starting to do unary rpc...")
	req := &blogpb.UpdateBlogRequest{
		Blog: &blogpb.Blog{
			Id:       "5e80ab03dc07445c58faf713",
			AuthorId: "SuliA",
			Title:    "Updated",
			Content:  "The content is fuck off",
		},
	}

	fmt.Println("Updating Blog")
	res, err := c.UpdateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to read blog: %v", err)
	}
	fmt.Printf("Blog Found: %v", res.Blog)
}

func doDeleteBlog(c blogpb.BlogServiceClient) {
	fmt.Println("starting to do unary rpc...")
	req := &blogpb.DeleteBlogRequest{
		BlogId: "5e80a9b5dc07445c58faf710",
	}

	fmt.Println("Deleting Blog")
	res, err := c.DeleteBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to read blog: %v", err)
	}
	fmt.Printf("Blog Deleted: %v", res.BlogId)
}
