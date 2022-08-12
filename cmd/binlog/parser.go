/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package binlog

import (
	"crane/app/binlog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// mgrCmd represents the mgr command
var parserCmd = &cobra.Command{
	Use:   "parser",
	Short: "解析binlog文件",
	Long:  `解析指定的mysql binlog文件,指定输出位置,默认空为标准输出`,
	Run: func(cmd *cobra.Command, args []string) {
		if parserOptions.SourceFile == "" && len(args) > 0 {
			parserOptions.SourceFile = args[0]
		}
		p := binlog.NewParser(parserOptions)
		err := p.ToFile()
		if err != nil {
			log.Fatal().Err(err).Msg("解析失败")
			return
		}
	},
}

var (
	parserOptions = new(binlog.ParserOptions)
)

func init() {
	binlogCmd.AddCommand(parserCmd)
	parserCmd.Flags().StringVarP(&parserOptions.OutputFile, "output", "o", "", "输出文件")
	parserCmd.Flags().IntVarP(&parserOptions.Limit, "limit", "l", 0, "限制Event数量")
	binlogCmd.Flags().AddFlagSet(binlogCmd.Flags())
	parserOptions.BaseOptions = baseOptions
}
