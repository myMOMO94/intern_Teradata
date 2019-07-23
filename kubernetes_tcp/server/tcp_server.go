package main

import "net"
import "fmt"
import "log"
import "bufio"
//import "os"
import "flag"
//import "strconv"
//import "io"

// convert string bytes info to integer
func Atoi (s string) int {
  var (
    n uint64
    i int
    v byte
  )

  for ; i < len(s); i++ {
    d := s[i]
    if '0' <= d && d <= '9' {
      v = d - '0'
    } else {
      break
    }
    n *= uint64(10)
    n += uint64(v)

  }
  //fmt.Println("int number: ", int(n))
  return int(n)
}

func main() {
  fmt.Println("Launching Server...")

  /*if (len(os.Args) != 2) {
    fmt.Println("Please follow the pattern ./tcp_server -interactive=true, not the exact number of arguments as expected!")
    return
  }*/

  /*numruns, err_run := strconv.Atoi(os.Args[3])
  if (err_run != nil) {
    log.Fatal(err_run)
  }*/

  //parse interactive flag
  interactivePtr := flag.Bool("interactive", false, "a bool")
  verbosePtr := flag.Bool("verbose", false, "a bool")
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
    if *verbosePtr == true {
      fmt.Println("SERVER: You are in the interactive mode, which allows you to type in a payload.")
    }
    for {
      // will listen for message to process ending in newline delim
      message, err_rec := bufio.NewReader(conn).ReadString('\n')
      if (err_rec != nil) {
        log.Fatal(err_rec)
      }
      if *verbosePtr == true {
        fmt.Print("SERVER: Message received: ", string(message))
      }

      //send message back to client
      _, err_send := conn.Write([]byte(string(message) + "\n"))
      if (err_send != nil) {
        log.Fatal(err_send)
      }
    }
  } else {
    if *verbosePtr == true {
      fmt.Println("SERVER: You are in the non-interactive mode!")
    }
    for{
      //will listen for message from client
      //get number of bytes info first
      nbytes, err_bytes := bufio.NewReader(conn).ReadString('\n')
      if (err_bytes != nil) {
        log.Fatal(err_bytes)
      }
      numbytes := Atoi(string(nbytes))
      if *verbosePtr == true {
        fmt.Println("SERVER: Received bytes info:  ", numbytes, " bytes. ")
      }

      // tell client that server received bytes info
      _, err_byt := conn.Write([]byte(string("Received bytes info.") + "\n"))
      if (err_byt != nil) {
        log.Fatal(err_byt)
      }
      if *verbosePtr == true {
        fmt.Println("SERVER: Sent received bytes info message back.")
      }

      //will listen for message from client
      data := make([]byte, numbytes)//32*1024)
      numBytes, err_read := conn.Read(data)
      if (err_read != nil) {
        log.Fatal(err_read)
      }
      for ; numBytes < numbytes; {
        n, err := conn.Read(data)
        if (err != nil) {
          log.Fatal(err)
        }
        numBytes += n
      }
      data = nil
      if *verbosePtr == true {
        fmt.Println("SERVER: ", numBytes, " bytes data received from client")
      }

      //send response message back to client
      message := make([]byte, numBytes)
      numbyte, err_sent := conn.Write(message)
      if (err_sent != nil) {
        log.Fatal(err_sent)
      }
      message = nil
      if *verbosePtr == true {
        fmt.Println("SERVER: sent ", numbyte, " bytes back to client.")
      }
    }
  }
}
