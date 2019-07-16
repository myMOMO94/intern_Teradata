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

*You can visit following URL:*
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

