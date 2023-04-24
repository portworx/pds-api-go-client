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

// checks if the ModelsAuthorizerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsAuthorizerResponse{}

// ModelsAuthorizerResponse struct for ModelsAuthorizerResponse
type ModelsAuthorizerResponse struct {
	Allow *bool `json:"allow,omitempty"`
}

// NewModelsAuthorizerResponse instantiates a new ModelsAuthorizerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsAuthorizerResponse() *ModelsAuthorizerResponse {
	this := ModelsAuthorizerResponse{}
	return &this
}

// NewModelsAuthorizerResponseWithDefaults instantiates a new ModelsAuthorizerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsAuthorizerResponseWithDefaults() *ModelsAuthorizerResponse {
	this := ModelsAuthorizerResponse{}
	return &this
}

// GetAllow returns the Allow field value if set, zero value otherwise.
func (o *ModelsAuthorizerResponse) GetAllow() bool {
	if o == nil || IsNil(o.Allow) {
		var ret bool
		return ret
	}
	return *o.Allow
}

// GetAllowOk returns a tuple with the Allow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAuthorizerResponse) GetAllowOk() (*bool, bool) {
	if o == nil || IsNil(o.Allow) {
		return nil, false
	}
	return o.Allow, true
}

// HasAllow returns a boolean if a field has been set.
func (o *ModelsAuthorizerResponse) HasAllow() bool {
	if o != nil && !IsNil(o.Allow) {
		return true
	}

	return false
}

// SetAllow gets a reference to the given bool and assigns it to the Allow field.
func (o *ModelsAuthorizerResponse) SetAllow(v bool) {
	o.Allow = &v
}

func (o ModelsAuthorizerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsAuthorizerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Allow) {
		toSerialize["allow"] = o.Allow
	}
	return toSerialize, nil
}

type NullableModelsAuthorizerResponse struct {
	value *ModelsAuthorizerResponse
	isSet bool
}

func (v NullableModelsAuthorizerResponse) Get() *ModelsAuthorizerResponse {
	return v.value
}

func (v *NullableModelsAuthorizerResponse) Set(val *ModelsAuthorizerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsAuthorizerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsAuthorizerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsAuthorizerResponse(val *ModelsAuthorizerResponse) *NullableModelsAuthorizerResponse {
	return &NullableModelsAuthorizerResponse{value: val, isSet: true}
}

func (v NullableModelsAuthorizerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsAuthorizerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


