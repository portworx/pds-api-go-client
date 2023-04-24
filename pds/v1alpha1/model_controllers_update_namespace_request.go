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

// checks if the ControllersUpdateNamespaceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControllersUpdateNamespaceRequest{}

// ControllersUpdateNamespaceRequest struct for ControllersUpdateNamespaceRequest
type ControllersUpdateNamespaceRequest struct {
	Status *string `json:"status,omitempty"`
}

// NewControllersUpdateNamespaceRequest instantiates a new ControllersUpdateNamespaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersUpdateNamespaceRequest() *ControllersUpdateNamespaceRequest {
	this := ControllersUpdateNamespaceRequest{}
	return &this
}

// NewControllersUpdateNamespaceRequestWithDefaults instantiates a new ControllersUpdateNamespaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersUpdateNamespaceRequestWithDefaults() *ControllersUpdateNamespaceRequest {
	this := ControllersUpdateNamespaceRequest{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ControllersUpdateNamespaceRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersUpdateNamespaceRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ControllersUpdateNamespaceRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ControllersUpdateNamespaceRequest) SetStatus(v string) {
	o.Status = &v
}

func (o ControllersUpdateNamespaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControllersUpdateNamespaceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableControllersUpdateNamespaceRequest struct {
	value *ControllersUpdateNamespaceRequest
	isSet bool
}

func (v NullableControllersUpdateNamespaceRequest) Get() *ControllersUpdateNamespaceRequest {
	return v.value
}

func (v *NullableControllersUpdateNamespaceRequest) Set(val *ControllersUpdateNamespaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersUpdateNamespaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersUpdateNamespaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersUpdateNamespaceRequest(val *ControllersUpdateNamespaceRequest) *NullableControllersUpdateNamespaceRequest {
	return &NullableControllersUpdateNamespaceRequest{value: val, isSet: true}
}

func (v NullableControllersUpdateNamespaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersUpdateNamespaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


