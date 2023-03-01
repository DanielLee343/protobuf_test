package main

import (
	"context"
	"log"
	"protobuf_test/chat"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":1097", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	msg := chat.Message{
		Body: "Hello from the client!",
		Uuid: "wfewfwbwgow",
	}
	response, err := c.SayHello(context.Background(), &msg)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from Server: %s", response.Body)
}
