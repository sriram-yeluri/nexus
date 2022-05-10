package nxrm

import (
	"github.com/go-resty/resty/v2"
)

var (
	Client *resty.Client
)

func SetRestyClient() *resty.Client {
	Client = resty.New()
	Client.SetBasicAuth(NXRMConfig.USERNAME, NXRMConfig.PASSWORD)
	Client.SetHeader("Accept", "application/json")
	Client.EnableTrace()
	return Client
}

func init() {

	NXRMConfig = LoadConfig()
	Client = SetRestyClient()
}
