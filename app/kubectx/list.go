package kubectx

import "crane/pkg/ui"

func (c *KubeCtx) List() error {
	data := make([][]string, 0, len(c.metadata)+1)
	data = append(data, []string{"host", "name", "namespace", "cluster", "user"})
	for _, v := range c.metadata {
		line := []string{v.Host, v.Name, v.Ctx.Namespace, v.Ctx.Cluster, v.Ctx.User}
		data = append(data, line)
	}
	ui.Table(data)
	return nil
}
