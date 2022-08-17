/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package mgr

import (
	"crane/app/mgr/show"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// metaCmd represents the mgrMeta command
var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "操作mgr的cluster资源",
	Long:  `操作mgr的cluster资源`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			_ = cmd.Help()
			return errors.New("miss args of name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		clusterShowOptions.Name = name
		s := show.NewClusterShow(clusterShowOptions)
		err := s.Show()
		if err != nil {
			log.Fatal().Err(err).Msg("fail")
		}
	},
}

var (
	clusterShowOptions = new(show.ClusterShowOptions)
)

func init() {
	mgrCmd.AddCommand(clusterCmd)
	clusterShowOptions.BaseShowOptions = baseShowOptions
}
