/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package kubectx

import (
	"crane/app/kubectx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// addCmd represents the etcd command
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"install"},
	Short:   "添加 [kubectl context] 资源",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 && addOpts.Host == "" {
			_ = cmd.Help()
			return
		}
		if len(args) > 0 {
			addOpts.ParserUri(args[0])
		}
		kc := kubectx.NewKubeCtx()
		err := kc.Add(addOpts)
		if err != nil {
			log.Err(err).Send()
			return
		}
		log.Info().Msg("ok")
	},
}

var addOpts = new(kubectx.AddOptions)

func init() {
	ctxCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&addOpts.Host, "host", "H", "", "the sever's host")
	addCmd.Flags().StringVarP(&addOpts.Port, "port", "P", "22", "the sever's host port")
	addCmd.Flags().StringVarP(&addOpts.Username, "username", "u", "root", "the sever's user")
	addCmd.Flags().StringVarP(&addOpts.Password, "password", "p", "", "the sever's password")
	addCmd.Flags().StringVarP(&addOpts.PrivateKey, "key", "k", "", "the sever's key")
	addCmd.Flags().StringVarP(&addOpts.Name, "name", "n", "", "rename the nameof this context")
	addCmd.Flags().StringVarP(&addOpts.Namespace, "namespace", "N", "", "reset the namespace of this context")
}
