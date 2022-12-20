package kubectx

type SetOptions struct {
	Target    string
	Name      string
	Namespace string
}

func (c *KubeCtx) Set(opts *SetOptions) error {
	if opts.Target == "" {
		opts.Target = c.metadata.Current
	}
	host, err := c.getHostByTarget(opts.Target)
	if err != nil {
		return err
	}
	filePath := c.metadata.Contexts[host].Path
	kubeConfig, err := LoadKubeConfig(filePath)
	if err != nil {
		return err
	}
	if opts.Name != "" {
		kubeConfig.CurrentContext = opts.Name
		kubeConfig.Contexts[0].Name = opts.Name
	}
	if opts.Namespace != "" {
		kubeConfig.Contexts[0].Context.Namespace = opts.Namespace
	}
	err = StoreKubeConfig(filePath, kubeConfig)
	if err != nil {
		return err
	}
	if opts.Name != "" {
		c.metadata.Contexts[host].Name = opts.Name
	}
	c.metadata.Contexts[host].Namespace = opts.Namespace
	c.StoreMetadata()
	return nil
}
