package sshx

import (
	"github.com/pkg/sftp"
	"io"
	"os"
)

type SftpClient struct {
	*sftp.Client
	*SshClient
}

func NewSftpClient(opts *Options) (*SftpClient, error) {
	client, err := NewSShClient(opts)
	if err != nil {
		return nil, err
	}
	return NewSftpClientFromSshClient(client)
}

func NewSftpClientFromSshClient(sshClient *SshClient) (*SftpClient, error) {
	client, err := sftp.NewClient(sshClient.Client)
	if err != nil {
		return nil, err
	}
	return &SftpClient{
		Client:    client,
		SshClient: sshClient,
	}, nil
}

func (c *SftpClient) Upload(sourcePath, targetPath string) error {
	source, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer source.Close()

	stat, err := source.Stat()
	if err != nil {
		return err
	}
	target, err := c.Create(targetPath)
	if err != nil {
		return err
	}
	defer target.Close()
	target.Chmod(stat.Mode())

	_, err = io.Copy(target, source)
	if err != nil {
		return err
	}
	return nil
}

func (c *SftpClient) Download(sourcePath, targetPath string) error {
	source, err := c.Open(sourcePath)
	if err != nil {
		return err
	}
	defer source.Close()

	stat, err := source.Stat()
	if err != nil {
		return err
	}
	target, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer target.Close()
	target.Chmod(stat.Mode())

	_, err = io.Copy(target, source)
	if err != nil {
		return err
	}
	return nil
}
