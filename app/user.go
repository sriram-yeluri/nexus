package app

import (
	"fmt"
	"log"
)

func GetUserList() {

	endpoint := fmt.Sprintf("%s/service/rest/v1/security/users", NXRMConfig.URL)
	InfoLogger.Println(endpoint)
	resp, err := client.R().Get(endpoint)
	if err != nil {
		ErrorLogger.Println(err)
	} else {
		log.Println(string(resp.Body()))
	}

}
