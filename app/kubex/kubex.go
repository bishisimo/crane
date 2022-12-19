package kubex

import (
	"context"
	"os"
	"os/exec"
	"time"
)

type Options struct {
	Namespace    string        `json:"namespace"`
	Kind         string        `json:"kind"`
	Name         string        `json:"name"`
	Contains     string        `json:"contains"`
	OutFormat    string        `json:"out_format"`
	Timeout      time.Duration `json:"timeout"`
	Affirm       bool          `json:"affirm"`
	Force        bool          `json:"force"`
	AllNamespace bool          `json:"allNamespace"`
}

type Kubex struct {
	*Options
	cmd       *exec.Cmd
	ctx       context.Context
	cancel    context.CancelFunc
	RawOut    []byte
	resources []string
}

func NewWorker(opts *Options) *Kubex {
	return &Kubex{
		Options: opts,
	}
}

func (k *Kubex) run(args []string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), k.Timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, "kubectl", args...)
	return cmd.CombinedOutput()
}

func (k *Kubex) process(args []string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd := exec.CommandContext(ctx, "kubectl", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
