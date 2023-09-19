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

// ControllersWhoAmIServiceIdentity struct for ControllersWhoAmIServiceIdentity
type ControllersWhoAmIServiceIdentity struct {
	AccountId *string `json:"account_id,omitempty"`
	// CreatedAt is autogenerated on creation
	CreatedAt *string `json:"created_at,omitempty"`
	// ID is auto generated on creation
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// UpdatedAt is autogenerated on update
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewControllersWhoAmIServiceIdentity instantiates a new ControllersWhoAmIServiceIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersWhoAmIServiceIdentity() *ControllersWhoAmIServiceIdentity {
	this := ControllersWhoAmIServiceIdentity{}
	return &this
}

// NewControllersWhoAmIServiceIdentityWithDefaults instantiates a new ControllersWhoAmIServiceIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersWhoAmIServiceIdentityWithDefaults() *ControllersWhoAmIServiceIdentity {
	this := ControllersWhoAmIServiceIdentity{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ControllersWhoAmIServiceIdentity) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersWhoAmIServiceIdentity) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ControllersWhoAmIServiceIdentity) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ControllersWhoAmIServiceIdentity) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ControllersWhoAmIServiceIdentity) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersWhoAmIServiceIdentity) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ControllersWhoAmIServiceIdentity) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ControllersWhoAmIServiceIdentity) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ControllersWhoAmIServiceIdentity) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersWhoAmIServiceIdentity) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ControllersWhoAmIServiceIdentity) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ControllersWhoAmIServiceIdentity) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ControllersWhoAmIServiceIdentity) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersWhoAmIServiceIdentity) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ControllersWhoAmIServiceIdentity) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ControllersWhoAmIServiceIdentity) SetName(v string) {
	o.Name = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ControllersWhoAmIServiceIdentity) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersWhoAmIServiceIdentity) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ControllersWhoAmIServiceIdentity) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ControllersWhoAmIServiceIdentity) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ControllersWhoAmIServiceIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableControllersWhoAmIServiceIdentity struct {
	value *ControllersWhoAmIServiceIdentity
	isSet bool
}

func (v NullableControllersWhoAmIServiceIdentity) Get() *ControllersWhoAmIServiceIdentity {
	return v.value
}

func (v *NullableControllersWhoAmIServiceIdentity) Set(val *ControllersWhoAmIServiceIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersWhoAmIServiceIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersWhoAmIServiceIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersWhoAmIServiceIdentity(val *ControllersWhoAmIServiceIdentity) *NullableControllersWhoAmIServiceIdentity {
	return &NullableControllersWhoAmIServiceIdentity{value: val, isSet: true}
}

func (v NullableControllersWhoAmIServiceIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersWhoAmIServiceIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


