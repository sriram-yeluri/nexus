/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	nxrm "github.com/sriram-yeluri/nexus/pkg/nxrm/status"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the health status of Nexus Repository Manager",
	Long:  `Check the complete status of Nexus Repository Manager`,
	Run: func(cmd *cobra.Command, args []string) {
		nxrm.Check()
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
