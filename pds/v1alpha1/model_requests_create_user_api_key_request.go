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

// RequestsCreateUserAPIKeyRequest struct for RequestsCreateUserAPIKeyRequest
type RequestsCreateUserAPIKeyRequest struct {
	// Time when the key expires.
	ExpiresAt *string `json:"expires_at,omitempty"`
	// Name of the UserAPIKey. Must be unique for the given User.
	Name *string `json:"name,omitempty"`
}

// NewRequestsCreateUserAPIKeyRequest instantiates a new RequestsCreateUserAPIKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestsCreateUserAPIKeyRequest() *RequestsCreateUserAPIKeyRequest {
	this := RequestsCreateUserAPIKeyRequest{}
	return &this
}

// NewRequestsCreateUserAPIKeyRequestWithDefaults instantiates a new RequestsCreateUserAPIKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestsCreateUserAPIKeyRequestWithDefaults() *RequestsCreateUserAPIKeyRequest {
	this := RequestsCreateUserAPIKeyRequest{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *RequestsCreateUserAPIKeyRequest) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateUserAPIKeyRequest) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *RequestsCreateUserAPIKeyRequest) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *RequestsCreateUserAPIKeyRequest) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RequestsCreateUserAPIKeyRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateUserAPIKeyRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RequestsCreateUserAPIKeyRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RequestsCreateUserAPIKeyRequest) SetName(v string) {
	o.Name = &v
}

func (o RequestsCreateUserAPIKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableRequestsCreateUserAPIKeyRequest struct {
	value *RequestsCreateUserAPIKeyRequest
	isSet bool
}

func (v NullableRequestsCreateUserAPIKeyRequest) Get() *RequestsCreateUserAPIKeyRequest {
	return v.value
}

func (v *NullableRequestsCreateUserAPIKeyRequest) Set(val *RequestsCreateUserAPIKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestsCreateUserAPIKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestsCreateUserAPIKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestsCreateUserAPIKeyRequest(val *RequestsCreateUserAPIKeyRequest) *NullableRequestsCreateUserAPIKeyRequest {
	return &NullableRequestsCreateUserAPIKeyRequest{value: val, isSet: true}
}

func (v NullableRequestsCreateUserAPIKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestsCreateUserAPIKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


