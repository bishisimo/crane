// Describe:
package remoteexec

import (
	"bytes"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

type RemoteExec struct {
	Config *rest.Config
	Client rest.Interface
	*rest.Request
	Container string
}

// NewRemoteExec
//  @Description: 实例化一个远程执行器
//  @param Config
//  @param namespace
//  @param podName
//  @return *RemoteExec
//  @return error
func NewRemoteExec(config *rest.Config, namespace string, resource string, name string, container string) (*RemoteExec, error) {
	client, err := rest.RESTClientFor(config)
	if err != nil {
		return nil, err
	}
	req := client.Post().
		Resource(resource).
		Namespace(namespace).
		Name(name).
		SubResource("exec")
	exec := &RemoteExec{
		Config:    config,
		Client:    client,
		Request:   req,
		Container: container,
	}
	return exec, nil
}

// Exec
//  @Description: 在指定的远程容器执行命令
//  @receiver e
//  @param cmd
//  @return stdout
//  @return stderr
//  @return err
func (e RemoteExec) Exec(cmd []string) (stdout *bytes.Buffer, stderr *bytes.Buffer, err error) {
	e.VersionedParams(&corev1.PodExecOptions{
		Container: e.Container,
		Command:   cmd,
		Stdout:    true,
		Stderr:    true,
	}, scheme.ParameterCodec)
	exec, err := remotecommand.NewSPDYExecutor(e.Config, "POST", e.URL())
	if err != nil {
		return
	}
	stdout = new(bytes.Buffer)
	stderr = new(bytes.Buffer)
	err = exec.Stream(remotecommand.StreamOptions{
		Stdout: stdout,
		Stderr: stderr,
		Tty:    true,
	})
	return
}
