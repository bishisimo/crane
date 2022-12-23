package kubex

import (
	"crane/pkg/errorx"
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

func (k *Kubex) Get() error {
	if k.AllNamespace && k.Name != "" {
		err := k.preGetAllNamespace()
		if errorx.IsNotFound(err) {
			log.Info().Msg("not found")
			return nil
		}
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
