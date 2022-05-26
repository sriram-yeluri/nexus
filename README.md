# Nexus CLI
nexus cli is created to manage Nexus OSS

## How to user the cli 

```sh
./nexus -h
Nexus Repository Manager OSS CLI.

Usage:
  nexus [command]

Available Commands:
  add         Add resources to Nexus
  check       Check the health status of Nexus Repository Manager
  help        Help about any command
  repo        Manage Nexus Repositories
  status      Get the status of Nexus Repository Manager
  user        Manage Nexus Users

Flags:
  -h, --help   help for nexus

Use "nexus [command] --help" for more information about a command.
```

## Prerequisites
* [Hosting Nexus OSS](docs/nexusOSS.md)
* [Installing Cobra-cli](docs/prerequisites.md)

### References for libraries used in creating nexus CLI

* [Cobra](https://pkg.go.dev/github.com/spf13/cobra@v1.4.0)
* [Cobra-cli README](https://github.com/spf13/cobra-cli/blob/main/README.md)
* [resty Simple HTTP and REST client](https://github.com/go-resty/resty)


