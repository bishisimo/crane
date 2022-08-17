package errorx

import "github.com/pkg/errors"

var (
	NotFound          = errors.New("not found")
	Canceled          = errors.New("canceled")
	UnexpectedNewLine = errors.New("unexpected newline")
)
