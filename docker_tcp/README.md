# http-server-client with docker network

*You can pull the images here:*
```
> momomengyu/tcp-client:part2
> momomengyu/tcp-server:part2
```
*Or, you can build the images by yourself*
```
> docker build -t server . (inside server directory)
> docker build -t client . (inside client directory)
```

*To run, please type the following command on your terminal:*\
*At first, please create a Docker user defined bridge network:*
```
> docker network create --driver bridge tcp-net
```
**(You can change the name of the network if you want)**

*Then, please run the following command (this will attach containers to the network you created above):*
```
> docker run -dit --name tcp-server --network tcp-net -e NUMRUNS="2" momomengyu/tcp-server:part2 -verbose=false
```
**(Please do not change the server container name, however, you can change the NUMRUNS if you want, just keep it consistent with NUMRUNS below.)**
```
> docker run -it --name tcp-client --network tcp-net -e NUMBYTES="512" -e NUMRUNS="2" momomengyu/tcp-client:part2 -verbose=false
```
