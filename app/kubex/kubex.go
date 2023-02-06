package kubex

import (
	"context"
	"crane/pkg/errorx"
	"os"
	"os/exec"
	"strings"
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
	resources []*Metadata
}

type Metadata struct {
	Namespace string
	Name      string
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

func (k *Kubex) ParseResources() error {
	s := strings.TrimSpace(string(k.RawOut))
	if strings.HasPrefix(s, "No resources found") {
		return errorx.NotFound
	}
	lines := strings.Split(s, "\n")[1:]
	resources := make([]*Metadata, 0, len(lines))
	namespaceIndex := 0
	nameIndex := 0
	if k.AllNamespace {
		nameIndex = 1
	}
	for _, line := range lines {
		items := strings.Fields(line)
		meta := &Metadata{
			Name: items[nameIndex],
		}
		if namespaceIndex != nameIndex {
			meta.Namespace = items[namespaceIndex]
		}
		resources = append(resources, meta)
	}
	if k.Contains == "" {
		k.resources = resources
		return nil
	}
	k.resources = make([]*Metadata, 0, len(lines))
	for _, meta := range resources {
		if strings.Contains(meta.Name, k.Contains) {
			k.resources = append(k.resources, meta)
		}
	}
	return nil
}
