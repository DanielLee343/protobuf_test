package chat

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s with uuid: %s \n", message.Body, message.Uuid)
	return &Message{Body: "Hello from the Servier!", Uuid: message.Uuid}, nil
}
