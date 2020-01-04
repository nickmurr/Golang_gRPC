package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go_grpc_server/blog/blogpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello, i'm a blog client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// createBlogItem(c)
	// readBlogItem(c)
	readAllBlogItem(c)
}

func createBlogItem(c blogpb.BlogServiceClient) {
	fmt.Println("Starting to do unary RPC")

	req := &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			Author:  "Allen",
			Title:   "Third Article",
			Content: "Hello, it's a third article",
		},
	}

	res, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	bs, err := json.MarshalIndent(res.Blog, "", " ")
	if err != nil {
		log.Fatalf("Error while unmarshal", err)
	}

	log.Printf("Response from Greet: %v", fmt.Sprintf(string(bs)))
}

func readBlogItem(c blogpb.BlogServiceClient) {
	fmt.Println("Start Read blog")

	req := &blogpb.ReadBlogRequest{
		BlogId: "5e0e7311a275e780cdd7ba85",
	}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while reading a blog: %v", err)
	}

	log.Printf("Blog: %v", res)
}

func readAllBlogItem(c blogpb.BlogServiceClient) {
	fmt.Println("Start Read blog")

	req := &blogpb.ReadAllBlogRequest{
		Search: "",
		Page:   1,
	}

	res, err := c.ReadAllBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while reading a blog: %v", err)
	}

	form, _ := json.MarshalIndent(res, "", " ")

	log.Printf("Blog: %v", string(form))
}
