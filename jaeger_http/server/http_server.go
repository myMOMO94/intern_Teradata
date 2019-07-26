package main

import (
    //"encoding/json"
    "log"
    "net/http"
    "io/ioutil"
    "fmt"
    "strconv"
    "context"

    opentracing "github.com/opentracing/opentracing-go"
    "github.com/opentracing/opentracing-go/ext"
    tracinglog "github.com/opentracing/opentracing-go/log"
    "intern_Teradata/jaeger_tracer_init"
)

type TracerContext struct {
    Tracer opentracing.Tracer
}

func (tctx *TracerContext) test(rw http.ResponseWriter, req *http.Request) {
    spanCtx, _ := tctx.Tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
    span := tctx.Tracer.StartSpan("Listening", ext.RPCServerOption(spanCtx))
    defer span.Finish()

    ctx := opentracing.ContextWithSpan(context.Background(), span)

    switch req.Method {
    case "GET":
        fmt.Fprintf(rw, "Just ouput something for test.")
    case "POST":
        body := readRequest(req, ctx)
        sendResponse(rw, body, ctx)
    }

    span.LogFields(
        tracinglog.String("event", "Received request and sent response back"),
        tracinglog.String("value", "Done"),
    )
}

func readRequest(req *http.Request, ctx context.Context) (body []byte) {
    span, ctx := opentracing.StartSpanFromContext(ctx, "readRequest")
    defer span.Finish()

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        panic(err)
    }
    span.LogFields(
        tracinglog.String("event", "Received Message length: "),
        tracinglog.String("value", strconv.Itoa(len(body))),
    )

    return body
}

func sendResponse(rw http.ResponseWriter, body []byte, ctx context.Context) {
    span, ctx := opentracing.StartSpanFromContext(ctx, "sendReponse")
    defer span.Finish()

    fmt.Fprintf(rw, string(body))

    span.LogFields(
        tracinglog.String("event", "Sent Response back to client"),
        tracinglog.String("value", "Done"),
    )
}

func main() {
    tracer, closer := tracerinit.InitJaeger("server-listen-and-serve")
    defer closer.Close()
    opentracing.SetGlobalTracer(tracer)

    myTracer := &TracerContext{Tracer: tracer}

    fmt.Println("Starting server for testing HTTP request...")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "This is a website server by a Go HTTP server.")
    })
    http.HandleFunc("/test", myTracer.test)
    log.Fatal(http.ListenAndServe(":80", nil))

}
