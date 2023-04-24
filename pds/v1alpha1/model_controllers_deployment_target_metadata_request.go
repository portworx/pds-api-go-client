/*
PDS API

Portworx Data Services API Server

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"encoding/json"
)

// checks if the ControllersDeploymentTargetMetadataRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControllersDeploymentTargetMetadataRequest{}

// ControllersDeploymentTargetMetadataRequest struct for ControllersDeploymentTargetMetadataRequest
type ControllersDeploymentTargetMetadataRequest struct {
	ClusterId *string `json:"cluster_id,omitempty"`
	KubeApiVersion *string `json:"kube_api_version,omitempty"`
	KubePlatform *string `json:"kube_platform,omitempty"`
	PdsChartVersion *string `json:"pds_chart_version,omitempty"`
	PxCsiEnabled *string `json:"px_csi_enabled,omitempty"`
	PxServiceNamespace *string `json:"px_service_namespace,omitempty"`
	PxVersion *string `json:"px_version,omitempty"`
}

// NewControllersDeploymentTargetMetadataRequest instantiates a new ControllersDeploymentTargetMetadataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersDeploymentTargetMetadataRequest() *ControllersDeploymentTargetMetadataRequest {
	this := ControllersDeploymentTargetMetadataRequest{}
	return &this
}

// NewControllersDeploymentTargetMetadataRequestWithDefaults instantiates a new ControllersDeploymentTargetMetadataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersDeploymentTargetMetadataRequestWithDefaults() *ControllersDeploymentTargetMetadataRequest {
	this := ControllersDeploymentTargetMetadataRequest{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *ControllersDeploymentTargetMetadataRequest) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersDeploymentTargetMetadataRequest) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *ControllersDeploymentTargetMetadataRequest) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *ControllersDeploymentTargetMetadataRequest) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetKubeApiVersion returns the KubeApiVersion field value if set, zero value otherwise.
func (o *ControllersDeploymentTargetMetadataRequest) GetKubeApiVersion() string {
	if o == nil || IsNil(o.KubeApiVersion) {
		var ret string
		return ret
	}
	return *o.KubeApiVersion
}

// GetKubeApiVersionOk returns a tuple with the KubeApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersDeploymentTargetMetadataRequest) GetKubeApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.KubeApiVersion) {
		return nil, false
	}
	return o.KubeApiVersion, true
}

// HasKubeApiVersion returns a boolean if a field has been set.
func (o *ControllersDeploymentTargetMetadataRequest) HasKubeApiVersion() bool {
	if o != nil && !IsNil(o.KubeApiVersion) {
		return true
	}

	return false
}

// SetKubeApiVersion gets a reference to the given string and assigns it to the KubeApiVersion field.
func (o *ControllersDeploymentTargetMetadataRequest) SetKubeApiVersion(v string) {
	o.KubeApiVersion = &v
}

// GetKubePlatform returns the KubePlatform field value if set, zero value otherwise.
func (o *ControllersDeploymentTargetMetadataRequest) GetKubePlatform() string {
	if o == nil || IsNil(o.KubePlatform) {
		var ret string
		return ret
	}
	return *o.KubePlatform
}

// GetKubePlatformOk returns a tuple with the KubePlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersDeploymentTargetMetadataRequest) GetKubePlatformOk() (*string, bool) {
	if o == nil || IsNil(o.KubePlatform) {
		return nil, false
	}
	return o.KubePlatform, true
}

// HasKubePlatform returns a boolean if a field has been set.
func (o *ControllersDeploymentTargetMetadataRequest) HasKubePlatform() bool {
	if o != nil && !IsNil(o.KubePlatform) {
		return true
	}

	return false
}

// SetKubePlatform gets a reference to the given string and assigns it to the KubePlatform field.
func (o *ControllersDeploymentTargetMetadataRequest) SetKubePlatform(v string) {
	o.KubePlatform = &v
}

// GetPdsChartVersion returns the PdsChartVersion field value if set, zero value otherwise.
func (o *ControllersDeploymentTargetMetadataRequest) GetPdsChartVersion() string {
	if o == nil || IsNil(o.PdsChartVersion) {
		var ret string
		return ret
	}
	return *o.PdsChartVersion
}

// GetPdsChartVersionOk returns a tuple with the PdsChartVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersDeploymentTargetMetadataRequest) GetPdsChartVersionOk() (*string, bool) {
	if o == nil || IsNil(o.PdsChartVersion) {
		return nil, false
	}
	return o.PdsChartVersion, true
}

// HasPdsChartVersion returns a boolean if a field has been set.
func (o *ControllersDeploymentTargetMetadataRequest) HasPdsChartVersion() bool {
	if o != nil && !IsNil(o.PdsChartVersion) {
		return true
	}

	return false
}

// SetPdsChartVersion gets a reference to the given string and assigns it to the PdsChartVersion field.
func (o *ControllersDeploymentTargetMetadataRequest) SetPdsChartVersion(v string) {
	o.PdsChartVersion = &v
}

// GetPxCsiEnabled returns the PxCsiEnabled field value if set, zero value otherwise.
func (o *ControllersDeploymentTargetMetadataRequest) GetPxCsiEnabled() string {
	if o == nil || IsNil(o.PxCsiEnabled) {
		var ret string
		return ret
	}
	return *o.PxCsiEnabled
}

// GetPxCsiEnabledOk returns a tuple with the PxCsiEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersDeploymentTargetMetadataRequest) GetPxCsiEnabledOk() (*string, bool) {
	if o == nil || IsNil(o.PxCsiEnabled) {
		return nil, false
	}
	return o.PxCsiEnabled, true
}

// HasPxCsiEnabled returns a boolean if a field has been set.
func (o *ControllersDeploymentTargetMetadataRequest) HasPxCsiEnabled() bool {
	if o != nil && !IsNil(o.PxCsiEnabled) {
		return true
	}

	return false
}

// SetPxCsiEnabled gets a reference to the given string and assigns it to the PxCsiEnabled field.
func (o *ControllersDeploymentTargetMetadataRequest) SetPxCsiEnabled(v string) {
	o.PxCsiEnabled = &v
}

// GetPxServiceNamespace returns the PxServiceNamespace field value if set, zero value otherwise.
func (o *ControllersDeploymentTargetMetadataRequest) GetPxServiceNamespace() string {
	if o == nil || IsNil(o.PxServiceNamespace) {
		var ret string
		return ret
	}
	return *o.PxServiceNamespace
}

// GetPxServiceNamespaceOk returns a tuple with the PxServiceNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersDeploymentTargetMetadataRequest) GetPxServiceNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.PxServiceNamespace) {
		return nil, false
	}
	return o.PxServiceNamespace, true
}

// HasPxServiceNamespace returns a boolean if a field has been set.
func (o *ControllersDeploymentTargetMetadataRequest) HasPxServiceNamespace() bool {
	if o != nil && !IsNil(o.PxServiceNamespace) {
		return true
	}

	return false
}

// SetPxServiceNamespace gets a reference to the given string and assigns it to the PxServiceNamespace field.
func (o *ControllersDeploymentTargetMetadataRequest) SetPxServiceNamespace(v string) {
	o.PxServiceNamespace = &v
}

// GetPxVersion returns the PxVersion field value if set, zero value otherwise.
func (o *ControllersDeploymentTargetMetadataRequest) GetPxVersion() string {
	if o == nil || IsNil(o.PxVersion) {
		var ret string
		return ret
	}
	return *o.PxVersion
}

// GetPxVersionOk returns a tuple with the PxVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersDeploymentTargetMetadataRequest) GetPxVersionOk() (*string, bool) {
	if o == nil || IsNil(o.PxVersion) {
		return nil, false
	}
	return o.PxVersion, true
}

// HasPxVersion returns a boolean if a field has been set.
func (o *ControllersDeploymentTargetMetadataRequest) HasPxVersion() bool {
	if o != nil && !IsNil(o.PxVersion) {
		return true
	}

	return false
}

// SetPxVersion gets a reference to the given string and assigns it to the PxVersion field.
func (o *ControllersDeploymentTargetMetadataRequest) SetPxVersion(v string) {
	o.PxVersion = &v
}

func (o ControllersDeploymentTargetMetadataRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControllersDeploymentTargetMetadataRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClusterId) {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if !IsNil(o.KubeApiVersion) {
		toSerialize["kube_api_version"] = o.KubeApiVersion
	}
	if !IsNil(o.KubePlatform) {
		toSerialize["kube_platform"] = o.KubePlatform
	}
	if !IsNil(o.PdsChartVersion) {
		toSerialize["pds_chart_version"] = o.PdsChartVersion
	}
	if !IsNil(o.PxCsiEnabled) {
		toSerialize["px_csi_enabled"] = o.PxCsiEnabled
	}
	if !IsNil(o.PxServiceNamespace) {
		toSerialize["px_service_namespace"] = o.PxServiceNamespace
	}
	if !IsNil(o.PxVersion) {
		toSerialize["px_version"] = o.PxVersion
	}
	return toSerialize, nil
}

type NullableControllersDeploymentTargetMetadataRequest struct {
	value *ControllersDeploymentTargetMetadataRequest
	isSet bool
}

func (v NullableControllersDeploymentTargetMetadataRequest) Get() *ControllersDeploymentTargetMetadataRequest {
	return v.value
}

func (v *NullableControllersDeploymentTargetMetadataRequest) Set(val *ControllersDeploymentTargetMetadataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersDeploymentTargetMetadataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersDeploymentTargetMetadataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersDeploymentTargetMetadataRequest(val *ControllersDeploymentTargetMetadataRequest) *NullableControllersDeploymentTargetMetadataRequest {
	return &NullableControllersDeploymentTargetMetadataRequest{value: val, isSet: true}
}

func (v NullableControllersDeploymentTargetMetadataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersDeploymentTargetMetadataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


