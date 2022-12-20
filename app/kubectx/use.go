package kubectx

import (
	"crane/util"
	"os"
)

type UseOptions struct {
	Target string
}

func (c *KubeCtx) Use(opts *UseOptions) error {
	key, err := c.getKeyByTarget(opts.Target)
	if err != nil {
		return err
	}
	err = c.useFile(key)
	if err != nil {
		return err
	}
	c.metadata.Current = key
	err = c.StoreMetadata()
	if err != nil {
		return err
	}
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
