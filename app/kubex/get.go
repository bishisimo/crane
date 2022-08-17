package kubex

import (
	"crane/pkg/errorx"
	"fmt"
	"strings"
)

func (w *Worker) Get(show bool) error {
	args := []string{"get", w.Kind}
	if w.Name != "" {
		args = append(args, w.Name)
	}
	if w.Namespace != "" {
		args = append(args, "-n", w.Namespace)
	}
	rawOut, err := w.run(args)
	if err != nil {
		return err
	}
	w.rawOut = rawOut
	if show {
		fmt.Println(string(w.rawOut))
	}
	return nil
}

func (w *Worker) ParseResources() error {
	err := w.Get(false)
	if err != nil {
		return err
	}
	s := strings.TrimSpace(string(w.rawOut))
	if strings.HasPrefix(s, "No resources found") {
		return errorx.NotFound
	}
	lines := strings.Split(s, "\n")[1:]
	resources := make([]string, 0, len(lines))
	for _, line := range lines {
		items := strings.Split(line, " ")
		resources = append(resources, items[0])
	}
	for _, name := range resources {
		if w.Contains == "" || strings.Contains(name, w.Contains) {
			w.resources = append(w.resources, name)
		}
	}
	return nil
}
