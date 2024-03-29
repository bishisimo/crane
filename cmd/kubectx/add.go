/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package kubectx

import (
	"crane/app/kubectx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"install"},
	Short:   "添加 [kubectl context] 资源",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 && addOpts.Host == "" && addOpts.AcpUrl == "" {
			_ = cmd.Help()
			return
		}
		if len(args) > 0 && addOpts.AcpUrl == "" {
			err := addOpts.ParserUri(args[0])
			if err != nil {
				log.Warn().Err(err).Send()
				return
			}
		}
		kc := kubectx.NewKubeCtx()
		err := kc.Add(addOpts)
		if err != nil {
			log.Warn().Err(err).Send()
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
	addCmd.Flags().StringVarP(&addOpts.AcpUrl, "acp", "", "", "the domain url for acp, also need cluster")
	addCmd.Flags().StringVarP(&addOpts.Cluster, "cluster", "", "", "the cluster name for acp, also need acp")
	addCmd.Flags().StringVarP(&addOpts.Name, "name", "", "", "rename the nameof this context")
	addCmd.Flags().StringVarP(&addOpts.Namespace, "namespace", "n", "", "reset the namespace of this context")
}
