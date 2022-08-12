package binlogx

import (
	"github.com/go-mysql-org/go-mysql/replication"
	"io"
)

type Parser struct {
	Path   string
	Parser *replication.BinlogParser
}

func NewParser(binlogPath string) *Parser {
	parser := replication.NewBinlogParser()
	return &Parser{
		Path:   binlogPath,
		Parser: parser,
	}
}

func (p *Parser) DumpBinlogTopN(writer io.Writer, topN int) error {
	i := 0
	err := p.Parser.ParseFile(p.Path, 0, func(event *replication.BinlogEvent) error {
		event.Dump(writer)
		i++
		if topN > 0 && i >= topN {
			p.Parser.Stop()
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
