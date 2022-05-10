/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/sriram-yeluri/nexus/pkg/nxrm/user"
	"github.com/sriram-yeluri/nexus/pkg/nxrm/util"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "users",
	Short: "Get the list of all the Nexus Users",
	Long:  `Get the ist of all the Nexus Users`,
	Run: func(cmd *cobra.Command, args []string) {
		um := user.NewUsersManager(resty.New())
		users, err := um.GetUsers()
		if err != nil {
			util.Error().Println(err)
			return
		}
		fmt.Println(*users)
	},
}

func init() {
	listCmd.AddCommand(userCmd)
}
