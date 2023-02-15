package kubectx

import (
	"gopkg.in/yaml.v3"
	"os"
)

type KubeConfig struct {
	ApiVersion     string         `json:"apiVersion" yaml:"apiVersion"`
	Kind           string         `json:"kind" yaml:"kind"`
	CurrentContext string         `json:"current-context" yaml:"current-context"`
	Clusters       []*Cluster     `json:"clusters" yaml:"clusters"`
	Users          []*User        `json:"users" yaml:"users"`
	Contexts       []*Context     `json:"contexts" yaml:"contexts"`
	Preferences    map[string]any `json:"preferences" yaml:"preferences"`
}

type Cluster struct {
	Name    string       `json:"name" yaml:"name"`
	Cluster *ClusterInfo `json:"cluster" yaml:"cluster"`
}

type ClusterInfo struct {
	Server                   string `json:"server" yaml:"server"`
	CertificateAuthorityData string `json:"certificate-authority-data" yaml:"certificate-authority-data"`
}

type User struct {
	Name string    `json:"name" yaml:"name"`
	User *UserInfo `json:"user" yaml:"user"`
}

type UserInfo struct {
	ClientCertificateData string `json:"client-certificate-data" yaml:"client-certificate-data"`
	ClientKeyData         string `json:"client-key-data" yaml:"client-key-data"`
	Token                 string `json:"token" yaml:"token"`
}

type Context struct {
	Name    string      `json:"name" yaml:"name"`
	Context ContextInfo `json:"context" yaml:"context"`
}

type ContextInfo struct {
	Cluster   string `json:"cluster" yaml:"cluster"`
	User      string `json:"user" yaml:"user"`
	Namespace string `json:"namespace" yaml:"namespace"`
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
