package app

import (
	"fmt"
	"log"
)

func Status() {

	endpoint := fmt.Sprintf("%s/service/rest/v1/status/check", NXRMConfig.URL)

	InfoLogger.Println(endpoint)
	//client := resty.New()

	resp, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetBasicAuth(NXRMConfig.USERNAME, NXRMConfig.PASSWORD).
		Get(endpoint)
	if err != nil {
		ErrorLogger.Println(err)
	} else {
		log.Println(string(resp.Body()))
	}
}
