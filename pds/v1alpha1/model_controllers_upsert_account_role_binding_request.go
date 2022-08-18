/*
PDS API

Portworx Data Services API Server

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"encoding/json"
)

// ControllersUpsertAccountRoleBindingRequest struct for ControllersUpsertAccountRoleBindingRequest
type ControllersUpsertAccountRoleBindingRequest struct {
	ActorId *string `json:"actor_id,omitempty"`
	ActorType *string `json:"actor_type,omitempty"`
	RoleName *string `json:"role_name,omitempty"`
}

// NewControllersUpsertAccountRoleBindingRequest instantiates a new ControllersUpsertAccountRoleBindingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersUpsertAccountRoleBindingRequest() *ControllersUpsertAccountRoleBindingRequest {
	this := ControllersUpsertAccountRoleBindingRequest{}
	return &this
}

// NewControllersUpsertAccountRoleBindingRequestWithDefaults instantiates a new ControllersUpsertAccountRoleBindingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersUpsertAccountRoleBindingRequestWithDefaults() *ControllersUpsertAccountRoleBindingRequest {
	this := ControllersUpsertAccountRoleBindingRequest{}
	return &this
}

// GetActorId returns the ActorId field value if set, zero value otherwise.
func (o *ControllersUpsertAccountRoleBindingRequest) GetActorId() string {
	if o == nil || o.ActorId == nil {
		var ret string
		return ret
	}
	return *o.ActorId
}

// GetActorIdOk returns a tuple with the ActorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersUpsertAccountRoleBindingRequest) GetActorIdOk() (*string, bool) {
	if o == nil || o.ActorId == nil {
		return nil, false
	}
	return o.ActorId, true
}

// HasActorId returns a boolean if a field has been set.
func (o *ControllersUpsertAccountRoleBindingRequest) HasActorId() bool {
	if o != nil && o.ActorId != nil {
		return true
	}

	return false
}

// SetActorId gets a reference to the given string and assigns it to the ActorId field.
func (o *ControllersUpsertAccountRoleBindingRequest) SetActorId(v string) {
	o.ActorId = &v
}

// GetActorType returns the ActorType field value if set, zero value otherwise.
func (o *ControllersUpsertAccountRoleBindingRequest) GetActorType() string {
	if o == nil || o.ActorType == nil {
		var ret string
		return ret
	}
	return *o.ActorType
}

// GetActorTypeOk returns a tuple with the ActorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersUpsertAccountRoleBindingRequest) GetActorTypeOk() (*string, bool) {
	if o == nil || o.ActorType == nil {
		return nil, false
	}
	return o.ActorType, true
}

// HasActorType returns a boolean if a field has been set.
func (o *ControllersUpsertAccountRoleBindingRequest) HasActorType() bool {
	if o != nil && o.ActorType != nil {
		return true
	}

	return false
}

// SetActorType gets a reference to the given string and assigns it to the ActorType field.
func (o *ControllersUpsertAccountRoleBindingRequest) SetActorType(v string) {
	o.ActorType = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *ControllersUpsertAccountRoleBindingRequest) GetRoleName() string {
	if o == nil || o.RoleName == nil {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersUpsertAccountRoleBindingRequest) GetRoleNameOk() (*string, bool) {
	if o == nil || o.RoleName == nil {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *ControllersUpsertAccountRoleBindingRequest) HasRoleName() bool {
	if o != nil && o.RoleName != nil {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *ControllersUpsertAccountRoleBindingRequest) SetRoleName(v string) {
	o.RoleName = &v
}

func (o ControllersUpsertAccountRoleBindingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActorId != nil {
		toSerialize["actor_id"] = o.ActorId
	}
	if o.ActorType != nil {
		toSerialize["actor_type"] = o.ActorType
	}
	if o.RoleName != nil {
		toSerialize["role_name"] = o.RoleName
	}
	return json.Marshal(toSerialize)
}

type NullableControllersUpsertAccountRoleBindingRequest struct {
	value *ControllersUpsertAccountRoleBindingRequest
	isSet bool
}

func (v NullableControllersUpsertAccountRoleBindingRequest) Get() *ControllersUpsertAccountRoleBindingRequest {
	return v.value
}

func (v *NullableControllersUpsertAccountRoleBindingRequest) Set(val *ControllersUpsertAccountRoleBindingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersUpsertAccountRoleBindingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersUpsertAccountRoleBindingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersUpsertAccountRoleBindingRequest(val *ControllersUpsertAccountRoleBindingRequest) *NullableControllersUpsertAccountRoleBindingRequest {
	return &NullableControllersUpsertAccountRoleBindingRequest{value: val, isSet: true}
}

func (v NullableControllersUpsertAccountRoleBindingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersUpsertAccountRoleBindingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


