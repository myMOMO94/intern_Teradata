package tracerinit

import (
    "fmt"
    "io"
    "time"

    opentracing "github.com/opentracing/opentracing-go"
    jaeger "github.com/uber/jaeger-client-go"
    config "github.com/uber/jaeger-client-go/config"
)

// InitJaeger returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func InitJaeger(service string, jaegerAgent string) (opentracing.Tracer, io.Closer) {
    cfg := config.Configuration{
        Sampler: &config.SamplerConfig{
            Type:  "const",
            Param: 1,
        },
        Reporter: &config.ReporterConfig{
            LogSpans: true,
            BufferFlushInterval: 1 * time.Second,
            LocalAgentHostPort: jaegerAgent,
        },
    }
    tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
    if err != nil {
        panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
    }
    return tracer, closer
}

