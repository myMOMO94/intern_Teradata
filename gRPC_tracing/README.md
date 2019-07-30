# gRPC-server-client with Jaeger opentracing

*You can pull the images here:*
```
> momomengyu/kubernetes-grpctracingserver:part0
```
*Or, you can build the images by yourself*
```
> docker build -t server -f server/Dockerfile . (inside gRPC_tracing directory)
```
**If you choose to do this, please change the image name in file "server/server.yaml"**

*To enable Jaeger tracing UI, please run the Jaeger all-in-one image:*
```
> docker run -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest
```

*To run, please type the following command on your terminal:*\
*(Inside gRPC_tracing directory)*
```
> kubectl apply -f server/server.yaml
> kubectl apply -f server/server-svc.yaml
> kubectl port-forward service/grpcserver 8080
```
*This will create a kubernetes service called grpcserver, then forward container port 8080 to localhost port 8080*

*To let client-server talk to each other, please type the following command on your terminal:*
```
> go build gRPC_client.go (inside client directory)
> ./gRPC_client -verbose=false 512 2
```
**(You will see messages on your screen)**

*Now, you can browse http://localhost:16686 to view all your traces.*
*Now, you can clean up everything:*
```
> kubectl delete -f server/ (inside gRPC_tracing directory)
```
