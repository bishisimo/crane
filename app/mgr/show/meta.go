package show

import (
	"context"
	mgrv1a1 "crane/app/mgr/api/v1alpha1"
	"crane/app/mgr/common"
	"crane/app/mgr/controller/meta"
	"crane/pkg/ui"
	"crane/pkg/ui/list"
	"fmt"
	"k8s.io/apimachinery/pkg/types"
)

type MetaShowOptions struct {
	*BaseShowOptions
}

type MetaShow struct {
	Options        *MetaShowOptions
	RawChan        chan *mgrv1a1.MySQLMeta
	loadingContext context.Context
	loadingCancel  context.CancelFunc
}

func NewMetaShow(options *MetaShowOptions) *MetaShow {
	c := make(chan *mgrv1a1.MySQLMeta, 1)
	ctx, cancel := context.WithCancel(context.Background())
	return &MetaShow{
		Options:        options,
		RawChan:        c,
		loadingContext: ctx,
		loadingCancel:  cancel,
	}
}

func (s MetaShow) Show() error {
	err := s.loading(true)
	if err != nil {
		return err
	}
	switch s.Options.OutFormat {
	case "html":
		return s.ShowByHtml()
	default:
		return s.List()
	}
}

func (s MetaShow) ShowByHtml() error {
	//e := gin.Default()
	//e.GET("/meta", func(c *gin.Context) {
	//	t, err := template.ParseFiles("")
	//	if err != nil {
	//		return
	//	}
	//	c.HTML(200,"meta",t.)
	//})
	return nil
}

func (s MetaShow) loading(show bool) error {
	if show {
		loading := ui.NewLoading(s.loadingContext)
		go func() {
			loading.Show()
		}()
	}
	key := types.NamespacedName{
		Namespace: s.Options.Namespace,
		Name:      s.Options.Name,
	}
	reconciler := meta.NewMetaReconciler(s.RawChan, key)
	err := reconciler.SetupWithManager()
	if err != nil {
		return err
	}
	common.AsyncStartMgr()
	return nil
}

func (s MetaShow) List() error {
	dataChan := make(chan []string, 1)
	go s.parseBackupInfo(dataChan)
	<-s.loadingContext.Done()
	l := list.NewList("Meta")
	return l.Show(dataChan)
}

func (s MetaShow) parseBackupInfo(dataChan chan []string) {
	for c := range s.RawChan {
		status := c.Status
		rows := make([]string, 0, len(status.BackupInfos))
		for i, info := range status.BackupInfos {
			t := "^"
			if info.Spec.Type == mgrv1a1.TypeOfBackupFull {
				t = "#"
			}
			state := ""
			colorStr := ""
			switch info.Status.State {
			case common.StateOfRunning:
				state = ">"
				colorStr = "(fg:blue)"
			case common.StateOfFailed:
				state = "x"
				colorStr = "(fg:red)"
			case common.StateOfSucceeded:
				state = "o"
				colorStr += "(fg:green)"
			case common.StateOfDeleted:
				state = "-"
				colorStr += "(fg:black)"
			case common.StateOfUnknown:
				state = "?"
				colorStr = "(fg:yellow)"
			}
			str := fmt.Sprintf("%-5d %v [%v]%v %v", i, t, state, colorStr, info.Name)
			rows = append(rows, str)
		}
		if s.loadingCancel != nil {
			s.loadingCancel()
			s.loadingCancel = nil
		}
		dataChan <- rows
	}
	return
}
