package main

import (
	"io"
	"log"
	"time"

	pb "github.com/kontentski/grpc1/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("received request with name %s", req.Name)

		res := &pb.HelloResponse{Message: "Hello, " + req.Name}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
}
