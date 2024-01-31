/*
public/portworx/pds/tasks/apiv1/tasks.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pdsclient

import (
	"encoding/json"
	"fmt"
)

// StatusPhase1 Enum for Phase of the Deployment.   - PHASE_UNSPECIFIED: Phase is unspecified.  - QUEUED: Backup is Queued.  - COMPLETED: Backup is completed.  - FAILED: Backup is failed.
type StatusPhase1 string

// List of StatusPhase1
const (
	STATUSPHASE1_PHASE_UNSPECIFIED StatusPhase1 = "PHASE_UNSPECIFIED"
	STATUSPHASE1_QUEUED StatusPhase1 = "QUEUED"
	STATUSPHASE1_COMPLETED StatusPhase1 = "COMPLETED"
	STATUSPHASE1_FAILED StatusPhase1 = "FAILED"
)

// All allowed values of StatusPhase1 enum
var AllowedStatusPhase1EnumValues = []StatusPhase1{
	"PHASE_UNSPECIFIED",
	"QUEUED",
	"COMPLETED",
	"FAILED",
}

func (v *StatusPhase1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StatusPhase1(value)
	for _, existing := range AllowedStatusPhase1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StatusPhase1", value)
}

// NewStatusPhase1FromValue returns a pointer to a valid StatusPhase1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatusPhase1FromValue(v string) (*StatusPhase1, error) {
	ev := StatusPhase1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StatusPhase1: valid values are %v", v, AllowedStatusPhase1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StatusPhase1) IsValid() bool {
	for _, existing := range AllowedStatusPhase1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatusPhase1 value
func (v StatusPhase1) Ptr() *StatusPhase1 {
	return &v
}

type NullableStatusPhase1 struct {
	value *StatusPhase1
	isSet bool
}

func (v NullableStatusPhase1) Get() *StatusPhase1 {
	return v.value
}

func (v *NullableStatusPhase1) Set(val *StatusPhase1) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusPhase1) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusPhase1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusPhase1(val *StatusPhase1) *NullableStatusPhase1 {
	return &NullableStatusPhase1{value: val, isSet: true}
}

func (v NullableStatusPhase1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusPhase1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

