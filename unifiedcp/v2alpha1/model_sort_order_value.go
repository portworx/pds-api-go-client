/*
public/portworx/pds/backupconfig/apiv1/backupconfig.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// SortOrderValue Value of sort order for the list of resources.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending.
type SortOrderValue string

// List of SortOrderValue
const (
	VALUE_UNSPECIFIED SortOrderValue = "VALUE_UNSPECIFIED"
	ASC SortOrderValue = "ASC"
	DESC SortOrderValue = "DESC"
)

// All allowed values of SortOrderValue enum
var AllowedSortOrderValueEnumValues = []SortOrderValue{
	"VALUE_UNSPECIFIED",
	"ASC",
	"DESC",
}

func (v *SortOrderValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SortOrderValue(value)
	for _, existing := range AllowedSortOrderValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SortOrderValue", value)
}

// NewSortOrderValueFromValue returns a pointer to a valid SortOrderValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSortOrderValueFromValue(v string) (*SortOrderValue, error) {
	ev := SortOrderValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SortOrderValue: valid values are %v", v, AllowedSortOrderValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SortOrderValue) IsValid() bool {
	for _, existing := range AllowedSortOrderValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SortOrderValue value
func (v SortOrderValue) Ptr() *SortOrderValue {
	return &v
}

type NullableSortOrderValue struct {
	value *SortOrderValue
	isSet bool
}

func (v NullableSortOrderValue) Get() *SortOrderValue {
	return v.value
}

func (v *NullableSortOrderValue) Set(val *SortOrderValue) {
	v.value = val
	v.isSet = true
}

func (v NullableSortOrderValue) IsSet() bool {
	return v.isSet
}

func (v *NullableSortOrderValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortOrderValue(val *SortOrderValue) *NullableSortOrderValue {
	return &NullableSortOrderValue{value: val, isSet: true}
}

func (v NullableSortOrderValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortOrderValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

