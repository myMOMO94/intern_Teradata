package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "google.golang.org/grpc"
    "intern_Teradata/gRPC_tracing/bytearray"

    //gwrapper "intern_Teradata/jaeger_gRPC_wrapper"
    gwrapper "intern_Teradata/gRPC_tracing/jaeger_gRPC_wrapper"
)

type byteArrayServer struct {
    bytearray.UnimplementedByteArrayServer
}

func newServer() *byteArrayServer{
    s := &byteArrayServer{}
    return s
}

func (s *byteArrayServer) SendRequest(ctx context.Context, req *bytearray.ByteRequest) (*bytearray.ByteReply, error) {
    fmt.Println("Received ", len(req.Request), "bytes request from client, sending it back.")
    return &bytearray.ByteReply{Reply: req.Request}, nil
}

func main() {
    fmt.Println("Launching server...")

    var servOpts []grpc.ServerOption
    tracer, closer := gwrapper.InitTracer("server", "127.0.0.1:6831")//"153.64.12.154:6831")
    defer closer.Close()

    if tracer != nil {
        servOpts = append(servOpts, gwrapper.ServerOption(tracer))
    }

    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
    if err != nil {
        log.Fatal("failed to listen: %v", err)
    }
    //var opts []grpc.ServerOption
    grpcServer := grpc.NewServer(servOpts...)
    bytearray.RegisterByteArrayServer(grpcServer, newServer())
    grpcServer.Serve(lis)
}
