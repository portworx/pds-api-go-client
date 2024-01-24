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

// Pdsdeploymentconfigupdatev1Status struct for Pdsdeploymentconfigupdatev1Status
type Pdsdeploymentconfigupdatev1Status struct {
	// Error Code is a short string that represents the error.
	ErrorCode *string `json:"errorCode,omitempty"`
	// Error Message is a description of the error.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	RetryCount *int32 `json:"retryCount,omitempty"`
	Phase *Pdsdeploymentconfigupdatev1StatusPhase `json:"phase,omitempty"`
}

// NewPdsdeploymentconfigupdatev1Status instantiates a new Pdsdeploymentconfigupdatev1Status object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPdsdeploymentconfigupdatev1Status() *Pdsdeploymentconfigupdatev1Status {
	this := Pdsdeploymentconfigupdatev1Status{}
	var phase Pdsdeploymentconfigupdatev1StatusPhase = PHASE_UNSPECIFIED
	this.Phase = &phase
	return &this
}

// NewPdsdeploymentconfigupdatev1StatusWithDefaults instantiates a new Pdsdeploymentconfigupdatev1Status object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPdsdeploymentconfigupdatev1StatusWithDefaults() *Pdsdeploymentconfigupdatev1Status {
	this := Pdsdeploymentconfigupdatev1Status{}
	var phase Pdsdeploymentconfigupdatev1StatusPhase = PHASE_UNSPECIFIED
	this.Phase = &phase
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *Pdsdeploymentconfigupdatev1Status) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pdsdeploymentconfigupdatev1Status) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *Pdsdeploymentconfigupdatev1Status) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *Pdsdeploymentconfigupdatev1Status) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *Pdsdeploymentconfigupdatev1Status) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pdsdeploymentconfigupdatev1Status) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *Pdsdeploymentconfigupdatev1Status) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *Pdsdeploymentconfigupdatev1Status) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetRetryCount returns the RetryCount field value if set, zero value otherwise.
func (o *Pdsdeploymentconfigupdatev1Status) GetRetryCount() int32 {
	if o == nil || o.RetryCount == nil {
		var ret int32
		return ret
	}
	return *o.RetryCount
}

// GetRetryCountOk returns a tuple with the RetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pdsdeploymentconfigupdatev1Status) GetRetryCountOk() (*int32, bool) {
	if o == nil || o.RetryCount == nil {
		return nil, false
	}
	return o.RetryCount, true
}

// HasRetryCount returns a boolean if a field has been set.
func (o *Pdsdeploymentconfigupdatev1Status) HasRetryCount() bool {
	if o != nil && o.RetryCount != nil {
		return true
	}

	return false
}

// SetRetryCount gets a reference to the given int32 and assigns it to the RetryCount field.
func (o *Pdsdeploymentconfigupdatev1Status) SetRetryCount(v int32) {
	o.RetryCount = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *Pdsdeploymentconfigupdatev1Status) GetPhase() Pdsdeploymentconfigupdatev1StatusPhase {
	if o == nil || o.Phase == nil {
		var ret Pdsdeploymentconfigupdatev1StatusPhase
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pdsdeploymentconfigupdatev1Status) GetPhaseOk() (*Pdsdeploymentconfigupdatev1StatusPhase, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *Pdsdeploymentconfigupdatev1Status) HasPhase() bool {
	if o != nil && o.Phase != nil {
		return true
	}

	return false
}

// SetPhase gets a reference to the given Pdsdeploymentconfigupdatev1StatusPhase and assigns it to the Phase field.
func (o *Pdsdeploymentconfigupdatev1Status) SetPhase(v Pdsdeploymentconfigupdatev1StatusPhase) {
	o.Phase = &v
}

func (o Pdsdeploymentconfigupdatev1Status) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.RetryCount != nil {
		toSerialize["retryCount"] = o.RetryCount
	}
	if o.Phase != nil {
		toSerialize["phase"] = o.Phase
	}
	return json.Marshal(toSerialize)
}

type NullablePdsdeploymentconfigupdatev1Status struct {
	value *Pdsdeploymentconfigupdatev1Status
	isSet bool
}

func (v NullablePdsdeploymentconfigupdatev1Status) Get() *Pdsdeploymentconfigupdatev1Status {
	return v.value
}

func (v *NullablePdsdeploymentconfigupdatev1Status) Set(val *Pdsdeploymentconfigupdatev1Status) {
	v.value = val
	v.isSet = true
}

func (v NullablePdsdeploymentconfigupdatev1Status) IsSet() bool {
	return v.isSet
}

func (v *NullablePdsdeploymentconfigupdatev1Status) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePdsdeploymentconfigupdatev1Status(val *Pdsdeploymentconfigupdatev1Status) *NullablePdsdeploymentconfigupdatev1Status {
	return &NullablePdsdeploymentconfigupdatev1Status{value: val, isSet: true}
}

func (v NullablePdsdeploymentconfigupdatev1Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePdsdeploymentconfigupdatev1Status) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


