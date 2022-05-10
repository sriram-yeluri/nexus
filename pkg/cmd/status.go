/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sriram-yeluri/nexus/pkg/nxrm/status"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the status of Nexus Repository Manager",
	Long:  `Get the status of Nexus Repository Manager`,
	Run: func(cmd *cobra.Command, args []string) {
		status.Status()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
