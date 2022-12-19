package kubectx

import (
	"crane/util"
	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"path"
	"strings"
)

type Options struct {
}

type Ctx struct {
	Name      string
	User      string
	Namespace string
	Cluster   string
}

type Metadata struct {
	Host string
	Name string
	Path string
	Ctx  *Ctx
}

type KubeCtx struct {
	*Options
	Workspace     string
	workContext   string
	mainContext   string
	backupContext string
	metaPath      string
	viper         *viper.Viper
	metadata      map[string]*Metadata
}

func NewKubeCtx(opts *Options) *KubeCtx {
	workspace := viper.GetString("crane.kubeSpace")
	k := &KubeCtx{
		Options:       opts,
		Workspace:     workspace,
		workContext:   path.Join(viper.GetString("HOME"), ".kube", "config"),
		mainContext:   path.Join(workspace, ".config"),
		backupContext: path.Join(workspace, ".config") + ".bk",
		metaPath:      path.Join(workspace, "metadata"),
		viper:         viper.New(),
		metadata:      make(map[string]*Metadata),
	}
	k.LoadMetadata()
	return k
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
	metadata := &Metadata{
		Host: host,
		Name: name,
		Path: filePath,
		Ctx: &Ctx{
			Name:      kubeConfig.CurrentContext,
			Cluster:   cluster,
			User:      user,
			Namespace: namespace,
		},
	}
	util.Println(metadata)
	c.metadata[host] = metadata
	err = c.StoreMetadata()
	if err != nil {
		return err
	}
	return nil
}

func (c *KubeCtx) DeleteMetadata(key string) error {
	delete(c.metadata, key)
	return nil
}

func (c *KubeCtx) LoadMetadata() error {
	if !util.IsFileExists(c.metaPath) {
		return nil
	}
	data, err := os.ReadFile(c.metaPath)
	if err != nil {
		return err
	}
	err = toml.Unmarshal(data, &c.metadata)
	if err != nil {
		return err
	}
	return nil
}

func (c *KubeCtx) StoreMetadata() error {
	data, err := toml.Marshal(c.metadata)
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

func (c *KubeCtx) getHostByTarget(target string) (string, error) {
	for k, v := range c.metadata {
		if k == target {
			return k, nil
		}
		if v.Name == target {
			return k, nil
		}
	}
	return "", errors.New("Not Found")
}
