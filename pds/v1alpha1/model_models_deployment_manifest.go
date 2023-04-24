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

// checks if the ModelsDeploymentManifest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsDeploymentManifest{}

// ModelsDeploymentManifest struct for ModelsDeploymentManifest
type ModelsDeploymentManifest struct {
	// CreatedAt is autogenerated on creation
	CreatedAt *string `json:"created_at,omitempty"`
	Deployment *ModelsDeployment `json:"deployment,omitempty"`
	DeploymentId *string `json:"deployment_id,omitempty"`
	// ID is auto generated on creation
	Id *string `json:"id,omitempty"`
	Manifest map[string]interface{} `json:"manifest,omitempty"`
	ReadyReplicas *int32 `json:"readyReplicas,omitempty"`
	Replicas *int32 `json:"replicas,omitempty"`
	Status *string `json:"status,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	// UpdatedAt is autogenerated on update
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewModelsDeploymentManifest instantiates a new ModelsDeploymentManifest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsDeploymentManifest() *ModelsDeploymentManifest {
	this := ModelsDeploymentManifest{}
	return &this
}

// NewModelsDeploymentManifestWithDefaults instantiates a new ModelsDeploymentManifest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsDeploymentManifestWithDefaults() *ModelsDeploymentManifest {
	this := ModelsDeploymentManifest{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ModelsDeploymentManifest) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeployment returns the Deployment field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetDeployment() ModelsDeployment {
	if o == nil || IsNil(o.Deployment) {
		var ret ModelsDeployment
		return ret
	}
	return *o.Deployment
}

// GetDeploymentOk returns a tuple with the Deployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetDeploymentOk() (*ModelsDeployment, bool) {
	if o == nil || IsNil(o.Deployment) {
		return nil, false
	}
	return o.Deployment, true
}

// HasDeployment returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasDeployment() bool {
	if o != nil && !IsNil(o.Deployment) {
		return true
	}

	return false
}

// SetDeployment gets a reference to the given ModelsDeployment and assigns it to the Deployment field.
func (o *ModelsDeploymentManifest) SetDeployment(v ModelsDeployment) {
	o.Deployment = &v
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetDeploymentId() string {
	if o == nil || IsNil(o.DeploymentId) {
		var ret string
		return ret
	}
	return *o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetDeploymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeploymentId) {
		return nil, false
	}
	return o.DeploymentId, true
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasDeploymentId() bool {
	if o != nil && !IsNil(o.DeploymentId) {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given string and assigns it to the DeploymentId field.
func (o *ModelsDeploymentManifest) SetDeploymentId(v string) {
	o.DeploymentId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelsDeploymentManifest) SetId(v string) {
	o.Id = &v
}

// GetManifest returns the Manifest field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetManifest() map[string]interface{} {
	if o == nil || IsNil(o.Manifest) {
		var ret map[string]interface{}
		return ret
	}
	return o.Manifest
}

// GetManifestOk returns a tuple with the Manifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetManifestOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Manifest) {
		return map[string]interface{}{}, false
	}
	return o.Manifest, true
}

// HasManifest returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasManifest() bool {
	if o != nil && !IsNil(o.Manifest) {
		return true
	}

	return false
}

// SetManifest gets a reference to the given map[string]interface{} and assigns it to the Manifest field.
func (o *ModelsDeploymentManifest) SetManifest(v map[string]interface{}) {
	o.Manifest = v
}

// GetReadyReplicas returns the ReadyReplicas field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetReadyReplicas() int32 {
	if o == nil || IsNil(o.ReadyReplicas) {
		var ret int32
		return ret
	}
	return *o.ReadyReplicas
}

// GetReadyReplicasOk returns a tuple with the ReadyReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetReadyReplicasOk() (*int32, bool) {
	if o == nil || IsNil(o.ReadyReplicas) {
		return nil, false
	}
	return o.ReadyReplicas, true
}

// HasReadyReplicas returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasReadyReplicas() bool {
	if o != nil && !IsNil(o.ReadyReplicas) {
		return true
	}

	return false
}

// SetReadyReplicas gets a reference to the given int32 and assigns it to the ReadyReplicas field.
func (o *ModelsDeploymentManifest) SetReadyReplicas(v int32) {
	o.ReadyReplicas = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetReplicas() int32 {
	if o == nil || IsNil(o.Replicas) {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetReplicasOk() (*int32, bool) {
	if o == nil || IsNil(o.Replicas) {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasReplicas() bool {
	if o != nil && !IsNil(o.Replicas) {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *ModelsDeploymentManifest) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ModelsDeploymentManifest) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *ModelsDeploymentManifest) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ModelsDeploymentManifest) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ModelsDeploymentManifest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsDeploymentManifest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Deployment) {
		toSerialize["deployment"] = o.Deployment
	}
	if !IsNil(o.DeploymentId) {
		toSerialize["deployment_id"] = o.DeploymentId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Manifest) {
		toSerialize["manifest"] = o.Manifest
	}
	if !IsNil(o.ReadyReplicas) {
		toSerialize["readyReplicas"] = o.ReadyReplicas
	}
	if !IsNil(o.Replicas) {
		toSerialize["replicas"] = o.Replicas
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableModelsDeploymentManifest struct {
	value *ModelsDeploymentManifest
	isSet bool
}

func (v NullableModelsDeploymentManifest) Get() *ModelsDeploymentManifest {
	return v.value
}

func (v *NullableModelsDeploymentManifest) Set(val *ModelsDeploymentManifest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsDeploymentManifest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsDeploymentManifest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsDeploymentManifest(val *ModelsDeploymentManifest) *NullableModelsDeploymentManifest {
	return &NullableModelsDeploymentManifest{value: val, isSet: true}
}

func (v NullableModelsDeploymentManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsDeploymentManifest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


