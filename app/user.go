package app

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
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
	uri := fmt.Sprintf("%s/service/rest/v1/security/users", NXRMConfig.URL)
	InfoLogger.Println(uri)
	resp, err := client.R().SetResult(users).Get(uri)
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
