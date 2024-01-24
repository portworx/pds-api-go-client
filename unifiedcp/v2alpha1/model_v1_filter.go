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

// V1Filter Filter clubs the infra and application resource details, to filter the requested list of projects.
type V1Filter struct {
	AppResource *V1ApplicationResourceType `json:"appResource,omitempty"`
	InfraResource *V1InfraResourceType `json:"infraResource,omitempty"`
	// ID of the resource for which projects to be listed.
	ResourceId *string `json:"resourceId,omitempty"`
}

// NewV1Filter instantiates a new V1Filter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Filter() *V1Filter {
	this := V1Filter{}
	var appResource V1ApplicationResourceType = TYPE_UNSPECIFIED
	this.AppResource = &appResource
	var infraResource V1InfraResourceType = TYPE_UNSPECIFIED
	this.InfraResource = &infraResource
	return &this
}

// NewV1FilterWithDefaults instantiates a new V1Filter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1FilterWithDefaults() *V1Filter {
	this := V1Filter{}
	var appResource V1ApplicationResourceType = TYPE_UNSPECIFIED
	this.AppResource = &appResource
	var infraResource V1InfraResourceType = TYPE_UNSPECIFIED
	this.InfraResource = &infraResource
	return &this
}

// GetAppResource returns the AppResource field value if set, zero value otherwise.
func (o *V1Filter) GetAppResource() V1ApplicationResourceType {
	if o == nil || o.AppResource == nil {
		var ret V1ApplicationResourceType
		return ret
	}
	return *o.AppResource
}

// GetAppResourceOk returns a tuple with the AppResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Filter) GetAppResourceOk() (*V1ApplicationResourceType, bool) {
	if o == nil || o.AppResource == nil {
		return nil, false
	}
	return o.AppResource, true
}

// HasAppResource returns a boolean if a field has been set.
func (o *V1Filter) HasAppResource() bool {
	if o != nil && o.AppResource != nil {
		return true
	}

	return false
}

// SetAppResource gets a reference to the given V1ApplicationResourceType and assigns it to the AppResource field.
func (o *V1Filter) SetAppResource(v V1ApplicationResourceType) {
	o.AppResource = &v
}

// GetInfraResource returns the InfraResource field value if set, zero value otherwise.
func (o *V1Filter) GetInfraResource() V1InfraResourceType {
	if o == nil || o.InfraResource == nil {
		var ret V1InfraResourceType
		return ret
	}
	return *o.InfraResource
}

// GetInfraResourceOk returns a tuple with the InfraResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Filter) GetInfraResourceOk() (*V1InfraResourceType, bool) {
	if o == nil || o.InfraResource == nil {
		return nil, false
	}
	return o.InfraResource, true
}

// HasInfraResource returns a boolean if a field has been set.
func (o *V1Filter) HasInfraResource() bool {
	if o != nil && o.InfraResource != nil {
		return true
	}

	return false
}

// SetInfraResource gets a reference to the given V1InfraResourceType and assigns it to the InfraResource field.
func (o *V1Filter) SetInfraResource(v V1InfraResourceType) {
	o.InfraResource = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *V1Filter) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Filter) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *V1Filter) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *V1Filter) SetResourceId(v string) {
	o.ResourceId = &v
}

func (o V1Filter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppResource != nil {
		toSerialize["appResource"] = o.AppResource
	}
	if o.InfraResource != nil {
		toSerialize["infraResource"] = o.InfraResource
	}
	if o.ResourceId != nil {
		toSerialize["resourceId"] = o.ResourceId
	}
	return json.Marshal(toSerialize)
}

type NullableV1Filter struct {
	value *V1Filter
	isSet bool
}

func (v NullableV1Filter) Get() *V1Filter {
	return v.value
}

func (v *NullableV1Filter) Set(val *V1Filter) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Filter) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Filter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Filter(val *V1Filter) *NullableV1Filter {
	return &NullableV1Filter{value: val, isSet: true}
}

func (v NullableV1Filter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Filter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


