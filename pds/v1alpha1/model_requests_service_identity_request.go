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

// RequestsServiceIdentityRequest struct for RequestsServiceIdentityRequest
type RequestsServiceIdentityRequest struct {
	Description *string `json:"description,omitempty"`
	Name string `json:"name"`
}

// NewRequestsServiceIdentityRequest instantiates a new RequestsServiceIdentityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestsServiceIdentityRequest(name string) *RequestsServiceIdentityRequest {
	this := RequestsServiceIdentityRequest{}
	this.Name = name
	return &this
}

// NewRequestsServiceIdentityRequestWithDefaults instantiates a new RequestsServiceIdentityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestsServiceIdentityRequestWithDefaults() *RequestsServiceIdentityRequest {
	this := RequestsServiceIdentityRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RequestsServiceIdentityRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsServiceIdentityRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestsServiceIdentityRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RequestsServiceIdentityRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *RequestsServiceIdentityRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RequestsServiceIdentityRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RequestsServiceIdentityRequest) SetName(v string) {
	o.Name = v
}

func (o RequestsServiceIdentityRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableRequestsServiceIdentityRequest struct {
	value *RequestsServiceIdentityRequest
	isSet bool
}

func (v NullableRequestsServiceIdentityRequest) Get() *RequestsServiceIdentityRequest {
	return v.value
}

func (v *NullableRequestsServiceIdentityRequest) Set(val *RequestsServiceIdentityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestsServiceIdentityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestsServiceIdentityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestsServiceIdentityRequest(val *RequestsServiceIdentityRequest) *NullableRequestsServiceIdentityRequest {
	return &NullableRequestsServiceIdentityRequest{value: val, isSet: true}
}

func (v NullableRequestsServiceIdentityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestsServiceIdentityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

