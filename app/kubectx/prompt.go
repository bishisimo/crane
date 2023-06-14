package kubectx

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/pkg/errors"
)

type PromptOptions struct {
	Prefix    string
	Suffix    string
	Icon      string
	Separator string
	Divider   string
}

func (c *KubeCtx) Prompt(opts *PromptOptions) error {
	key, err := c.getKeyByTarget(c.metadata.Current)
	if err != nil {
		return err
	}
	info := c.metadata.Contexts[key]
	if info == nil {
		return errors.New("not found")
	}
	fmt.Printf("%v%v%v%v%v%v%v", opts.Prefix, color.Cyan.Render(opts.Icon), opts.Separator, color.Red.Render(info.Name), opts.Divider, color.Blue.Render(info.Namespace), opts.Suffix)
	return nil
}
