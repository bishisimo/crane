/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package kubectx

import (
	"crane/app/kubectx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// getCmd represents the etcd command
var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"info"},
	Short:   "查看指定 [kubectl context] 资源",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && getOptions.Target == "" {
			getOptions.Target = args[0]
		}
		kc := kubectx.NewKubeCtx()
		err := kc.Get(getOptions)
		if err != nil {
			log.Err(err).Send()
			return
		}
	},
}

var getOptions = new(kubectx.GetOptions)

func init() {
	ctxCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&getOptions.Target, "target", "t", "", "查看指定资源的信息")
}
