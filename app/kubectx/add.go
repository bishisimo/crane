package kubectx

import (
	"crane/pkg/sshx"
	"crane/pkg/ui"
	"crane/util"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
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
	Name       string
	Namespace  string
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
	addr := fmt.Sprintf("%v:%v", opts.Host, opts.Port)
	if opts.PrivateKey != "" && !util.IsFileExists(opts.PrivateKey) {
		opts.PrivateKey = path.Join(viper.GetString("sys.sshKeyDir"), opts.PrivateKey)
	}
	if opts.Password == "" && opts.PrivateKey == "" {
		password := ui.Input("input password:")
		opts.Password = password
	}

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
	if !util.IsFileExists(c.Workspace) {
		os.MkdirAll(c.Workspace, os.ModePerm)
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
