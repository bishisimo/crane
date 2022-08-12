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

// MetaSpec defines the desired state of MySQLMeta
type MetaSpec struct {
}

type TimeRange struct {
	// StartTime is the start time of backup time in this scope.
	StartTime *metav1.Time `json:"startTime,omitempty"`
	// StopTime is the stop time of backup time in this scope.
	StopTime *metav1.Time `json:"stopTime,omitempty"`
	// StartIndex is the start index of backup info in this scope.
	StartIndex int `json:"startIndex,omitempty"`
	// StopIndex is the stop index of backup info in this scope
	StopIndex int `json:"stopIndex,omitempty"`
}

type BackupInfo struct {
	// Name is the name of Backup CR.
	Name string `json:"name,omitempty"`
	// Version is the mark of continues backup.
	Version int `json:"version,omitempty"`
	// Spec is the spec of Backup CR.
	Spec *BackupSpec `json:"spec,omitempty"`
	// Status is the status of backup CR.
	Status *BackupStatus `json:"status,omitempty"`
}

// MetaStatus defines the observed state of MySQLMeta
type MetaStatus struct {
	// Cluster is the name of managed cluster.
	Cluster string `json:"cluster,omitempty"`
	// Version is the mark of continues backup.
	Version int `json:"version,omitempty"`
	// ValidTimeRanges store the time range of recoverable data.
	ValidTimeRanges []*TimeRange `json:"validTimeRange,omitempty"`
	// BackupInfos are record items of backup metadata.
	BackupInfos []*BackupInfo `json:"backupInfos,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MySQLMeta is the Schema for the meta API
type MySQLMeta struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MetaSpec   `json:"spec,omitempty"`
	Status MetaStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MySQLMetaList contains a list of MySQLMeta
type MySQLMetaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLMeta `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MySQLMeta{}, &MySQLMetaList{})
}
