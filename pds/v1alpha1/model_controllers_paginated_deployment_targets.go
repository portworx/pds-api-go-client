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

// ControllersPaginatedDeploymentTargets struct for ControllersPaginatedDeploymentTargets
type ControllersPaginatedDeploymentTargets struct {
	Data []ModelsDeploymentTarget `json:"data,omitempty"`
	Pagination *ConstraintPagination `json:"pagination,omitempty"`
}

// NewControllersPaginatedDeploymentTargets instantiates a new ControllersPaginatedDeploymentTargets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersPaginatedDeploymentTargets() *ControllersPaginatedDeploymentTargets {
	this := ControllersPaginatedDeploymentTargets{}
	return &this
}

// NewControllersPaginatedDeploymentTargetsWithDefaults instantiates a new ControllersPaginatedDeploymentTargets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersPaginatedDeploymentTargetsWithDefaults() *ControllersPaginatedDeploymentTargets {
	this := ControllersPaginatedDeploymentTargets{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ControllersPaginatedDeploymentTargets) GetData() []ModelsDeploymentTarget {
	if o == nil || o.Data == nil {
		var ret []ModelsDeploymentTarget
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersPaginatedDeploymentTargets) GetDataOk() ([]ModelsDeploymentTarget, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ControllersPaginatedDeploymentTargets) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ModelsDeploymentTarget and assigns it to the Data field.
func (o *ControllersPaginatedDeploymentTargets) SetData(v []ModelsDeploymentTarget) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ControllersPaginatedDeploymentTargets) GetPagination() ConstraintPagination {
	if o == nil || o.Pagination == nil {
		var ret ConstraintPagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersPaginatedDeploymentTargets) GetPaginationOk() (*ConstraintPagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ControllersPaginatedDeploymentTargets) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given ConstraintPagination and assigns it to the Pagination field.
func (o *ControllersPaginatedDeploymentTargets) SetPagination(v ConstraintPagination) {
	o.Pagination = &v
}

func (o ControllersPaginatedDeploymentTargets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableControllersPaginatedDeploymentTargets struct {
	value *ControllersPaginatedDeploymentTargets
	isSet bool
}

func (v NullableControllersPaginatedDeploymentTargets) Get() *ControllersPaginatedDeploymentTargets {
	return v.value
}

func (v *NullableControllersPaginatedDeploymentTargets) Set(val *ControllersPaginatedDeploymentTargets) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersPaginatedDeploymentTargets) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersPaginatedDeploymentTargets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersPaginatedDeploymentTargets(val *ControllersPaginatedDeploymentTargets) *NullableControllersPaginatedDeploymentTargets {
	return &NullableControllersPaginatedDeploymentTargets{value: val, isSet: true}
}

func (v NullableControllersPaginatedDeploymentTargets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersPaginatedDeploymentTargets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


