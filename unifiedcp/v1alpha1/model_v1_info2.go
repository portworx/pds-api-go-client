/*
public/portworx/pds/tasks/apiv1/tasks.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pdsclient

import (
	"encoding/json"
)

// checks if the V1Info2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1Info2{}

// V1Info2 Information related to the data service version.
type V1Info2 struct {
	// Enabled indicates if the version is enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewV1Info2 instantiates a new V1Info2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Info2() *V1Info2 {
	this := V1Info2{}
	return &this
}

// NewV1Info2WithDefaults instantiates a new V1Info2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1Info2WithDefaults() *V1Info2 {
	this := V1Info2{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *V1Info2) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Info2) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *V1Info2) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *V1Info2) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o V1Info2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1Info2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableV1Info2 struct {
	value *V1Info2
	isSet bool
}

func (v NullableV1Info2) Get() *V1Info2 {
	return v.value
}

func (v *NullableV1Info2) Set(val *V1Info2) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Info2) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Info2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Info2(val *V1Info2) *NullableV1Info2 {
	return &NullableV1Info2{value: val, isSet: true}
}

func (v NullableV1Info2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Info2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


