/*
PDS API

Portworx Data Services API Server

API version: 121
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"encoding/json"
)

// ControllersCompatibleVersionsResponse struct for ControllersCompatibleVersionsResponse
type ControllersCompatibleVersionsResponse struct {
	Data []CompatibilityCompatibleVersions `json:"data,omitempty"`
}

// NewControllersCompatibleVersionsResponse instantiates a new ControllersCompatibleVersionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersCompatibleVersionsResponse() *ControllersCompatibleVersionsResponse {
	this := ControllersCompatibleVersionsResponse{}
	return &this
}

// NewControllersCompatibleVersionsResponseWithDefaults instantiates a new ControllersCompatibleVersionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersCompatibleVersionsResponseWithDefaults() *ControllersCompatibleVersionsResponse {
	this := ControllersCompatibleVersionsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ControllersCompatibleVersionsResponse) GetData() []CompatibilityCompatibleVersions {
	if o == nil || o.Data == nil {
		var ret []CompatibilityCompatibleVersions
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCompatibleVersionsResponse) GetDataOk() ([]CompatibilityCompatibleVersions, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ControllersCompatibleVersionsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []CompatibilityCompatibleVersions and assigns it to the Data field.
func (o *ControllersCompatibleVersionsResponse) SetData(v []CompatibilityCompatibleVersions) {
	o.Data = v
}

func (o ControllersCompatibleVersionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableControllersCompatibleVersionsResponse struct {
	value *ControllersCompatibleVersionsResponse
	isSet bool
}

func (v NullableControllersCompatibleVersionsResponse) Get() *ControllersCompatibleVersionsResponse {
	return v.value
}

func (v *NullableControllersCompatibleVersionsResponse) Set(val *ControllersCompatibleVersionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersCompatibleVersionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersCompatibleVersionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersCompatibleVersionsResponse(val *ControllersCompatibleVersionsResponse) *NullableControllersCompatibleVersionsResponse {
	return &NullableControllersCompatibleVersionsResponse{value: val, isSet: true}
}

func (v NullableControllersCompatibleVersionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersCompatibleVersionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


