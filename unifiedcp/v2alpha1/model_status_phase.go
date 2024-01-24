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

// StatusPhase Enum for Phases of the backup configuration.   - PHASE_UNSPECIFIED: Unspecified state.  - PENDING: Pending state.  - FAILED: Failed state.  - CREATED: Created state.  - SUSPENDED: Suspended state.
type StatusPhase string

// List of StatusPhase
const (
	PHASE_UNSPECIFIED StatusPhase = "PHASE_UNSPECIFIED"
	PENDING StatusPhase = "PENDING"
	FAILED StatusPhase = "FAILED"
	CREATED StatusPhase = "CREATED"
	SUSPENDED StatusPhase = "SUSPENDED"
)

// All allowed values of StatusPhase enum
var AllowedStatusPhaseEnumValues = []StatusPhase{
	"PHASE_UNSPECIFIED",
	"PENDING",
	"FAILED",
	"CREATED",
	"SUSPENDED",
}

func (v *StatusPhase) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StatusPhase(value)
	for _, existing := range AllowedStatusPhaseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StatusPhase", value)
}

// NewStatusPhaseFromValue returns a pointer to a valid StatusPhase
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatusPhaseFromValue(v string) (*StatusPhase, error) {
	ev := StatusPhase(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StatusPhase: valid values are %v", v, AllowedStatusPhaseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StatusPhase) IsValid() bool {
	for _, existing := range AllowedStatusPhaseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatusPhase value
func (v StatusPhase) Ptr() *StatusPhase {
	return &v
}

type NullableStatusPhase struct {
	value *StatusPhase
	isSet bool
}

func (v NullableStatusPhase) Get() *StatusPhase {
	return v.value
}

func (v *NullableStatusPhase) Set(val *StatusPhase) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusPhase) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusPhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusPhase(val *StatusPhase) *NullableStatusPhase {
	return &NullableStatusPhase{value: val, isSet: true}
}

func (v NullableStatusPhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusPhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

