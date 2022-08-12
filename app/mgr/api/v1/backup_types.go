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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type S3 struct {
	// Secret is the secret which store the information for s3
	Secret string `json:"secret,omitempty"`
	// BaseDir is the base dir of the file stored in bucket
	BaseDir string `json:"baseDir,omitempty"`
}

type Storage struct {
	// S3 means s3 compatible object storage
	S3 *S3 `json:"s3,omitempty"`
}

type BackupWay string

const (
	WayOfScheduled = BackupWay("scheduled")
	WayOfManual    = BackupWay("manual")
)

type BackupType string

const (
	TypeOfBackupFull      = BackupType("full")
	TypeOfBackupIncrement = BackupType("increment")
)

// BackupSpec defines the desired state of Backup
type BackupSpec struct {
	// Way means schedule or manual backup.
	Way BackupWay `json:"way,omitempty"`
	// Type means full or increment backup.
	Type BackupType `json:"type,omitempty"`
	// Storage is the storage information of backup for.
	Storage *Storage `json:"storage,omitempty"`
	// Cluster is the name of the data backup source.
	Cluster string `json:"cluster,omitempty"`
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
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Backup is the Schema for the backups API
type Backup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupSpec   `json:"spec,omitempty"`
	Status BackupStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BackupList contains a list of Backup
type BackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Backup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Backup{}, &BackupList{})
}
