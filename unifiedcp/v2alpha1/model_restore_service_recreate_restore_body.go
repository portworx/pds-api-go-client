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

// RestoreServiceRecreateRestoreBody struct for RestoreServiceRecreateRestoreBody
type RestoreServiceRecreateRestoreBody struct {
	TargetClusterId *string `json:"targetClusterId,omitempty"`
	Name *string `json:"name,omitempty"`
	NamespaceId *string `json:"namespaceId,omitempty"`
}

// NewRestoreServiceRecreateRestoreBody instantiates a new RestoreServiceRecreateRestoreBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreServiceRecreateRestoreBody() *RestoreServiceRecreateRestoreBody {
	this := RestoreServiceRecreateRestoreBody{}
	return &this
}

// NewRestoreServiceRecreateRestoreBodyWithDefaults instantiates a new RestoreServiceRecreateRestoreBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreServiceRecreateRestoreBodyWithDefaults() *RestoreServiceRecreateRestoreBody {
	this := RestoreServiceRecreateRestoreBody{}
	return &this
}

// GetTargetClusterId returns the TargetClusterId field value if set, zero value otherwise.
func (o *RestoreServiceRecreateRestoreBody) GetTargetClusterId() string {
	if o == nil || o.TargetClusterId == nil {
		var ret string
		return ret
	}
	return *o.TargetClusterId
}

// GetTargetClusterIdOk returns a tuple with the TargetClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreServiceRecreateRestoreBody) GetTargetClusterIdOk() (*string, bool) {
	if o == nil || o.TargetClusterId == nil {
		return nil, false
	}
	return o.TargetClusterId, true
}

// HasTargetClusterId returns a boolean if a field has been set.
func (o *RestoreServiceRecreateRestoreBody) HasTargetClusterId() bool {
	if o != nil && o.TargetClusterId != nil {
		return true
	}

	return false
}

// SetTargetClusterId gets a reference to the given string and assigns it to the TargetClusterId field.
func (o *RestoreServiceRecreateRestoreBody) SetTargetClusterId(v string) {
	o.TargetClusterId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RestoreServiceRecreateRestoreBody) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreServiceRecreateRestoreBody) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RestoreServiceRecreateRestoreBody) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RestoreServiceRecreateRestoreBody) SetName(v string) {
	o.Name = &v
}

// GetNamespaceId returns the NamespaceId field value if set, zero value otherwise.
func (o *RestoreServiceRecreateRestoreBody) GetNamespaceId() string {
	if o == nil || o.NamespaceId == nil {
		var ret string
		return ret
	}
	return *o.NamespaceId
}

// GetNamespaceIdOk returns a tuple with the NamespaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreServiceRecreateRestoreBody) GetNamespaceIdOk() (*string, bool) {
	if o == nil || o.NamespaceId == nil {
		return nil, false
	}
	return o.NamespaceId, true
}

// HasNamespaceId returns a boolean if a field has been set.
func (o *RestoreServiceRecreateRestoreBody) HasNamespaceId() bool {
	if o != nil && o.NamespaceId != nil {
		return true
	}

	return false
}

// SetNamespaceId gets a reference to the given string and assigns it to the NamespaceId field.
func (o *RestoreServiceRecreateRestoreBody) SetNamespaceId(v string) {
	o.NamespaceId = &v
}

func (o RestoreServiceRecreateRestoreBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TargetClusterId != nil {
		toSerialize["targetClusterId"] = o.TargetClusterId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NamespaceId != nil {
		toSerialize["namespaceId"] = o.NamespaceId
	}
	return json.Marshal(toSerialize)
}

type NullableRestoreServiceRecreateRestoreBody struct {
	value *RestoreServiceRecreateRestoreBody
	isSet bool
}

func (v NullableRestoreServiceRecreateRestoreBody) Get() *RestoreServiceRecreateRestoreBody {
	return v.value
}

func (v *NullableRestoreServiceRecreateRestoreBody) Set(val *RestoreServiceRecreateRestoreBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreServiceRecreateRestoreBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreServiceRecreateRestoreBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreServiceRecreateRestoreBody(val *RestoreServiceRecreateRestoreBody) *NullableRestoreServiceRecreateRestoreBody {
	return &NullableRestoreServiceRecreateRestoreBody{value: val, isSet: true}
}

func (v NullableRestoreServiceRecreateRestoreBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreServiceRecreateRestoreBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


