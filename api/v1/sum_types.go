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

// SumSpec defines the desired state of Sum
type SumSpec struct {
	NumberOne int `json:"numberOne,omitempty"`
	NumberTwo int `json:"numberTwo,omitempty"`
}

// SumStatus defines the observed state of Sum
type SumStatus struct {
	Result int `json:"result,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Number One",type="integer",JSONPath=".spec.numberOne",description="Input number one"
//+kubebuilder:printcolumn:name="Number Two",type="integer",JSONPath=".spec.numberTwo",description="Input number two"
//+kubebuilder:printcolumn:name="Result",type="integer",JSONPath=".status.result",description="Sum of two numbers"

// Sum is the Schema for the sums API
type Sum struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SumSpec   `json:"spec,omitempty"`
	Status SumStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SumList contains a list of Sum
type SumList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Sum `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Sum{}, &SumList{})
}
