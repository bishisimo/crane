/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package kubectx

import (
	"github.com/spf13/cobra"
)

var Cmd = ctxCmd

// ctxCmd represents the etcd command
var ctxCmd = &cobra.Command{
	Use: "context",
	Aliases: []string{
		"ctx",
	},
	Short: "管理kubectl使用的context",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {

}
