package list

import ui "github.com/gizak/termui/v3"

func (l *List) Show(data chan []string) error {
	if err := ui.Init(); err != nil {
		return err
	}
	defer ui.Close()
	go l.showController()
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

func (l *List) showController() {
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
