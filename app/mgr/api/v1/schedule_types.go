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

type Plan struct {
	// Enable is enable or disable schedule backup.
	Enable bool `json:"enable,omitempty"`
	// KeepCopies is the number of the latest backup data hold.
	KeepCopies int `json:"keepCopies,omitempty"`
	// Cron is the scheduled backup time expressed using cron.
	Cron string `json:"cron,omitempty"`
}

// ScheduleSpec defines the desired state of Schedule
type ScheduleSpec struct {
	// Cluster is the name of managed cluster.
	Cluster string `json:"cluster,omitempty"`
	// Enable is global enable or disable schedule backup.
	Enable bool `json:"enable,omitempty"`
	// Full is the full backup schedule.
	Full *Plan `json:"full,omitempty"`
	// Increment is the increment backup schedule.
	Increment *Plan `json:"increment,omitempty"`
	// Storage is the place where the backup data store to.
	Storage *Storage `json:"storage,omitempty"`
}

// ScheduleStatus defines the observed state of Schedule
type ScheduleStatus struct {
	// FullLatestTime is the full backup latest run time.
	FullLatestTime *metav1.Time `json:"fullLatestTime,omitempty"`
	// IncrementLatestTime is the increment backup latest run time.
	IncrementLatestTime *metav1.Time `json:"incrementLatestTime,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Schedule is the Schema for the schedules API
type Schedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ScheduleSpec   `json:"spec,omitempty"`
	Status ScheduleStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ScheduleList contains a list of Schedule
type ScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Schedule `json:"items"`
}
