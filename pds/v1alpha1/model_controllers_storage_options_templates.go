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

// ControllersStorageOptionsTemplates struct for ControllersStorageOptionsTemplates
type ControllersStorageOptionsTemplates struct {
	Data []ModelsStorageOptionsTemplate `json:"data,omitempty"`
}

// NewControllersStorageOptionsTemplates instantiates a new ControllersStorageOptionsTemplates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersStorageOptionsTemplates() *ControllersStorageOptionsTemplates {
	this := ControllersStorageOptionsTemplates{}
	return &this
}

// NewControllersStorageOptionsTemplatesWithDefaults instantiates a new ControllersStorageOptionsTemplates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersStorageOptionsTemplatesWithDefaults() *ControllersStorageOptionsTemplates {
	this := ControllersStorageOptionsTemplates{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ControllersStorageOptionsTemplates) GetData() []ModelsStorageOptionsTemplate {
	if o == nil || o.Data == nil {
		var ret []ModelsStorageOptionsTemplate
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersStorageOptionsTemplates) GetDataOk() ([]ModelsStorageOptionsTemplate, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ControllersStorageOptionsTemplates) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ModelsStorageOptionsTemplate and assigns it to the Data field.
func (o *ControllersStorageOptionsTemplates) SetData(v []ModelsStorageOptionsTemplate) {
	o.Data = v
}

func (o ControllersStorageOptionsTemplates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableControllersStorageOptionsTemplates struct {
	value *ControllersStorageOptionsTemplates
	isSet bool
}

func (v NullableControllersStorageOptionsTemplates) Get() *ControllersStorageOptionsTemplates {
	return v.value
}

func (v *NullableControllersStorageOptionsTemplates) Set(val *ControllersStorageOptionsTemplates) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersStorageOptionsTemplates) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersStorageOptionsTemplates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersStorageOptionsTemplates(val *ControllersStorageOptionsTemplates) *NullableControllersStorageOptionsTemplates {
	return &NullableControllersStorageOptionsTemplates{value: val, isSet: true}
}

func (v NullableControllersStorageOptionsTemplates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersStorageOptionsTemplates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


