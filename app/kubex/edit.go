package kubex

import (
	"crane/pkg/errorx"
	"crane/pkg/ui/list"
)

func (w *Worker) Edit() error {
	if w.Name != "" {
		return w.editOneByName(w.Name)
	}
	err := w.Get(false)
	if err != nil {
		return err
	}
	err = w.ParseResources()
	if err != nil {
		return err
	}
	name, err := w.affirmEdit()
	if err != nil {
		return err
	}
	err = w.editOneByName(name)
	if err != nil {
		return err
	}
	return nil
}

func (w *Worker) editOneByName(name string) error {
	args, err := NewArgument("edit", w.Options).WithKind().WithName(name).WithNamespace().get()
	if err != nil {
		return err
	}
	err = w.process(args)
	if err != nil {
		return err
	}
	return nil
}

func (w *Worker) affirmEdit() (string, error) {
	if len(w.resources) == 0 {
		return "", errorx.NotFound
	}
	if len(w.resources) == 1 {
		return w.resources[0], nil
	}
	l := list.NewList("edit option")
	s, err := l.Select(w.resources)
	if err != nil {
		return "", err
	}
	return w.resources[s], nil
}
