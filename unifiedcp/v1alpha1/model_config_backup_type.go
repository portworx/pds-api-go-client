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

// ConfigBackupType Enum for Backup Type of the backup configuration.   - BACKUP_TYPE_UNSPECIFIED: Unspecified type.  - ADHOC: Adhoc Backup.  - SCHEDULED: Scheduled Backup.
type ConfigBackupType string

// List of ConfigBackupType
const (
	CONFIGBACKUPTYPE_BACKUP_TYPE_UNSPECIFIED ConfigBackupType = "BACKUP_TYPE_UNSPECIFIED"
	CONFIGBACKUPTYPE_ADHOC ConfigBackupType = "ADHOC"
	CONFIGBACKUPTYPE_SCHEDULED ConfigBackupType = "SCHEDULED"
)

// All allowed values of ConfigBackupType enum
var AllowedConfigBackupTypeEnumValues = []ConfigBackupType{
	"BACKUP_TYPE_UNSPECIFIED",
	"ADHOC",
	"SCHEDULED",
}

func (v *ConfigBackupType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConfigBackupType(value)
	for _, existing := range AllowedConfigBackupTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConfigBackupType", value)
}

// NewConfigBackupTypeFromValue returns a pointer to a valid ConfigBackupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConfigBackupTypeFromValue(v string) (*ConfigBackupType, error) {
	ev := ConfigBackupType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConfigBackupType: valid values are %v", v, AllowedConfigBackupTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConfigBackupType) IsValid() bool {
	for _, existing := range AllowedConfigBackupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConfigBackupType value
func (v ConfigBackupType) Ptr() *ConfigBackupType {
	return &v
}

type NullableConfigBackupType struct {
	value *ConfigBackupType
	isSet bool
}

func (v NullableConfigBackupType) Get() *ConfigBackupType {
	return v.value
}

func (v *NullableConfigBackupType) Set(val *ConfigBackupType) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigBackupType) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigBackupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigBackupType(val *ConfigBackupType) *NullableConfigBackupType {
	return &NullableConfigBackupType{value: val, isSet: true}
}

func (v NullableConfigBackupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigBackupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

