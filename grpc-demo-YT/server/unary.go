package main

import (
	"context"

	pb "github.com/daniellee343/grpc-demo-yt/proto"
)

// The server overwrites the Sayhello interface defined in the generated proto file
func(s *helloServer) SayHello(ctx context.Context, req *pb.NoParam)(*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}