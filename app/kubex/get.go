package kubex

import "fmt"

func (w *Worker) Get(show bool) error {
	args := []string{"get", w.Kind}
	if w.Name != "" {
		args = append(args, w.Name)
	}
	if w.Namespace != "" {
		args = append(args, "-o", w.Namespace)
	}
	rawOut, err := w.run(args)
	if err != nil {
		return err
	}
	w.rawOut = rawOut
	if show {
		fmt.Println(string(w.rawOut))
	}
	return nil
}
