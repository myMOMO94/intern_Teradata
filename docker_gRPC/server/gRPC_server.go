package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "google.golang.org/grpc"
    "intern_Teradata/docker_gRPC/bytearray"
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
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
    if err != nil {
        log.Fatal("failed to listen: %v", err)
    }
    var opts []grpc.ServerOption
    grpcServer := grpc.NewServer(opts...)
    bytearray.RegisterByteArrayServer(grpcServer, newServer())
    grpcServer.Serve(lis)
}
