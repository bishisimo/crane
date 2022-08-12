package kubex

import (
	"context"
	"os/exec"
	"time"
)

type Options struct {
	Namespace string        `json:"namespace"`
	Kind      string        `json:"kind"`
	Name      string        `json:"name"`
	Timeout   time.Duration `json:"timeout"`
	Contains  string        `json:"contains"`
}

type Worker struct {
	*Options
	cmd       *exec.Cmd
	ctx       context.Context
	cancel    context.CancelFunc
	rawOut    []byte
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
