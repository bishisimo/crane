package kubectx

import (
	"crane/util"
	"os"
)

type UseOptions struct {
	Target string
}

func (c *KubeCtx) Use(opts *UseOptions) error {
	host, err := c.getHostByTarget(opts.Target)
	if err != nil {
		return err
	}
	err = c.useFile(host)
	if err != nil {
		return err
	}
	return nil
}

func (c *KubeCtx) useFile(host string) error {
	filePath := c.metadata[host].Path
	if util.IsFileExists(c.workContext) {
		err := os.Remove(c.workContext)
		if err != nil {
			return err
		}
	}
	err := os.Symlink(filePath, c.workContext)
	if err != nil {
		return err
	}
	return nil
}
