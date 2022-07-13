/*
PDS API

Portworx Data Services API Server

API version: 3.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"encoding/json"
)

// ModelsDeploymentTargetMetadata struct for ModelsDeploymentTargetMetadata
type ModelsDeploymentTargetMetadata struct {
	KubeApiVersion *string `json:"kube_api_version,omitempty"`
	PdsChartVersion *string `json:"pds_chart_version,omitempty"`
	PxServiceNamespace *string `json:"px_service_namespace,omitempty"`
}

// NewModelsDeploymentTargetMetadata instantiates a new ModelsDeploymentTargetMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsDeploymentTargetMetadata() *ModelsDeploymentTargetMetadata {
	this := ModelsDeploymentTargetMetadata{}
	return &this
}

// NewModelsDeploymentTargetMetadataWithDefaults instantiates a new ModelsDeploymentTargetMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsDeploymentTargetMetadataWithDefaults() *ModelsDeploymentTargetMetadata {
	this := ModelsDeploymentTargetMetadata{}
	return &this
}

// GetKubeApiVersion returns the KubeApiVersion field value if set, zero value otherwise.
func (o *ModelsDeploymentTargetMetadata) GetKubeApiVersion() string {
	if o == nil || o.KubeApiVersion == nil {
		var ret string
		return ret
	}
	return *o.KubeApiVersion
}

// GetKubeApiVersionOk returns a tuple with the KubeApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentTargetMetadata) GetKubeApiVersionOk() (*string, bool) {
	if o == nil || o.KubeApiVersion == nil {
		return nil, false
	}
	return o.KubeApiVersion, true
}

// HasKubeApiVersion returns a boolean if a field has been set.
func (o *ModelsDeploymentTargetMetadata) HasKubeApiVersion() bool {
	if o != nil && o.KubeApiVersion != nil {
		return true
	}

	return false
}

// SetKubeApiVersion gets a reference to the given string and assigns it to the KubeApiVersion field.
func (o *ModelsDeploymentTargetMetadata) SetKubeApiVersion(v string) {
	o.KubeApiVersion = &v
}

// GetPdsChartVersion returns the PdsChartVersion field value if set, zero value otherwise.
func (o *ModelsDeploymentTargetMetadata) GetPdsChartVersion() string {
	if o == nil || o.PdsChartVersion == nil {
		var ret string
		return ret
	}
	return *o.PdsChartVersion
}

// GetPdsChartVersionOk returns a tuple with the PdsChartVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentTargetMetadata) GetPdsChartVersionOk() (*string, bool) {
	if o == nil || o.PdsChartVersion == nil {
		return nil, false
	}
	return o.PdsChartVersion, true
}

// HasPdsChartVersion returns a boolean if a field has been set.
func (o *ModelsDeploymentTargetMetadata) HasPdsChartVersion() bool {
	if o != nil && o.PdsChartVersion != nil {
		return true
	}

	return false
}

// SetPdsChartVersion gets a reference to the given string and assigns it to the PdsChartVersion field.
func (o *ModelsDeploymentTargetMetadata) SetPdsChartVersion(v string) {
	o.PdsChartVersion = &v
}

// GetPxServiceNamespace returns the PxServiceNamespace field value if set, zero value otherwise.
func (o *ModelsDeploymentTargetMetadata) GetPxServiceNamespace() string {
	if o == nil || o.PxServiceNamespace == nil {
		var ret string
		return ret
	}
	return *o.PxServiceNamespace
}

// GetPxServiceNamespaceOk returns a tuple with the PxServiceNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentTargetMetadata) GetPxServiceNamespaceOk() (*string, bool) {
	if o == nil || o.PxServiceNamespace == nil {
		return nil, false
	}
	return o.PxServiceNamespace, true
}

// HasPxServiceNamespace returns a boolean if a field has been set.
func (o *ModelsDeploymentTargetMetadata) HasPxServiceNamespace() bool {
	if o != nil && o.PxServiceNamespace != nil {
		return true
	}

	return false
}

// SetPxServiceNamespace gets a reference to the given string and assigns it to the PxServiceNamespace field.
func (o *ModelsDeploymentTargetMetadata) SetPxServiceNamespace(v string) {
	o.PxServiceNamespace = &v
}

func (o ModelsDeploymentTargetMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KubeApiVersion != nil {
		toSerialize["kube_api_version"] = o.KubeApiVersion
	}
	if o.PdsChartVersion != nil {
		toSerialize["pds_chart_version"] = o.PdsChartVersion
	}
	if o.PxServiceNamespace != nil {
		toSerialize["px_service_namespace"] = o.PxServiceNamespace
	}
	return json.Marshal(toSerialize)
}

type NullableModelsDeploymentTargetMetadata struct {
	value *ModelsDeploymentTargetMetadata
	isSet bool
}

func (v NullableModelsDeploymentTargetMetadata) Get() *ModelsDeploymentTargetMetadata {
	return v.value
}

func (v *NullableModelsDeploymentTargetMetadata) Set(val *ModelsDeploymentTargetMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsDeploymentTargetMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsDeploymentTargetMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsDeploymentTargetMetadata(val *ModelsDeploymentTargetMetadata) *NullableModelsDeploymentTargetMetadata {
	return &NullableModelsDeploymentTargetMetadata{value: val, isSet: true}
}

func (v NullableModelsDeploymentTargetMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsDeploymentTargetMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


