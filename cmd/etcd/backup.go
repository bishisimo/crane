/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package etcd

import (
	"github.com/rs/zerolog/log"
	"os"

	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "etcd 备份",
	Long:  `etcd 备份`,
	Run: func(cmd *cobra.Command, args []string) {
		if snapshotSave.SavePath == "" && len(args) > 0 {
			snapshotSave.SavePath = args[0]
		}
		log.Info().Str("save", snapshotSave.SavePath).Send()
		err := snapshotSave.Run()
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	etcdCmd.AddCommand(backupCmd)
	backupCmd.Flags().AddFlagSet(snapshotSaveCmd.Flags())
}
