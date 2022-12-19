package kubectx

import (
	"crane/pkg/ui"
	"crane/util"
	"github.com/duke-git/lancet/v2/fileutil"
)

func (c *KubeCtx) InitMainConfig() error {
	if util.IsFileExists(c.workContext) {
		if !util.IsRegularFile(c.workContext) {
			return nil
		}
		if !ui.Confirm("kube config file is init, are you sure cover that?") {
			return nil
		}
	}

	err := fileutil.CopyFile(c.workContext, c.mainContext)
	if err != nil {
		return err
	}
	if !util.IsFileExists(c.backupContext) {
		fileutil.CopyFile(c.mainContext, c.backupContext)
	}
	return nil
}
