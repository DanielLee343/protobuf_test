package main

import (
	"context"
	"log"
	"time"

	pb "github.com/daniellee343/grpc-demo-yt/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})// here is calling the SayHello in server unary.go
	if err != nil {
		log.Fatalf("Could not greet: %s", err)
	}
	log.Printf("%s", res.Message)
}

