/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package mgr

import (
	"crane/app/mgr/show"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// metaCmd represents the mgrMeta command
var metaCmd = &cobra.Command{
	Use:   "meta",
	Short: "操作mgr的meta资源",
	Long:  `操作mgr的meta资源`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
		}
		name := args[0]
		metaShowOptions.Name = name
		s := show.NewMetaShow(metaShowOptions)
		err := s.Show()
		if err != nil {
			log.Fatal().Err(err).Msg("fail")
		}
	},
}
var (
	metaShowOptions = new(show.MetaShowOptions)
)

func init() {
	mgrCmd.AddCommand(metaCmd)
	metaCmd.Flags().AddFlagSet(mgrCmd.Flags())
	metaShowOptions.BaseShowOptions = baseShowOptions
}
