package main

import (
	"log"
	"net"

	pb "github.com/kontentski/grpc1/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("gRPC server started on port %s", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start the grpc server: %v", err)
	}
}
