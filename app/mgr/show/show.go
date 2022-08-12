package show

type BaseShowOptions struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

type Show interface {
	Show() error
}
