package main

import (
	"io"
	"log"

	pb "github.com/kontentski/grpc1/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Message: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("got request with name %s", req.Name)
		messages = append(messages, "Hello", req.Name)
	}
}
