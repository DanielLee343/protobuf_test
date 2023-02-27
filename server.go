package main

import (
	"log"
	"net"

	"protobuf_test/chat"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":1097")
	if err != nil {
		log.Fatalf("Failed to listen on port 1097: %v", err)
	}
	grpcServer := grpc.NewServer()
	s := chat.Server{}

	chat.RegisterChatServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
