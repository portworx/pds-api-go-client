/*
PDS API

Portworx Data Services API Server

API version: 3.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"encoding/json"
)

// ControllersCreateBackupCredentialsRequest struct for ControllersCreateBackupCredentialsRequest
type ControllersCreateBackupCredentialsRequest struct {
	Credentials *ControllersCredentials `json:"credentials,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewControllersCreateBackupCredentialsRequest instantiates a new ControllersCreateBackupCredentialsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersCreateBackupCredentialsRequest() *ControllersCreateBackupCredentialsRequest {
	this := ControllersCreateBackupCredentialsRequest{}
	return &this
}

// NewControllersCreateBackupCredentialsRequestWithDefaults instantiates a new ControllersCreateBackupCredentialsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersCreateBackupCredentialsRequestWithDefaults() *ControllersCreateBackupCredentialsRequest {
	this := ControllersCreateBackupCredentialsRequest{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *ControllersCreateBackupCredentialsRequest) GetCredentials() ControllersCredentials {
	if o == nil || o.Credentials == nil {
		var ret ControllersCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateBackupCredentialsRequest) GetCredentialsOk() (*ControllersCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *ControllersCreateBackupCredentialsRequest) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given ControllersCredentials and assigns it to the Credentials field.
func (o *ControllersCreateBackupCredentialsRequest) SetCredentials(v ControllersCredentials) {
	o.Credentials = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ControllersCreateBackupCredentialsRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateBackupCredentialsRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ControllersCreateBackupCredentialsRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ControllersCreateBackupCredentialsRequest) SetName(v string) {
	o.Name = &v
}

func (o ControllersCreateBackupCredentialsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableControllersCreateBackupCredentialsRequest struct {
	value *ControllersCreateBackupCredentialsRequest
	isSet bool
}

func (v NullableControllersCreateBackupCredentialsRequest) Get() *ControllersCreateBackupCredentialsRequest {
	return v.value
}

func (v *NullableControllersCreateBackupCredentialsRequest) Set(val *ControllersCreateBackupCredentialsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersCreateBackupCredentialsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersCreateBackupCredentialsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersCreateBackupCredentialsRequest(val *ControllersCreateBackupCredentialsRequest) *NullableControllersCreateBackupCredentialsRequest {
	return &NullableControllersCreateBackupCredentialsRequest{value: val, isSet: true}
}

func (v NullableControllersCreateBackupCredentialsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersCreateBackupCredentialsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


