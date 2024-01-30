/*
public/portworx/pds/images/apiv1/images.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package images

import (
	"encoding/json"
)

// checks if the ProtobufAny type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProtobufAny{}

// ProtobufAny struct for ProtobufAny
type ProtobufAny struct {
	Type *string `json:"@type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProtobufAny ProtobufAny

// NewProtobufAny instantiates a new ProtobufAny object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtobufAny() *ProtobufAny {
	this := ProtobufAny{}
	return &this
}

// NewProtobufAnyWithDefaults instantiates a new ProtobufAny object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtobufAnyWithDefaults() *ProtobufAny {
	this := ProtobufAny{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProtobufAny) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtobufAny) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProtobufAny) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProtobufAny) SetType(v string) {
	o.Type = &v
}

func (o ProtobufAny) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProtobufAny) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["@type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProtobufAny) UnmarshalJSON(data []byte) (err error) {
	varProtobufAny := _ProtobufAny{}

	err = json.Unmarshal(data, &varProtobufAny)

	if err != nil {
		return err
	}

	*o = ProtobufAny(varProtobufAny)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "@type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProtobufAny struct {
	value *ProtobufAny
	isSet bool
}

func (v NullableProtobufAny) Get() *ProtobufAny {
	return v.value
}

func (v *NullableProtobufAny) Set(val *ProtobufAny) {
	v.value = val
	v.isSet = true
}

func (v NullableProtobufAny) IsSet() bool {
	return v.isSet
}

func (v *NullableProtobufAny) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtobufAny(val *ProtobufAny) *NullableProtobufAny {
	return &NullableProtobufAny{value: val, isSet: true}
}

func (v NullableProtobufAny) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtobufAny) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


