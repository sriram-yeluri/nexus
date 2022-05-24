/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage Nexus Users",
	Long:  `Manage Nexus Users`,
	// Omitting the Run (and RunE) field from the cobra.Command
}

func init() {
	rootCmd.AddCommand(userCmd)
}
