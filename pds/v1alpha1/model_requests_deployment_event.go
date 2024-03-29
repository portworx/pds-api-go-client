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

// RequestsDeploymentEvent struct for RequestsDeploymentEvent
type RequestsDeploymentEvent struct {
	// DeploymentID of the event.
	DeploymentId string `json:"deployment_id"`
	// List of k8s events for the deployment
	Events []RequestsK8sEvent `json:"events"`
	// Namespace in which the deployment events were generated.
	Namespace string `json:"namespace"`
}

// NewRequestsDeploymentEvent instantiates a new RequestsDeploymentEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestsDeploymentEvent(deploymentId string, events []RequestsK8sEvent, namespace string) *RequestsDeploymentEvent {
	this := RequestsDeploymentEvent{}
	this.DeploymentId = deploymentId
	this.Events = events
	this.Namespace = namespace
	return &this
}

// NewRequestsDeploymentEventWithDefaults instantiates a new RequestsDeploymentEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestsDeploymentEventWithDefaults() *RequestsDeploymentEvent {
	this := RequestsDeploymentEvent{}
	return &this
}

// GetDeploymentId returns the DeploymentId field value
func (o *RequestsDeploymentEvent) GetDeploymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value
// and a boolean to check if the value has been set.
func (o *RequestsDeploymentEvent) GetDeploymentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DeploymentId, true
}

// SetDeploymentId sets field value
func (o *RequestsDeploymentEvent) SetDeploymentId(v string) {
	o.DeploymentId = v
}

// GetEvents returns the Events field value
func (o *RequestsDeploymentEvent) GetEvents() []RequestsK8sEvent {
	if o == nil {
		var ret []RequestsK8sEvent
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *RequestsDeploymentEvent) GetEventsOk() ([]RequestsK8sEvent, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *RequestsDeploymentEvent) SetEvents(v []RequestsK8sEvent) {
	o.Events = v
}

// GetNamespace returns the Namespace field value
func (o *RequestsDeploymentEvent) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *RequestsDeploymentEvent) GetNamespaceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *RequestsDeploymentEvent) SetNamespace(v string) {
	o.Namespace = v
}

func (o RequestsDeploymentEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["deployment_id"] = o.DeploymentId
	}
	if true {
		toSerialize["events"] = o.Events
	}
	if true {
		toSerialize["namespace"] = o.Namespace
	}
	return json.Marshal(toSerialize)
}

type NullableRequestsDeploymentEvent struct {
	value *RequestsDeploymentEvent
	isSet bool
}

func (v NullableRequestsDeploymentEvent) Get() *RequestsDeploymentEvent {
	return v.value
}

func (v *NullableRequestsDeploymentEvent) Set(val *RequestsDeploymentEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestsDeploymentEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestsDeploymentEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestsDeploymentEvent(val *RequestsDeploymentEvent) *NullableRequestsDeploymentEvent {
	return &NullableRequestsDeploymentEvent{value: val, isSet: true}
}

func (v NullableRequestsDeploymentEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestsDeploymentEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


