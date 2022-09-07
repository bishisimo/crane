package ui

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

type errMsg error

type Loading struct {
	spinner  spinner.Model
	quitting bool
	err      error
	context.Context
}

func NewLoading(ctx context.Context) *Loading {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return &Loading{
		spinner: s,
		Context: ctx,
	}
}

func (l Loading) Show() {
	p := tea.NewProgram(l)
	go func() {
		<-l.Context.Done()
		p.Send(tea.KeyMsg{
			Type:  -1,
			Runes: []rune{'o', 'k'},
		})
	}()
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (l Loading) Init() tea.Cmd {
	return l.spinner.Tick
}

func (l Loading) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			os.Exit(0)
			return nil, nil
		case "ok":
			l.quitting = true
			return l, tea.Quit
		default:
			return l, nil
		}
	case errMsg:
		l.err = msg
		return l, nil

	default:
		var cmd tea.Cmd
		l.spinner, cmd = l.spinner.Update(msg)
		return l, cmd
	}

}

func (l Loading) View() string {
	if l.err != nil {
		return l.err.Error()
	}
	str := fmt.Sprintf("\n\n   %s loading ...press q to quit\n\n", l.spinner.View())
	if l.quitting {
		return str + "\n"
	}
	return str
}
