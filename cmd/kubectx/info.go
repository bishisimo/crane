/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package kubectx

import (
	"crane/app/kubectx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:     "info",
	Aliases: []string{},
	Short:   "查看指定 [kubectl context] 资源",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && getOptions.Target == "" {
			getOptions.Target = args[0]
		}
		kc := kubectx.NewKubeCtx()
		err := kc.Get(getOptions)
		if err != nil {
			log.Warn().Err(err).Send()
			return
		}
	},
}

var getOptions = new(kubectx.GetOptions)

func init() {
	ctxCmd.AddCommand(infoCmd)
	infoCmd.Flags().StringVarP(&getOptions.Target, "target", "t", "", "查看指定资源的信息")
}
