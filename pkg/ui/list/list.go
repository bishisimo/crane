package list

import (
	"context"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Presenter struct {
	*widgets.List
	ctx      context.Context
	cancel   context.CancelFunc
	isCancel bool
}

func NewList(title string) *Presenter {
	l := widgets.NewList()
	l.Title = title
	l.TextStyle = ui.NewStyle(ui.ColorClear)
	l.WrapText = false
	l.SetRect(0, 0, 100, 12)
	l.SelectedRowStyle = ui.NewStyle(ui.ColorCyan)
	ctx, cancel := context.WithCancel(context.Background())
	return &Presenter{
		List:   l,
		ctx:    ctx,
		cancel: cancel,
	}
}
