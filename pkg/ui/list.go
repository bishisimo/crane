package ui

import (
	"context"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type List struct {
	*widgets.List
	ctx    context.Context
	cancel context.CancelFunc
}

func NewList() *List {
	l := widgets.NewList()
	l.Title = "Meta"
	l.TextStyle = ui.NewStyle(ui.ColorClear)
	l.WrapText = false
	l.SetRect(0, 0, 100, 12)
	l.SelectedRowStyle = ui.NewStyle(ui.ColorCyan)
	ctx, cancel := context.WithCancel(context.Background())
	return &List{
		List:   l,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (l List) Show(data chan []string) error {
	if err := ui.Init(); err != nil {
		return err
	}
	defer ui.Close()
	go l.controller()
	for {
		select {
		case <-l.ctx.Done():
			return nil
		case rows := <-data:
			l.Rows = rows
			ui.Render(l)
		}
	}
}

func (l List) controller() {
	previousKey := ""
	go func() {
		uiEvents := ui.PollEvents()
		for {
			e := <-uiEvents
			switch e.ID {
			case "q", "<Chan-c>":
				l.cancel()
				return
			case "j", "<Down>":
				l.ScrollDown()
			case "k", "<Up>":
				l.ScrollUp()
			case "<Chan-d>":
				l.ScrollHalfPageDown()
			case "<Chan-u>":
				l.ScrollHalfPageUp()
			case "<Chan-f>":
				l.ScrollPageDown()
			case "<Chan-b>":
				l.ScrollPageUp()
			case "g":
				if previousKey == "g" {
					l.ScrollTop()
				}
			case "<Home>":
				l.ScrollTop()
			case "G", "<End>":
				l.ScrollBottom()
			}

			if previousKey == "g" {
				previousKey = ""
			} else {
				previousKey = e.ID
			}

			ui.Render(l)
		}
	}()
}
