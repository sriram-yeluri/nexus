
### Hosting the Nexus OSS instance

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

## Prerequisites

### Install Cobra library for CLI

```sh
go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest
```


### References for libraries i used in creating this CLI

* [Cobra](https://pkg.go.dev/github.com/spf13/cobra@v1.4.0)
* [Cobra-cli README](https://github.com/spf13/cobra-cli/blob/main/README.md)
* [resty Simple HTTP and REST client](https://github.com/go-resty/resty)


