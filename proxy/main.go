package main

import (
	"context"
    "log"
    "net"
    "net/http"
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"
    gen "github.com/Alnomrosi/sovd-grpc-gateway/gen/go/applications_manager"
)

func main() {
	// creating mux for gRPC gateway. This will multiplex or route request different gRPC service
	//  The mux knows how to translate HTTP requests into gRPC calls.
    mux:=runtime.NewServeMux()

	// setting up a dail up for gRPC service by specifying endpoint/target url
    err := gen.RegisterAppLauncherHandlerFromEndpoint(context.Background(), mux, "localhost:50052", []grpc.DialOption{grpc.WithInsecure()})
    if err != nil {
        log.Fatal(err)
    }

    // Creating a SOVD HTTP server
    server:=http.Server{
        Handler: mux,
    }

	// creating a listener for SOVD server 
    l,err:=net.Listen("tcp",":8081")
    if err!=nil {
        log.Fatal(err)
    }

	// start SOVD server
    err = server.Serve(l)
    if err != nil {
        log.Fatal(err)
    }
}