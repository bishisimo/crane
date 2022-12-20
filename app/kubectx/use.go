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
	c.metadata.Current = host
	c.StoreMetadata()
	return nil
}

func (c *KubeCtx) useFile(host string) error {
	filePath := c.metadata.Contexts[host].Path
	if util.IsRegularFile(c.mainContext) && !util.IsFileExists(c.wardContext) {
		err := c.InitMainConfig()
		if err != nil {
			return err
		}
	}
	err := os.Remove(c.mainContext)
	if err != nil {
		return err
	}
	err = os.Symlink(filePath, c.mainContext)
	if err != nil {
		return err
	}
	return nil
}
