package util

import (
	"github.com/k0kubun/pp"
	"runtime"
)

func getCallerName(stackNum ...int) string {
	stackSkip := 2
	if len(stackNum) == 1 {
		stackSkip += stackNum[0]
	}
	pc, _, _, _ := runtime.Caller(stackSkip)
	return runtime.FuncForPC(pc).Name()
}

func Println(a ...interface{}) {
	_, _ = pp.Println(getCallerName(), a)
}

func Printf(format string, a ...interface{}) {
	_, _ = pp.Printf(getCallerName(), format, a)
}
