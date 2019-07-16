package main

import (
    //"encoding/json"
    "log"
    "net/http"
    "io/ioutil"
    "fmt"
)

/*type test_struct struct {
    Test string
}*/

func test(rw http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case "GET":
        fmt.Fprintf(rw, "Just ouput something for test.")
    case "POST":
        body, err := ioutil.ReadAll(req.Body)
        if err != nil {
            panic(err)
        }
        //log.Println(string(body))
        fmt.Fprintf(rw, string(body))
        /*var t test_struct
        err = json.Unmarshal(body, &t)
        if err != nil {
            panic(err)
        }
        log.Println(t.Test)*/
    }
}

func main() {
    fmt.Println("Starting server for testing HTTP request...")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "This is a website server by a Go HTTP server.")
    })
    http.HandleFunc("/test", test)
    log.Fatal(http.ListenAndServe(":80", nil))

}
