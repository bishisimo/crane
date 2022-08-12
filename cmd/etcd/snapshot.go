/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package etcd

import (
	"crane/app/etcd"
	"github.com/spf13/cobra"
)

// snapshotCmd represents the snapshot command
var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "etcd 快照",
	Long:  `将etcd数据进行快照备份`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var (
	snapshot = new(etcd.Snapshot)
)

func init() {
	etcdCmd.AddCommand(snapshotCmd)
	snapshotCmd.Flags().AddFlagSet(etcdCmd.Flags())
	snapshot.Etcd = etcdEntity
}
