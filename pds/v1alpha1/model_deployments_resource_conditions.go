/*
PDS API

Portworx Data Services API Server

API version: 3.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"encoding/json"
)

// DeploymentsResourceConditions struct for DeploymentsResourceConditions
type DeploymentsResourceConditions struct {
	Conditions []DeploymentsCondition `json:"conditions,omitempty"`
	Resource *V1TypedLocalObjectReference `json:"resource,omitempty"`
}

// NewDeploymentsResourceConditions instantiates a new DeploymentsResourceConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentsResourceConditions() *DeploymentsResourceConditions {
	this := DeploymentsResourceConditions{}
	return &this
}

// NewDeploymentsResourceConditionsWithDefaults instantiates a new DeploymentsResourceConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentsResourceConditionsWithDefaults() *DeploymentsResourceConditions {
	this := DeploymentsResourceConditions{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *DeploymentsResourceConditions) GetConditions() []DeploymentsCondition {
	if o == nil || o.Conditions == nil {
		var ret []DeploymentsCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentsResourceConditions) GetConditionsOk() ([]DeploymentsCondition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *DeploymentsResourceConditions) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []DeploymentsCondition and assigns it to the Conditions field.
func (o *DeploymentsResourceConditions) SetConditions(v []DeploymentsCondition) {
	o.Conditions = v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *DeploymentsResourceConditions) GetResource() V1TypedLocalObjectReference {
	if o == nil || o.Resource == nil {
		var ret V1TypedLocalObjectReference
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentsResourceConditions) GetResourceOk() (*V1TypedLocalObjectReference, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *DeploymentsResourceConditions) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given V1TypedLocalObjectReference and assigns it to the Resource field.
func (o *DeploymentsResourceConditions) SetResource(v V1TypedLocalObjectReference) {
	o.Resource = &v
}

func (o DeploymentsResourceConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentsResourceConditions struct {
	value *DeploymentsResourceConditions
	isSet bool
}

func (v NullableDeploymentsResourceConditions) Get() *DeploymentsResourceConditions {
	return v.value
}

func (v *NullableDeploymentsResourceConditions) Set(val *DeploymentsResourceConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentsResourceConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentsResourceConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentsResourceConditions(val *DeploymentsResourceConditions) *NullableDeploymentsResourceConditions {
	return &NullableDeploymentsResourceConditions{value: val, isSet: true}
}

func (v NullableDeploymentsResourceConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentsResourceConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


