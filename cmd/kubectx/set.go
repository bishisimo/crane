/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package kubectx

import (
	"crane/app/kubectx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:     "set",
	Aliases: []string{"edit", "e"},
	Short:   "设置 [kubectl context] 资源",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		if setOptions.Target == "" && len(args) > 0 {
			setOptions.Target = args[0]
		}
		kc := kubectx.NewKubeCtx()
		err := kc.Set(setOptions)
		if err != nil {
			log.Warn().Err(err).Send()
			return
		}
	},
}

var setOptions = new(kubectx.SetOptions)

func init() {
	ctxCmd.AddCommand(setCmd)
	setCmd.Flags().StringVarP(&setOptions.Target, "target", "t", "", "指定context资源")
	setCmd.Flags().StringVarP(&setOptions.Name, "name", "", "", "设置名称")
	setCmd.Flags().StringVarP(&setOptions.Namespace, "namespace", "n", "", "设置默认命名空间")
}
