/*
public/portworx/pds/tasks/apiv1/tasks.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pdsclient

import (
	"encoding/json"
)

// checks if the V1References2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1References2{}

// V1References2 References to other resources.
type V1References2 struct {
	// UID of the target cluster in which Data Service will be deployed.
	TargetClusterId *string `json:"targetClusterId,omitempty"`
	// UID of the image to be used for the Data Service Deployment.
	ImageId *string `json:"imageId,omitempty"`
	// UID of the project to which DataService Deployment associated.
	ProjectId *string `json:"projectId,omitempty"`
	// UID of the restore id for the Deployment.
	RestoreId *string `json:"restoreId,omitempty"`
}

// NewV1References2 instantiates a new V1References2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1References2() *V1References2 {
	this := V1References2{}
	return &this
}

// NewV1References2WithDefaults instantiates a new V1References2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1References2WithDefaults() *V1References2 {
	this := V1References2{}
	return &this
}

// GetTargetClusterId returns the TargetClusterId field value if set, zero value otherwise.
func (o *V1References2) GetTargetClusterId() string {
	if o == nil || IsNil(o.TargetClusterId) {
		var ret string
		return ret
	}
	return *o.TargetClusterId
}

// GetTargetClusterIdOk returns a tuple with the TargetClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1References2) GetTargetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetClusterId) {
		return nil, false
	}
	return o.TargetClusterId, true
}

// HasTargetClusterId returns a boolean if a field has been set.
func (o *V1References2) HasTargetClusterId() bool {
	if o != nil && !IsNil(o.TargetClusterId) {
		return true
	}

	return false
}

// SetTargetClusterId gets a reference to the given string and assigns it to the TargetClusterId field.
func (o *V1References2) SetTargetClusterId(v string) {
	o.TargetClusterId = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *V1References2) GetImageId() string {
	if o == nil || IsNil(o.ImageId) {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1References2) GetImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *V1References2) HasImageId() bool {
	if o != nil && !IsNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *V1References2) SetImageId(v string) {
	o.ImageId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *V1References2) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1References2) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *V1References2) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *V1References2) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetRestoreId returns the RestoreId field value if set, zero value otherwise.
func (o *V1References2) GetRestoreId() string {
	if o == nil || IsNil(o.RestoreId) {
		var ret string
		return ret
	}
	return *o.RestoreId
}

// GetRestoreIdOk returns a tuple with the RestoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1References2) GetRestoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.RestoreId) {
		return nil, false
	}
	return o.RestoreId, true
}

// HasRestoreId returns a boolean if a field has been set.
func (o *V1References2) HasRestoreId() bool {
	if o != nil && !IsNil(o.RestoreId) {
		return true
	}

	return false
}

// SetRestoreId gets a reference to the given string and assigns it to the RestoreId field.
func (o *V1References2) SetRestoreId(v string) {
	o.RestoreId = &v
}

func (o V1References2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1References2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetClusterId) {
		toSerialize["targetClusterId"] = o.TargetClusterId
	}
	if !IsNil(o.ImageId) {
		toSerialize["imageId"] = o.ImageId
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.RestoreId) {
		toSerialize["restoreId"] = o.RestoreId
	}
	return toSerialize, nil
}

type NullableV1References2 struct {
	value *V1References2
	isSet bool
}

func (v NullableV1References2) Get() *V1References2 {
	return v.value
}

func (v *NullableV1References2) Set(val *V1References2) {
	v.value = val
	v.isSet = true
}

func (v NullableV1References2) IsSet() bool {
	return v.isSet
}

func (v *NullableV1References2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1References2(val *V1References2) *NullableV1References2 {
	return &NullableV1References2{value: val, isSet: true}
}

func (v NullableV1References2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1References2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


