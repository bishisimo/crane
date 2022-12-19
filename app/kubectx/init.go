package kubectx

import (
	"crane/pkg/ui"
	"crane/util"
	"github.com/duke-git/lancet/v2/fileutil"
	"os"
	"path"
)

func (c *KubeCtx) InitMainConfig() error {
	mainConfigPath := path.Join(c.Workspace, ".config")
	if util.IsFileExists(c.workContext) {
		stat, err := os.Lstat(c.workContext)
		if err != nil {
			return err
		}
		if !stat.Mode().Type().IsRegular() {
			return nil
		}
		if !ui.Confirm("kube config file is init, are you sure cover that?") {
			return nil
		}
	}

	err := fileutil.CopyFile(c.workContext, mainConfigPath)
	if err != nil {
		return err
	}
	bkPath := mainConfigPath + ".bk"
	if !util.IsFileExists(bkPath) {
		fileutil.CopyFile(mainConfigPath, bkPath)
	}
	return nil
}
