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

// ModelsPaginatedResultModelsNamespace struct for ModelsPaginatedResultModelsNamespace
type ModelsPaginatedResultModelsNamespace struct {
	Data []ModelsNamespace `json:"data,omitempty"`
	Pagination *ConstraintPagination `json:"pagination,omitempty"`
}

// NewModelsPaginatedResultModelsNamespace instantiates a new ModelsPaginatedResultModelsNamespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsPaginatedResultModelsNamespace() *ModelsPaginatedResultModelsNamespace {
	this := ModelsPaginatedResultModelsNamespace{}
	return &this
}

// NewModelsPaginatedResultModelsNamespaceWithDefaults instantiates a new ModelsPaginatedResultModelsNamespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsPaginatedResultModelsNamespaceWithDefaults() *ModelsPaginatedResultModelsNamespace {
	this := ModelsPaginatedResultModelsNamespace{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModelsPaginatedResultModelsNamespace) GetData() []ModelsNamespace {
	if o == nil || o.Data == nil {
		var ret []ModelsNamespace
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPaginatedResultModelsNamespace) GetDataOk() ([]ModelsNamespace, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModelsPaginatedResultModelsNamespace) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ModelsNamespace and assigns it to the Data field.
func (o *ModelsPaginatedResultModelsNamespace) SetData(v []ModelsNamespace) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ModelsPaginatedResultModelsNamespace) GetPagination() ConstraintPagination {
	if o == nil || o.Pagination == nil {
		var ret ConstraintPagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPaginatedResultModelsNamespace) GetPaginationOk() (*ConstraintPagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ModelsPaginatedResultModelsNamespace) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given ConstraintPagination and assigns it to the Pagination field.
func (o *ModelsPaginatedResultModelsNamespace) SetPagination(v ConstraintPagination) {
	o.Pagination = &v
}

func (o ModelsPaginatedResultModelsNamespace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableModelsPaginatedResultModelsNamespace struct {
	value *ModelsPaginatedResultModelsNamespace
	isSet bool
}

func (v NullableModelsPaginatedResultModelsNamespace) Get() *ModelsPaginatedResultModelsNamespace {
	return v.value
}

func (v *NullableModelsPaginatedResultModelsNamespace) Set(val *ModelsPaginatedResultModelsNamespace) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsPaginatedResultModelsNamespace) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsPaginatedResultModelsNamespace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsPaginatedResultModelsNamespace(val *ModelsPaginatedResultModelsNamespace) *NullableModelsPaginatedResultModelsNamespace {
	return &NullableModelsPaginatedResultModelsNamespace{value: val, isSet: true}
}

func (v NullableModelsPaginatedResultModelsNamespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsPaginatedResultModelsNamespace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


