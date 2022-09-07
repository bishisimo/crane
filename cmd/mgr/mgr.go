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
	Short: "mgr tools",
	Long:  `辅助调试mgr`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var (
	baseShowOptions = show.NewBaseShowOptions()
)

func init() {
	mgrCmd.Flags().StringVarP(&baseShowOptions.Namespace, "namespace", "n", "", "资源所在命名空间")
	mgrCmd.Flags().StringVarP(&baseShowOptions.OutFormat, "out", "o", "terminal", "展示形式")
	clusterCmd.Flags().AddFlagSet(mgrCmd.Flags())
	metaCmd.Flags().AddFlagSet(mgrCmd.Flags())
}
