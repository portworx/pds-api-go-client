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

// ControllersPaginatedBackupCredentials struct for ControllersPaginatedBackupCredentials
type ControllersPaginatedBackupCredentials struct {
	Data []ModelsBackupCredentials `json:"data,omitempty"`
	Pagination *ConstraintPagination `json:"pagination,omitempty"`
}

// NewControllersPaginatedBackupCredentials instantiates a new ControllersPaginatedBackupCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersPaginatedBackupCredentials() *ControllersPaginatedBackupCredentials {
	this := ControllersPaginatedBackupCredentials{}
	return &this
}

// NewControllersPaginatedBackupCredentialsWithDefaults instantiates a new ControllersPaginatedBackupCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersPaginatedBackupCredentialsWithDefaults() *ControllersPaginatedBackupCredentials {
	this := ControllersPaginatedBackupCredentials{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ControllersPaginatedBackupCredentials) GetData() []ModelsBackupCredentials {
	if o == nil || o.Data == nil {
		var ret []ModelsBackupCredentials
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersPaginatedBackupCredentials) GetDataOk() ([]ModelsBackupCredentials, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ControllersPaginatedBackupCredentials) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ModelsBackupCredentials and assigns it to the Data field.
func (o *ControllersPaginatedBackupCredentials) SetData(v []ModelsBackupCredentials) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ControllersPaginatedBackupCredentials) GetPagination() ConstraintPagination {
	if o == nil || o.Pagination == nil {
		var ret ConstraintPagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersPaginatedBackupCredentials) GetPaginationOk() (*ConstraintPagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ControllersPaginatedBackupCredentials) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given ConstraintPagination and assigns it to the Pagination field.
func (o *ControllersPaginatedBackupCredentials) SetPagination(v ConstraintPagination) {
	o.Pagination = &v
}

func (o ControllersPaginatedBackupCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableControllersPaginatedBackupCredentials struct {
	value *ControllersPaginatedBackupCredentials
	isSet bool
}

func (v NullableControllersPaginatedBackupCredentials) Get() *ControllersPaginatedBackupCredentials {
	return v.value
}

func (v *NullableControllersPaginatedBackupCredentials) Set(val *ControllersPaginatedBackupCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersPaginatedBackupCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersPaginatedBackupCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersPaginatedBackupCredentials(val *ControllersPaginatedBackupCredentials) *NullableControllersPaginatedBackupCredentials {
	return &NullableControllersPaginatedBackupCredentials{value: val, isSet: true}
}

func (v NullableControllersPaginatedBackupCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersPaginatedBackupCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

