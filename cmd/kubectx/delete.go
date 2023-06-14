/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package kubectx

import (
	"crane/app/kubectx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"remove", "rm"},
	Short:   "删除 [kubectl context] 资源",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 && deleteOpts.Target == "" {
			_ = cmd.Help()
			return
		}
		if len(args) > 0 && deleteOpts.Target == "" {
			deleteOpts.Target = args[0]
		}
		kc := kubectx.NewKubeCtx()
		err := kc.Delete(deleteOpts)
		if err != nil {
			log.Warn().Err(err).Send()
			return
		}
	},
}

var deleteOpts = new(kubectx.DeleteOptions)

func init() {
	ctxCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&deleteOpts.Target, "target", "t", "", "删除指定的context资源")
}
