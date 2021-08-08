package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zazin/test-proto-grpc-gw/gateway"
	helloworldpb "github.com/zazin/test-proto-grpc-gw/proto"
	"log"
	"net/http"
)

var (
	helloworldEndpoint = flag.String("hello_endpoint", "localhost:50051", "endpoint of GreeterServer")
	port               = flag.Int("p", 8081, "port of the service")
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{Message: in.Name + " world"}, nil
}

func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux, err := gateway.New(ctx, *helloworldEndpoint)
	if err != nil {
		log.Printf("Setting up the gateway: %s", err.Error())
		return
	}

	srvAddr := fmt.Sprintf(":%d", *port)

	s := &http.Server{
		Addr:    srvAddr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		log.Printf("Shutting down the http server")
		if err := s.Shutdown(context.Background()); err != nil {
			log.Printf("Failed to shutdown http server: %v", err)
		}
	}()

	log.Printf("Starting listening at %s", srvAddr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("Failed to listen and serve: %v", err)
	}
}
