package kubectx

import (
	"crane/pkg/ui"
	"crane/util"
	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"path"
	"strings"
)

type ContextMetadata struct {
	Host      string `toml:"host"`
	Name      string `toml:"name"`
	Namespace string `toml:"namespace"`
	Path      string `toml:"path"`
	User      string `toml:"user"`
	Cluster   string `toml:"cluster"`
}

type Metadata struct {
	Current  string                      `toml:"current"`
	Contexts map[string]*ContextMetadata `toml:"contexts"`
}

type KubeCtx struct {
	Workspace     string
	mainContext   string
	wardContext   string
	backupContext string
	metaPath      string
	viper         *viper.Viper
	metadata      *Metadata
}

func NewKubeCtx() *KubeCtx {
	workspace := viper.GetString("crane.kubeSpace")
	metadata := &Metadata{
		Current:  "",
		Contexts: make(map[string]*ContextMetadata),
	}
	k := &KubeCtx{
		Workspace:     workspace,
		mainContext:   path.Join(viper.GetString("HOME"), ".kube", "config"),
		wardContext:   path.Join(workspace, ".config"),
		backupContext: path.Join(workspace, ".config") + ".bk",
		metaPath:      path.Join(workspace, "metadata"),
		viper:         viper.New(),
		metadata:      metadata,
	}
	err := k.LoadMetadata()
	if err != nil {
		return nil
	}
	return k
}

func (c *KubeCtx) LoadMetadata() error {
	if !util.IsFileExists(c.metaPath) {
		return nil
	}
	data, err := os.ReadFile(c.metaPath)
	if err != nil {
		return err
	}
	err = toml.Unmarshal(data, c.metadata)
	if err != nil {
		return err
	}
	return nil
}

func (c *KubeCtx) StoreMetadata() error {
	data, err := toml.Marshal(*c.metadata)
	if err != nil {
		return err
	}
	f, err := os.Create(c.metaPath)
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write(data)
	return nil
}

func (c *KubeCtx) AddMetadata(filePath string) error {
	host := path.Base(filePath)
	kubeConfig, err := LoadKubeConfig(filePath)
	if err != nil {
		return err
	}
	user := kubeConfig.Users[0].Name
	cluster := kubeConfig.Clusters[0].Name
	name := kubeConfig.CurrentContext
	name = strings.TrimPrefix(name, user+"@")
	namespace := "default"
	if kubeConfig.Contexts[0].Context.Namespace != "" {
		namespace = kubeConfig.Contexts[0].Context.Namespace
	}
	metadata := &ContextMetadata{
		Host:      host,
		Name:      name,
		Namespace: namespace,
		Path:      filePath,
		Cluster:   cluster,
		User:      user,
	}
	util.Println(metadata)
	c.metadata.Contexts[host] = metadata
	err = c.StoreMetadata()
	if err != nil {
		return err
	}
	return nil
}

func (c *KubeCtx) DeleteMetadata(key string) error {
	delete(c.metadata.Contexts, key)
	return nil
}

func (c *KubeCtx) getKeyByTarget(target string) (string, error) {
	candidate := make([]string, 0)
	for k, v := range c.metadata.Contexts {
		if k == target {
			candidate = append(candidate, k)
		}
		if v.Name == target {
			candidate = append(candidate, k)
		}
	}
	if len(candidate) == 1 {
		return candidate[0], nil
	}
	if len(candidate) > 1 {
		data := make([]string, 0, len(candidate))
		for _, key := range candidate {
			line := []string{key, c.metadata.Contexts[key].Name, c.metadata.Contexts[key].Namespace}
			data = append(data, strings.Join(line, " | "))
		}
		i, err := ui.Select(data)
		if err != nil {
			return "", err
		}
		return candidate[i], nil
	}
	return "", errors.New("Not Found")
}
