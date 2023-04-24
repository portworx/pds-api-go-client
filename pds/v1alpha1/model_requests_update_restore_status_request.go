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

// checks if the RequestsUpdateRestoreStatusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestsUpdateRestoreStatusRequest{}

// RequestsUpdateRestoreStatusRequest struct for RequestsUpdateRestoreStatusRequest
type RequestsUpdateRestoreStatusRequest struct {
	// Status of the restore process on the Target Cluster
	CompletionStatus *string `json:"completion_status,omitempty"`
	// CompletionTime of the restore process.
	CompletionTime *string `json:"completion_time,omitempty"`
	// Error code of the restore from Target Cluster
	ErrorCode *string `json:"error_code,omitempty"`
	// StartTime of the restore process.
	StartTime *string `json:"start_time,omitempty"`
}

// NewRequestsUpdateRestoreStatusRequest instantiates a new RequestsUpdateRestoreStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestsUpdateRestoreStatusRequest() *RequestsUpdateRestoreStatusRequest {
	this := RequestsUpdateRestoreStatusRequest{}
	return &this
}

// NewRequestsUpdateRestoreStatusRequestWithDefaults instantiates a new RequestsUpdateRestoreStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestsUpdateRestoreStatusRequestWithDefaults() *RequestsUpdateRestoreStatusRequest {
	this := RequestsUpdateRestoreStatusRequest{}
	return &this
}

// GetCompletionStatus returns the CompletionStatus field value if set, zero value otherwise.
func (o *RequestsUpdateRestoreStatusRequest) GetCompletionStatus() string {
	if o == nil || IsNil(o.CompletionStatus) {
		var ret string
		return ret
	}
	return *o.CompletionStatus
}

// GetCompletionStatusOk returns a tuple with the CompletionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsUpdateRestoreStatusRequest) GetCompletionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CompletionStatus) {
		return nil, false
	}
	return o.CompletionStatus, true
}

// HasCompletionStatus returns a boolean if a field has been set.
func (o *RequestsUpdateRestoreStatusRequest) HasCompletionStatus() bool {
	if o != nil && !IsNil(o.CompletionStatus) {
		return true
	}

	return false
}

// SetCompletionStatus gets a reference to the given string and assigns it to the CompletionStatus field.
func (o *RequestsUpdateRestoreStatusRequest) SetCompletionStatus(v string) {
	o.CompletionStatus = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *RequestsUpdateRestoreStatusRequest) GetCompletionTime() string {
	if o == nil || IsNil(o.CompletionTime) {
		var ret string
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsUpdateRestoreStatusRequest) GetCompletionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CompletionTime) {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *RequestsUpdateRestoreStatusRequest) HasCompletionTime() bool {
	if o != nil && !IsNil(o.CompletionTime) {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given string and assigns it to the CompletionTime field.
func (o *RequestsUpdateRestoreStatusRequest) SetCompletionTime(v string) {
	o.CompletionTime = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *RequestsUpdateRestoreStatusRequest) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsUpdateRestoreStatusRequest) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *RequestsUpdateRestoreStatusRequest) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *RequestsUpdateRestoreStatusRequest) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *RequestsUpdateRestoreStatusRequest) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsUpdateRestoreStatusRequest) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *RequestsUpdateRestoreStatusRequest) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *RequestsUpdateRestoreStatusRequest) SetStartTime(v string) {
	o.StartTime = &v
}

func (o RequestsUpdateRestoreStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestsUpdateRestoreStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompletionStatus) {
		toSerialize["completion_status"] = o.CompletionStatus
	}
	if !IsNil(o.CompletionTime) {
		toSerialize["completion_time"] = o.CompletionTime
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	return toSerialize, nil
}

type NullableRequestsUpdateRestoreStatusRequest struct {
	value *RequestsUpdateRestoreStatusRequest
	isSet bool
}

func (v NullableRequestsUpdateRestoreStatusRequest) Get() *RequestsUpdateRestoreStatusRequest {
	return v.value
}

func (v *NullableRequestsUpdateRestoreStatusRequest) Set(val *RequestsUpdateRestoreStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestsUpdateRestoreStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestsUpdateRestoreStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestsUpdateRestoreStatusRequest(val *RequestsUpdateRestoreStatusRequest) *NullableRequestsUpdateRestoreStatusRequest {
	return &NullableRequestsUpdateRestoreStatusRequest{value: val, isSet: true}
}

func (v NullableRequestsUpdateRestoreStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestsUpdateRestoreStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


