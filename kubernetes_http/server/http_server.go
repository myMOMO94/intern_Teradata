package main

import (
    //"encoding/json"
    "log"
    "net/http"
    "io/ioutil"
    "fmt"
    "text/template"
    "github.com/yosssi/gohtml"
    "os"
)

/*type test_struct struct {
    Test string
}*/

func test(rw http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case "GET":
        tpl, err := template.New("test").Parse("<html><h3>Hello!</h3><b>Hostname:</b> {{.hostname}}</html>")
        if err != nil {
	    panic(err)
	}
        addrs, err := os.Hostname()
        //data["hostname"] = addrs
	data := map[string]interface{}{"hostname": addrs}

        file, err := os.Create("form.html")
        if (err != nil) {
            log.Fatal(err)
        }
        defer file.Close()
	err = tpl.Execute(gohtml.NewWriter(file), data)

	if err != nil {
	    panic(err)
	}
        http.ServeFile(rw, req, "form.html")
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
    log.Fatal(http.ListenAndServe(":8080", nil))

}
