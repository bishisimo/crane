package kubectx

import (
	"crane/pkg/ui"
	"strings"
)

func (c *KubeCtx) Select() error {
	data := make([]string, 0, len(c.metadata))
	hosts := make([]string, 0, len(c.metadata))
	for k, v := range c.metadata {
		line := []string{k, v.Name, v.Ctx.Namespace}
		data = append(data, strings.Join(line, " | "))
		hosts = append(hosts, k)
	}
	i, err := ui.Select(data)
	if err != nil {
		return err
	}
	c.useFile(hosts[i])
	return nil
}
