package app

import (
	"fmt"
	"log"
)

func Status() {

	endpoint := fmt.Sprintf("%s/service/rest/v1/status", NXRMConfig.URL)

	InfoLogger.Println(endpoint)
	resp, err := client.R().Get(endpoint)
	if err != nil {
		ErrorLogger.Println(err)
	} else {
		log.Println(string(resp.Status()))
	}
}

func Check() {
	endpoint := fmt.Sprintf("%s/service/rest/v1/status/check", NXRMConfig.URL)
	InfoLogger.Println(endpoint)
	resp, err := client.R().Get(endpoint)

	if err != nil {
		ErrorLogger.Println(err)
	} else {
		log.Println(string(resp.Body()))
		InfoLogger.Println(resp.StatusCode())
	}
}
