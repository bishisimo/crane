package kubectx

import (
	"crane/pkg/errorx"
	"crane/pkg/sshx"
	"crane/pkg/ui"
	"crane/util"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

type AddOptions struct {
	Host       string
	Port       string
	Username   string
	Password   string
	PrivateKey string
	SourcePath string
	AcpUrl     string
	Cluster    string
	Name       string
	Namespace  string
	token      string
}

func (o *AddOptions) ParserUri(uri string) error {
	prefix := "ssh://"
	if !strings.HasPrefix(uri, prefix) {
		uri = prefix + uri
	}
	parse, err := url.Parse(uri)
	if err != nil {
		return err
	}
	if parse.Host != "" {
		o.Host = parse.Host
	}
	if parse.Port() != "" {
		o.Port = parse.Port()
	}
	if parse.User.Username() != "" {
		o.Username = parse.User.Username()
	}
	password, _ := parse.User.Password()
	if password != "" {
		o.Password = password
	}
	if parse.Path != "" {
		o.SourcePath = parse.Path
	}
	return nil
}

func (c *KubeCtx) Add(opts *AddOptions) error {
	if !util.IsFileExists(c.Workspace) {
		os.MkdirAll(c.Workspace, os.ModePerm)
	}
	if opts.AcpUrl != "" {
		return c.addByAcp(opts)
	}
	err := c.addBySftp(opts)
	if err != nil {
		return err
	}
	return nil
}

func (c *KubeCtx) addBySftp(opts *AddOptions) error {
	if opts.PrivateKey != "" && !util.IsFileExists(opts.PrivateKey) {
		opts.PrivateKey = path.Join(viper.GetString("sys.sshKeyDir"), opts.PrivateKey)
	}
	if opts.Password == "" && opts.PrivateKey == "" {
		password := ui.Input("input password:")
		opts.Password = password
	}
	addr := fmt.Sprintf("%v:%v", opts.Host, opts.Port)
	sftpOpts := &sshx.Options{
		Addr:       addr,
		Username:   opts.Username,
		Password:   opts.Password,
		PrivateKey: opts.PrivateKey,
	}

	client, err := sshx.NewSftpClient(sftpOpts)
	if err != nil {
		return err
	}
	if opts.SourcePath == "" {
		home, err := client.GetHome()
		if err != nil {
			return err
		}
		opts.SourcePath = path.Join(home, ".kube/config")
		log.Info().Str("path", opts.SourcePath).Msg("use default path")
	}
	targetPath := path.Join(c.Workspace, opts.Host)
	err = client.Download(opts.SourcePath, targetPath)
	if err != nil {
		return err
	}
	err = c.AddMetadata(targetPath)
	if err != nil {
		return err
	}
	if opts.Name != "" || opts.Namespace != "" {
		setOpts := &SetOptions{
			Target:    opts.Host,
			Name:      opts.Name,
			Namespace: opts.Namespace,
		}
		err = c.Set(setOpts)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *KubeCtx) addByAcp(opts *AddOptions) error {
	if !strings.HasPrefix(opts.AcpUrl, "https://") {
		opts.AcpUrl = `https://` + opts.AcpUrl
	}
	opts.token = ui.Input("input token:")
	err := c.ensureCluster(opts)
	if err != nil {
		return err
	}
	configPath := path.Join("/auth/v1/clusters", opts.Cluster, "kubeconfig")
	configUrl := opts.AcpUrl + configPath
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().
		SetAuthToken(opts.token).
		Get(configUrl)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New("fetch config url fail")
	}
	kubeConfig := new(KubeConfig)
	err = json.Unmarshal(resp.Body(), kubeConfig)
	if err != nil {
		return err
	}
	var targetCluster *Cluster
	for _, cluster := range kubeConfig.Clusters {
		if strings.Contains(cluster.Cluster.Server, ":") {
			targetCluster = cluster
			data := targetCluster.Cluster.CertificateAuthorityData
			desc := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
			base64.StdEncoding.Encode(desc, []byte(data))
			if err != nil {
				return err
			}
			targetCluster.Cluster.CertificateAuthorityData = string(desc)
		}
	}
	if targetCluster == nil {
		return errorx.NotFound
	}
	var targetContext *Context
	for _, ctx := range kubeConfig.Contexts {
		if ctx.Context.Cluster == targetCluster.Name {
			targetContext = ctx
		}
	}
	if targetContext == nil {
		return errorx.NotFound
	}
	kubeConfig.Clusters = []*Cluster{targetCluster}
	kubeConfig.Contexts = []*Context{targetContext}
	kubeConfig.CurrentContext = targetContext.Name
	yData, err := yaml.Marshal(kubeConfig)
	if err != nil {
		return err
	}
	server := kubeConfig.Clusters[0].Cluster.Server
	host := strings.TrimPrefix(server, "https://")
	host = strings.Split(host, ":")[0]
	host = strings.Split(host, "/")[0]
	targetPath := path.Join(c.Workspace, host)
	f, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(yData)
	if err != nil {
		return err
	}
	err = c.AddMetadata(targetPath)
	if err != nil {
		return err
	}
	if opts.Name != "" || opts.Namespace != "" {
		setOpts := &SetOptions{
			Target:    opts.Host,
			Name:      opts.Name,
			Namespace: opts.Namespace,
		}
		err = c.Set(setOpts)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *KubeCtx) ensureCluster(opts *AddOptions) error {
	if opts.Cluster != "" {
		return nil
	}
	clustersUrl := opts.AcpUrl + "/v1/query/list"
	body := map[string]any{"query": map[string]any{
		"apiVersion": "cluster.alauda.io/v1alpha1",
		"kind":       "ClusterModule",
		"limit":      -1,
	}}
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().
		SetAuthToken(opts.token).
		SetBody(body).
		Post(clustersUrl)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New("fetch cluster url fail")
	}
	ci := new(ClusterItems)
	err = json.Unmarshal(resp.Body(), ci)
	if err != nil {
		return err
	}
	clusterList := make([]string, 0, len(ci.Items))
	for _, item := range ci.Items {
		clusterList = append(clusterList, item["metadata"].(map[string]any)["name"].(string))
	}
	i, err := ui.Select(clusterList)
	if err != nil {
		return err
	}
	opts.Cluster = clusterList[i]
	return nil
}

type ClusterItems struct {
	Items []map[string]any `json:"items"`
}
