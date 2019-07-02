package main

import "net"
import "fmt"
import "log"
import "bufio"
import "io/ioutil"
import "flag"

func main() {
  fmt.Println("Launching Server...")

  //parse interactive flag
  interactivePtr := flag.Bool("interactive", false, "a bool")
  flag.Parse()

  //listen on all interfaces
  ln, err := net.Listen("tcp", ":8080")
  if (err != nil) {
    // handle error
    log.Fatal(err)
  }

  //accept connection on port
  conn, err_conn := ln.Accept()
  if (err_conn != nil) {
    // handle error
    log.Fatal(err_conn)
  }

  if *interactivePtr == true {
    for {
      // will listen for message to process ending in newline delim
      message, err_rec := bufio.NewReader(conn).ReadString('\n')
      if (err_rec != nil) {
        log.Fatal(err_rec)
      }
      fmt.Print("Message received: ", string(message))

      //send message back to client
      _, err_send := conn.Write([]byte(string(message) + "\n"))
      if (err_send != nil) {
        log.Fatal(err_send)
      }
      /*const MaxUint = ^uint(0)
      const MaxInt = int(MaxUint >> 1)
      data := make([]byte, MaxInt)
      numBytes, err_read := conn.Read(data)
      if (err_read != nil) {
        log.Fatal(err_read)
      }
      data = data[:0]
      fmt.Println(numBytes, "bytes data received from client")

      //send response message back to client
      message := make([]byte, numBytes)
      _, err_sent := conn.Write(message)
      if (err_sent != nil) {
        log.Fatal(err_sent)
      }
      message = nil*/
    }
  } else {
    reader := bufio.NewReader(conn)
    b, err := ioutil.ReadAll(reader)
    if (err != nil) {
      log.Fatal(err)
    }
    fmt.Println(len(b), " bytes data received from client.")

    //sent response back to client
    data := make([]byte, len(b))
    _, err_sent := conn.Write(data)
    if (err_sent != nil) {
      log.Fatal(err_sent)
    }
  }
}
