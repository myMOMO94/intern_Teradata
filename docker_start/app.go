package main

import "github.com/go-redis/redis"
import "fmt"
import "log"
import "net/http"
import "text/template"
//import "net"
import "os"
import "github.com/yosssi/gohtml"

func main() {
    client := redis.NewClient(&redis.Options{
        Network:  "tcp",
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
    })

    _, err := client.Ping().Result()
    //fmt.Println(pong)
    var visits string = ""
    if (err != nil) {
        visits = "<i>cannot connect to Redis, counter disabled</i>"
    } else {
        _, err := client.Incr("counter").Result()
        if (err != nil) {
            visits = "<i>cannot connect to Redis, counter disabled</i>"
        }
    }
    fmt.Println(visits)
    http.HandleFunc("/", hello)

    fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe(":80", nil); err != nil {
        log.Fatal(err)
    }
}

func hello(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    visits := "<i>cannot connect to Redis, counter disabled</i>"

    switch r.Method {
    case "GET":     
         //http.ServeFile(w, r, "form.html")
        tpl, err := template.New("test").Parse("<html><h3>Hello {{.name}}!</h3><b>Hostname:</b> {{.hostname}}<br/><b>Visits:</b> {{.visits}}</html>")
        if err != nil {
	    panic(err)
	}
	data := map[string]interface{}{"name": os.Getenv("NAME")}
        //addrs, err := net.LookupHost(hostname)
        addrs, err := os.Hostname()
        data["hostname"] = addrs
        data["visits"] = visits

        file, err := os.Create("form.html")
        if (err != nil) {
            log.Fatal(err)
        }
        defer file.Close()
	err = tpl.Execute(gohtml.NewWriter(file), data)

	if err != nil {
	    panic(err)
	}
        http.ServeFile(w, r, "form.html")
        //WriteToFile("form.html", )
    case "POST":
        // Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }
        fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
        name := r.FormValue("name")
        address := r.FormValue("address")
        fmt.Fprintf(w, "Name = %s\n", name)
        fmt.Fprintf(w, "Address = %s\n", address)
    default:
        fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}

/*func WriteToFile(filename string, data string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = io.WriteString(file, data)
    if err != nil {
        return err
    }
    return file.Sync()
}*/
