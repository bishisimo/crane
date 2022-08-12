/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package mgr

import (
	"crane/app/mgr/show"
	"github.com/spf13/cobra"
)

var MgrCmd = mgrCmd

// mgrCmd represents the mgr command
var mgrCmd = &cobra.Command{
	Use:   "mgr",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var (
	baseShowOptions = new(show.BaseShowOptions)
)

func init() {
	mgrCmd.Flags().StringVarP(&baseShowOptions.Namespace, "namespace", "n", "", "资源所在命名空间")
	mgrCmd.Flags().StringVarP(&baseShowOptions.Name, "target", "t", "", "目标资源")
}
