/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// versionCmd represents the test command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version of crane",
	Long:  `Print build version for crane`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v\n", getVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func getVersion() string {
	return "0.0.0"
}
