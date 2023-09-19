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

// ControllersGenerateTokenRequest struct for ControllersGenerateTokenRequest
type ControllersGenerateTokenRequest struct {
	ClientId *string `json:"client_id,omitempty"`
	ClientToken *string `json:"client_token,omitempty"`
}

// NewControllersGenerateTokenRequest instantiates a new ControllersGenerateTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersGenerateTokenRequest() *ControllersGenerateTokenRequest {
	this := ControllersGenerateTokenRequest{}
	return &this
}

// NewControllersGenerateTokenRequestWithDefaults instantiates a new ControllersGenerateTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersGenerateTokenRequestWithDefaults() *ControllersGenerateTokenRequest {
	this := ControllersGenerateTokenRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ControllersGenerateTokenRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersGenerateTokenRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ControllersGenerateTokenRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ControllersGenerateTokenRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientToken returns the ClientToken field value if set, zero value otherwise.
func (o *ControllersGenerateTokenRequest) GetClientToken() string {
	if o == nil || o.ClientToken == nil {
		var ret string
		return ret
	}
	return *o.ClientToken
}

// GetClientTokenOk returns a tuple with the ClientToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersGenerateTokenRequest) GetClientTokenOk() (*string, bool) {
	if o == nil || o.ClientToken == nil {
		return nil, false
	}
	return o.ClientToken, true
}

// HasClientToken returns a boolean if a field has been set.
func (o *ControllersGenerateTokenRequest) HasClientToken() bool {
	if o != nil && o.ClientToken != nil {
		return true
	}

	return false
}

// SetClientToken gets a reference to the given string and assigns it to the ClientToken field.
func (o *ControllersGenerateTokenRequest) SetClientToken(v string) {
	o.ClientToken = &v
}

func (o ControllersGenerateTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientToken != nil {
		toSerialize["client_token"] = o.ClientToken
	}
	return json.Marshal(toSerialize)
}

type NullableControllersGenerateTokenRequest struct {
	value *ControllersGenerateTokenRequest
	isSet bool
}

func (v NullableControllersGenerateTokenRequest) Get() *ControllersGenerateTokenRequest {
	return v.value
}

func (v *NullableControllersGenerateTokenRequest) Set(val *ControllersGenerateTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersGenerateTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersGenerateTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersGenerateTokenRequest(val *ControllersGenerateTokenRequest) *NullableControllersGenerateTokenRequest {
	return &NullableControllersGenerateTokenRequest{value: val, isSet: true}
}

func (v NullableControllersGenerateTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersGenerateTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


