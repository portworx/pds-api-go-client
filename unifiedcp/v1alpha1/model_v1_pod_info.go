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

// checks if the V1PodInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1PodInfo{}

// V1PodInfo PodInfo contains information about a pod.
type V1PodInfo struct {
	// The IP of a pod.
	Ip *string `json:"ip,omitempty"`
	// Name is the Hostname of a pod.
	Name *string `json:"name,omitempty"`
	// Node that hosts a particular pod.
	WorkerNode *string `json:"workerNode,omitempty"`
}

// NewV1PodInfo instantiates a new V1PodInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1PodInfo() *V1PodInfo {
	this := V1PodInfo{}
	return &this
}

// NewV1PodInfoWithDefaults instantiates a new V1PodInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1PodInfoWithDefaults() *V1PodInfo {
	this := V1PodInfo{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *V1PodInfo) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodInfo) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *V1PodInfo) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *V1PodInfo) SetIp(v string) {
	o.Ip = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1PodInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1PodInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1PodInfo) SetName(v string) {
	o.Name = &v
}

// GetWorkerNode returns the WorkerNode field value if set, zero value otherwise.
func (o *V1PodInfo) GetWorkerNode() string {
	if o == nil || IsNil(o.WorkerNode) {
		var ret string
		return ret
	}
	return *o.WorkerNode
}

// GetWorkerNodeOk returns a tuple with the WorkerNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodInfo) GetWorkerNodeOk() (*string, bool) {
	if o == nil || IsNil(o.WorkerNode) {
		return nil, false
	}
	return o.WorkerNode, true
}

// HasWorkerNode returns a boolean if a field has been set.
func (o *V1PodInfo) HasWorkerNode() bool {
	if o != nil && !IsNil(o.WorkerNode) {
		return true
	}

	return false
}

// SetWorkerNode gets a reference to the given string and assigns it to the WorkerNode field.
func (o *V1PodInfo) SetWorkerNode(v string) {
	o.WorkerNode = &v
}

func (o V1PodInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1PodInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.WorkerNode) {
		toSerialize["workerNode"] = o.WorkerNode
	}
	return toSerialize, nil
}

type NullableV1PodInfo struct {
	value *V1PodInfo
	isSet bool
}

func (v NullableV1PodInfo) Get() *V1PodInfo {
	return v.value
}

func (v *NullableV1PodInfo) Set(val *V1PodInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableV1PodInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableV1PodInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1PodInfo(val *V1PodInfo) *NullableV1PodInfo {
	return &NullableV1PodInfo{value: val, isSet: true}
}

func (v NullableV1PodInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1PodInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


