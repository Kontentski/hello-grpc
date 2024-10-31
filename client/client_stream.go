package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kontentski/grpc1/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not start client streaming: %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("could not send name: %v", err)
		}
		log.Printf("sent the request with the name %s", name)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("client streaming finished")
	if err != nil {
		log.Fatalf("could not receive response: %v", err)
	}
	log.Printf("%s", res.Message)
}
