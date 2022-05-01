package app

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/sriram-yeluri/nxrm3/config"
)

func Status() {

	endpoint := fmt.Sprintf("%s/service/rest/v1/status/check", config.NXRMConfig.URL)

	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetBasicAuth(config.NXRMConfig.USERNAME, config.NXRMConfig.PASSWORD).
		Get(endpoint)
	if err != nil {

	} else {
		log.Println(string(resp.Body()))
	}
}
