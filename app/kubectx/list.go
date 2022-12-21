package kubectx

import "crane/pkg/ui"

func (c *KubeCtx) List() error {
	data := make([][]string, 0, len(c.metadata.Contexts)+1)
	data = append(data, []string{"host", "name", "namespace", "cluster", "user"})
	for _, v := range c.metadata.Contexts {
		flag := "  "
		if v.Host == c.metadata.Current {
			flag = "* "
		}
		line := []string{flag + v.Host, v.Name, v.Namespace, v.Cluster, v.User}
		data = append(data, line)
	}
	ui.Table(data)
	return nil
}
