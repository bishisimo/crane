package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/input/confirm"
)

func Confirm(title string) bool {
	keyMap := confirm.KeyMap{
		Quit: components.InterruptKey,
		Yes: key.NewBinding(
			key.WithKeys("y", "Y"),
			key.WithHelp("y/Y", "yes"),
		),
		No: key.NewBinding(
			key.WithKeys("n", "N"),
			key.WithHelp("n/N", "no"),
		)}
	c := infinite.NewConfirm(
		confirm.WithPrompt(title),
		confirm.WithNotice(" [y/n] "),
		confirm.WithKeyMap(keyMap),
	)

	c.Display()

	if c.Value() {
		return true
	}
	return false
}
