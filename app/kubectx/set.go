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
	key, err := c.getKeyByTarget(opts.Target)
	if err != nil {
		return err
	}
	filePath := c.metadata.Contexts[key].Path
	kubeConfig, err := LoadKubeConfig(filePath)
	if err != nil {
		return err
	}
	if opts.Name != "" {
		kubeConfig.CurrentContext = opts.Name
		kubeConfig.Contexts[0].Name = opts.Name
		c.metadata.Contexts[key].Name = opts.Name
	}
	if opts.Namespace != "" {
		kubeConfig.Contexts[0].Context.Namespace = opts.Namespace
		c.metadata.Contexts[key].Namespace = opts.Namespace
	}
	err = StoreKubeConfig(filePath, kubeConfig)
	if err != nil {
		return err
	}

	err = c.StoreMetadata()
	if err != nil {
		return err
	}
	return nil
}
