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

// V1NodeRestrictions struct for V1NodeRestrictions
type V1NodeRestrictions struct {
	// List of all allowed node sizes.
	AllowedSizes []string `json:"allowedSizes,omitempty"`
	// Flag to determine whether downsizing the cluster is enabled during editing.
	DownsizeEnabled *bool `json:"downsizeEnabled,omitempty"`
	// List of all node sizes that are not allowed to change during editing.
	ResizeDisabledSizes []string `json:"resizeDisabledSizes,omitempty"`
}

// NewV1NodeRestrictions instantiates a new V1NodeRestrictions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1NodeRestrictions() *V1NodeRestrictions {
	this := V1NodeRestrictions{}
	return &this
}

// NewV1NodeRestrictionsWithDefaults instantiates a new V1NodeRestrictions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1NodeRestrictionsWithDefaults() *V1NodeRestrictions {
	this := V1NodeRestrictions{}
	return &this
}

// GetAllowedSizes returns the AllowedSizes field value if set, zero value otherwise.
func (o *V1NodeRestrictions) GetAllowedSizes() []string {
	if o == nil || o.AllowedSizes == nil {
		var ret []string
		return ret
	}
	return o.AllowedSizes
}

// GetAllowedSizesOk returns a tuple with the AllowedSizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NodeRestrictions) GetAllowedSizesOk() ([]string, bool) {
	if o == nil || o.AllowedSizes == nil {
		return nil, false
	}
	return o.AllowedSizes, true
}

// HasAllowedSizes returns a boolean if a field has been set.
func (o *V1NodeRestrictions) HasAllowedSizes() bool {
	if o != nil && o.AllowedSizes != nil {
		return true
	}

	return false
}

// SetAllowedSizes gets a reference to the given []string and assigns it to the AllowedSizes field.
func (o *V1NodeRestrictions) SetAllowedSizes(v []string) {
	o.AllowedSizes = v
}

// GetDownsizeEnabled returns the DownsizeEnabled field value if set, zero value otherwise.
func (o *V1NodeRestrictions) GetDownsizeEnabled() bool {
	if o == nil || o.DownsizeEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DownsizeEnabled
}

// GetDownsizeEnabledOk returns a tuple with the DownsizeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NodeRestrictions) GetDownsizeEnabledOk() (*bool, bool) {
	if o == nil || o.DownsizeEnabled == nil {
		return nil, false
	}
	return o.DownsizeEnabled, true
}

// HasDownsizeEnabled returns a boolean if a field has been set.
func (o *V1NodeRestrictions) HasDownsizeEnabled() bool {
	if o != nil && o.DownsizeEnabled != nil {
		return true
	}

	return false
}

// SetDownsizeEnabled gets a reference to the given bool and assigns it to the DownsizeEnabled field.
func (o *V1NodeRestrictions) SetDownsizeEnabled(v bool) {
	o.DownsizeEnabled = &v
}

// GetResizeDisabledSizes returns the ResizeDisabledSizes field value if set, zero value otherwise.
func (o *V1NodeRestrictions) GetResizeDisabledSizes() []string {
	if o == nil || o.ResizeDisabledSizes == nil {
		var ret []string
		return ret
	}
	return o.ResizeDisabledSizes
}

// GetResizeDisabledSizesOk returns a tuple with the ResizeDisabledSizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NodeRestrictions) GetResizeDisabledSizesOk() ([]string, bool) {
	if o == nil || o.ResizeDisabledSizes == nil {
		return nil, false
	}
	return o.ResizeDisabledSizes, true
}

// HasResizeDisabledSizes returns a boolean if a field has been set.
func (o *V1NodeRestrictions) HasResizeDisabledSizes() bool {
	if o != nil && o.ResizeDisabledSizes != nil {
		return true
	}

	return false
}

// SetResizeDisabledSizes gets a reference to the given []string and assigns it to the ResizeDisabledSizes field.
func (o *V1NodeRestrictions) SetResizeDisabledSizes(v []string) {
	o.ResizeDisabledSizes = v
}

func (o V1NodeRestrictions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedSizes != nil {
		toSerialize["allowedSizes"] = o.AllowedSizes
	}
	if o.DownsizeEnabled != nil {
		toSerialize["downsizeEnabled"] = o.DownsizeEnabled
	}
	if o.ResizeDisabledSizes != nil {
		toSerialize["resizeDisabledSizes"] = o.ResizeDisabledSizes
	}
	return json.Marshal(toSerialize)
}

type NullableV1NodeRestrictions struct {
	value *V1NodeRestrictions
	isSet bool
}

func (v NullableV1NodeRestrictions) Get() *V1NodeRestrictions {
	return v.value
}

func (v *NullableV1NodeRestrictions) Set(val *V1NodeRestrictions) {
	v.value = val
	v.isSet = true
}

func (v NullableV1NodeRestrictions) IsSet() bool {
	return v.isSet
}

func (v *NullableV1NodeRestrictions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1NodeRestrictions(val *V1NodeRestrictions) *NullableV1NodeRestrictions {
	return &NullableV1NodeRestrictions{value: val, isSet: true}
}

func (v NullableV1NodeRestrictions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1NodeRestrictions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


