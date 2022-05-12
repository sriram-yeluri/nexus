# Prerequisites

## Cobra library for CLI

```sh
go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest

# Initialization
cd nexus
go mod init github.com/sriram-yeluri/nexus
cobra-cli init
cobra-cli add status
cobra-cli add list
cobra-cli add users -p 'configList'
```
