package kubectx

import (
	"crane/pkg/ui"
	"strings"
)

func (c *KubeCtx) Select() error {
	data := make([]string, 0, len(c.metadata.Contexts))
	hosts := make([]string, 0, len(c.metadata.Contexts))
	for k, v := range c.metadata.Contexts {
		line := []string{k, v.Name, v.Namespace}
		data = append(data, strings.Join(line, " | "))
		hosts = append(hosts, k)
	}
	i, err := ui.Select(data)
	if err != nil {
		return err
	}
	err = c.useFile(hosts[i])
	if err != nil {
		return err
	}
	c.metadata.Current = hosts[i]
	c.StoreMetadata()
	return nil
}
