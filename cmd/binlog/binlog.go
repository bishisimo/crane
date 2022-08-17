/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package binlog

import (
	"crane/app/binlog"
	"github.com/spf13/cobra"
)

var BinlogCmd = binlogCmd

// mgrCmd represents the mgr command
var binlogCmd = &cobra.Command{
	Use:   "binlog",
	Short: "binlog 命令集合",
	Long:  `binlog 命令集合`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var (
	baseOptions = new(binlog.BaseOptions)
)

func init() {
	binlogCmd.Flags().StringVarP(&baseOptions.SourceFile, "source", "s", "", "源文件")
	binlogCmd.Flags().AddFlagSet(binlogCmd.Flags())
}
