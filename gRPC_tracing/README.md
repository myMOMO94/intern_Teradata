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
![jaegerui](https://github.com/myMOMO94/intern_Teradata/blob/master/gRPC_tracing/jaegerUI.png)

*Now, you can clean up everything:*
```
> kubectl delete -f server/ (inside gRPC_tracing directory)
```

## Part 2 Deploy grpc-server to kubernetes

*You can pull the images here:*
```
> momomengyu/kubernetes-grpctracingserver:part1
> momomengyu/kubernetes-grpctracingclient:part1
```

*At first, use Jaeger Operator to depoly Jaeger to Kubernetes Cluster as a DaemonSet:*
```
> kubectl apply -f daemonset.yaml
```

*Find the Jaeger UI ip address:*
```
>minikube ip
```
**(If you are using minikube, otherwise, this will not work. Also, if you want your local client or client on other machine connect to Jaeger, please change the jaeger agent host address inside client program with this address.)**

*Make sure the ingress ip is show up here before you go next:*
```
> kubectl get ing
```

*Deploy grpc-server to a Kubernetes Pod:*
```
> kubectl apply -f server/server.yaml
```
*Create a grpc-server services:*
```
> kubectl apply -f server/server-svc.yaml
```

*Run client locally:*
```
> go build gRPC_client.go
> ./gRPC_client -verbose=false 512 2
```
*Or, you can run the client on Kubernetes as well:*
```
> kubectl apply -f client/client.yaml
> kubectl apply -f client/client-svc.yaml
```

*To access the Jaeger UI to find all your traces: browse the minikube ip address on your browser.*
