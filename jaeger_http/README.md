# http-server-client with jaeger opentracing

*To enable Jaeger tracing UI, please run the Jaeger all-in-one image:*
```
> docker run -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest
```

*Then, you can run the server by:*
```
> go build server/http_server.go
> ./http_server -verbose=false
```
*Run the client by:*
```
> go build client/http_client.go
> ./http_client -verbose=false 512 2
```

*Now, you can browse http://localhost:16686 to view all your traces.*

