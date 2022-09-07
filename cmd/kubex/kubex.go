/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package kubex

import (
	"crane/app/kubex"
	"github.com/spf13/cobra"
	"time"
)

var KubexCmd = kubexCmd

// kubexCmd represents the etcd command
var kubexCmd = &cobra.Command{
	Use:   "kubex",
	Short: "kubectl 扩展",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var (
	kubexOptions = new(kubex.Options)
)

func init() {
	kubexCmd.Flags().StringVarP(&kubexOptions.Namespace, "namespace", "n", "", "指定资源命名空间")
	kubexCmd.Flags().StringVarP(&kubexOptions.Contains, "contains", "c", "", "指定资源包含字符串")
	kubexCmd.Flags().DurationVarP(&kubexOptions.Timeout, "timeout", "t", 10*time.Second, "指定超时时间")
	kubexCmd.Flags().BoolVarP(&kubexOptions.Affirm, "affirm", "y", false, "无需确认")
	kubexCmd.Flags().BoolVarP(&kubexOptions.Force, "force", "f", false, "强制执行")

	getCmd.Flags().AddFlagSet(kubexCmd.Flags())
	deleteCmd.Flags().AddFlagSet(kubexCmd.Flags())
	editCmd.Flags().AddFlagSet(kubexCmd.Flags())
}
