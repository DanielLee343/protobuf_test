package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/daniellee343/grpc-demo-yt/proto"
)
func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}
	waitc := make(chan struct{})
	go func ()  {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("some error handling streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{//create pb request
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending %v", err)
		}
		log.Printf("Sent the request with name: %s", name)
		time.Sleep(2 * time.Second)
	}
	
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional streaming finished")
}