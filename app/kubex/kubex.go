package kubex

import (
	"context"
	"os"
	"os/exec"
	"time"
)

type Options struct {
	Namespace string        `json:"namespace"`
	Kind      string        `json:"kind"`
	Name      string        `json:"name"`
	Contains  string        `json:"contains"`
	OutFormat string        `json:"out_format"`
	Timeout   time.Duration `json:"timeout"`
	Affirm    bool          `json:"affirm,omitempty"`
	Force     bool          `json:"force"`
}

type Worker struct {
	*Options
	cmd       *exec.Cmd
	ctx       context.Context
	cancel    context.CancelFunc
	RawOut    []byte
	resources []string
}

func NewWorker(opts *Options) *Worker {
	return &Worker{
		Options: opts,
	}
}

func (w *Worker) run(args []string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), w.Timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, "kubectl", args...)
	return cmd.CombinedOutput()
}

func (w *Worker) process(args []string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd := exec.CommandContext(ctx, "kubectl", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
