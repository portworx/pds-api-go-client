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

// checks if the ControllersPartialS3CompatibleCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControllersPartialS3CompatibleCredentials{}

// ControllersPartialS3CompatibleCredentials struct for ControllersPartialS3CompatibleCredentials
type ControllersPartialS3CompatibleCredentials struct {
	// Access key for S3-compatible user.
	AccessKey *string `json:"access_key,omitempty"`
	// Endpoint of S3-compatible storage service.
	Endpoint *string `json:"endpoint,omitempty"`
}

// NewControllersPartialS3CompatibleCredentials instantiates a new ControllersPartialS3CompatibleCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersPartialS3CompatibleCredentials() *ControllersPartialS3CompatibleCredentials {
	this := ControllersPartialS3CompatibleCredentials{}
	return &this
}

// NewControllersPartialS3CompatibleCredentialsWithDefaults instantiates a new ControllersPartialS3CompatibleCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersPartialS3CompatibleCredentialsWithDefaults() *ControllersPartialS3CompatibleCredentials {
	this := ControllersPartialS3CompatibleCredentials{}
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *ControllersPartialS3CompatibleCredentials) GetAccessKey() string {
	if o == nil || IsNil(o.AccessKey) {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersPartialS3CompatibleCredentials) GetAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKey) {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *ControllersPartialS3CompatibleCredentials) HasAccessKey() bool {
	if o != nil && !IsNil(o.AccessKey) {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *ControllersPartialS3CompatibleCredentials) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *ControllersPartialS3CompatibleCredentials) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersPartialS3CompatibleCredentials) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *ControllersPartialS3CompatibleCredentials) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *ControllersPartialS3CompatibleCredentials) SetEndpoint(v string) {
	o.Endpoint = &v
}

func (o ControllersPartialS3CompatibleCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControllersPartialS3CompatibleCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessKey) {
		toSerialize["access_key"] = o.AccessKey
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	return toSerialize, nil
}

type NullableControllersPartialS3CompatibleCredentials struct {
	value *ControllersPartialS3CompatibleCredentials
	isSet bool
}

func (v NullableControllersPartialS3CompatibleCredentials) Get() *ControllersPartialS3CompatibleCredentials {
	return v.value
}

func (v *NullableControllersPartialS3CompatibleCredentials) Set(val *ControllersPartialS3CompatibleCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersPartialS3CompatibleCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersPartialS3CompatibleCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersPartialS3CompatibleCredentials(val *ControllersPartialS3CompatibleCredentials) *NullableControllersPartialS3CompatibleCredentials {
	return &NullableControllersPartialS3CompatibleCredentials{value: val, isSet: true}
}

func (v NullableControllersPartialS3CompatibleCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersPartialS3CompatibleCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


