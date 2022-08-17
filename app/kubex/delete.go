package kubex

import (
	"crane/pkg/errorx"
	"fmt"
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
		if w.Contains != "" && !strings.Contains(name, w.Contains) || !w.affirmDelete(name) {
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

func (w *Worker) affirmDelete(name string) bool {
	if w.Force {
		return true
	}
	fmt.Printf("确认删除[%v]: %v ? Y/[N]", w.Kind, name)
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
