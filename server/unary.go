package main

import (
	"context"
	"fmt"

	pb "github.com/kontentski/grpc1/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.Empty) (*pb.HelloResponse, error) {
	fmt.Println("resieved request")
	return &pb.HelloResponse{
		Message: "Hello, World",
	}, nil
}
