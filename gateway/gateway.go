package gateway

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	helloworldpb "github.com/zazin/test-proto-grpc-gw/proto"
	"google.golang.org/grpc"
	"io"
	"net/http"
)

func homePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("grpc post message")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"alive": "yes"}`)
	}
}

func New(ctx context.Context, helloEndpoint string) (http.Handler, error) {
	gw := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := helloworldpb.RegisterGreeterHandlerFromEndpoint(ctx, gw, helloEndpoint, opts); err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.Handle("/test/", homePage())
	mux.Handle("/", gw)
	return mux, nil
}
