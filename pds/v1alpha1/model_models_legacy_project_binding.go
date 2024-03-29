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

// ModelsLegacyProjectBinding struct for ModelsLegacyProjectBinding
type ModelsLegacyProjectBinding struct {
	ActorId *string `json:"actor_id,omitempty"`
	ActorType *string `json:"actor_type,omitempty"`
	ProjectId *string `json:"project_id,omitempty"`
	RoleName *string `json:"role_name,omitempty"`
}

// NewModelsLegacyProjectBinding instantiates a new ModelsLegacyProjectBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsLegacyProjectBinding() *ModelsLegacyProjectBinding {
	this := ModelsLegacyProjectBinding{}
	return &this
}

// NewModelsLegacyProjectBindingWithDefaults instantiates a new ModelsLegacyProjectBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsLegacyProjectBindingWithDefaults() *ModelsLegacyProjectBinding {
	this := ModelsLegacyProjectBinding{}
	return &this
}

// GetActorId returns the ActorId field value if set, zero value otherwise.
func (o *ModelsLegacyProjectBinding) GetActorId() string {
	if o == nil || o.ActorId == nil {
		var ret string
		return ret
	}
	return *o.ActorId
}

// GetActorIdOk returns a tuple with the ActorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsLegacyProjectBinding) GetActorIdOk() (*string, bool) {
	if o == nil || o.ActorId == nil {
		return nil, false
	}
	return o.ActorId, true
}

// HasActorId returns a boolean if a field has been set.
func (o *ModelsLegacyProjectBinding) HasActorId() bool {
	if o != nil && o.ActorId != nil {
		return true
	}

	return false
}

// SetActorId gets a reference to the given string and assigns it to the ActorId field.
func (o *ModelsLegacyProjectBinding) SetActorId(v string) {
	o.ActorId = &v
}

// GetActorType returns the ActorType field value if set, zero value otherwise.
func (o *ModelsLegacyProjectBinding) GetActorType() string {
	if o == nil || o.ActorType == nil {
		var ret string
		return ret
	}
	return *o.ActorType
}

// GetActorTypeOk returns a tuple with the ActorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsLegacyProjectBinding) GetActorTypeOk() (*string, bool) {
	if o == nil || o.ActorType == nil {
		return nil, false
	}
	return o.ActorType, true
}

// HasActorType returns a boolean if a field has been set.
func (o *ModelsLegacyProjectBinding) HasActorType() bool {
	if o != nil && o.ActorType != nil {
		return true
	}

	return false
}

// SetActorType gets a reference to the given string and assigns it to the ActorType field.
func (o *ModelsLegacyProjectBinding) SetActorType(v string) {
	o.ActorType = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ModelsLegacyProjectBinding) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsLegacyProjectBinding) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ModelsLegacyProjectBinding) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *ModelsLegacyProjectBinding) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *ModelsLegacyProjectBinding) GetRoleName() string {
	if o == nil || o.RoleName == nil {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsLegacyProjectBinding) GetRoleNameOk() (*string, bool) {
	if o == nil || o.RoleName == nil {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *ModelsLegacyProjectBinding) HasRoleName() bool {
	if o != nil && o.RoleName != nil {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *ModelsLegacyProjectBinding) SetRoleName(v string) {
	o.RoleName = &v
}

func (o ModelsLegacyProjectBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActorId != nil {
		toSerialize["actor_id"] = o.ActorId
	}
	if o.ActorType != nil {
		toSerialize["actor_type"] = o.ActorType
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.RoleName != nil {
		toSerialize["role_name"] = o.RoleName
	}
	return json.Marshal(toSerialize)
}

type NullableModelsLegacyProjectBinding struct {
	value *ModelsLegacyProjectBinding
	isSet bool
}

func (v NullableModelsLegacyProjectBinding) Get() *ModelsLegacyProjectBinding {
	return v.value
}

func (v *NullableModelsLegacyProjectBinding) Set(val *ModelsLegacyProjectBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsLegacyProjectBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsLegacyProjectBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsLegacyProjectBinding(val *ModelsLegacyProjectBinding) *NullableModelsLegacyProjectBinding {
	return &NullableModelsLegacyProjectBinding{value: val, isSet: true}
}

func (v NullableModelsLegacyProjectBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsLegacyProjectBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


