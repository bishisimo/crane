package list

import (
	"crane/pkg/errorx"
	ui "github.com/gizak/termui/v3"
)

func (p *Presenter) Select(rows []string) (int, error) {
	if err := ui.Init(); err != nil {
		return 0, err
	}
	defer ui.Close()
	go p.selectController()
	p.Rows = rows
	ui.Render(p)
	<-p.ctx.Done()
	if p.isCancel {
		return 0, errorx.Canceled
	}
	return p.SelectedRow, nil
}

func (p *Presenter) selectController() {
	go func() {
		previousKey := ""
		uiEvents := ui.PollEvents()
		for {
			e := <-uiEvents
			switch e.ID {
			case "q", "Escape", "<C-c>":
				p.isCancel = true
				p.cancel()
				return
			case "j", "<Down>":
				p.ScrollDown()
			case "k", "<Up>":
				p.ScrollUp()
			case "g":
				if previousKey == "g" {
					p.ScrollTop()
				}
			case "<Home>":
				p.ScrollTop()
			case "G", "<End>":
				p.ScrollBottom()
			case "<Enter>":
				p.cancel()
				return
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
