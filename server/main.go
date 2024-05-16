package main

import (
	"context"
	"examples/grpc/proto/blog"

	"log"
	"net"

	"google.golang.org/protobuf/types/known/timestamppb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	convertedBlogs := make(map[int32]*blog.BlogPost)
	for key, value := range blogs {
		convertedBlogs[int32(key)] = value
	}

	blog.RegisterBlogServiceServer(grpcServer, Server{blogs: convertedBlogs})
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}

type Server struct {
	blog.UnimplementedBlogServiceServer
	blogs map[int32]*blog.BlogPost
}

func (s Server) GetBlog(ctx context.Context, request *blog.GetBlogRequest) (*blog.GetBlogResponse, error) {
	b, ok := s.blogs[int32(request.BlogId)]
	if !ok {
		return nil, status.Error(codes.NotFound, "blog not found")
	}
	response := &blog.GetBlogResponse{
		BlogPost: b,
	}
	return response, nil
}

func (s Server) UpsertBlog(ctx context.Context, request *blog.UpsertBlogRequest) (*blog.UpsertBlogResponse, error) {
	s.blogs[int32(request.BlogPost.BlogId)] = request.BlogPost
	response := &blog.UpsertBlogResponse{
		BlogPost: s.blogs[int32(request.BlogPost.BlogId)],
	}
	return response, nil
}

func (s Server) DeleteBLog(ctx context.Context, request *blog.DeleteBlogRequest) (*blog.DeleteBlogResponse, error) {
	if _, ok := s.blogs[int32(request.BlogId)]; !ok {
		return nil, status.Error(codes.NotFound, "blog not found")
	}
	delete(s.blogs, int32(request.BlogId))
	return &blog.DeleteBlogResponse{}, nil
}

var blogs = map[uint32]*blog.BlogPost{
	1: {
		BlogId:       1,
		AuthorId:     123,
		TourId:       456,
		Title:        "First Blog Post",
		Description:  "This is the first blog post.",
		CreationDate: timestamppb.Now(),
		ImageUrls:    "https://example.com/image.jpg",
		Status:       blog.BlogPostStatus_PUBLISHED,
	},
	2: {
		BlogId:       2,
		AuthorId:     456,
		TourId:       789,
		Title:        "Second Blog Post",
		Description:  "This is the second blog post.",
		CreationDate: timestamppb.Now(),
		ImageUrls:    "https://example.com/image2.jpg",
		Status:       blog.BlogPostStatus_DRAFT,
	},
}
