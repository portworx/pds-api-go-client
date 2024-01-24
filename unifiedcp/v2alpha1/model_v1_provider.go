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

// V1Provider struct for V1Provider
type V1Provider struct {
	CloudProvider *V1ProviderType `json:"cloudProvider,omitempty"`
}

// NewV1Provider instantiates a new V1Provider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Provider() *V1Provider {
	this := V1Provider{}
	var cloudProvider V1ProviderType = TYPE_UNSPECIFIED
	this.CloudProvider = &cloudProvider
	return &this
}

// NewV1ProviderWithDefaults instantiates a new V1Provider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ProviderWithDefaults() *V1Provider {
	this := V1Provider{}
	var cloudProvider V1ProviderType = TYPE_UNSPECIFIED
	this.CloudProvider = &cloudProvider
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *V1Provider) GetCloudProvider() V1ProviderType {
	if o == nil || o.CloudProvider == nil {
		var ret V1ProviderType
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Provider) GetCloudProviderOk() (*V1ProviderType, bool) {
	if o == nil || o.CloudProvider == nil {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *V1Provider) HasCloudProvider() bool {
	if o != nil && o.CloudProvider != nil {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given V1ProviderType and assigns it to the CloudProvider field.
func (o *V1Provider) SetCloudProvider(v V1ProviderType) {
	o.CloudProvider = &v
}

func (o V1Provider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudProvider != nil {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	return json.Marshal(toSerialize)
}

type NullableV1Provider struct {
	value *V1Provider
	isSet bool
}

func (v NullableV1Provider) Get() *V1Provider {
	return v.value
}

func (v *NullableV1Provider) Set(val *V1Provider) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Provider) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Provider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Provider(val *V1Provider) *NullableV1Provider {
	return &NullableV1Provider{value: val, isSet: true}
}

func (v NullableV1Provider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Provider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


