package main

import "net"
import "os"
import "bufio"
import "fmt"
import "log"

func main() {
  //initialize connection
  /*address_reader := bufio.NewReader(os.Stdin)
  fmt.Print("Ip and port to connect: (follow the pattern 127.0.0.1:8080): ")
  address, err_add := address_reader.ReadString('\n')
  if (err_add != nil) {
    log.Fatal(err_add)
  }
  fmt.Print("Connecting to Ip: " + string(address))*/

  scanner := bufio.NewScanner(os.Stdin)
  var address string
  fmt.Print("Ip and port to connect(follow the pattern 127.0.0.1:8080): ")
  scanner.Scan()
  address = scanner.Text()
  conn, err := net.Dial("tcp", address) //"127.0.0.1:8080")
  fmt.Println("Connected to Ip: " + address)
  if (err != nil) {
    log.Fatal(err)
  }

  for {
    //read input message from stdin
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send: ")
    text, err_text := reader.ReadString('\n')
    if (err_text != nil) {
      log.Fatal(err_text)
    }

    //send message to server
    numbytes, err_sent := fmt.Fprintf(conn, text + "\n")
    if (err_sent != nil) {
      log.Fatal(err_sent)
    }
    fmt.Print("Sent ", numbytes, " bytes Message to server: " + text)

    //Listen for sever reply
    message, err_read := bufio.NewReader(conn).ReadString('\n')
    if (err_read != nil) {
      log.Fatal(err_read)
    }
    fmt.Print("Message from server: " + message)
  }
}
