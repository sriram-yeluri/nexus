/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Manage Nexus Repositories",
	Long:  `Manage Nexus Repositories. Supports list, add, delete`,
}

func init() {
	rootCmd.AddCommand(repoCmd)
}
