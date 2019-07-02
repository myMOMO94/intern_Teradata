package main

import "net"
import "os"
import "bufio"
import "fmt"
import "log"
import "flag"
import "strconv"

func main() {
  //initialize connection
  //choose the mode, either interactive or not interactive
  interactivePtr := flag.Bool("interactive", false, "a bool")
  flag.Parse()

  //find the ip address and num of bytes
  address := os.Args[2]
  numbytes,_ := strconv.Atoi(os.Args[3])

  //connect to server
  conn, err := net.Dial("tcp", address) //"127.0.0.1:8080")
  fmt.Println("Connected to Ip: " + address)
  if (err != nil) {
    log.Fatal(err)
  }

  //interactive flag is true
  if *interactivePtr == true {
    for {
      //read input message from stdin
      reader := bufio.NewReader(os.Stdin)
      fmt.Print("Text to send: ")
      text, err_text := reader.ReadString('\n')
      if (err_text != nil) {
        log.Fatal(err_text)
      }

      //send message to server
      nbytes, err_sent := fmt.Fprintf(conn, text + "\n")
      if (err_sent != nil) {
        log.Fatal(err_sent)
      }
      fmt.Print("Sent ", nbytes, " bytes Message to server: " + text)

      //Listen for sever reply
      message, err_read := bufio.NewReader(conn).ReadString('\n')
      if (err_read != nil) {
        log.Fatal(err_read)
      }
      fmt.Print("Message from server: " + message)
    }
  } else {
    //interactive flag is false
    data := make([]byte, numbytes)
    nbytes, err_sent := conn.Write(data)
    if (err_sent != nil) {
      log.Fatal(err_sent)
    }
    fmt.Println("Sent ", nbytes, " bytes Message to server")

    //Listen for server reply
    /*_, err_read := bufio.NewReader(conn).ReadString('\n')
    if (err_read != nil) {
      log.Fatal(err_read)
    }*/
  }
}
