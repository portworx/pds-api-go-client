/*
public/portworx/pds/backupconfig/apiv1/backupconfig.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Serviceaccountv1Status Status of the service account.
type Serviceaccountv1Status struct {
	// Represents how many times the service account secret has been rotated.
	SecretGenerationCount *int32 `json:"secretGenerationCount,omitempty"`
	// When last time the secrets of the service account has been updated.
	LastSecretUpdateTime *time.Time `json:"lastSecretUpdateTime,omitempty"`
}

// NewServiceaccountv1Status instantiates a new Serviceaccountv1Status object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceaccountv1Status() *Serviceaccountv1Status {
	this := Serviceaccountv1Status{}
	return &this
}

// NewServiceaccountv1StatusWithDefaults instantiates a new Serviceaccountv1Status object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceaccountv1StatusWithDefaults() *Serviceaccountv1Status {
	this := Serviceaccountv1Status{}
	return &this
}

// GetSecretGenerationCount returns the SecretGenerationCount field value if set, zero value otherwise.
func (o *Serviceaccountv1Status) GetSecretGenerationCount() int32 {
	if o == nil || o.SecretGenerationCount == nil {
		var ret int32
		return ret
	}
	return *o.SecretGenerationCount
}

// GetSecretGenerationCountOk returns a tuple with the SecretGenerationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Serviceaccountv1Status) GetSecretGenerationCountOk() (*int32, bool) {
	if o == nil || o.SecretGenerationCount == nil {
		return nil, false
	}
	return o.SecretGenerationCount, true
}

// HasSecretGenerationCount returns a boolean if a field has been set.
func (o *Serviceaccountv1Status) HasSecretGenerationCount() bool {
	if o != nil && o.SecretGenerationCount != nil {
		return true
	}

	return false
}

// SetSecretGenerationCount gets a reference to the given int32 and assigns it to the SecretGenerationCount field.
func (o *Serviceaccountv1Status) SetSecretGenerationCount(v int32) {
	o.SecretGenerationCount = &v
}

// GetLastSecretUpdateTime returns the LastSecretUpdateTime field value if set, zero value otherwise.
func (o *Serviceaccountv1Status) GetLastSecretUpdateTime() time.Time {
	if o == nil || o.LastSecretUpdateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSecretUpdateTime
}

// GetLastSecretUpdateTimeOk returns a tuple with the LastSecretUpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Serviceaccountv1Status) GetLastSecretUpdateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastSecretUpdateTime == nil {
		return nil, false
	}
	return o.LastSecretUpdateTime, true
}

// HasLastSecretUpdateTime returns a boolean if a field has been set.
func (o *Serviceaccountv1Status) HasLastSecretUpdateTime() bool {
	if o != nil && o.LastSecretUpdateTime != nil {
		return true
	}

	return false
}

// SetLastSecretUpdateTime gets a reference to the given time.Time and assigns it to the LastSecretUpdateTime field.
func (o *Serviceaccountv1Status) SetLastSecretUpdateTime(v time.Time) {
	o.LastSecretUpdateTime = &v
}

func (o Serviceaccountv1Status) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SecretGenerationCount != nil {
		toSerialize["secretGenerationCount"] = o.SecretGenerationCount
	}
	if o.LastSecretUpdateTime != nil {
		toSerialize["lastSecretUpdateTime"] = o.LastSecretUpdateTime
	}
	return json.Marshal(toSerialize)
}

type NullableServiceaccountv1Status struct {
	value *Serviceaccountv1Status
	isSet bool
}

func (v NullableServiceaccountv1Status) Get() *Serviceaccountv1Status {
	return v.value
}

func (v *NullableServiceaccountv1Status) Set(val *Serviceaccountv1Status) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceaccountv1Status) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceaccountv1Status) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceaccountv1Status(val *Serviceaccountv1Status) *NullableServiceaccountv1Status {
	return &NullableServiceaccountv1Status{value: val, isSet: true}
}

func (v NullableServiceaccountv1Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceaccountv1Status) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


