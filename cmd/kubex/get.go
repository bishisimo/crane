/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package kubex

import (
	"crane/app/kubex"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// etcdCmd represents the etcd command
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
	getCmd.Flags().AddFlagSet(kubexCmd.Flags())
}
