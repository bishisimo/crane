/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package kubex

import (
	"crane/app/kubex"
	"crane/pkg/errorx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// deleteCmd represents the etcd command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"remove", "rm"},
	Short:   "kubectl 删除资源",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			return
		}
		kind := args[0]
		kubexOptions.Kind = kind
		if len(args) > 1 {
			kubexOptions.Name = args[1]
		}
		k := kubex.NewWorker(kubexOptions)
		err := k.Delete()
		if errorx.IsNotFound(err) {
			log.Warn().Err(err).Str("Kind", kubexOptions.Kind).Send()
			return
		}
		if err != nil {
			log.Warn().Err(err).Send()
		}
	},
}

func init() {
	kubexCmd.AddCommand(deleteCmd)
}
