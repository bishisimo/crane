package cluster

import (
	"context"
	mgrv1a1 "crane/app/mgr/api/v1alpha1"
	"crane/app/mgr/common"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// Reconciler reconciles a Meta object
type Reconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Chan   chan *mgrv1a1.MySQLCluster
	Key    client.ObjectKey
}

func NewMetaReconciler(c chan *mgrv1a1.MySQLCluster, key client.ObjectKey) *Reconciler {
	mgr := common.GetManger()
	return &Reconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
		Chan:   c,
		Key:    key,
	}
}

//+kubebuilder:rbac:groups=mgr.my.domain,resources=metas,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=mgr.my.domain,resources=metas/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=mgr.my.domain,resources=metas/finalizers,verbs=update

func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx).WithValues("namespace", req.Namespace, "name", req.Name)
	if !r.FilterByTypes(req) {
		return ctrl.Result{}, nil
	}
	// prepare
	clusterCr := new(mgrv1a1.MySQLCluster)
	err := r.Client.Get(ctx, req.NamespacedName, clusterCr)
	if apierrors.IsNotFound(err) {
		logger.Info("api-resource loss, do nothing")
		return ctrl.Result{}, nil
	} else if err != nil {
		return ctrl.Result{}, err
	}
	// run
	r.Send(clusterCr)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *Reconciler) SetupWithManager() error {
	mgr := common.GetManger()
	return ctrl.NewControllerManagedBy(mgr).
		For(&mgrv1a1.MySQLCluster{}).
		Complete(r)
}

func (r *Reconciler) FilterByTypes(req ctrl.Request) bool {
	if r.Key.Namespace != "" && req.Namespace != r.Key.Namespace {
		return false
	}
	if r.Key.Name != "" && req.Name != r.Key.Name {
		return false
	}
	return true
}

func (r *Reconciler) Send(clusterCr *mgrv1a1.MySQLCluster) {
	select {
	case r.Chan <- clusterCr:
	default:
	}
}
