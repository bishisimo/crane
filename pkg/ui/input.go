package ui

import (
	"github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/input/text"
	"github.com/fzdwx/infinite/theme"
)

func Input(title string) string {
	i := infinite.NewText(
		text.WithPrompt(title),
		text.WithPromptStyle(theme.DefaultTheme.PromptStyle),
	)

	val, _ := i.Display()

	return val
}
