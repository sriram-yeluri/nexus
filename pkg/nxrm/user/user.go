package user

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/sriram-yeluri/nexus/pkg/nxrm/nxrm"
	"github.com/sriram-yeluri/nexus/pkg/nxrm/util"
)

type User struct {
	UserId        string        `json:"userId"`
	FirstName     string        `json:"firstName"`
	LastName      string        `json:"lastName"`
	EmailAddress  string        `json:"emailAddress"`
	Source        string        `json:"source"`
	Status        string        `json:"status"`
	ReadOnly      bool          `json:"readOnly"`
	Roles         []string      `json:"roles"`
	ExternalRoles []interface{} `json:"externalRoles"`
}

type UsersManager struct {
	Client *resty.Client
}

func NewUsersManager(client *resty.Client) UsersManager {
	return UsersManager{
		client,
	}
}

func (um *UsersManager) GetUsers() (*[]User, error) {
	users := new([]User)
	uri := fmt.Sprintf("%s/service/rest/v1/security/users", nxrm.NXRMConfig.URL)
	util.Info().Println(uri)
	resp, err := nxrm.Client.R().SetResult(users).Get(uri)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode() {

	case http.StatusOK:
		return users, nil
	default:
		return nil, errors.New("request to nexus failed...")
	}
}
