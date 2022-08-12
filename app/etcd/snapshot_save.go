// Describe:
package etcd

import (
	"crane/pkg/validate"
	"crane/util"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"os/exec"
	"path"
)

type SnapshotSave struct {
	*Snapshot
	SavePath string `json:"savePath" validate:"omitempty,file" label:"save"`
}

func (s SnapshotSave) validate() error {
	err := validate.Validator.Struct(s)
	if err != nil {
		return err
	}
	saveDir := path.Dir(s.SavePath)
	log.Debug().Str("saveDir", saveDir).Send()
	if !util.IsFileExists(saveDir) {
		err := os.MkdirAll(saveDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s SnapshotSave) Run() error {
	err := s.validate()
	if err != nil {
		return err
	}
	program := "ETCDCTL_API=3 etcdctl"
	format := `--endpoints=%v --cacert=%v --cert=%v --key=%v save %v`
	args := fmt.Sprintf(format, s.Endpoints, s.CaPath, s.CertPath, s.KeyPath, s.SavePath)
	cmd := exec.Command(program, args)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
