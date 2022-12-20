/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package kubectx

import (
	"crane/app/kubectx"
	"github.com/spf13/cobra"
)

// listCmd represents the etcd command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "展示 [kubectl context] 资源",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		kc := kubectx.NewKubeCtx()
		err := kc.List()
		if err != nil {
			return
		}
	},
}

func init() {
	ctxCmd.AddCommand(listCmd)
}
