package list

import (
	"crane/pkg/errorx"
	ui "github.com/gizak/termui/v3"
)

func (l *List) Select(rows []string) (int, error) {
	if err := ui.Init(); err != nil {
		return 0, err
	}
	defer ui.Close()
	go l.selectController()
	l.Rows = rows
	ui.Render(l)
	<-l.ctx.Done()
	if l.isCancel {
		return 0, errorx.Canceled
	}
	return l.SelectedRow, nil
}

func (l *List) selectController() {
	go func() {
		previousKey := ""
		uiEvents := ui.PollEvents()
		for {
			e := <-uiEvents
			switch e.ID {
			case "q", "<Chan-c>":
				l.isCancel = true
				l.cancel()
				return
			case "j", "<Down>":
				l.ScrollDown()
			case "k", "<Up>":
				l.ScrollUp()
			case "g":
				if previousKey == "g" {
					l.ScrollTop()
				}
			case "<Home>":
				l.ScrollTop()
			case "G", "<End>":
				l.ScrollBottom()
			case "<Enter>":
				l.cancel()
				return
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
