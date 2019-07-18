# gRPC-server-client with docker network

*You can pull the images here:*
```
> momomengyu/grpc-client:part2
> momomengyu/grpc-server:part1
```
*Or, you can build the images by yourself*
```
> docker build -t server -f server/Dockerfile . (inside docker_gRPC directory)
> docker build -t client -f client/Dockerfile . (inside docker_gRPC directory)
```

*To run, please type the following command on your terminal:*\
*At first, please create a Docker user defined bridge network:*
```
> docker network create --driver bridge grpc-net
```
**(You can change the name of the network if you want)**

*Then, please run the following command (this will attach containers to the network you created above):*
```
> docker run -dit --name gRPC-server --network grpc-net momomengyu/grpc-server:part1
```
**(Please do not change the server container name)**
```
> docker run -it --name gRPC-client --network grpc-net -e NUMBYTES="512" -e NUMRUNS="2" momomengyu/grpc-client:part2 -verbose=false
```
**(You can change the NUMBYTES and NUMRUNS if you want)**

*Now, you can stop the container:*
```
> docker container stop <CONTAINER ID or NAME>
```
