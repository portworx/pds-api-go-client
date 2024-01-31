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

// checks if the Backupconfigv1Status type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Backupconfigv1Status{}

// Backupconfigv1Status Status for backup configuration.
type Backupconfigv1Status struct {
	Phase *StatusPhase `json:"phase,omitempty"`
	// Custom Resource Name is the kubernetes resource name for the backup that is built from ID.
	CustomResourceName *string `json:"customResourceName,omitempty"`
	// Flag to check if the schedule is synchronized or not.
	IsScheduleSynchronized *bool `json:"isScheduleSynchronized,omitempty"`
	DeploymentMetaData *V1DeploymentMetaData `json:"deploymentMetaData,omitempty"`
}

// NewBackupconfigv1Status instantiates a new Backupconfigv1Status object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupconfigv1Status() *Backupconfigv1Status {
	this := Backupconfigv1Status{}
	var phase StatusPhase = STATUSPHASE_PHASE_UNSPECIFIED
	this.Phase = &phase
	return &this
}

// NewBackupconfigv1StatusWithDefaults instantiates a new Backupconfigv1Status object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupconfigv1StatusWithDefaults() *Backupconfigv1Status {
	this := Backupconfigv1Status{}
	var phase StatusPhase = STATUSPHASE_PHASE_UNSPECIFIED
	this.Phase = &phase
	return &this
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *Backupconfigv1Status) GetPhase() StatusPhase {
	if o == nil || IsNil(o.Phase) {
		var ret StatusPhase
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backupconfigv1Status) GetPhaseOk() (*StatusPhase, bool) {
	if o == nil || IsNil(o.Phase) {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *Backupconfigv1Status) HasPhase() bool {
	if o != nil && !IsNil(o.Phase) {
		return true
	}

	return false
}

// SetPhase gets a reference to the given StatusPhase and assigns it to the Phase field.
func (o *Backupconfigv1Status) SetPhase(v StatusPhase) {
	o.Phase = &v
}

// GetCustomResourceName returns the CustomResourceName field value if set, zero value otherwise.
func (o *Backupconfigv1Status) GetCustomResourceName() string {
	if o == nil || IsNil(o.CustomResourceName) {
		var ret string
		return ret
	}
	return *o.CustomResourceName
}

// GetCustomResourceNameOk returns a tuple with the CustomResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backupconfigv1Status) GetCustomResourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.CustomResourceName) {
		return nil, false
	}
	return o.CustomResourceName, true
}

// HasCustomResourceName returns a boolean if a field has been set.
func (o *Backupconfigv1Status) HasCustomResourceName() bool {
	if o != nil && !IsNil(o.CustomResourceName) {
		return true
	}

	return false
}

// SetCustomResourceName gets a reference to the given string and assigns it to the CustomResourceName field.
func (o *Backupconfigv1Status) SetCustomResourceName(v string) {
	o.CustomResourceName = &v
}

// GetIsScheduleSynchronized returns the IsScheduleSynchronized field value if set, zero value otherwise.
func (o *Backupconfigv1Status) GetIsScheduleSynchronized() bool {
	if o == nil || IsNil(o.IsScheduleSynchronized) {
		var ret bool
		return ret
	}
	return *o.IsScheduleSynchronized
}

// GetIsScheduleSynchronizedOk returns a tuple with the IsScheduleSynchronized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backupconfigv1Status) GetIsScheduleSynchronizedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsScheduleSynchronized) {
		return nil, false
	}
	return o.IsScheduleSynchronized, true
}

// HasIsScheduleSynchronized returns a boolean if a field has been set.
func (o *Backupconfigv1Status) HasIsScheduleSynchronized() bool {
	if o != nil && !IsNil(o.IsScheduleSynchronized) {
		return true
	}

	return false
}

// SetIsScheduleSynchronized gets a reference to the given bool and assigns it to the IsScheduleSynchronized field.
func (o *Backupconfigv1Status) SetIsScheduleSynchronized(v bool) {
	o.IsScheduleSynchronized = &v
}

// GetDeploymentMetaData returns the DeploymentMetaData field value if set, zero value otherwise.
func (o *Backupconfigv1Status) GetDeploymentMetaData() V1DeploymentMetaData {
	if o == nil || IsNil(o.DeploymentMetaData) {
		var ret V1DeploymentMetaData
		return ret
	}
	return *o.DeploymentMetaData
}

// GetDeploymentMetaDataOk returns a tuple with the DeploymentMetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backupconfigv1Status) GetDeploymentMetaDataOk() (*V1DeploymentMetaData, bool) {
	if o == nil || IsNil(o.DeploymentMetaData) {
		return nil, false
	}
	return o.DeploymentMetaData, true
}

// HasDeploymentMetaData returns a boolean if a field has been set.
func (o *Backupconfigv1Status) HasDeploymentMetaData() bool {
	if o != nil && !IsNil(o.DeploymentMetaData) {
		return true
	}

	return false
}

// SetDeploymentMetaData gets a reference to the given V1DeploymentMetaData and assigns it to the DeploymentMetaData field.
func (o *Backupconfigv1Status) SetDeploymentMetaData(v V1DeploymentMetaData) {
	o.DeploymentMetaData = &v
}

func (o Backupconfigv1Status) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Backupconfigv1Status) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Phase) {
		toSerialize["phase"] = o.Phase
	}
	if !IsNil(o.CustomResourceName) {
		toSerialize["customResourceName"] = o.CustomResourceName
	}
	if !IsNil(o.IsScheduleSynchronized) {
		toSerialize["isScheduleSynchronized"] = o.IsScheduleSynchronized
	}
	if !IsNil(o.DeploymentMetaData) {
		toSerialize["deploymentMetaData"] = o.DeploymentMetaData
	}
	return toSerialize, nil
}

type NullableBackupconfigv1Status struct {
	value *Backupconfigv1Status
	isSet bool
}

func (v NullableBackupconfigv1Status) Get() *Backupconfigv1Status {
	return v.value
}

func (v *NullableBackupconfigv1Status) Set(val *Backupconfigv1Status) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupconfigv1Status) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupconfigv1Status) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupconfigv1Status(val *Backupconfigv1Status) *NullableBackupconfigv1Status {
	return &NullableBackupconfigv1Status{value: val, isSet: true}
}

func (v NullableBackupconfigv1Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupconfigv1Status) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


