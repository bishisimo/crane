// Describe:
package etcd

import (
	"crane/pkg/validate"
)

type Snapshot struct {
	*Etcd
}

func (s Snapshot) validate() error {
	err := validate.Validator.Struct(s)
	if err != nil {
		return err
	}
	return nil
}
