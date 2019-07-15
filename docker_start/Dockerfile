# Use an official Python runtime as a parent image
FROM golang

LABEL maintainer="Mengyu Mo <mengyu.mo@teradata.com>"

# Set the working directory to /app
WORKDIR $GOPATH/src/intern_Teradata/docker_start

# Copy the current directory contents into the container at /app
COPY . .

# Install any needed packages specified in requirements.txt
#RUN pip install --trusted-host pypi.python.org -r requirements.txt
RUN go get -u github.com/go-redis/redis
RUN go get -u github.com/yosssi/gohtml

RUN go get ./...
RUN go install ./...

RUN go build ./app.go

# Make port 80 available to the world outside this container
EXPOSE 80

# Define environment variable
ENV NAME World
ENV http_proxy host:port
ENV https_proxy host:port

# Run app.py when the container launches
#CMD ["go", "run", "app.go"]
CMD ["./app"]
