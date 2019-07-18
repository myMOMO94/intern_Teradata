package main

import (
    "context"
    "log"
    "fmt"
    "google.golang.org/grpc"
    "intern_Teradata/docker_gRPC/bytearray"
    "time"
    "os"
    "flag"
    "strconv"
)

func getReply(client bytearray.ByteArrayClient, req *bytearray.ByteRequest, numbytes int, verbose bool) {
    if verbose == true {
        fmt.Println("Getting echo message for", numbytes, " bytes request from server")
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()
    //data := make([]byte, 512)
    reply, err := client.SendRequest(ctx, req)//&bytearray.ByteRequest{Request: data})
    if err != nil {
        log.Fatalf("%v.SendRequest(_) = _, %v", client, err)
    }
    fmt.Println("Received ", len(reply.Reply), "bytes reply from server.")
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

    var conn *grpc.ClientConn

    conn, err := grpc.Dial("gRPC-server:8080", grpc.WithInsecure())
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