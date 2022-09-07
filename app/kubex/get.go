package kubex

import (
	"crane/pkg/errorx"
	"fmt"
	"strings"
)

func (w *Worker) Get(show bool) error {
	args, err := NewArgument("get", w.Options).WithKind().WithName().WithNamespace().WithOutFormat().get()
	if err != nil {
		return err
	}
	rawOut, err := w.run(args)
	if err != nil {
		return err
	}
	w.RawOut = rawOut
	if show {
		return w.ShowGet()
	}
	return nil
}

func (w *Worker) ShowGet() error {
	s := string(w.RawOut)
	sp := strings.Split(s, "\n")
	result := make([]string, 0)
	for i, line := range sp {
		if i == 0 || w.Contains == "" || strings.Contains(line, w.Contains) {
			result = append(result, line)
		}
	}
	fmt.Println(strings.Join(result, "\n"))
	return nil
}

func (w *Worker) ParseResources() error {
	s := strings.TrimSpace(string(w.RawOut))
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
