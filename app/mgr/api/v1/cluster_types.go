/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type CertConfig struct {
	Duration   *metav1.Duration `json:"duration,omitempty"`
	DNSNames   []string         `json:"DNSNames,omitempty"`
	SecretName string           `json:"secretName,omitempty"`
}

type Mysql struct {
	// Version defines the MySQL Docker image version.
	Version string `json:"version,omitempty"`
	// Repository defines the image repository from which to pull the MySQL server image.
	Repository string `json:"repository,omitempty"`
	// ImagePullSecret defines the name of the secret that contains the required credentials for pulling from the specified Repository.
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecret,omitempty"`
	// Members defines the number of MySQL instances in a cluster
	Members int32 `json:"members,omitempty"`
	// BaseServerID defines the base number used to create unique server_id for MySQL instances in the cluster.
	// Valid range 1 to 4294967286. If omitted in the manifest file (or set to 0) defaultBaseServerID value will be used.
	BaseServerID uint32 `json:"baseServerId,omitempty"`
	// MultiMaster defines the mode of the MySQL cluster.
	// If set to true, all instances will be R/W. If false (the default), only a single instance will be R/W and the rest will be R/O.
	MultiMaster bool `json:"multiMaster,omitempty"`
	// NodeSelector is a selector which must be true for the pod to fit on a node.
	// Selector which must match a node's labels for the pod to be scheduled on that node.
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// If specified, affinity will define the pod's scheduling constraints
	Affinity *corev1.Affinity `json:"affinity,omitempty"`
	// EnableStorage  allow a user to enable mysql storage
	EnableStorage bool `json:"enableStorage,omitempty"`
	// VolumeClaimTemplate allows a user to specify how volumes inside a MySQL cluster
	VolumeClaimTemplate *corev1.PersistentVolumeClaim `json:"volumeClaimTempla te,omitempty"`
	// BackupVolumeClaimTemplate allows a user to specify a volume to temporarily store the data for a backup prior to it being shipped to object storage.
	BackupVolumeClaimTemplate *corev1.PersistentVolumeClaim `json:"backupVolumeClaimTemplate,omitempty"`
	// If defined, we use this secret for configuring the MYSQL_ROOT_PASSWORD. If it is not set we generate a secret dynamically
	RootPasswordSecret *corev1.LocalObjectReference `json:"rootPasswordSecret,omitempty"`
	// Config allows a user to specify a custom configuration file for MySQL.
	Config *corev1.LocalObjectReference `json:"config,omitempty"`
	// SSLSecret allows a user to specify custom CA certificate, server certificate and server key for group replication SSL.
	SSLSecret *corev1.LocalObjectReference `json:"sslSecret,omitempty"`
	// CertConfig is the ssl cert config
	CertConfig *CertConfig `json:"certConfig,omitempty"`
	// Configuration is the Mysql configuration
	Configuration string `json:"configuration,omitempty"`
	// SecurityContext holds Pod-level security attributes and common Container settings.
	SecurityContext *corev1.PodSecurityContext `json:"securityContext,omitempty"`
	// Tolerations allows specifying a list of tolerations for controlling which set of Nodes a Pod can be scheduled on
	Tolerations *[]corev1.Toleration `json:"tolerations,omitempty"`
	// Resources holds ResourceRequirements for the mysql Containers
	MySQLResources *corev1.ResourceRequirements `json:"mySqlResources,omitempty"`
	// Resources holds ResourceRequirements for the agent Containers
	AgentResources *corev1.ResourceRequirements `json:"agentResources,omitempty"`
	// PodCIDR is the limit of ip access
	PodCIDR string `json:"podCIDR,omitempty"`
}

type Router struct {
	// MySQLUser defines the user that connect to MySQL.
	MySQLUser string `json:"mysqlUser,omitempty"`
	// MySQLPasswordSecret defines the secret name of MySQL user.
	MySQLPasswordSecret string `json:"mysqlPasswordSecret,omitempty"`
	// Replicas defines the replicas of the MySQL Router.
	Replicas int32 `json:"replicas,omitempty"`
	// Resources holds ResourceRequirements for the MySQL Agent & Server Containers
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
	// Tolerations allows specifying a list of tolerations for controlling which set of Nodes a Pod can be scheduled on
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`
	// NodeSelector is a selector which must be true for the pod to fit on a node.
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// Affinity if specified, affinity will define the pod's scheduling constraints
	Affinity *corev1.Affinity `json:"affinity,omitempty"`
	// SecurityContext holds Pod-level security attributes and common Container settings.
	SecurityContext *corev1.PodSecurityContext `json:"securityContext,omitempty"`
}

type Exporter struct {
	// Enable means enable monitor
	Enable bool `json:"enable,omitempty"`
	// Image means the image of the exporter
	Image string `json:"image,omitempty"`
	// Resources holds ResourceRequirements for the MySQL Agent & Server Containers
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	// Mysql defines the MGR Mysql
	Mysql *Mysql `json:"mysql,omitempty"`
	// Router defines the MGR Router
	Router *Router `json:"router,omitempty"`
	// Exporter define MGR Exporter
	Exporter *Exporter `json:"monitor,omitempty"`
}

// ClusterStatus defines the observed state of Cluster
type ClusterStatus struct {
	// Condition describes the observed state at a certain point.
	Conditions []*corev1.PodCondition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Cluster is the Schema for the clusters API
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterList contains a list of Cluster
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
