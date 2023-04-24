/*
PDS API

Portworx Data Services API Server

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"encoding/json"
)

// checks if the ModelsBackupPolicySample type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsBackupPolicySample{}

// ModelsBackupPolicySample struct for ModelsBackupPolicySample
type ModelsBackupPolicySample struct {
	Created *string `json:"created,omitempty"`
	// Name of the backup policy. Must be unique for the given tenant.
	Name *string `json:"name,omitempty"`
	// A list of the backup schedules. Must be non-empty.
	Schedules []ModelsBackupSchedule `json:"schedules,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewModelsBackupPolicySample instantiates a new ModelsBackupPolicySample object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsBackupPolicySample() *ModelsBackupPolicySample {
	this := ModelsBackupPolicySample{}
	return &this
}

// NewModelsBackupPolicySampleWithDefaults instantiates a new ModelsBackupPolicySample object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsBackupPolicySampleWithDefaults() *ModelsBackupPolicySample {
	this := ModelsBackupPolicySample{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ModelsBackupPolicySample) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupPolicySample) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ModelsBackupPolicySample) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *ModelsBackupPolicySample) SetCreated(v string) {
	o.Created = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelsBackupPolicySample) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupPolicySample) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelsBackupPolicySample) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelsBackupPolicySample) SetName(v string) {
	o.Name = &v
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *ModelsBackupPolicySample) GetSchedules() []ModelsBackupSchedule {
	if o == nil || IsNil(o.Schedules) {
		var ret []ModelsBackupSchedule
		return ret
	}
	return o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupPolicySample) GetSchedulesOk() ([]ModelsBackupSchedule, bool) {
	if o == nil || IsNil(o.Schedules) {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *ModelsBackupPolicySample) HasSchedules() bool {
	if o != nil && !IsNil(o.Schedules) {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given []ModelsBackupSchedule and assigns it to the Schedules field.
func (o *ModelsBackupPolicySample) SetSchedules(v []ModelsBackupSchedule) {
	o.Schedules = v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ModelsBackupPolicySample) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupPolicySample) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ModelsBackupPolicySample) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *ModelsBackupPolicySample) SetUpdated(v string) {
	o.Updated = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ModelsBackupPolicySample) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupPolicySample) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ModelsBackupPolicySample) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ModelsBackupPolicySample) SetVersion(v int32) {
	o.Version = &v
}

func (o ModelsBackupPolicySample) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsBackupPolicySample) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Schedules) {
		toSerialize["schedules"] = o.Schedules
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableModelsBackupPolicySample struct {
	value *ModelsBackupPolicySample
	isSet bool
}

func (v NullableModelsBackupPolicySample) Get() *ModelsBackupPolicySample {
	return v.value
}

func (v *NullableModelsBackupPolicySample) Set(val *ModelsBackupPolicySample) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsBackupPolicySample) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsBackupPolicySample) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsBackupPolicySample(val *ModelsBackupPolicySample) *NullableModelsBackupPolicySample {
	return &NullableModelsBackupPolicySample{value: val, isSet: true}
}

func (v NullableModelsBackupPolicySample) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsBackupPolicySample) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


