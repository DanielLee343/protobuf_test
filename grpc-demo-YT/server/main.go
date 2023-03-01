package main

import (
	"log"
	"net"

	pb "github.com/daniellee343/grpc-demo-yt/proto" // 1
	"google.golang.org/grpc"
)

const(
	port = ":1097"
)

type helloServer struct { //3
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen to port: %s", err)
	}
	//create a grpc server
	grpcServer := grpc.NewServer()
	//use protobuf to resiger the service
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{}) // 2
	log.Printf("server started at %v\n", lis.Addr())
	//start the server with the port
	if err := grpcServer.Serve(lis); err != nil { //4
		log.Fatalf("Failed to start the server: %s", err)
	}
}