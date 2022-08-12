package common

import (
	"crane/app/mgr/api/v1alpha1"
	"github.com/rs/zerolog/log"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"sync"
)

var (
	mgr         ctrl.Manager
	scheme      = runtime.NewScheme()
	managerOnce sync.Once
	asyncOnce   sync.Once
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	//utilruntime.Must(v1.AddToScheme(scheme))
}

func GetManger() ctrl.Manager {
	managerOnce.Do(func() {
		var err error
		mgr, err = ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
			MetricsBindAddress: ":18080",
			Scheme:             scheme,
			Port:               9443,
			LeaderElectionID:   "a8881db0.v1alpha1.mysql.middleware.alauda.io",
		})
		if err != nil {
			log.Fatal().Err(err).Msg("create manage fail")
		}
	})
	return mgr
}

func SyncStartMgr() {
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		log.Err(err).Msg("Start signal handler")
		os.Exit(1)
	}
}

func AsyncStartMgr() {
	asyncOnce.Do(func() {
		go func() {
			log.Info().Msg("async run mgr ...")
			SyncStartMgr()
		}()
	})
	log.Info().Msg("mgr is running ...")
}
