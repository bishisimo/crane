package kubectx

import (
	"crane/pkg/ui"
	"crane/util"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/pkg/errors"
	"os"
)

func (c *KubeCtx) Restore() error {
	if !util.IsFileExists(c.mainContext) {
		if !util.IsFileExists(c.backupContext) {
			return errors.New("not found")
		}
		err := fileutil.CopyFile(c.backupContext, c.wardContext)
		if err != nil {
			return err
		}
	}

	if util.IsFileExists(c.mainContext) {
		if util.IsRegularFile(c.mainContext) {
			if !ui.Confirm("kube config file is exist, are you sure cover that?") {
				return nil
			}
		}
		err := os.Remove(c.mainContext)
		if err != nil {
			return err
		}
	}

	err := fileutil.CopyFile(c.wardContext, c.mainContext)
	if err != nil {
		return err
	}
	return nil
}
