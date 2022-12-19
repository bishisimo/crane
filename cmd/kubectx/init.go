/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package kubectx

import (
	"crane/app/kubectx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// initCmd represents the etcd command
var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{""},
	Short:   "初始化 [kubectl context] 资源",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		kc := kubectx.NewKubeCtx(nil)
		err := kc.InitMainConfig()
		if err != nil {
			log.Err(err).Send()
			return
		}
	},
}

func init() {
	ctxCmd.AddCommand(initCmd)
}
