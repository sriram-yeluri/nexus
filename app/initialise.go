package app

import (
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/sriram-yeluri/nxrm3/util"
)

var (
	NXRMConfig  util.Config
	client      *resty.Client
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func SetRestyClient() *resty.Client {
	client = resty.New()
	client.SetBasicAuth(NXRMConfig.USERNAME, NXRMConfig.PASSWORD)
	client.SetHeader("Accept", "application/json")
	client.EnableTrace()
	return client
}

func init() {
	InfoLogger = log.New(os.Stdout, "[INFO] : ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "[ERROR] : ", log.Ldate|log.Ltime|log.Lshortfile)
	NXRMConfig = util.LoadConfig()
	client = SetRestyClient()
}
