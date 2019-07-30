# Start from golang v1.11 base image
FROM golang

# ARG interactive=false
# ARG numbytes=512
# ARG numruns=1

# Add Maintainer Info
LABEL maintainer="Mengyu Mo <mengyu.mo@teradata.com>"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/intern_Teradata/gRPC_tracing/server

RUN mkdir $GOPATH/src/intern_Teradata/gRPC_tracing/bytearray
RUN mkdir $GOPATH/src/intern_Teradata/gRPC_tracing/jaeger_gRPC_wrapper
RUN mkdir $GOPATH/src/intern_Teradata/gRPC_tracing/jaeger_tracer_init
#RUN mkdir $GOPATH/src/intern_Teradata/http_client_server/client
#RUN cd $GOPATH/src/intern_Teradata/http_client_server/server

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
#COPY . .
COPY ./jaeger_tracer_init/tracerinit.go ../jaeger_tracer_init/tracerinit.go
COPY ./jaeger_gRPC_wrapper/wrapper.go ../jaeger_gRPC_wrapper/wrapper.go
COPY ./server/gRPC_server.go gRPC_server.go
COPY ./bytearray/bytearray.pb.go ../bytearray/bytearray.pb.go
COPY ./server/start.sh start.sh
#COPY http_server.go http_server.go
#COPY http_client.go ./client/http_client.go

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get ./...

# Install the package
RUN go install ./...

RUN go build gRPC_server.go

#RUN go build ./client/http_client.go

# This container exposes port 8080 to the outside world
EXPOSE 8080

ENTRYPOINT ["./start.sh"]

# Run the executable
#CMD ["tcp_client_server"]
#CMD ./start.sh
#CMD ["./http_server", "-verbose=false"]
CMD []
