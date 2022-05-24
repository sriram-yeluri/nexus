/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sriram-yeluri/nexus/pkg/nxrm/nxrm"
	"github.com/sriram-yeluri/nexus/pkg/nxrm/user"
	"github.com/sriram-yeluri/nexus/pkg/nxrm/util"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the list of all the Nexus Users",
	Long:  `Get the list of all the Nexus Users`,
	Run: func(cmd *cobra.Command, args []string) {
		um := user.NewUsersManager(nxrm.Client)
		users, err := um.GetUsers()
		if err != nil {
			util.Error().Println(err)
			return
		}

		fmt.Println(" --------- Nexus Users ---------- ")
		for _, user := range *users {
			fmt.Println(user.UserId)
		}
	},
}

// listRepoCmd represents the list command for repositories
var listRepoCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the list of all the Nexus Repositories",
	Long:  `Get the list of all the Nexus Repositories`,
	Run: func(cmd *cobra.Command, args []string) {
		um := user.NewUsersManager(nxrm.Client)
		users, err := um.GetUsers()
		if err != nil {
			util.Error().Println(err)
			return
		}

		fmt.Println(" --------- Nexus Users ---------- ")
		for _, user := range *users {
			fmt.Println(user.UserId)
		}
	},
}

func init() {
	userCmd.AddCommand(listCmd)
	repoCmd.AddCommand(listRepoCmd)
}
