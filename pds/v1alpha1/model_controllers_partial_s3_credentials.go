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

// checks if the ControllersPartialS3Credentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControllersPartialS3Credentials{}

// ControllersPartialS3Credentials struct for ControllersPartialS3Credentials
type ControllersPartialS3Credentials struct {
	// Access key for the AWS IAM user.
	AccessKey *string `json:"access_key,omitempty"`
	// Endpoint of S3 storage service.
	Endpoint *string `json:"endpoint,omitempty"`
}

// NewControllersPartialS3Credentials instantiates a new ControllersPartialS3Credentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersPartialS3Credentials() *ControllersPartialS3Credentials {
	this := ControllersPartialS3Credentials{}
	return &this
}

// NewControllersPartialS3CredentialsWithDefaults instantiates a new ControllersPartialS3Credentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersPartialS3CredentialsWithDefaults() *ControllersPartialS3Credentials {
	this := ControllersPartialS3Credentials{}
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *ControllersPartialS3Credentials) GetAccessKey() string {
	if o == nil || IsNil(o.AccessKey) {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersPartialS3Credentials) GetAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKey) {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *ControllersPartialS3Credentials) HasAccessKey() bool {
	if o != nil && !IsNil(o.AccessKey) {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *ControllersPartialS3Credentials) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *ControllersPartialS3Credentials) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersPartialS3Credentials) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *ControllersPartialS3Credentials) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *ControllersPartialS3Credentials) SetEndpoint(v string) {
	o.Endpoint = &v
}

func (o ControllersPartialS3Credentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControllersPartialS3Credentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessKey) {
		toSerialize["access_key"] = o.AccessKey
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	return toSerialize, nil
}

type NullableControllersPartialS3Credentials struct {
	value *ControllersPartialS3Credentials
	isSet bool
}

func (v NullableControllersPartialS3Credentials) Get() *ControllersPartialS3Credentials {
	return v.value
}

func (v *NullableControllersPartialS3Credentials) Set(val *ControllersPartialS3Credentials) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersPartialS3Credentials) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersPartialS3Credentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersPartialS3Credentials(val *ControllersPartialS3Credentials) *NullableControllersPartialS3Credentials {
	return &NullableControllersPartialS3Credentials{value: val, isSet: true}
}

func (v NullableControllersPartialS3Credentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersPartialS3Credentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


