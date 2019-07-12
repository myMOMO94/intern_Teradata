package main

import (
    "net/http"
    "log"
    "io/ioutil"
    "bytes"
    //"encoding/json"
    "fmt"
    "flag"
    "os"
    "strconv"
    "time"
)

func main() {
    verbosePtr := flag.Bool("verbose", false, "a bool")
    flag.Parse()

    numbytes, err_size := strconv.Atoi(os.Args[2])
    if err_size != nil {
        log.Fatal(err_size)
    }

    numruns, err_runs := strconv.Atoi(os.Args[3])
    if err_runs != nil {
        log.Fatal(err_runs)
    }

    url := "http://localhost:8082/test"
    if *verbosePtr == true {
        fmt.Println("URL:>", url)
    }

    //var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
    for i := 0; i < numruns; i++ {
        data := make([]byte, numbytes)
        start := time.Now()
        req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))//bytes.NewBuffer(jsonStr))
        req.Header.Set("X-Custom-Header", "myvalue")
        //req.Header.Set("Content-Type", "application/json")

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            panic(err)
        }
        defer resp.Body.Close()
        end := time.Since(start)
        if *verbosePtr == true {
            fmt.Println("response Status:", resp.Status)
            fmt.Println("response Headers:", resp.Header)
        }

        body, _ := ioutil.ReadAll(resp.Body)
        //fmt.Println("response Body:", string(body))
        fmt.Println("Response body size: ", len(body), " total time to send request till received response is: ", end)
    }
}
