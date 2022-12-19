package kubex

import (
	"crane/pkg/errorx"
	"fmt"
	"strings"
)

func (k *Kubex) Get(show bool) error {
	args, err := NewArgument("get", k.Options).WithKind().WithName().WithNamespace().WithOutFormat().get()
	if err != nil {
		return err
	}
	rawOut, err := k.run(args)
	if err != nil {
		return err
	}
	k.RawOut = rawOut
	if show {
		return k.ShowGet()
	}
	return nil
}

func (k *Kubex) ShowGet() error {
	s := string(k.RawOut)
	sp := strings.Split(s, "\n")
	result := make([]string, 0)
	for i, line := range sp {
		if i == 0 || k.Contains == "" || strings.Contains(line, k.Contains) {
			result = append(result, line)
		}
	}
	fmt.Println(strings.Join(result, "\n"))
	return nil
}

func (k *Kubex) ParseResources() error {
	s := strings.TrimSpace(string(k.RawOut))
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
		if k.Contains == "" || strings.Contains(name, k.Contains) {
			k.resources = append(k.resources, name)
		}
	}
	return nil
}
