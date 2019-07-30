package jaegerwrapper

import (
    "context"
    "io"

    "github.com/opentracing/opentracing-go"
    "github.com/opentracing/opentracing-go/ext"
    "github.com/opentracing/opentracing-go/log"
    "intern_Teradata/gRPC_tracing/jaeger_tracer_init"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
    "google.golang.org/grpc/metadata"
)

type TextMapReader struct {
    metadata.MD
}

// ForeachKey returns TextMap contents cia repeated calls to the 'handler' function.
// If any call to 'handler' returns a non-nil error, ForeachKey terminates and returns that error.
// TextMapReader is the Extract() carrier for the TextMap builtin format.
// The default TextMapCarrier allows the use of regular map[string]string
// For gRPC metadata, it uses map[string][]string
// Therefore, we re-implement this method here.
func (c TextMapReader) ForeachKey(handler func(key, val string) error) error {
    for k, vs := range c.MD {
        for _, v := range vs {
            if err := handler(k, v); err != nil {
                return err
            }
        }
    }

    return nil
}

type TextMapWriter struct {
    metadata.MD
}

// Set a key:value pair to the carrier. Mutiple calls to Set() for the same key leads to undefinde behavior
// Set implements Set() of opentracing.TextMapWriter
// TextMapWriter is the Inject() carrier for the TextMap builtin format.
func (c TextMapWriter) Set(key, val string) {
    //key = strings.ToLower(key)
    c.MD[key] = append(c.MD[key], val)
}

func InitTracer (service string, jagerAgent string) (tracer opentracing.Tracer, closer io.Closer) {
    tracer, closer = tracerinit.InitJaeger(service, jagerAgent)
    opentracing.SetGlobalTracer(tracer)
    return
}

func ClientInterceptor(tracer opentracing.Tracer) grpc.UnaryClientInterceptor {
    return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
        var parentCtx opentracing.SpanContext
        parentSpan := opentracing.SpanFromContext(ctx)
        if parentSpan != nil {
            parentCtx = parentSpan.Context()
        }

        //tracer := opentracing.GlobalTracer()
        span := tracer.StartSpan(
                method,
                opentracing.ChildOf(parentCtx),
                opentracing.Tag{Key: string(ext.Component), Value: "gRPC"},
                ext.SpanKindRPCClient)
        defer span.Finish()

        md, ok := metadata.FromIncomingContext(ctx)
        if !ok {
            md = metadata.New(nil)
        } else {
            md = md.Copy()
        }

        mdWriter := TextMapWriter{md}
        e := tracer.Inject(
               span.Context(),
               opentracing.TextMap,
               mdWriter)
        if e != nil {
            span.LogFields(log.String("inject error", e.Error()))
        }

        newCtx := metadata.NewOutgoingContext(ctx, md)
        err := invoker(newCtx, method, req, reply, cc, opts...)
        if err != nil {
            span.LogFields(log.String("call error", err.Error()))
        }

        span.LogFields(
            log.String("event", "Sending request then wait for receive response from server"),
            log.String("value", "Done"),
        )

        return err
    }
}

func ServerInterceptor(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
    return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
        md, ok := metadata.FromIncomingContext(ctx)
        if !ok {
            md = metadata.New(nil)
        }

        //tracer := opentracing.GlobalTracer()
        spanContext, err := tracer.Extract(opentracing.TextMap, TextMapReader{md})
        if err != nil && err != opentracing.ErrSpanContextNotFound {
            grpclog.Errorf("extract from metadata err : %v", err)
        } else {
            span := tracer.StartSpan(
                    info.FullMethod,
                    ext.RPCServerOption(spanContext),
                    opentracing.Tag{Key: string(ext.Component), Value: "gRPC"},
                    ext.SpanKindRPCServer,
            )
            defer span.Finish()

            span.LogFields(
                log.String("event", "Received request then sending response to client"),
                log.String("value", "Done"),
            )

            ctx = opentracing.ContextWithSpan(ctx, span)
        }

        return handler(ctx, req)
    }
}

func DialOption(tracer opentracing.Tracer) grpc.DialOption {
    return grpc.WithUnaryInterceptor(ClientInterceptor(tracer))
}

func ServerOption(tracer opentracing.Tracer) grpc.ServerOption {
    return grpc.UnaryInterceptor(ServerInterceptor(tracer))
}
