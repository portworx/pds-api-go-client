/*
public/portworx/pds/backupconfig/apiv1/backupconfig.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// V1Config1 Desired configuration of the Backup.
type V1Config1 struct {
	References *V1References1 `json:"references,omitempty"`
	// BackupCapability of the deployment target when the snapshot was created.
	BackupCapability *string `json:"backupCapability,omitempty"`
}

// NewV1Config1 instantiates a new V1Config1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Config1() *V1Config1 {
	this := V1Config1{}
	return &this
}

// NewV1Config1WithDefaults instantiates a new V1Config1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1Config1WithDefaults() *V1Config1 {
	this := V1Config1{}
	return &this
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *V1Config1) GetReferences() V1References1 {
	if o == nil || o.References == nil {
		var ret V1References1
		return ret
	}
	return *o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Config1) GetReferencesOk() (*V1References1, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *V1Config1) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given V1References1 and assigns it to the References field.
func (o *V1Config1) SetReferences(v V1References1) {
	o.References = &v
}

// GetBackupCapability returns the BackupCapability field value if set, zero value otherwise.
func (o *V1Config1) GetBackupCapability() string {
	if o == nil || o.BackupCapability == nil {
		var ret string
		return ret
	}
	return *o.BackupCapability
}

// GetBackupCapabilityOk returns a tuple with the BackupCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Config1) GetBackupCapabilityOk() (*string, bool) {
	if o == nil || o.BackupCapability == nil {
		return nil, false
	}
	return o.BackupCapability, true
}

// HasBackupCapability returns a boolean if a field has been set.
func (o *V1Config1) HasBackupCapability() bool {
	if o != nil && o.BackupCapability != nil {
		return true
	}

	return false
}

// SetBackupCapability gets a reference to the given string and assigns it to the BackupCapability field.
func (o *V1Config1) SetBackupCapability(v string) {
	o.BackupCapability = &v
}

func (o V1Config1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	if o.BackupCapability != nil {
		toSerialize["backupCapability"] = o.BackupCapability
	}
	return json.Marshal(toSerialize)
}

type NullableV1Config1 struct {
	value *V1Config1
	isSet bool
}

func (v NullableV1Config1) Get() *V1Config1 {
	return v.value
}

func (v *NullableV1Config1) Set(val *V1Config1) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Config1) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Config1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Config1(val *V1Config1) *NullableV1Config1 {
	return &NullableV1Config1{value: val, isSet: true}
}

func (v NullableV1Config1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Config1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


