// Describe:
package etcd

import (
	"crane/pkg/validate"
)

type SnapshotRestore struct {
	*Snapshot
	DataDir                  string `json:"dataDir" validate:"omitempty,dir" label:"data-dir"`
	HostName                 string `json:"hostName" validate:"omitempty,ip" label:"name"`
	InitialCluster           string `json:"initialCluster" validate:"omitempty,ip" label:"initial-cluster"`
	InitialAdvertisePeerUrls string `json:"initialAdvertisePeerUrls" validate:"omitempty,url" label:"initial-advertise-peer-urls"`
}

func (s SnapshotRestore) validate() error {
	err := validate.Validator.Struct(s)
	if err != nil {
		return err
	}
	return nil
}
