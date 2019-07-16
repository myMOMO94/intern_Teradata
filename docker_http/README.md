# http-server-client with docker network

*You can pull the images here:*
```
> momomengyu/http-client:part2
> momomengyu/http-server:part2
```
*Or, you can build the images by yourself*
```
> docker build -t server . (inside server directory)
> docker build -t client . (inside client directory)
```

*To run, please type the following command on your terminal:*
*At first, please create a Docker user defined bridge network:*
```
> docker network create --driver bridge http-net
```
**(You can change the name of the network if you want)**

*Then, please run the following command (this will attach containers to the network you created above):*
```
> docker run -dit --name http-server --network http-net momomengyu/http-server:part2 -verbose=false
```
**(Please do not change the server container name)**
```
> docker run -it --name http-client --network http-net -e NUMBYTES="512" -e NUMRUNS="2" momomengyu/http-client:part2 -verbose=false
```
