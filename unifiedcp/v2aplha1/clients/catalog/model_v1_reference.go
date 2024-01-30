/*
public/portworx/pds/catalog/dataservices/apiv1/dataservices.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package catalog

import (
	"encoding/json"
)

// checks if the V1Reference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1Reference{}

// V1Reference Reference identifies the resource type, version of the uid and the resource.
type V1Reference struct {
	// API group of the resource.
	Type *string `json:"type,omitempty"`
	// Version of the API.
	Version *string `json:"version,omitempty"`
	// UID of the resource.
	Uid *string `json:"uid,omitempty"`
}

// NewV1Reference instantiates a new V1Reference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Reference() *V1Reference {
	this := V1Reference{}
	return &this
}

// NewV1ReferenceWithDefaults instantiates a new V1Reference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ReferenceWithDefaults() *V1Reference {
	this := V1Reference{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1Reference) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Reference) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1Reference) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V1Reference) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *V1Reference) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Reference) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *V1Reference) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *V1Reference) SetVersion(v string) {
	o.Version = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *V1Reference) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Reference) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *V1Reference) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *V1Reference) SetUid(v string) {
	o.Uid = &v
}

func (o V1Reference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1Reference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	return toSerialize, nil
}

type NullableV1Reference struct {
	value *V1Reference
	isSet bool
}

func (v NullableV1Reference) Get() *V1Reference {
	return v.value
}

func (v *NullableV1Reference) Set(val *V1Reference) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Reference) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Reference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Reference(val *V1Reference) *NullableV1Reference {
	return &NullableV1Reference{value: val, isSet: true}
}

func (v NullableV1Reference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Reference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


