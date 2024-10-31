package main

import (
	"log"

	pb "github.com/kontentski/grpc1/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to %v:", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	/* 	names := &pb.NamesList{
		Names: []string{"John", "Doe", "bruh.."},
	} */

	callSayHello(client)
	//	callSayHelloServerStream(client, names)
	//	callSayHelloClientStream(client, names)
	//	callSayHelloBidirectionalStreaming(client, names)

}
