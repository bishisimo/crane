package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Plan struct {
	// Enable is enable or disable schedule backup.
	Enable bool `json:"enable,omitempty"`
	// KeepCopies is the number of the latest backup data hold.
	KeepCopies int `json:"keepCopies,omitempty"`
	// Cron is the scheduled backup time expressed using cron.
	Cron string `json:"cron,omitempty"`
}

type ScheduleSpec struct {
	// Cluster is the name of managed cluster.
	Cluster *corev1.LocalObjectReference `json:"cluster,omitempty"`
	// Enable is global enable or disable schedule backup.
	Enable bool `json:"enable,omitempty"`
	// Full is the full backup schedule.
	Full *Plan `json:"full,omitempty"`
	// Increment is the increment backup schedule.
	Increment *Plan `json:"increment,omitempty"`
	// StorageProvider is the place where the backup data store to.
	StorageProvider *StorageProvider `json:"storageProvider,omitempty"`
}

type ScheduleStatus struct {
	// FullLatestTime is the full backup latest run time.
	FullLatestTime *metav1.Time `json:"fullLatestTime,omitempty"`
	// IncrementLatestTime is the increment backup latest run time.
	IncrementLatestTime *metav1.Time `json:"incrementLatestTime,omitempty"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLSchedule is a backup schedule for a Cluster.
type MySQLSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   ScheduleSpec   `json:"spec,omitempty"`
	Status ScheduleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLScheduleList is a list of MySQLSchedule.
type MySQLScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MySQLSchedule `json:"items"`
}
