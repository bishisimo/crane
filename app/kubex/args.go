package kubex

import (
	"crane/pkg/errorx"
	"github.com/pkg/errors"
)

type argument struct {
	*Options
	data []string
	err  error
}

func NewArgument(cmd string, opts *Options) *argument {
	return &argument{
		Options: opts,
		data:    []string{cmd},
	}
}

func (a *argument) get() ([]string, error) {
	return a.data, a.err
}

func (a *argument) WithKind() *argument {
	if a.err != nil {
		return a
	}
	if a.Kind == "" {
		a.err = errors.Wrap(errorx.NotFound, "kind")
		return a
	}
	a.data = append(a.data[:1], append([]string{a.Kind}, a.data[1:]...)...)
	return a
}

func (a *argument) WithNamespace() *argument {
	if a.err == nil && a.Namespace != "" {
		a.data = append(a.data, "-n", a.Namespace)
	}
	return a
}

func (a *argument) WithName(name ...string) *argument {
	if a.err != nil {
		return a
	}
	if len(name) > 0 {
		a.data = append(a.data, name[0])
	} else if a.Name != "" {
		a.data = append(a.data, a.Name)
	}
	return a
}

func (a *argument) WithOutFormat() *argument {
	if a.err == nil && a.OutFormat != "" {
		a.data = append(a.data, "-o", a.OutFormat)
	}
	return a
}

func (a *argument) WithForce() *argument {
	if a.err == nil && a.Force {
		a.data = append(a.data, "--force")
	}
	return a
}
