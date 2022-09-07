/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package kubex

import (
	"crane/app/kubex"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// getCmd represents the etcd command
var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"list", "ls"},
	Short:   "kubectl 查看资源",
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
		err := k.Get(true)
		if err != nil {
			log.Fatal().Err(err).Msg("fail")
		}
	},
}

func init() {
	kubexCmd.AddCommand(getCmd)
	kubexCmd.Flags().StringVarP(&kubexOptions.OutFormat, "out", "o", "", "指定资源输出格式")
}
