package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/selection/singleselect"
	"github.com/fzdwx/infinite/style"
)

func Select(data []string) (int, error) {
	input := components.NewInput()
	input.Prompt = "Filtering: "
	input.PromptStyle = style.New().Bold().Italic().Fg(color.Cyan)
	keyBinding := components.DefaultSingleKeyMap()
	keyBinding.Choice = key.NewBinding(
		key.WithKeys("tab", "enter"),
		key.WithHelp("tab/enter", "choice it"),
	)
	keyBinding.Confirm = key.NewBinding(
		key.WithKeys("tab", "enter"),
		key.WithHelp("tab/enter", "finish selection"),
	)

	return infinite.NewSingleSelect(
		data,
		singleselect.WithFilterInput(input),
		singleselect.WithKeyBinding(keyBinding),
		singleselect.WithChoiceTextStyle(style.New().Bold().Italic().Fg(color.Green)),
		singleselect.WithValueStyle(style.New().Bold().Italic().Fg(color.Green)),
	).Display("select context!")
}
