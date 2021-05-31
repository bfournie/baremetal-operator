/*


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
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HostFirmwareSettingsSpec defines the desired state of HostFirmwareSettings
type HostFirmwareSettingsSpec struct {
	Settings map[string]string `json:"settings" required:"true"`
}

type SchemaReference struct {
	// `namespace` is the namespace of the where the schema is stored.
	Namespace string `json:"namespace"`
	// `name` is the reference to the schema.
	Name string `json:"name"`
}

// HostFirmwareSettingsStatus defines the observed state of HostFirmwareSettings
type HostFirmwareSettingsStatus struct {
	// FirmwareSchema describes each firmware Setting
	FirmwareSchema SchemaReference `json:"schema,omitempty"`

	// Settings are the firmware settings stored as name/value pairs
	// +patchStrategy=merge
	Settings map[string]string `json:"settings" required:"true"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// HostFirmwareSettings is the Schema for the hostfirmwaresettings API
type HostFirmwareSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HostFirmwareSettingsSpec   `json:"spec,omitempty"`
	Status HostFirmwareSettingsStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HostFirmwareSettingsList contains a list of HostFirmwareSettings
type HostFirmwareSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostFirmwareSettings `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HostFirmwareSettings{}, &HostFirmwareSettingsList{})
}
