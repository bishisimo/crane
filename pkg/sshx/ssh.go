package sshx

import (
	"bytes"
	"golang.org/x/crypto/ssh"
	"os"
	"strings"
)

type Options struct {
	Addr       string
	Username   string
	Password   string
	PrivateKey string
}

type SshClient struct {
	*ssh.Client
}

func NewSShClient(opts *Options) (*SshClient, error) {
	authMethod := []ssh.AuthMethod{ssh.Password(opts.Password)}
	if opts.PrivateKey != "" {
		key, err := os.ReadFile(opts.PrivateKey)
		if err != nil {
			return nil, err
		}
		// Create the Signer for this private key.
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return nil, err
		}
		authMethod = append(authMethod, ssh.PublicKeys(signer))
	}
	config := &ssh.ClientConfig{
		User:            opts.Username,
		Auth:            authMethod,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", opts.Addr, config)
	if err != nil {
		return nil, err
	}

	return &SshClient{client}, nil
}

func (s *SshClient) SftpClient() (*SftpClient, error) {
	return NewSftpClientFromSshClient(s)
}

func (s *SshClient) Exec(cmd string) (*bytes.Buffer, error) {
	session, err := s.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	stdout := new(bytes.Buffer)
	session.Stdout = stdout

	err = session.Run(cmd)
	if err != nil {
		return nil, err
	}
	return stdout, nil
}

func (s *SshClient) GetHome() (string, error) {
	cmd := "echo $HOME"
	out, err := s.Exec(cmd)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}
