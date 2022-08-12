// Describe:
package init

import (
	"gitee.com/bishisimo/log-parser/console"
	"gitee.com/bishisimo/log-parser/parser"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func initLog() {
	zerolog.CallerSkipFrameCount = 2 //这里根据实际，另外获取的是Msg调用处的文件路径和行号
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	out := console.NewConsoleWriter(parser.NewZeroLogParser("crane"))
	log.Logger = log.Output(out).With().Caller().Stack().Logger()
}
