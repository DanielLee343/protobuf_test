package main

import (
	"log"

	pb "github.com/daniellee343/grpc-demo-yt/proto" // 1
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const(
	port = ":1097"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	//create a client obj
	client := pb.NewGreetServiceClient(conn)
	names := &pb.NamesList{
		Names: []string{"A", "B", "C"},
	}
	// callSayHello(client)
	// callSayHelloServerStream(client, names)
	// callSayHelloClientStream(client, names)
	callSayHelloBidirectionalStream(client, names)
}