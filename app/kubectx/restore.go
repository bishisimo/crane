package kubectx

import (
	"crane/pkg/ui"
	"crane/util"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/pkg/errors"
	"os"
	"path"
)

func (c *KubeCtx) Restore() error {
	mainConfigPath := path.Join(c.Workspace, ".config")
	if !util.IsFileExists(c.workContext) {
		bkPath := mainConfigPath + ".bk"
		if !util.IsFileExists(bkPath) {
			return errors.New("not found")
		}
		err := fileutil.CopyFile(bkPath, mainConfigPath)
		if err != nil {
			return err
		}
	}

	if util.IsFileExists(c.workContext) {
		stat, err := os.Lstat(c.workContext)
		if err != nil {
			return err
		}
		if stat.Mode().Type().IsRegular() {
			if !ui.Confirm("kube config file is exist, are you sure cover that?") {
				return nil
			}
		}
		err = os.Remove(c.workContext)
		if err != nil {
			return err
		}
	}

	err := fileutil.CopyFile(mainConfigPath, c.workContext)
	if err != nil {
		return err
	}
	return nil
}
