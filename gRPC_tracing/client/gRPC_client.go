package main

import (
    "context"
    "log"
    "fmt"
    "google.golang.org/grpc"
    "time"
    "os"
    "flag"
    "strconv"

    //gwrapper "intern_Teradata/jaeger_gRPC_wrapper"
    "intern_Teradata/gRPC_tracing/bytearray"
    gwrapper "intern_Teradata/gRPC_tracing/jaeger_gRPC_wrapper"
)

func getReply(client bytearray.ByteArrayClient, req *bytearray.ByteRequest, numbytes int, verbose bool) {
    if verbose == true {
        fmt.Println("Getting echo message for", numbytes, " bytes request from server")
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()
    //data := make([]byte, 512)
    start := time.Now()
    reply, err := client.SendRequest(ctx, req)//&bytearray.ByteRequest{Request: data})
    if err != nil {
        log.Fatalf("%v.SendRequest(_) = _, %v", client, err)
    }
    end := time.Since(start)
    fmt.Println("Received ", len(reply.Reply), "bytes reply from server within ", end)
}

func main() {
    verbosePtr := flag.Bool("verbose", false, "a bool")
    flag.Parse()

    numbytes, err_size := strconv.Atoi(os.Args[2])
    if err_size != nil {
        log.Fatalf("fail to read numbytes: %v", err_size)
    }

    numruns, err_runs := strconv.Atoi(os.Args[3])
    if err_runs != nil {
        log.Fatalf("fail to read numrunsL %v", err_runs)
    }

    dialOpts := []grpc.DialOption{grpc.WithInsecure()}
    jaeger_agent := string(os.Getenv("JAEGER_AGENT_HOST") + ":6831")
    fmt.Println(jaeger_agent)
    tracer, closer := gwrapper.InitTracer("client", "192.168.99.104:6831")//jaeger_agent)/*os.Getenv("JAEGER_AGENT_HOST"))*/"127.0.0.1:6831")
    defer closer.Close()

    if tracer != nil {
        dialOpts = append(dialOpts, gwrapper.DialOption(tracer))
    } else {
        fmt.Println("tracer is nil")
    }

    var conn *grpc.ClientConn

    //conn, err := grpc.Dial("gRPC-server:8080", dialOpts...)
    conn, err := grpc.Dial("127.0.0.1:8080"/*"grpcserver:8080"*/, dialOpts...)
    if err != nil {
        log.Fatalf("fail to dial: %v", err)
    }
    defer conn.Close()

    client := bytearray.NewByteArrayClient(conn)
    if *verbosePtr == true {
        //addrs, err := os.Hostname()
        addrs := "gRPC_server"
        //if err != nil {
            fmt.Println("Connected to server: ", addrs)
        //} else {
            //log.Fatal(err)
        //}
    }
    data := make([]byte, numbytes)
    for i := 0; i < numruns; i++ {
      getReply(client, &bytearray.ByteRequest{Request: data}, numbytes, *verbosePtr)
    }
}
