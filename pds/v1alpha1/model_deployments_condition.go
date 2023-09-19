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

// DeploymentsCondition struct for DeploymentsCondition
type DeploymentsCondition struct {
	Message *string `json:"message,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Status *string `json:"status,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewDeploymentsCondition instantiates a new DeploymentsCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentsCondition() *DeploymentsCondition {
	this := DeploymentsCondition{}
	return &this
}

// NewDeploymentsConditionWithDefaults instantiates a new DeploymentsCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentsConditionWithDefaults() *DeploymentsCondition {
	this := DeploymentsCondition{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DeploymentsCondition) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentsCondition) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DeploymentsCondition) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DeploymentsCondition) SetMessage(v string) {
	o.Message = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *DeploymentsCondition) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentsCondition) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *DeploymentsCondition) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *DeploymentsCondition) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentsCondition) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentsCondition) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentsCondition) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeploymentsCondition) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeploymentsCondition) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentsCondition) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeploymentsCondition) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DeploymentsCondition) SetType(v string) {
	o.Type = &v
}

func (o DeploymentsCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentsCondition struct {
	value *DeploymentsCondition
	isSet bool
}

func (v NullableDeploymentsCondition) Get() *DeploymentsCondition {
	return v.value
}

func (v *NullableDeploymentsCondition) Set(val *DeploymentsCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentsCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentsCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentsCondition(val *DeploymentsCondition) *NullableDeploymentsCondition {
	return &NullableDeploymentsCondition{value: val, isSet: true}
}

func (v NullableDeploymentsCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentsCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


