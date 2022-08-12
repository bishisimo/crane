package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Resources holds ResourceRequirements for the MySQL Agent & Server Containers
type Resources struct {
	Agent  *corev1.ResourceRequirements `json:"agent,omitempty"`
	Server *corev1.ResourceRequirements `json:"server,omitempty"`
}

// ClusterSpec defines the attributes a user can specify when creating a cluster
type ClusterSpec struct {
	// Version defines the MySQL Docker image version.
	Version string `json:"version,omitempty"`
	// Repository defines the image repository from which to pull the MySQL server image.
	Repository string `json:"repository,omitempty"`
	// ImagePullSecret defines the name of the secret that contains the
	// required credentials for pulling from the specified Repository.
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecret,omitempty"`
	// Members defines the number of MySQL instances in a cluster
	Members int32 `json:"members,omitempty"`
	// BaseServerID defines the base number used to create unique server_id
	// for MySQL instances in the cluster. Valid range 1 to 4294967286.
	// If omitted in the manifest file (or set to 0) defaultBaseServerID
	// value will be used.
	BaseServerID uint32 `json:"baseServerId,omitempty"`
	// MultiMaster defines the mode of the MySQL cluster. If set to true,
	// all instances will be R/W. If false (the default), only a single instance
	// will be R/W and the rest will be R/O.
	MultiMaster bool `json:"multiMaster,omitempty"`
	// Router defines the MySQL Router
	Router MySQLRouter `json:"router,omitempty"`
	// NodeSelector is a selector which must be true for the pod to fit on a node.
	// Selector which must match a node's labels for the pod to be scheduled on that node.
	// More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// If specified, affinity will define the pod's scheduling constraints
	// +optional
	Affinity *corev1.Affinity `json:"affinity,omitempty"`
	// EnableStorage  allow a user to enable mysql storage
	// +optional
	EnableStorage bool `json:"enableStorage,omitempty"`
	// VolumeClaimTemplate allows a user to specify how volumes inside a MySQL cluster
	// +optional
	VolumeClaimTemplate *corev1.PersistentVolumeClaim `json:"volumeClaimTemplate,omitempty"`
	// BackupVolumeClaimTemplate allows a user to specify a volume to temporarily store the
	// data for a backup prior to it being shipped to object storage.
	// +optional
	BackupVolumeClaimTemplate *corev1.PersistentVolumeClaim `json:"backupVolumeClaimTemplate,omitempty"`
	// If defined, we use this secret for configuring the MYSQL_ROOT_PASSWORD
	// If it is not set we generate a secret dynamically
	// +optional
	RootPasswordSecret *corev1.LocalObjectReference `json:"rootPasswordSecret,omitempty"`
	// Config allows a user to specify a custom configuration file for MySQL.
	// +optional
	Config *corev1.LocalObjectReference `json:"config,omitempty"`
	// SSLSecret allows a user to specify custom CA certificate, server certificate
	// and server key for group replication SSL.
	// +optional
	SSLSecret *corev1.LocalObjectReference `json:"sslSecret,omitempty"`

	CertConfig *CertConfig `json:"certConfig,omitempty"`

	// Mysql configuration
	Configuration string `json:"configuration,omitempty"`

	// SecurityContext holds Pod-level security attributes and common Container settings.
	SecurityContext *corev1.PodSecurityContext `json:"securityContext,omitempty"`
	// Tolerations allows specifying a list of tolerations for controlling which
	// set of Nodes a Pod can be scheduled on
	Tolerations *[]corev1.Toleration `json:"tolerations,omitempty"`
	// Resources holds ResourceRequirements for the MySQL Agent & Server Containers
	Resources *Resources `json:"resources,omitempty"`
	// Monitor define MySQL monitor spec
	// +optional
	Monitor *Monitor `json:"monitor,omitempty"`

	PodCIDR string `json:"podCIDR,omitempty"`
}

type CertConfig struct {
	Duration   *metav1.Duration `json:"duration,omitempty"`
	DNSNames   []string         `json:"DNSNames,omitempty"`
	SecretName string           `json:"secretName,omitempty"`
}

// MySQLRouter describes the spec of the MySQL Router.
type MySQLRouter struct {
	// MySQLUser defines the user that connect to MySQL.
	// +optional
	MySQLUser string `json:"mysqlUser,omitempty"`
	// MySQLPasswordSecret defines the secret name of MySQL user.
	// +optional
	MySQLPasswordSecret string `json:"mysqlPasswordSecret,omitempty"`
	// Replicas defines the replicas of the MySQL Router.
	// +optional
	Replicas int32 `json:"replicas,omitempty"`
	// Resources holds ResourceRequirements for the MySQL Agent & Server Containers
	// +optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
	// Tolerations allows specifying a list of tolerations for controlling which
	// set of Nodes a Pod can be scheduled on
	// +optional
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`
	// NodeSelector is a selector which must be true for the pod to fit on a node.
	// Selector which must match a node's labels for the pod to be scheduled on that node.
	// More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// If specified, affinity will define the pod's scheduling constraints
	// +optional
	Affinity *corev1.Affinity `json:"affinity,omitempty"`

	SecurityContext *corev1.PodSecurityContext `json:"securityContext,omitempty"`
}

// Monitor define MySQL monitor spec
type Monitor struct {
	// Enable means enable monitor
	// +optional
	Enable bool `json:"enable,omitempty"`
	// Exporter define MySQL exporter spec
	// +optional
	Exporter MySQLExporter `json:"exporter,omitempty"`
}

type MySQLExporter struct {
	// Image means the image of the exporter
	// +optional
	Image string `json:"image,omitempty"`
	// Resources holds ResourceRequirements for the MySQL Agent & Server Containers
	// +optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}

// ClusterConditionType represents a valid condition of a Cluster.
type ClusterConditionType string

// ClusterCondition describes the observed state of a Cluster at a certain point.
type ClusterCondition struct {
	Type   ClusterConditionType   `json:"type,omitempty"`
	Status corev1.ConditionStatus `json:"status,omitempty"`
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// +optional
	Reason string `json:"reason,omitempty"`
	// +optional
	Message string `json:"message,omitempty"`
}

// ClusterStatus defines the current status of a MySQL cluster
// propagating useful information back to the cluster admin
type ClusterStatus struct {
	// +optional
	Conditions []ClusterCondition `json:"conditions,omitempty"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLCluster represents a cluster spec and associated metadata
type MySQLCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLClusterList is a placeholder type for a list of MySQL clusters
type MySQLClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MySQLCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MySQLCluster{}, &MySQLClusterList{})
}
