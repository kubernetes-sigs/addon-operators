/*
Copyright 2018 The Kubernetes Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	addonv1alpha1 "sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/apis/v1alpha1"
)

// DashboardSpec defines the desired state of Dashboard
type DashboardSpec struct {
	addonv1alpha1.CommonSpec
}

// DashboardStatus defines the observed state of Dashboard
type DashboardStatus struct {
	addonv1alpha1.CommonStatus
}

var _ addonv1alpha1.CommonObject = &Dashboard{}

func (c *Dashboard) ComponentName() string {
	return "dashboard"
}

func (c *Dashboard) CommonSpec() addonv1alpha1.CommonSpec {
	return c.Spec.CommonSpec
}

func (c *Dashboard) GetCommonStatus() addonv1alpha1.CommonStatus {
	return c.Status.CommonStatus
}

func (c *Dashboard) SetCommonStatus(s addonv1alpha1.CommonStatus) {
	c.Status.CommonStatus = s
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Dashboard is the Schema for the dashboards API
// +k8s:openapi-gen=true
type Dashboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DashboardSpec   `json:"spec,omitempty"`
	Status DashboardStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DashboardList contains a list of Dashboard
type DashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dashboard `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Dashboard{}, &DashboardList{})
}
