package kubectx

import (
	"crane/pkg/ui"
	"sort"
	"strings"
)

func (c *KubeCtx) Select() error {
	data := make([]string, 1, len(c.metadata.Contexts))
	selectMap := make(map[string]string, len(c.metadata.Contexts))
	for k, v := range c.metadata.Contexts {
		flag := "  "
		if k == c.metadata.Current {
			flag = "🔥"
		}
		line := []string{flag, k, v.Name, v.Namespace, v.Cluster}
		showInfo := strings.Join(line, " | ")
		if k == c.metadata.Current {
			data[0] = showInfo
		} else {
			data = append(data, showInfo)
		}
		selectMap[showInfo] = k
	}
	sort.Strings(data[1:])
	i, err := ui.Select(data)
	if err != nil {
		return err
	}
	host := selectMap[data[i]]
	err = c.useFile(host)
	if err != nil {
		return err
	}
	c.metadata.Current = host
	err = c.StoreMetadata()
	if err != nil {
		return err
	}
	return nil
}
