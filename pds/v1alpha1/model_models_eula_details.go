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

// ModelsEULADetails struct for ModelsEULADetails
type ModelsEULADetails struct {
	Accepted *bool `json:"accepted,omitempty"`
	AcceptedVersion *string `json:"accepted_version,omitempty"`
}

// NewModelsEULADetails instantiates a new ModelsEULADetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsEULADetails() *ModelsEULADetails {
	this := ModelsEULADetails{}
	return &this
}

// NewModelsEULADetailsWithDefaults instantiates a new ModelsEULADetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsEULADetailsWithDefaults() *ModelsEULADetails {
	this := ModelsEULADetails{}
	return &this
}

// GetAccepted returns the Accepted field value if set, zero value otherwise.
func (o *ModelsEULADetails) GetAccepted() bool {
	if o == nil || o.Accepted == nil {
		var ret bool
		return ret
	}
	return *o.Accepted
}

// GetAcceptedOk returns a tuple with the Accepted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEULADetails) GetAcceptedOk() (*bool, bool) {
	if o == nil || o.Accepted == nil {
		return nil, false
	}
	return o.Accepted, true
}

// HasAccepted returns a boolean if a field has been set.
func (o *ModelsEULADetails) HasAccepted() bool {
	if o != nil && o.Accepted != nil {
		return true
	}

	return false
}

// SetAccepted gets a reference to the given bool and assigns it to the Accepted field.
func (o *ModelsEULADetails) SetAccepted(v bool) {
	o.Accepted = &v
}

// GetAcceptedVersion returns the AcceptedVersion field value if set, zero value otherwise.
func (o *ModelsEULADetails) GetAcceptedVersion() string {
	if o == nil || o.AcceptedVersion == nil {
		var ret string
		return ret
	}
	return *o.AcceptedVersion
}

// GetAcceptedVersionOk returns a tuple with the AcceptedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEULADetails) GetAcceptedVersionOk() (*string, bool) {
	if o == nil || o.AcceptedVersion == nil {
		return nil, false
	}
	return o.AcceptedVersion, true
}

// HasAcceptedVersion returns a boolean if a field has been set.
func (o *ModelsEULADetails) HasAcceptedVersion() bool {
	if o != nil && o.AcceptedVersion != nil {
		return true
	}

	return false
}

// SetAcceptedVersion gets a reference to the given string and assigns it to the AcceptedVersion field.
func (o *ModelsEULADetails) SetAcceptedVersion(v string) {
	o.AcceptedVersion = &v
}

func (o ModelsEULADetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accepted != nil {
		toSerialize["accepted"] = o.Accepted
	}
	if o.AcceptedVersion != nil {
		toSerialize["accepted_version"] = o.AcceptedVersion
	}
	return json.Marshal(toSerialize)
}

type NullableModelsEULADetails struct {
	value *ModelsEULADetails
	isSet bool
}

func (v NullableModelsEULADetails) Get() *ModelsEULADetails {
	return v.value
}

func (v *NullableModelsEULADetails) Set(val *ModelsEULADetails) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsEULADetails) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsEULADetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsEULADetails(val *ModelsEULADetails) *NullableModelsEULADetails {
	return &NullableModelsEULADetails{value: val, isSet: true}
}

func (v NullableModelsEULADetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsEULADetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


