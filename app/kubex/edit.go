package kubex

import (
	"crane/pkg/errorx"
	"crane/pkg/ui/list"
)

func (k *Kubex) Edit() error {
	if k.Namespace != "" && k.Name != "" {
		return k.editOneByName()
	}
	err := k.getPrecision()
	if err != nil {
		return err
	}
	err = k.affirmEdit()
	if err != nil {
		return err
	}
	err = k.editOneByName()
	if err != nil {
		return err
	}
	return nil
}

func (k *Kubex) editOneByName() error {
	args, err := NewArgument("edit", k.Options).WithKind().WithName().WithNamespace().get()
	if err != nil {
		return err
	}
	err = k.process(args)
	if err != nil {
		return err
	}
	return nil
}

func (k *Kubex) affirmEdit() error {
	if len(k.resources) == 0 {
		return errorx.NotFound
	}
	if len(k.resources) == 1 {
		return nil
	}
	data := make([]string, 0, len(k.resources))
	for _, meta := range k.resources {
		data = append(data, meta.Name)
	}
	l := list.NewList("edit option")
	i, err := l.Select(data)
	if err != nil {
		return err
	}
	k.Name = k.resources[i].Name
	if k.resources[i].Namespace != "" {
		k.Namespace = k.resources[i].Namespace
	}
	return nil
}
