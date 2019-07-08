package main

import "net"
import "os"
import "bufio"
import "fmt"
import "log"
import "flag"
import "strconv"
//import "io/ioutil"

func main() {
  //initialize connection
  //choose the mode, either interactive or not interactive
  interactivePtr := flag.Bool("interactive", false, "a bool")
  flag.Parse()

  if (len(os.Args) != 4) {
    fmt.Println("Please follow the pattern ./tcp_client -interactive=true 127.0.0.1:8080 256, not the exact number of arguments as expected!")
    return
  }

  //find the ip address and num of bytes
  address := os.Args[2]
  numbytes, err_size := strconv.Atoi(os.Args[3])
  /*if (numbytes > 32*1024) {
    fmt.Println("Out of range! Maximum allowed size is 32KB")
    return
  }*/
  if (err_size != nil) {
   //fmt.Println("Please check the data size want to be transferred, maximum allowed size is 32KB")
   log.Fatal(err_size)
  }

  //connect to server
  conn, err := net.Dial("tcp", address) //"127.0.0.1:8080")
  fmt.Println("Connected to Ip: " + address)
  if (err != nil) {
    log.Fatal(err)
  }

  //interactive flag is true
  if *interactivePtr == true {
    fmt.Println("You are in the interactive mode, which allows you to type in a payload.")
    fmt.Println("Press Enter if you want to exit.")
    for {
      //read input message from stdin
      reader := bufio.NewReader(os.Stdin)
      fmt.Print("Text to send: ")
      text, err_text := reader.ReadString('\n')
      if (err_text != nil) {
        log.Fatal(err_text)
      }
      //fmt.Println(string(text))
      //if there is nothing there, break the connection
      if (string(text) == "\n") {
        fmt.Println("Nothing want to send to server, close the connection! ")
        break
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
    //send message to server
    fmt.Println("You are in the non-interactive mode!")
    // send number of bytes info to server first
    str_numbytes := strconv.Itoa(numbytes)
    _, err_bytes := fmt.Fprintf(conn, str_numbytes + "\n")
    if (err_bytes != nil) {
      log.Fatal(err_bytes)
    }
    fmt.Println("Sent number of bytes info to server: " + str_numbytes)

    byte_message, err_byt := bufio.NewReader(conn).ReadString('\n')
    if (err_byt != nil) {
      log.Fatal(err_byt)
    }
    fmt.Print("server sent back: " + byte_message)

    // send message to server
    message := make([]byte, numbytes)
    nbytes, err_sent := conn.Write(message)
    if (err_sent != nil) {
      log.Fatal(err_sent)
    }
    fmt.Println("Sent ", nbytes, " bytes Message to server")

    data := make([]byte, numbytes)//32*1024)
    numBytes, err_read := conn.Read(data)
    if (err_read != nil) {
      log.Fatal(err_read)
    }
    data = nil
    fmt.Println(numBytes, "bytes data received from server")
  }
}
