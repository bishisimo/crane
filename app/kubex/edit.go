package kubex

import (
	"crane/pkg/ui/list"
)

func (w *Worker) Edit() error {
	if w.Name != "" {
		return w.editOneByName(w.Name)
	}
	err := w.ParseResources()
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
	args := []string{"edit", w.Kind, name}
	if w.Namespace != "" {
		args = append(args, "-n", w.Namespace)
	}
	err := w.process(args)
	if err != nil {
		return err
	}
	return nil
}

func (w *Worker) affirmEdit() (string, error) {
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
