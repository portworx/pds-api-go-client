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

// checks if the ControllersPaginatedTenantRoleBindings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControllersPaginatedTenantRoleBindings{}

// ControllersPaginatedTenantRoleBindings struct for ControllersPaginatedTenantRoleBindings
type ControllersPaginatedTenantRoleBindings struct {
	Data []ModelsTenantRoleBinding `json:"data,omitempty"`
}

// NewControllersPaginatedTenantRoleBindings instantiates a new ControllersPaginatedTenantRoleBindings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersPaginatedTenantRoleBindings() *ControllersPaginatedTenantRoleBindings {
	this := ControllersPaginatedTenantRoleBindings{}
	return &this
}

// NewControllersPaginatedTenantRoleBindingsWithDefaults instantiates a new ControllersPaginatedTenantRoleBindings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersPaginatedTenantRoleBindingsWithDefaults() *ControllersPaginatedTenantRoleBindings {
	this := ControllersPaginatedTenantRoleBindings{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ControllersPaginatedTenantRoleBindings) GetData() []ModelsTenantRoleBinding {
	if o == nil || IsNil(o.Data) {
		var ret []ModelsTenantRoleBinding
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersPaginatedTenantRoleBindings) GetDataOk() ([]ModelsTenantRoleBinding, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ControllersPaginatedTenantRoleBindings) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ModelsTenantRoleBinding and assigns it to the Data field.
func (o *ControllersPaginatedTenantRoleBindings) SetData(v []ModelsTenantRoleBinding) {
	o.Data = v
}

func (o ControllersPaginatedTenantRoleBindings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControllersPaginatedTenantRoleBindings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableControllersPaginatedTenantRoleBindings struct {
	value *ControllersPaginatedTenantRoleBindings
	isSet bool
}

func (v NullableControllersPaginatedTenantRoleBindings) Get() *ControllersPaginatedTenantRoleBindings {
	return v.value
}

func (v *NullableControllersPaginatedTenantRoleBindings) Set(val *ControllersPaginatedTenantRoleBindings) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersPaginatedTenantRoleBindings) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersPaginatedTenantRoleBindings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersPaginatedTenantRoleBindings(val *ControllersPaginatedTenantRoleBindings) *NullableControllersPaginatedTenantRoleBindings {
	return &NullableControllersPaginatedTenantRoleBindings{value: val, isSet: true}
}

func (v NullableControllersPaginatedTenantRoleBindings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersPaginatedTenantRoleBindings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


