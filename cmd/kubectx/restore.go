/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package kubectx

import (
	"crane/app/kubectx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// restoreCmd represents the etcd command
var restoreCmd = &cobra.Command{
	Use:     "restore",
	Aliases: []string{""},
	Short:   "恢复 kubectl context",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		kc := kubectx.NewKubeCtx(nil)
		err := kc.Restore()
		if err != nil {
			log.Err(err).Send()
			return
		}
	},
}

func init() {
	ctxCmd.AddCommand(restoreCmd)
}
