package main

import "net"
import "fmt"
import "log"
import "bufio"

func main() {
  fmt.Println("Launching Server...")
  
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
  }
}
