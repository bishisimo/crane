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

// RestoreSpec defines the desired state of Restore
type RestoreSpec struct {
	// Source is the cluster where the data comes form
	Source *corev1.LocalObjectReference `json:"source,omitempty"`
	// Backup is the name of Backup CR you want to restore.
	Backup *corev1.LocalObjectReference `json:"backup,omitempty"`
	// TimePoint is the deadline of data you want to restore.
	TimePoint *metav1.Time `json:"timePoint,omitempty"`
	// Cluster is the name of cluster you want to restore to.
	Cluster *corev1.LocalObjectReference `json:"cluster,omitempty"`
}

type RestoreInfo struct {
	// BackupInfo is the backup information for restore
	BackupInfo *BackupInfo `json:"backupInfo,omitempty"`
	// State is the state for restore
	State string `json:"state,omitempty"`
	// Message is the message of run result of restore
	Message string `json:"message,omitempty"`
}

// RestoreStatus defines the observed state of Restore
type RestoreStatus struct {
	// AllocatedPod is the name of pod allocate to.
	AllocatedPod string `json:"allocatedPod,omitempty"`
	// Progress is the current restore schedule in running.
	Progress int `json:"progress,omitempty"`
	// RestoreSchedule store this restore schedule to a list.
	RestoreSchedule []*RestoreInfo `json:"restoreSchedule,omitempty"`
	// State is the overall state for restore
	State string `json:"state,omitempty"`
	// Message is the message of run result of restore
	Message string `json:"message,omitempty"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLRestore is a MySQL Operator resource that represents the restoration of
// backup of a MySQL cluster.
type MySQLRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   RestoreSpec   `json:"spec,omitempty"`
	Status RestoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLRestoreList is a list of Restores.
type MySQLRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MySQLRestore `json:"items"`
}
