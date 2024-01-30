/*
public/portworx/pds/backup/apiv1/backup.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package backup

import (
	"encoding/json"
)

// checks if the V1Sort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1Sort{}

// V1Sort The details of the attribute for which the requested list of resource to be sorted.
type V1Sort struct {
	SortBy *SortByField `json:"sortBy,omitempty"`
	SortOrder *SortOrderValue `json:"sortOrder,omitempty"`
}

// NewV1Sort instantiates a new V1Sort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Sort() *V1Sort {
	this := V1Sort{}
	var sortBy SortByField = FIELD_UNSPECIFIED
	this.SortBy = &sortBy
	var sortOrder SortOrderValue = VALUE_UNSPECIFIED
	this.SortOrder = &sortOrder
	return &this
}

// NewV1SortWithDefaults instantiates a new V1Sort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1SortWithDefaults() *V1Sort {
	this := V1Sort{}
	var sortBy SortByField = FIELD_UNSPECIFIED
	this.SortBy = &sortBy
	var sortOrder SortOrderValue = VALUE_UNSPECIFIED
	this.SortOrder = &sortOrder
	return &this
}

// GetSortBy returns the SortBy field value if set, zero value otherwise.
func (o *V1Sort) GetSortBy() SortByField {
	if o == nil || IsNil(o.SortBy) {
		var ret SortByField
		return ret
	}
	return *o.SortBy
}

// GetSortByOk returns a tuple with the SortBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Sort) GetSortByOk() (*SortByField, bool) {
	if o == nil || IsNil(o.SortBy) {
		return nil, false
	}
	return o.SortBy, true
}

// HasSortBy returns a boolean if a field has been set.
func (o *V1Sort) HasSortBy() bool {
	if o != nil && !IsNil(o.SortBy) {
		return true
	}

	return false
}

// SetSortBy gets a reference to the given SortByField and assigns it to the SortBy field.
func (o *V1Sort) SetSortBy(v SortByField) {
	o.SortBy = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *V1Sort) GetSortOrder() SortOrderValue {
	if o == nil || IsNil(o.SortOrder) {
		var ret SortOrderValue
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Sort) GetSortOrderOk() (*SortOrderValue, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *V1Sort) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given SortOrderValue and assigns it to the SortOrder field.
func (o *V1Sort) SetSortOrder(v SortOrderValue) {
	o.SortOrder = &v
}

func (o V1Sort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1Sort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SortBy) {
		toSerialize["sortBy"] = o.SortBy
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sortOrder"] = o.SortOrder
	}
	return toSerialize, nil
}

type NullableV1Sort struct {
	value *V1Sort
	isSet bool
}

func (v NullableV1Sort) Get() *V1Sort {
	return v.value
}

func (v *NullableV1Sort) Set(val *V1Sort) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Sort) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Sort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Sort(val *V1Sort) *NullableV1Sort {
	return &NullableV1Sort{value: val, isSet: true}
}

func (v NullableV1Sort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Sort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


