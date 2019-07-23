# http-server-client with docker network

*You can pull the images here:*
```
> momomengyu/kubernetes-httpserver:part1
```
*Or, you can build the images by yourself*
```
> docker build -t server . (inside server directory)
```
**If you choose to do this, please change the image name in file "server/server.yaml"**

*To run, please type the following command on your terminal:o*\
*(Inside kubernetes_http directory)*
```
> kubectl apply -f server/server.yaml
> kubectl apply -f server/server-svc.yaml
> kubectl port-forward service/httpserver 8080
```
*This will create a kubernetes service called httpserver, then forward port 8080 to localhost*

*Then, you can browse http://localhost:8080/ or http://localhost:8080/test in your browser, you will see the difference.*

*To let client-server talk to each other, please type the following command on your terminal:*
```
> go build http_client.go (inside client directory)
> ./http_client -verbose=false 512 2
```
**(You will see messages on your screen)**

*Now, you can clean up everything:*
```
> kubectl delete -f server/ (inside kubernetes_http directory)
```
