package show

import "crane/app/mgr"

type BaseShowOptions struct {
	*mgr.BaseOptions
	OutFormat string `json:"out_format"`
}

func NewBaseShowOptions() *BaseShowOptions {
	return &BaseShowOptions{
		BaseOptions: &mgr.BaseOptions{
			Namespace: "",
			Name:      "",
		},
		OutFormat: "",
	}
}

type Show interface {
	Show() error
}
