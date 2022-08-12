package kubex

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"strings"
)

func (w *Worker) Delete() error {
	if w.Name != "" {
		return w.deleteOneByName(w.Name)
	}
	err := w.ParseResources()
	if err != nil {
		return err
	}
	for _, name := range w.resources {
		if w.Contains != "" && !strings.Contains(name, w.Contains) {
			continue
		}
		err = w.deleteOneByName(name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *Worker) deleteOneByName(name string) error {
	args := []string{"delete", w.Kind, name}
	if w.Namespace != "" {
		args = append(args, "-n", w.Namespace)
	}
	rawOut, err := w.run(args)
	if err != nil {
		return err
	}
	log.Info().Msg(string(rawOut))
	return nil
}

func (w *Worker) ParseResources() error {
	err := w.Get(false)
	if err != nil {
		return err
	}
	s := strings.TrimSpace(string(w.rawOut))
	if strings.HasPrefix(s, "No resources found") {
		return errors.New("No resources found")
	}
	lines := strings.Split(s, "\n")[1:]
	resources := make([]string, 0, len(lines))
	for _, line := range lines {
		items := strings.Split(line, " ")
		resources = append(resources, items[0])
	}
	w.resources = resources
	return nil
}
