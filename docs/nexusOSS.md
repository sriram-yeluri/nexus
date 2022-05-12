# Hosting the Nexus OSS instance locally

I am running Nexus OSS on Raspberry PI as a docker container. 

```sh
docker run -d -p 8081:8081 --name nexus klo2k/nexus3

# check if the nexus3 container is created successfully
docker container ls 

# wait till the nexus instance is fully started
docker logs nexus -f 

#access nexus3 at default port 8081 -> http://192.168.2.123:8081

#exec into the contaner to get the initial admin password
docker exec -it nexus /bin/bash
```
