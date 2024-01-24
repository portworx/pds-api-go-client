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

// V1StatusPhase Enum for Status of the Deployment.   - PHASE_UNSPECIFIED: Phase is unspecified.  - INITIALIZING: Deployment is initializing.  - RESTORING: Restore of Deployment is in progress.  - FAILED: Deployment is Failed.  - DELETING: Deployment is being deleted.  - PREPARING: Deployment is Preparing.  - DEPLOYED: Deployment is Deployed.  - TIMED_OUT: Deployment is Timeout.
type V1StatusPhase string

// List of v1StatusPhase
const (
	PHASE_UNSPECIFIED V1StatusPhase = "PHASE_UNSPECIFIED"
	INITIALIZING V1StatusPhase = "INITIALIZING"
	RESTORING V1StatusPhase = "RESTORING"
	FAILED V1StatusPhase = "FAILED"
	DELETING V1StatusPhase = "DELETING"
	PREPARING V1StatusPhase = "PREPARING"
	DEPLOYED V1StatusPhase = "DEPLOYED"
	TIMED_OUT V1StatusPhase = "TIMED_OUT"
)

// All allowed values of V1StatusPhase enum
var AllowedV1StatusPhaseEnumValues = []V1StatusPhase{
	"PHASE_UNSPECIFIED",
	"INITIALIZING",
	"RESTORING",
	"FAILED",
	"DELETING",
	"PREPARING",
	"DEPLOYED",
	"TIMED_OUT",
}

func (v *V1StatusPhase) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1StatusPhase(value)
	for _, existing := range AllowedV1StatusPhaseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1StatusPhase", value)
}

// NewV1StatusPhaseFromValue returns a pointer to a valid V1StatusPhase
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1StatusPhaseFromValue(v string) (*V1StatusPhase, error) {
	ev := V1StatusPhase(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1StatusPhase: valid values are %v", v, AllowedV1StatusPhaseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1StatusPhase) IsValid() bool {
	for _, existing := range AllowedV1StatusPhaseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1StatusPhase value
func (v V1StatusPhase) Ptr() *V1StatusPhase {
	return &v
}

type NullableV1StatusPhase struct {
	value *V1StatusPhase
	isSet bool
}

func (v NullableV1StatusPhase) Get() *V1StatusPhase {
	return v.value
}

func (v *NullableV1StatusPhase) Set(val *V1StatusPhase) {
	v.value = val
	v.isSet = true
}

func (v NullableV1StatusPhase) IsSet() bool {
	return v.isSet
}

func (v *NullableV1StatusPhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1StatusPhase(val *V1StatusPhase) *NullableV1StatusPhase {
	return &NullableV1StatusPhase{value: val, isSet: true}
}

func (v NullableV1StatusPhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1StatusPhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

