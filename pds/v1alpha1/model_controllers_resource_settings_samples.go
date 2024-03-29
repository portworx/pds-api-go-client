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

// ControllersResourceSettingsSamples struct for ControllersResourceSettingsSamples
type ControllersResourceSettingsSamples struct {
	Data []ModelsResourceSettingsSample `json:"data,omitempty"`
}

// NewControllersResourceSettingsSamples instantiates a new ControllersResourceSettingsSamples object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersResourceSettingsSamples() *ControllersResourceSettingsSamples {
	this := ControllersResourceSettingsSamples{}
	return &this
}

// NewControllersResourceSettingsSamplesWithDefaults instantiates a new ControllersResourceSettingsSamples object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersResourceSettingsSamplesWithDefaults() *ControllersResourceSettingsSamples {
	this := ControllersResourceSettingsSamples{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ControllersResourceSettingsSamples) GetData() []ModelsResourceSettingsSample {
	if o == nil || o.Data == nil {
		var ret []ModelsResourceSettingsSample
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersResourceSettingsSamples) GetDataOk() ([]ModelsResourceSettingsSample, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ControllersResourceSettingsSamples) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ModelsResourceSettingsSample and assigns it to the Data field.
func (o *ControllersResourceSettingsSamples) SetData(v []ModelsResourceSettingsSample) {
	o.Data = v
}

func (o ControllersResourceSettingsSamples) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableControllersResourceSettingsSamples struct {
	value *ControllersResourceSettingsSamples
	isSet bool
}

func (v NullableControllersResourceSettingsSamples) Get() *ControllersResourceSettingsSamples {
	return v.value
}

func (v *NullableControllersResourceSettingsSamples) Set(val *ControllersResourceSettingsSamples) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersResourceSettingsSamples) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersResourceSettingsSamples) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersResourceSettingsSamples(val *ControllersResourceSettingsSamples) *NullableControllersResourceSettingsSamples {
	return &NullableControllersResourceSettingsSamples{value: val, isSet: true}
}

func (v NullableControllersResourceSettingsSamples) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersResourceSettingsSamples) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


