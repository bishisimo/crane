/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package kubectx

import (
	"crane/app/kubectx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// selectCmd represents the etcd command
var selectCmd = &cobra.Command{
	Use:     "select",
	Aliases: []string{"s"},
	Short:   "选择 [kubectl context] 资源",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		kc := kubectx.NewKubeCtx()
		err := kc.Select()
		if err != nil {
			log.Err(err).Send()
			return
		}
	},
}

func init() {
	ctxCmd.AddCommand(selectCmd)
}
