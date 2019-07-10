package main

import "net"
import "os"
import "bufio"
import "fmt"
import "log"
import "flag"
import "strconv"
import "time"
//import "io/ioutil"

func main() {
  //initialize connection
  //choose the mode, either interactive or not interactive
  interactivePtr := flag.Bool("interactive", false, "a bool")
  verbosePtr := flag.Bool("verbose", false, "a bool")
  flag.Parse()

  if (len(os.Args) != 6) {
    fmt.Println("CLIENT: Please follow the pattern ./tcp_client -interactive=true 127.0.0.1:8080 256, not the exact number of arguments as expected!")
    return
  }

  //find the ip address and num of bytes
  address := os.Args[3]
  numbytes, err_size := strconv.Atoi(os.Args[4])
  /*if (numbytes > 32*1024) {
    fmt.Println("Out of range! Maximum allowed size is 32KB")
    return
  }*/
  if (err_size != nil) {
   //fmt.Println("Please check the data size want to be transferred, maximum allowed size is 32KB")
   log.Fatal(err_size)
  }

  numruns, err_run := strconv.Atoi(os.Args[5])
  if (err_run != nil) {
    log.Fatal(err_run)
  }

  //connect to server
  conn, err := net.Dial("tcp", address) //"127.0.0.1:8080")
  fmt.Println("CLIENT: Connected to Ip: " + address)
  if (err != nil) {
    log.Fatal(err)
  }

  //interactive flag is true
  if *interactivePtr == true {
    if *verbosePtr == true {
      fmt.Println("CLIENT: You are in the interactive mode, which allows you to type in a payload.")
      fmt.Println("CLIENT: Press Enter if you want to exit.")
    }
    for {
      //read input message from stdin
      reader := bufio.NewReader(os.Stdin)
      fmt.Print("CLIENT: Text to send: ")
      text, err_text := reader.ReadString('\n')
      if (err_text != nil) {
        log.Fatal(err_text)
      }
      //fmt.Println(string(text))
      //if there is nothing there, break the connection
      if (string(text) == "\n") {
        fmt.Println("CLIENT: Nothing want to send to server, close the connection! ")
        break
      }

      //send message to server
      nbytes, err_sent := fmt.Fprintf(conn, text + "\n")
      if (err_sent != nil) {
        log.Fatal(err_sent)
      }
      if *verbosePtr == true {
        fmt.Print("CLIENT: Sent ", nbytes, " bytes Message to server: " + text)
      }

      //Listen for sever reply
      message, err_read := bufio.NewReader(conn).ReadString('\n')
      if (err_read != nil) {
        log.Fatal(err_read)
      }
      if *verbosePtr == true {
        fmt.Print("CLIENT: Message from server: " + message)
      }
    }
  } else {
    //interactive flag is false
    if *verbosePtr == true {
      fmt.Println("CLIENT: You are in the non-interactive mode!")
    }
    for i := 0 ; i < numruns; i++ {
      //send message to server
      // send number of bytes info to server first
      str_numbytes := strconv.Itoa(numbytes)
      _, err_bytes := fmt.Fprintf(conn, str_numbytes + "\n")
      if (err_bytes != nil) {
        log.Fatal(err_bytes)
      }
      if *verbosePtr == true {
        fmt.Println("CLIENT: Sent number of bytes info to server: " + str_numbytes)
      }

      byte_message, err_byt := bufio.NewReader(conn).ReadString('\n')
      if (err_byt != nil) {
        log.Fatal(err_byt)
      }
      if *verbosePtr == true {
        fmt.Print("CLIENT: server sent back: " + byte_message)
      }

      start := time.Now()
      // send message to server
      message := make([]byte, numbytes)
      nbytes, err_sent := conn.Write(message)
      if (err_sent != nil) {
        log.Fatal(err_sent)
      }
      message = nil
      end := time.Since(start)
      fmt.Println("CLIENT: It took ", end, " to sent ", nbytes, " bytes Message to server.")
      //fmt.Println("CLIENT: Sent ", nbytes, " bytes Message to server")

      read_start := time.Now()
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
      read_end := time.Since(read_start)
      fmt.Println("CLIENT: It took ", read_end, " to receive ", numBytes, " bytes data from server after send message to server.")
      if *verbosePtr == true {
        fmt.Println("ClIENT: It took ", end + read_end, " total.")
      }
    }
  }
}
