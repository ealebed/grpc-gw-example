package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	hello "./proto/hello"
)

var (
	endpoint = flag.String("endpoint", "localhost:50051", "address")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := hello.RegisterSayHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":80", mux)
}

func main() {
	flag.Parse()

	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
