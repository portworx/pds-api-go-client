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

// checks if the V1DeploymentEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DeploymentEvent{}

// V1DeploymentEvent DeploymentEvent contains all the fields received from kubernetes event of a deployment.
type V1DeploymentEvent struct {
	Meta *V1Meta `json:"meta,omitempty"`
	Status *Deploymenteventsv1Status `json:"status,omitempty"`
}

// NewV1DeploymentEvent instantiates a new V1DeploymentEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DeploymentEvent() *V1DeploymentEvent {
	this := V1DeploymentEvent{}
	return &this
}

// NewV1DeploymentEventWithDefaults instantiates a new V1DeploymentEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DeploymentEventWithDefaults() *V1DeploymentEvent {
	this := V1DeploymentEvent{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V1DeploymentEvent) GetMeta() V1Meta {
	if o == nil || IsNil(o.Meta) {
		var ret V1Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeploymentEvent) GetMetaOk() (*V1Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V1DeploymentEvent) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V1Meta and assigns it to the Meta field.
func (o *V1DeploymentEvent) SetMeta(v V1Meta) {
	o.Meta = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V1DeploymentEvent) GetStatus() Deploymenteventsv1Status {
	if o == nil || IsNil(o.Status) {
		var ret Deploymenteventsv1Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeploymentEvent) GetStatusOk() (*Deploymenteventsv1Status, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V1DeploymentEvent) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Deploymenteventsv1Status and assigns it to the Status field.
func (o *V1DeploymentEvent) SetStatus(v Deploymenteventsv1Status) {
	o.Status = &v
}

func (o V1DeploymentEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DeploymentEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableV1DeploymentEvent struct {
	value *V1DeploymentEvent
	isSet bool
}

func (v NullableV1DeploymentEvent) Get() *V1DeploymentEvent {
	return v.value
}

func (v *NullableV1DeploymentEvent) Set(val *V1DeploymentEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DeploymentEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DeploymentEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DeploymentEvent(val *V1DeploymentEvent) *NullableV1DeploymentEvent {
	return &NullableV1DeploymentEvent{value: val, isSet: true}
}

func (v NullableV1DeploymentEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DeploymentEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

