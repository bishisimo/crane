// Copyright 2018 Oracle and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type S3 struct {
	// CredentialsSecret is a reference to the Secret containing the credentials authenticating with the S3 compatible storage service.
	CredentialsSecret *corev1.LocalObjectReference `json:"credentialsSecret,omitempty"`
	// Region in which the S3 compatible bucket is located.
	Region string `json:"region,omitempty"`
	// Endpoint (hostname only or fully qualified URI) of S3 compatible storage service.
	Endpoint string `json:"endpoint,omitempty"`
	// Bucket in which to store the Backup.
	Bucket string `json:"bucket,omitempty"`
}

type StorageProvider struct {
	// S3 means s3 compatible object storage
	S3 *S3 `json:"s3,omitempty"`
}

type BackupWay string

const (
	WayOfScheduled  = "scheduled"
	WayOfManual     = "manual"
	WayOfDeprecated = "deprecated"
	WayOfCompatible = "compatible"
)

type BackupType string

const (
	TypeOfBackupFull      = BackupType("full")
	TypeOfBackupIncrement = BackupType("increment")
)

type BackupExecutor struct {
	// UseBinlog 是否开启增量备份
	UseBinlog bool `json:"useBinlog,omitempty"`
}

// BackupSpec defines the desired state of Backup
type BackupSpec struct {
	// Way means schedule or manual backup.
	Way BackupWay `json:"way,omitempty"`
	// Type means full or increment backup.
	Type BackupType `json:"type,omitempty"`
	// StorageProvider is the storage information of backup for.
	StorageProvider *StorageProvider `json:"storageProvider,omitempty"`
	// Cluster is the name of the data backup source.
	Cluster *corev1.LocalObjectReference `json:"cluster,omitempty"`
	// Executor is the configuration of the tool that will produce the backup.
	Executor BackupExecutor `json:"executor,omitempty"`
}

// BackupOutcome describes the location of a Backup
type BackupOutcome struct {
	// Location is the Object StorageProvider network location of the Backup.
	Location string `json:"location,omitempty"`
	// State 返回状态：200-成功；204,205-部分成功； 500-失败
	State int `json:"state,omitempty"`
	// Error 返回错误信息
	Error string `json:"error,omitempty"`
}

// BackupStatus defines the observed state of Backup
type BackupStatus struct {
	// AllocatedPod is the name of pod allocate to.
	AllocatedPod string `json:"allocatedPod,omitempty"`
	// PreviousGtid is the gtid of previous backup.
	PreviousGtid string `json:"previousGtid,omitempty"`
	// RunTime is the start time of backup runs.
	RunTime *metav1.Time `json:"runTime,omitempty"`
	// Path is the path of the file stored in storage.
	Path string `json:"path,omitempty"`
	// DataStartTime is the start time of increment backup data.
	DataStartTime *metav1.Time `json:"dataStartTime,omitempty"`
	// DataStopTime is the stop time of increment backup or full backup data.
	DataStopTime *metav1.Time `json:"dataStopTime,omitempty"`
	// GtidStart is the start gtid of backup data
	GtidStart string `json:"gtidStart,omitempty"`
	// GtidStop is the stop gtid of backup data
	GtidStop string `json:"gtidStop,omitempty"`
	// State is the backup state and enumeration of success, fail or running.
	State string `json:"state,omitempty"`
	// Message is the message of run result of backup
	Message string `json:"message,omitempty"`
	// TimeStarted is the time at which the backup was started.
	TimeStarted *metav1.Time `json:"timeStarted,omitempty"`
	// TimeCompleted is the time at which the backup completed.
	TimeCompleted *metav1.Time `json:"timeCompleted,omitempty"`
	// Outcome holds the results of a successful backup.
	Outcome BackupOutcome `json:"outcome,omitempty"`
	// Conditions contains details for the current condition of this pod.
	Conditions []corev1.PodCondition `json:"conditions,omitempty"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLBackup is a backup of a Cluster.
type MySQLBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   BackupSpec   `json:"spec,omitempty"`
	Status BackupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLBackupList is a list of Backups.
type MySQLBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []MySQLBackup `json:"items"`
}
