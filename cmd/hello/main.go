package main

import (
	"context"
	helloworldpb "github.com/zazin/test-proto-grpc-gw/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const portServerHello = ":50051"

type server struct{}

func (s *server) mustEmbedUnimplementedGreeterServer() {
	panic("implement me")
}

func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{Message: in.Name + " world"}, nil
}

func main() {
	lis, err := net.Listen("tcp", portServerHello)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	s := grpc.NewServer()
	helloworldpb.RegisterGreeterServer(s, &server{})
	log.Println("Serving gRPC on 0.0.0.0:" + portServerHello)
	log.Fatal(s.Serve(lis))
}
