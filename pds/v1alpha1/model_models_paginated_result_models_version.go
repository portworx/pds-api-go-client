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

// ModelsPaginatedResultModelsVersion struct for ModelsPaginatedResultModelsVersion
type ModelsPaginatedResultModelsVersion struct {
	Data []ModelsVersion `json:"data,omitempty"`
	Pagination *ConstraintPagination `json:"pagination,omitempty"`
}

// NewModelsPaginatedResultModelsVersion instantiates a new ModelsPaginatedResultModelsVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsPaginatedResultModelsVersion() *ModelsPaginatedResultModelsVersion {
	this := ModelsPaginatedResultModelsVersion{}
	return &this
}

// NewModelsPaginatedResultModelsVersionWithDefaults instantiates a new ModelsPaginatedResultModelsVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsPaginatedResultModelsVersionWithDefaults() *ModelsPaginatedResultModelsVersion {
	this := ModelsPaginatedResultModelsVersion{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModelsPaginatedResultModelsVersion) GetData() []ModelsVersion {
	if o == nil || o.Data == nil {
		var ret []ModelsVersion
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPaginatedResultModelsVersion) GetDataOk() ([]ModelsVersion, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModelsPaginatedResultModelsVersion) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ModelsVersion and assigns it to the Data field.
func (o *ModelsPaginatedResultModelsVersion) SetData(v []ModelsVersion) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ModelsPaginatedResultModelsVersion) GetPagination() ConstraintPagination {
	if o == nil || o.Pagination == nil {
		var ret ConstraintPagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPaginatedResultModelsVersion) GetPaginationOk() (*ConstraintPagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ModelsPaginatedResultModelsVersion) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given ConstraintPagination and assigns it to the Pagination field.
func (o *ModelsPaginatedResultModelsVersion) SetPagination(v ConstraintPagination) {
	o.Pagination = &v
}

func (o ModelsPaginatedResultModelsVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableModelsPaginatedResultModelsVersion struct {
	value *ModelsPaginatedResultModelsVersion
	isSet bool
}

func (v NullableModelsPaginatedResultModelsVersion) Get() *ModelsPaginatedResultModelsVersion {
	return v.value
}

func (v *NullableModelsPaginatedResultModelsVersion) Set(val *ModelsPaginatedResultModelsVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsPaginatedResultModelsVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsPaginatedResultModelsVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsPaginatedResultModelsVersion(val *ModelsPaginatedResultModelsVersion) *NullableModelsPaginatedResultModelsVersion {
	return &NullableModelsPaginatedResultModelsVersion{value: val, isSet: true}
}

func (v NullableModelsPaginatedResultModelsVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsPaginatedResultModelsVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


