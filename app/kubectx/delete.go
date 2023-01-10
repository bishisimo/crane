package kubectx

import (
	"os"
	"path"
)

type DeleteOptions struct {
	Target string
}

func (c *KubeCtx) Delete(opts *DeleteOptions) error {
	key, err := c.getKeyByTarget(opts.Target)
	if err != nil {
		return err
	}
	err = c.DeleteMetadata(key)
	if err != nil {
		return err
	}
	err = c.StoreMetadata()
	if err != nil {
		return err
	}
	targetPath := path.Join(c.Workspace, key)
	err = os.Remove(targetPath)
	if err != nil {
		return err
	}
	return nil
}
