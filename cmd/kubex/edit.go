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

// editCmd represents the etcd command
var editCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e"},
	Short:   "kubectl 编辑资源",
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
		err := k.Edit()
		if err != nil {
			if errorx.IsNotFound(err) || errorx.IsCanceled(err) {
				log.Warn().Str("Kind", kubexOptions.Kind).Msg(err.Error())
			} else {
				log.Fatal().Err(err).Msg("fail")
			}
		}
	},
}

func init() {
	kubexCmd.AddCommand(editCmd)
}
