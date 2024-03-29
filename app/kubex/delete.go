package kubex

import (
	"crane/pkg/ui"
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

func (k *Kubex) Delete() error {
	if k.Namespace != "" && k.Name != "" {
		return k.deleteOneByName()
	}
	err := k.getPrecision()
	if err != nil {
		return err
	}
	for _, meta := range k.resources {
		if k.Contains != "" && !strings.Contains(meta.Name, k.Contains) || !k.affirmDelete(meta.Name) {
			continue
		}
		if meta.Namespace != "" {
			k.Namespace = meta.Namespace
		}
		k.Name = meta.Name
		err = k.deleteOneByName()
		if err != nil {
			return err
		}
	}
	return nil
}

func (k *Kubex) deleteOneByName() error {
	args, err := NewArgument("delete", k.Options).WithKind().WithName().WithNamespace().WithForce().get()
	if err != nil {
		return err
	}
	rawOut, err := k.run(args)
	if err != nil {
		return err
	}
	log.Info().Msg(string(rawOut))
	return nil
}

func (k *Kubex) affirmDelete(name string) bool {
	if k.Affirm {
		return true
	}
	title := fmt.Sprintf("确认删除[%v]: %v ?", k.Kind, name)
	confirm := ui.Confirm(title)
	return confirm
}
