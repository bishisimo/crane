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
	"crane/app/mgr/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RestoreSpec defines the desired state of Restore
type RestoreSpec struct {
	// Backup is the name of Backup CR you want to restore.
	Backup string `json:"backupName,omitempty"`
	// TimePoint is the deadline of data you want to restore.
	TimePoint *metav1.Time `json:"timePoint,omitempty"`
	// Cluster is the name of cluster you want to restore to.
	Cluster string `json:"cluster,omitempty"`
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

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Restore is the Schema for the restores API
type Restore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RestoreSpec   `json:"spec,omitempty"`
	Status RestoreStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RestoreList contains a list of Restore
type RestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Restore `json:"items"`
}

func init() {
	v1alpha1.SchemeBuilder.Register(&Restore{}, &RestoreList{})
}
