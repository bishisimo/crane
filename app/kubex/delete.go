package kubex

import (
	"crane/pkg/errorx"
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

func (k *Kubex) Delete() error {
	if k.Name != "" {
		return k.deleteOneByName(k.Name)
	}
	err := k.Get(false)
	if err != nil {
		return err
	}
	err = k.ParseResources()
	if err != nil {
		return err
	}
	for _, name := range k.resources {
		if k.Contains != "" && !strings.Contains(name, k.Contains) || !k.affirmDelete(name) {
			continue
		}

		err = k.deleteOneByName(name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k *Kubex) deleteOneByName(name string) error {
	args, err := NewArgument("delete", k.Options).WithKind().WithName(name).WithNamespace().WithForce().get()
	if err != nil {
		return err
	}
	rawOut, err := k.run(args)
	if err != nil {
		return err
	}
	log.Info().Msg(string(rawOut))
	return nil
}

func (k *Kubex) affirmDelete(name string) bool {
	if k.Affirm {
		return true
	}
	fmt.Printf("确认删除[%v]: %v ? Y/[N]", k.Kind, name)
	affirm := "N"
	_, err := fmt.Scanln(&affirm)
	if err != nil {
		if !errorx.IsUnexpectedNewLine(err) {
			log.Err(err).Send()
			return false
		}
	}
	if strings.ToUpper(affirm) != "Y" {
		return false
	}
	return true
}
