package main

import (
	"io"
	"log"

	pb "github.com/daniellee343/grpc-demo-yt/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) (error) {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Reach eof")
			return err
		}
		if err != nil {
			return nil
		}
		log.Printf("Got request with name: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello "+ req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}