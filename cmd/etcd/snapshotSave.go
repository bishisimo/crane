/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package etcd

import (
	"crane/app/etcd"
	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

// snapshotSaveCmd represents the etcdSnapshotSave command
var snapshotSaveCmd = &cobra.Command{
	Use:   "save",
	Short: "储存etcd的快照",
	Long:  `储存etcd快照`,
	Run: func(cmd *cobra.Command, args []string) {
		if snapshotSave.SavePath == "" && len(args) > 0 {
			snapshotSave.SavePath = args[0]
		}
		log.Info().Str("save", snapshotSave.SavePath).Send()
		err := snapshotSave.Run()
		if err != nil {
			return
		}
	},
}

var (
	snapshotSave = new(etcd.SnapshotSave)
)

func init() {
	snapshotCmd.AddCommand(snapshotSaveCmd)
	snapshotSaveCmd.Flags().AddFlagSet(etcdCmd.Flags())
	snapshotSaveCmd.Flags().AddFlagSet(snapshotCmd.Flags())
	snapshotSaveCmd.Flags().StringVarP(&snapshotSave.SavePath, "save", "s", "./", "etcd 快照存放路径")
	snapshotSave.Snapshot = snapshot
}
