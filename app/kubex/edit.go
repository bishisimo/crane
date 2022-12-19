package kubex

import (
	"crane/pkg/errorx"
	"crane/pkg/ui/list"
)

func (k *Kubex) Edit() error {
	if k.Name != "" {
		return k.editOneByName(k.Name)
	}
	err := k.Get(false)
	if err != nil {
		return err
	}
	err = k.ParseResources()
	if err != nil {
		return err
	}
	name, err := k.affirmEdit()
	if err != nil {
		return err
	}
	err = k.editOneByName(name)
	if err != nil {
		return err
	}
	return nil
}

func (k *Kubex) editOneByName(name string) error {
	args, err := NewArgument("edit", k.Options).WithKind().WithName(name).WithNamespace().get()
	if err != nil {
		return err
	}
	err = k.process(args)
	if err != nil {
		return err
	}
	return nil
}

func (k *Kubex) affirmEdit() (string, error) {
	if len(k.resources) == 0 {
		return "", errorx.NotFound
	}
	if len(k.resources) == 1 {
		return k.resources[0], nil
	}
	l := list.NewList("edit option")
	s, err := l.Select(k.resources)
	if err != nil {
		return "", err
	}
	return k.resources[s], nil
}
