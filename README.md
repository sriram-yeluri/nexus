

```sh
docker run -d -p 8081:8081 --name nexus klo2k/nexus3
docker container ls # check if the nexus3 container is created successfully
docker logs nexus -f # wait till the nexus instance is fully started
#access nexus3 at port 8081 on either localhost of the hosted server
# http://192.168.2.123:8081
docker exec -it nexus /bin/bash #exec into the contaner to get the initial admin password
```

## Prerequisites
## Install Cobra library for CLI

```sh
go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest
```


### References

[Cobra](https://pkg.go.dev/github.com/spf13/cobra@v1.4.0)
[Cobra-cli-README](https://github.com/spf13/cobra-cli/blob/main/README.md)


