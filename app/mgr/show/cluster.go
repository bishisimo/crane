package show

import (
	"context"
	v1a1 "crane/app/mgr/api/v1alpha1"
	"crane/app/mgr/common"
	"crane/app/mgr/controller/cluster"
	"crane/pkg/ui"
	"k8s.io/apimachinery/pkg/types"
)

type ClusterShowOptions struct {
	*BaseShowOptions
}

type ClusterShow struct {
	Options        *ClusterShowOptions
	RawChan        chan *v1a1.MySQLCluster
	loadingContext context.Context
	loadingCancel  context.CancelFunc
}

func NewClusterShow(options *ClusterShowOptions) *ClusterShow {
	c := make(chan *v1a1.MySQLCluster, 1)
	ctx, cancel := context.WithCancel(context.Background())
	return &ClusterShow{
		Options:        options,
		RawChan:        c,
		loadingContext: ctx,
		loadingCancel:  cancel,
	}
}

func (s *ClusterShow) Show() error {
	err := s.loading(true)
	if err != nil {
		return err
	}
	return nil
}

func (s ClusterShow) loading(show bool) error {
	if show {
		loading := ui.NewLoading(s.loadingContext)
		go func() {
			loading.Show()
		}()
	}
	key := types.NamespacedName{
		Namespace: s.Options.Namespace,
		Name:      s.Options.Name,
	}
	reconciler := cluster.NewMetaReconciler(s.RawChan, key)
	err := reconciler.SetupWithManager()
	if err != nil {
		return err
	}
	common.AsyncStartMgr()
	return nil
}
