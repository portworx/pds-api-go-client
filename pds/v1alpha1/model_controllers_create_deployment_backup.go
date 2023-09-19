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

// ControllersCreateDeploymentBackup struct for ControllersCreateDeploymentBackup
type ControllersCreateDeploymentBackup struct {
	BackupLevel *string `json:"backup_level,omitempty"`
	BackupTargetId *string `json:"backup_target_id,omitempty"`
	BackupType *string `json:"backup_type,omitempty"`
	JobHistoryLimit *int32 `json:"job_history_limit,omitempty"`
	Schedule *string `json:"schedule,omitempty"`
}

// NewControllersCreateDeploymentBackup instantiates a new ControllersCreateDeploymentBackup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersCreateDeploymentBackup() *ControllersCreateDeploymentBackup {
	this := ControllersCreateDeploymentBackup{}
	return &this
}

// NewControllersCreateDeploymentBackupWithDefaults instantiates a new ControllersCreateDeploymentBackup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersCreateDeploymentBackupWithDefaults() *ControllersCreateDeploymentBackup {
	this := ControllersCreateDeploymentBackup{}
	return &this
}

// GetBackupLevel returns the BackupLevel field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentBackup) GetBackupLevel() string {
	if o == nil || o.BackupLevel == nil {
		var ret string
		return ret
	}
	return *o.BackupLevel
}

// GetBackupLevelOk returns a tuple with the BackupLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentBackup) GetBackupLevelOk() (*string, bool) {
	if o == nil || o.BackupLevel == nil {
		return nil, false
	}
	return o.BackupLevel, true
}

// HasBackupLevel returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentBackup) HasBackupLevel() bool {
	if o != nil && o.BackupLevel != nil {
		return true
	}

	return false
}

// SetBackupLevel gets a reference to the given string and assigns it to the BackupLevel field.
func (o *ControllersCreateDeploymentBackup) SetBackupLevel(v string) {
	o.BackupLevel = &v
}

// GetBackupTargetId returns the BackupTargetId field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentBackup) GetBackupTargetId() string {
	if o == nil || o.BackupTargetId == nil {
		var ret string
		return ret
	}
	return *o.BackupTargetId
}

// GetBackupTargetIdOk returns a tuple with the BackupTargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentBackup) GetBackupTargetIdOk() (*string, bool) {
	if o == nil || o.BackupTargetId == nil {
		return nil, false
	}
	return o.BackupTargetId, true
}

// HasBackupTargetId returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentBackup) HasBackupTargetId() bool {
	if o != nil && o.BackupTargetId != nil {
		return true
	}

	return false
}

// SetBackupTargetId gets a reference to the given string and assigns it to the BackupTargetId field.
func (o *ControllersCreateDeploymentBackup) SetBackupTargetId(v string) {
	o.BackupTargetId = &v
}

// GetBackupType returns the BackupType field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentBackup) GetBackupType() string {
	if o == nil || o.BackupType == nil {
		var ret string
		return ret
	}
	return *o.BackupType
}

// GetBackupTypeOk returns a tuple with the BackupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentBackup) GetBackupTypeOk() (*string, bool) {
	if o == nil || o.BackupType == nil {
		return nil, false
	}
	return o.BackupType, true
}

// HasBackupType returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentBackup) HasBackupType() bool {
	if o != nil && o.BackupType != nil {
		return true
	}

	return false
}

// SetBackupType gets a reference to the given string and assigns it to the BackupType field.
func (o *ControllersCreateDeploymentBackup) SetBackupType(v string) {
	o.BackupType = &v
}

// GetJobHistoryLimit returns the JobHistoryLimit field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentBackup) GetJobHistoryLimit() int32 {
	if o == nil || o.JobHistoryLimit == nil {
		var ret int32
		return ret
	}
	return *o.JobHistoryLimit
}

// GetJobHistoryLimitOk returns a tuple with the JobHistoryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentBackup) GetJobHistoryLimitOk() (*int32, bool) {
	if o == nil || o.JobHistoryLimit == nil {
		return nil, false
	}
	return o.JobHistoryLimit, true
}

// HasJobHistoryLimit returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentBackup) HasJobHistoryLimit() bool {
	if o != nil && o.JobHistoryLimit != nil {
		return true
	}

	return false
}

// SetJobHistoryLimit gets a reference to the given int32 and assigns it to the JobHistoryLimit field.
func (o *ControllersCreateDeploymentBackup) SetJobHistoryLimit(v int32) {
	o.JobHistoryLimit = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentBackup) GetSchedule() string {
	if o == nil || o.Schedule == nil {
		var ret string
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentBackup) GetScheduleOk() (*string, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentBackup) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given string and assigns it to the Schedule field.
func (o *ControllersCreateDeploymentBackup) SetSchedule(v string) {
	o.Schedule = &v
}

func (o ControllersCreateDeploymentBackup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackupLevel != nil {
		toSerialize["backup_level"] = o.BackupLevel
	}
	if o.BackupTargetId != nil {
		toSerialize["backup_target_id"] = o.BackupTargetId
	}
	if o.BackupType != nil {
		toSerialize["backup_type"] = o.BackupType
	}
	if o.JobHistoryLimit != nil {
		toSerialize["job_history_limit"] = o.JobHistoryLimit
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}

type NullableControllersCreateDeploymentBackup struct {
	value *ControllersCreateDeploymentBackup
	isSet bool
}

func (v NullableControllersCreateDeploymentBackup) Get() *ControllersCreateDeploymentBackup {
	return v.value
}

func (v *NullableControllersCreateDeploymentBackup) Set(val *ControllersCreateDeploymentBackup) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersCreateDeploymentBackup) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersCreateDeploymentBackup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersCreateDeploymentBackup(val *ControllersCreateDeploymentBackup) *NullableControllersCreateDeploymentBackup {
	return &NullableControllersCreateDeploymentBackup{value: val, isSet: true}
}

func (v NullableControllersCreateDeploymentBackup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersCreateDeploymentBackup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


