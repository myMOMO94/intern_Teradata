package main

import (
    "net/http"
    "log"
    "io/ioutil"
    "bytes"
    "fmt"
    "flag"
    "os"
    "strconv"
    "time"
    "context"

    "github.com/opentracing/opentracing-go"
    "github.com/opentracing/opentracing-go/ext"
    tracinglog "github.com/opentracing/opentracing-go/log"
    "intern_Teradata/jaeger_tracer_init"
)

func main() {
    if len(os.Args) != 4 {
        log.Fatal("ERROR: Expecting four argument!")
    }

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

    tracer, closer := tracerinit.InitJaeger("client-to-server")
    defer closer.Close()
    opentracing.SetGlobalTracer(tracer)

    span := tracer.StartSpan("client-sent-bytearray")
    defer span.Finish()

    ctx := opentracing.ContextWithSpan(context.Background(), span)

    for i := 0; i < numruns; i++ {
        sendRequest(numbytes, *verbosePtr, ctx)
    }
}

func sendRequest(numbytes int, verbose bool, ctx context.Context) {
    span, ctx := opentracing.StartSpanFromContext(ctx, "prepareToSendDataAndWaitForResponse")
    defer span.Finish()

    url := "http://localhost:80/test"
    //url := "http://http-server:80/test"
    if verbose == true {
        fmt.Println("URL:>", url)
    }
    data := make([]byte, numbytes)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))//bytes.NewBuffer(jsonStr))
    if err != nil {
        log.Fatal(err)
    }

    req.Header.Set("X-Custom-Header", "myvalue")
    //req.Header.Set("Content-Type", "application/json")

    ext.SpanKindRPCClient.Set(span)
    ext.HTTPUrl.Set(span, url)
    ext.HTTPMethod.Set(span, "POST")
    span.Tracer().Inject(
        span.Context(),
        opentracing.HTTPHeaders,
        opentracing.HTTPHeadersCarrier(req.Header),
    )

    resp, dur := sending(req, verbose, url, ctx)

    readResponse(resp, verbose, dur, ctx)

    span.LogFields(
        tracinglog.String("event", "Sent request then received response from server"),
        tracinglog.String("value", "Done"),
    )
}

func sending(req *http.Request, verbose bool, url string, ctx context.Context) (resp *http.Response, end time.Duration){
    span, _ := opentracing.StartSpanFromContext(ctx, "sendRequest")
    defer span.Finish()

    start := time.Now()
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    end = time.Since(start)
    if verbose == true {
        fmt.Println("response Status:", resp.Status)
        fmt.Println("response Headers:", resp.Header)
    }

    span.LogFields(
        tracinglog.String("event", "Request sent to"),
        tracinglog.String("value", url),
    )

    return resp, end
}

func readResponse(resp *http.Response, verbose bool, dur time.Duration, ctx context.Context) {
    span, _ := opentracing.StartSpanFromContext(ctx, "readResponse")
    defer span.Finish()

    body, _ := ioutil.ReadAll(resp.Body)
    //fmt.Println("response Body:", string(body))
    if verbose == true {
        fmt.Println("Response body size: ", len(body), " total time to send request till received response is: ", dur)
    }

    span.LogFields(
        tracinglog.String("event", "Response received"),
        tracinglog.String("value", "Response body size: " + strconv.Itoa(len(body))),
    )
}
