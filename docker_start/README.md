# play around with docker container

*You can pull the image here:*
```
> momomengyu/docker_start:part1
```

*Or, you can build the image by yourself*
```
> docker build -t docker_start .
```

*Then, run following command:*
```
> docker run -p 80:80 -d momomengyu/docker_start:part1
```

*To see the result, you can visit following URL:*
```
http://localhost:80/
```
*Or, type the following command on your terminal:*
```
curl http://localhost:80/
```

*Now, you can stop the container:*
```
> docker container stop <CONTAINER ID or NAME>
```

*You can also define, run a service by run the following command:*
```
> docker swarm init
> docker stack deploy -c docker-compose.yml getstartedlab
```

*The following example command lets you view all services associated with the getstartedlab stack:*
```
> docker stack services getstartedlab
```

*A single container running in a service is called a task. Tasks are given unique IDs that numerically increment, up to the number of replicas you defined in docker-compose.yml. List the tasks for your service:*
```
> docker service ps getstartedlab_web
```

*You can scale the app by changing the replicas value in docker-compose.yml, saving the change, and re-running the docker stack deploy command:*
```
> docker stack deploy -c docker-compose.yml getstartedlab
```

*Take the app down with docker stack rm:*
```
> docker stack rm getstartedlab
```

*Take down the swarm:*
```
> docker swarm leave --force
```
