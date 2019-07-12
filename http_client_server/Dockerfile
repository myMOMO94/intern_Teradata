# Start from golang v1.11 base image
FROM golang

# ARG interactive=false
# ARG numbytes=512
# ARG numruns=1

# Add Maintainer Info
LABEL maintainer="Mengyu Mo <mengyu.mo@teradata.com>"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/intern_Teradata/http_client_server

RUN mkdir $GOPATH/src/intern_Teradata/http_client_server/server
RUN mkdir $GOPATH/src/intern_Teradata/http_client_server/client
RUN cd $GOPATH/src/intern_Teradata/http_client_server

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
#COPY . .
#COPY tcp_server tcp_server
#COPY tcp_client tcp_client
COPY start.sh start.sh
COPY http_server.go ./server/http_server.go
COPY http_client.go ./client/http_client.go

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get ./server/...

# Install the package
RUN go install ./server/...

RUN go build ./server/http_server.go

RUN go build ./client/http_client.go

# This container exposes port 8080 to the outside world
#EXPOSE 8080

ENTRYPOINT ["./start.sh"]

# Run the executable
#CMD ["tcp_client_server"]
#CMD ./start.sh
CMD []
