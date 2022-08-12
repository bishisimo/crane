package binlog

import (
	"crane/pkg/binlogx"
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
	"time"
)

var (
	fileNameFormat = "2006-01-02-15-04-05.sql"
)

type ParserOptions struct {
	*BaseOptions
	OutputFile string
	Limit      int
}

type Parser struct {
	*ParserOptions
}

func NewParser(opts *ParserOptions) *Parser {
	return &Parser{
		ParserOptions: opts,
	}
}

func (p *Parser) ToFile() error {
	err := p.fullOutput()
	if err != nil {
		return err
	}

	var out io.Writer
	if p.OutputFile == "" {
		out = os.Stdout
	} else {
		out, err = os.Create(p.OutputFile)
		if err != nil {
			return err
		}
	}

	parser := binlogx.NewParser(p.SourceFile)
	err = parser.DumpBinlogTopN(out, p.Limit)
	if err != nil {
		return err
	}
	return nil
}

func (p *Parser) fullOutput() error {
	if p.OutputFile == "" {
		return nil
	} else if _, err := os.Stat(p.OutputFile); !os.IsNotExist(err) {
		p.OutputFile = path.Join(p.OutputFile, time.Now().Format(fileNameFormat))
	} else {
		dir := path.Dir(p.OutputFile)
		if _, err = os.Stat(dir); !os.IsNotExist(err) {
			return errors.Wrap(err, "invalid path")
		}
	}
	return nil
}
