package status

import (
	"fmt"
	"log"

	"github.com/sriram-yeluri/nxrm3/pkg/nxrm/nxrm"
	"github.com/sriram-yeluri/nxrm3/pkg/nxrm/util"
)

func Status() {

	endpoint := fmt.Sprintf("%s/service/rest/v1/status", nxrm.NXRMConfig.URL)

	util.Info().Println(endpoint)
	resp, err := nxrm.Client.R().Get(endpoint)
	if err != nil {
		util.Error().Println(err)
	} else {
		util.Info().Println(string(resp.Status()))
	}
}

func Check() {
	endpoint := fmt.Sprintf("%s/service/rest/v1/status/check", nxrm.NXRMConfig.URL)
	util.Info().Println(endpoint)
	resp, err := nxrm.Client.R().Get(endpoint)

	if err != nil {
		util.Error().Println(err)
	} else {
		log.Println(string(resp.Body()))
		util.Info().Println(resp.StatusCode())
	}
}
