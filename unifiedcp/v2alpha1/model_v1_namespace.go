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

// V1Namespace A managed Kubernetes namespace.
type V1Namespace struct {
	Meta *V1Meta `json:"meta,omitempty"`
	Status *V1NamespaceStatus `json:"status,omitempty"`
}

// NewV1Namespace instantiates a new V1Namespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Namespace() *V1Namespace {
	this := V1Namespace{}
	return &this
}

// NewV1NamespaceWithDefaults instantiates a new V1Namespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1NamespaceWithDefaults() *V1Namespace {
	this := V1Namespace{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V1Namespace) GetMeta() V1Meta {
	if o == nil || o.Meta == nil {
		var ret V1Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Namespace) GetMetaOk() (*V1Meta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V1Namespace) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V1Meta and assigns it to the Meta field.
func (o *V1Namespace) SetMeta(v V1Meta) {
	o.Meta = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V1Namespace) GetStatus() V1NamespaceStatus {
	if o == nil || o.Status == nil {
		var ret V1NamespaceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Namespace) GetStatusOk() (*V1NamespaceStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V1Namespace) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given V1NamespaceStatus and assigns it to the Status field.
func (o *V1Namespace) SetStatus(v V1NamespaceStatus) {
	o.Status = &v
}

func (o V1Namespace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableV1Namespace struct {
	value *V1Namespace
	isSet bool
}

func (v NullableV1Namespace) Get() *V1Namespace {
	return v.value
}

func (v *NullableV1Namespace) Set(val *V1Namespace) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Namespace) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Namespace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Namespace(val *V1Namespace) *NullableV1Namespace {
	return &NullableV1Namespace{value: val, isSet: true}
}

func (v NullableV1Namespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Namespace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


