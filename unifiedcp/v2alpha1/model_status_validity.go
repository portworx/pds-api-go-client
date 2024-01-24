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

// StatusValidity - VALIDITY_UNSPECIFIED: Validity unspecified  - VALID: Valid backup location indicating its reachable  - INVALID: Invalid backup location indicating its unreachable  - NOT_APPLICABLE: NOT_APPLICABLE backup location specifies that status can not be defined for the backup location eg: S3Compatible location
type StatusValidity string

// List of StatusValidity
const (
	VALIDITY_UNSPECIFIED StatusValidity = "VALIDITY_UNSPECIFIED"
	VALID StatusValidity = "VALID"
	INVALID StatusValidity = "INVALID"
	NOT_APPLICABLE StatusValidity = "NOT_APPLICABLE"
)

// All allowed values of StatusValidity enum
var AllowedStatusValidityEnumValues = []StatusValidity{
	"VALIDITY_UNSPECIFIED",
	"VALID",
	"INVALID",
	"NOT_APPLICABLE",
}

func (v *StatusValidity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StatusValidity(value)
	for _, existing := range AllowedStatusValidityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StatusValidity", value)
}

// NewStatusValidityFromValue returns a pointer to a valid StatusValidity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatusValidityFromValue(v string) (*StatusValidity, error) {
	ev := StatusValidity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StatusValidity: valid values are %v", v, AllowedStatusValidityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StatusValidity) IsValid() bool {
	for _, existing := range AllowedStatusValidityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatusValidity value
func (v StatusValidity) Ptr() *StatusValidity {
	return &v
}

type NullableStatusValidity struct {
	value *StatusValidity
	isSet bool
}

func (v NullableStatusValidity) Get() *StatusValidity {
	return v.value
}

func (v *NullableStatusValidity) Set(val *StatusValidity) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusValidity) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusValidity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusValidity(val *StatusValidity) *NullableStatusValidity {
	return &NullableStatusValidity{value: val, isSet: true}
}

func (v NullableStatusValidity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusValidity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

