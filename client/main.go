package main

import (
	"context"
	"examples/grpc/proto/blog"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	blogService := blog.NewBlogServiceClient(conn)

	getResp, err := blogService.GetBlog(context.Background(), &blog.GetBlogRequest{BlogId: 1})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(getResp.BlogPost)
	}

	_, err = blogService.DeleteBlog(context.Background(), &blog.DeleteBlogRequest{BlogId: 1})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("blog deleted")
	}
}
