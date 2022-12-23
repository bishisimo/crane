package errorx

import (
	"github.com/pkg/errors"
	"strings"
)

func errContains(err error, str string) bool {
	if err == nil {
		return false
	}
	return strings.Contains(strings.ToLower(err.Error()), strings.ToLower(str))
}

func IsNotFound(err error) bool {
	return errors.Is(err, NotFound) || errContains(err, NotFound.Error())
}

func IsCanceled(err error) bool {
	return errors.Is(err, Canceled) || errContains(err, Canceled.Error())
}

func IsUnexpectedNewLine(err error) bool {
	return errors.Is(err, UnexpectedNewLine) || errContains(err, UnexpectedNewLine.Error())
}
