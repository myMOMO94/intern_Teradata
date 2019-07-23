# kubernetes-tcp-server-client

*You can pull the images here:*
```
> momomengyu/kubernetes-tcpserver:part2
```
*Or, you can build the images by yourself*
```
> docker build -t server . (inside server directory)
```
**If you choose to do this, please change the image name in file "server/server.yaml"**

*To run, please type the following command on your terminal:*\
*(Inside kubernetes_tcp directory)*
```
> kubectl apply -f server/server.yaml
> kubectl apply -f server/server-svc.yaml
> kubectl port-forward service/tcpserver 8888
```
*This will create a kubernetes service called tcpserver, then forward container port 8080 to localhost port 8888*

*Then, you can browse http://localhost:8888/ or http://localhost:8888/test in your browser, you will see the difference.*

*To let client-server talk to each other, please type the following command on your terminal:*
```
> go build tcp_client.go (inside client directory)
> ./tcp_client -interactive=false -verbose=false 127.0.0.1:8888 512 2
```
**(You will see messages on your screen, also, please keep the interactive flag to false, leave the ip address unchange)**

*Now, you can clean up everything:*
```
> kubectl delete -f server/ (inside kubernetes_tcp directory)
```
