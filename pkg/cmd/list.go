/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get list of users and repositories in Nexus",
	Long:  `Get list of users and repositories in Nexus`,
	// Omitting the Run (and RunE) field from the cobra.Command
	// will make it a requirement for a valid subcommand to be given
}

func init() {
	rootCmd.AddCommand(listCmd)
}
