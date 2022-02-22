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

// ControllersDeploymentTargetHeartbeatResponse struct for ControllersDeploymentTargetHeartbeatResponse
type ControllersDeploymentTargetHeartbeatResponse struct {
	ClusterId *string `json:"cluster_id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewControllersDeploymentTargetHeartbeatResponse instantiates a new ControllersDeploymentTargetHeartbeatResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersDeploymentTargetHeartbeatResponse() *ControllersDeploymentTargetHeartbeatResponse {
	this := ControllersDeploymentTargetHeartbeatResponse{}
	return &this
}

// NewControllersDeploymentTargetHeartbeatResponseWithDefaults instantiates a new ControllersDeploymentTargetHeartbeatResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersDeploymentTargetHeartbeatResponseWithDefaults() *ControllersDeploymentTargetHeartbeatResponse {
	this := ControllersDeploymentTargetHeartbeatResponse{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *ControllersDeploymentTargetHeartbeatResponse) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersDeploymentTargetHeartbeatResponse) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *ControllersDeploymentTargetHeartbeatResponse) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *ControllersDeploymentTargetHeartbeatResponse) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ControllersDeploymentTargetHeartbeatResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersDeploymentTargetHeartbeatResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ControllersDeploymentTargetHeartbeatResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ControllersDeploymentTargetHeartbeatResponse) SetName(v string) {
	o.Name = &v
}

func (o ControllersDeploymentTargetHeartbeatResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClusterId != nil {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableControllersDeploymentTargetHeartbeatResponse struct {
	value *ControllersDeploymentTargetHeartbeatResponse
	isSet bool
}

func (v NullableControllersDeploymentTargetHeartbeatResponse) Get() *ControllersDeploymentTargetHeartbeatResponse {
	return v.value
}

func (v *NullableControllersDeploymentTargetHeartbeatResponse) Set(val *ControllersDeploymentTargetHeartbeatResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersDeploymentTargetHeartbeatResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersDeploymentTargetHeartbeatResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersDeploymentTargetHeartbeatResponse(val *ControllersDeploymentTargetHeartbeatResponse) *NullableControllersDeploymentTargetHeartbeatResponse {
	return &NullableControllersDeploymentTargetHeartbeatResponse{value: val, isSet: true}
}

func (v NullableControllersDeploymentTargetHeartbeatResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersDeploymentTargetHeartbeatResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


