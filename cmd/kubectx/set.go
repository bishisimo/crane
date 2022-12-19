/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package kubectx

import (
	"crane/app/kubectx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// setCmd represents the etcd command
var setCmd = &cobra.Command{
	Use:     "set",
	Aliases: []string{"edit", "e"},
	Short:   "设置 [kubectl context] 资源",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		kc := kubectx.NewKubeCtx(nil)
		err := kc.Set(setOptions)
		if err != nil {
			log.Err(err).Send()
			return
		}
	},
}

var setOptions = new(kubectx.SetOptions)

func init() {
	ctxCmd.AddCommand(setCmd)
	setCmd.Flags().StringVarP(&setOptions.Target, "target", "t", "", "设置指定的context资源")
	setCmd.Flags().StringVarP(&setOptions.Name, "name", "n", "", "设置名称")
	setCmd.Flags().StringVarP(&setOptions.Namespace, "namespace", "N", "", "设置默认命名空间")
}
