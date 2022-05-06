/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/sriram-yeluri/nxrm3/app"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get list of users in Nexus",
	Long:  `Get list of users in Nexus`,
	Run: func(cmd *cobra.Command, args []string) {
		um := app.NewUsersManager(resty.New())
		users, err := um.GetUsers()
		if err != nil {
			app.ErrorLogger.Println(err)
			return
		}
		fmt.Println(*users)
	},
}

func init() {
	userCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
