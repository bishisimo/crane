package list

import (
	ui "github.com/gizak/termui/v3"
)

func (p *Presenter) Show(data []string) error {
	if err := ui.Init(); err != nil {
		return err
	}
	defer ui.Close()
	go p.showController()
	p.Rows = data
	ui.Render(p)
	<-p.ctx.Done()
	return nil
}

func (p *Presenter) showController() {
	previousKey := ""
	go func() {
		uiEvents := ui.PollEvents()
		for {
			e := <-uiEvents
			switch e.ID {
			case "q", "Escape", "<C-c>":
				p.cancel()
				return
			case "j", "<Down>":
				p.ScrollDown()
			case "k", "<Up>":
				p.ScrollUp()
			case "<Chan-d>":
				p.ScrollHalfPageDown()
			case "<Chan-u>":
				p.ScrollHalfPageUp()
			case "<Chan-f>":
				p.ScrollPageDown()
			case "<Chan-b>":
				p.ScrollPageUp()
			case "g":
				if previousKey == "g" {
					p.ScrollTop()
				}
			case "<Home>":
				p.ScrollTop()
			case "G", "<End>":
				p.ScrollBottom()
			}

			if previousKey == "g" {
				previousKey = ""
			} else {
				previousKey = e.ID
			}

			ui.Render(p)
		}
	}()
}
