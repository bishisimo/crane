package kubectx

import (
	"gopkg.in/yaml.v3"
	"os"
)

type KubeConfig struct {
	ApiVersion     string         `yaml:"apiVersion"`
	Kind           string         `yaml:"kind"`
	CurrentContext string         `yaml:"current-context"`
	Clusters       []*Cluster     `yaml:"clusters"`
	Users          []*User        `yaml:"users"`
	Contexts       []*Context     `yaml:"contexts"`
	Preferences    map[string]any `yaml:"preferences"`
}

type Cluster struct {
	Name    string       `yaml:"name"`
	Cluster *ClusterInfo `yaml:"cluster"`
}

type ClusterInfo struct {
	Server                   string `yaml:"server"`
	CertificateAuthorityData string `yaml:"certificate-authority-data"`
}

type User struct {
	Name string    `yaml:"name"`
	User *UserInfo `yaml:"user"`
}

type UserInfo struct {
	ClientCertificateData string `yaml:"client-certificate-data"`
	ClientKeyData         string `yaml:"client-key-data"`
}

type Context struct {
	Name    string      `yaml:"name"`
	Context ContextInfo `yaml:"context"`
}

type ContextInfo struct {
	Cluster   string `yaml:"cluster"`
	User      string `yaml:"user"`
	Namespace string `yaml:"namespace"`
}

func LoadKubeConfig(filePath string) (*KubeConfig, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	config := new(KubeConfig)
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func StoreKubeConfig(filePath string, kubeConfig *KubeConfig) error {
	data, err := yaml.Marshal(kubeConfig)
	if err != nil {
		return err
	}
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
