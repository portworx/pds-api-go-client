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

// ControllersPartialAzureCredentials struct for ControllersPartialAzureCredentials
type ControllersPartialAzureCredentials struct {
	// Name of the Azure Storage account.
	AccountName *string `json:"account_name,omitempty"`
}

// NewControllersPartialAzureCredentials instantiates a new ControllersPartialAzureCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersPartialAzureCredentials() *ControllersPartialAzureCredentials {
	this := ControllersPartialAzureCredentials{}
	return &this
}

// NewControllersPartialAzureCredentialsWithDefaults instantiates a new ControllersPartialAzureCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersPartialAzureCredentialsWithDefaults() *ControllersPartialAzureCredentials {
	this := ControllersPartialAzureCredentials{}
	return &this
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *ControllersPartialAzureCredentials) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersPartialAzureCredentials) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *ControllersPartialAzureCredentials) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *ControllersPartialAzureCredentials) SetAccountName(v string) {
	o.AccountName = &v
}

func (o ControllersPartialAzureCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountName != nil {
		toSerialize["account_name"] = o.AccountName
	}
	return json.Marshal(toSerialize)
}

type NullableControllersPartialAzureCredentials struct {
	value *ControllersPartialAzureCredentials
	isSet bool
}

func (v NullableControllersPartialAzureCredentials) Get() *ControllersPartialAzureCredentials {
	return v.value
}

func (v *NullableControllersPartialAzureCredentials) Set(val *ControllersPartialAzureCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersPartialAzureCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersPartialAzureCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersPartialAzureCredentials(val *ControllersPartialAzureCredentials) *NullableControllersPartialAzureCredentials {
	return &NullableControllersPartialAzureCredentials{value: val, isSet: true}
}

func (v NullableControllersPartialAzureCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersPartialAzureCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


