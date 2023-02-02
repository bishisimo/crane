package kubex

import (
	"crane/pkg/errorx"
	"crane/pkg/ui"
	"fmt"
	"strings"
)

func (k *Kubex) Get() error {
	if k.AllNamespace || k.Contains != "" {
		err := k.preAllData()
		if err != nil {
			return err
		}
	}
	err := k.get()
	if err != nil {
		return err
	}
	return k.ShowGet()
}

func (k *Kubex) preAllData() error {
	args, err := NewArgument("get", k.Options).WithKind().WithNamespace().get()
	if err != nil {
		return err
	}
	rawOut, err := k.run(args)
	if err != nil {
		return err
	}
	k.RawOut = rawOut
	err = k.ParseResources()
	if err != nil {
		return err
	}
	if k.Contains != "" {
		return k.preContains()
	}
	if k.Name != "" {
		for _, meta := range k.resources {
			if meta.Name == k.Name {
				k.AllNamespace = false
				k.Namespace = meta.Namespace
				return nil
			}
		}
	}
	return nil
}

func (k *Kubex) preGetAllNamespace() error {
	args, err := NewArgument("get", k.Options).WithKind().WithNamespace().get()
	if err != nil {
		return err
	}
	rawOut, err := k.run(args)
	if err != nil {
		return err
	}
	k.RawOut = rawOut
	err = k.ParseResources()
	if err != nil {
		return err
	}
	for _, meta := range k.resources {
		if meta.Name == k.Name {
			k.AllNamespace = false
			k.Namespace = meta.Namespace
			return nil
		}
	}
	return errorx.NotFound
}

func (k *Kubex) preContains() error {
	switch len(k.resources) {
	case 0:
		return errorx.NotFound
	case 1:
		k.Contains = ""
		k.AllNamespace = false
		k.Namespace = k.resources[0].Namespace
		k.Name = k.resources[0].Name
	default:
		if k.OutFormat == "" {
			return nil
		}
		data := make([]string, 0, len(k.resources))
		for _, meta := range k.resources {
			data = append(data, meta.Name)
		}
		i, err := ui.Select(data)
		if err != nil {
			return err
		}
		k.Contains = ""
		k.Name = data[i]
	}
	return nil
}

func (k *Kubex) get() error {
	args, err := NewArgument("get", k.Options).WithKind().WithName().WithNamespace().WithOutFormat().get()
	if err != nil {
		return err
	}
	rawOut, err := k.run(args)
	if err != nil {
		return err
	}
	k.RawOut = rawOut
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
