package kubectx

import (
	"crane/util"
	"github.com/pkg/errors"
)

type GetOptions struct {
	Target string
}

func (c *KubeCtx) Get(opts *GetOptions) error {
	if opts.Target == "" {
		opts.Target = c.metadata.Current
	}
	key, err := c.getKeyByTarget(opts.Target)
	if err != nil {
		return err
	}
	if info, ok := c.metadata.Contexts[key]; ok {
		util.Println(info)
		return nil
	}
	return errors.New("not found")
}
