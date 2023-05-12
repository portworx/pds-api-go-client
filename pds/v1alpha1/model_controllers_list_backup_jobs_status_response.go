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

// ControllersListBackupJobsStatusResponse struct for ControllersListBackupJobsStatusResponse
type ControllersListBackupJobsStatusResponse struct {
	Data []ModelsBackupJobStatusResponse `json:"data,omitempty"`
}

// NewControllersListBackupJobsStatusResponse instantiates a new ControllersListBackupJobsStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersListBackupJobsStatusResponse() *ControllersListBackupJobsStatusResponse {
	this := ControllersListBackupJobsStatusResponse{}
	return &this
}

// NewControllersListBackupJobsStatusResponseWithDefaults instantiates a new ControllersListBackupJobsStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersListBackupJobsStatusResponseWithDefaults() *ControllersListBackupJobsStatusResponse {
	this := ControllersListBackupJobsStatusResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ControllersListBackupJobsStatusResponse) GetData() []ModelsBackupJobStatusResponse {
	if o == nil || o.Data == nil {
		var ret []ModelsBackupJobStatusResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersListBackupJobsStatusResponse) GetDataOk() ([]ModelsBackupJobStatusResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ControllersListBackupJobsStatusResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ModelsBackupJobStatusResponse and assigns it to the Data field.
func (o *ControllersListBackupJobsStatusResponse) SetData(v []ModelsBackupJobStatusResponse) {
	o.Data = v
}

func (o ControllersListBackupJobsStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableControllersListBackupJobsStatusResponse struct {
	value *ControllersListBackupJobsStatusResponse
	isSet bool
}

func (v NullableControllersListBackupJobsStatusResponse) Get() *ControllersListBackupJobsStatusResponse {
	return v.value
}

func (v *NullableControllersListBackupJobsStatusResponse) Set(val *ControllersListBackupJobsStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersListBackupJobsStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersListBackupJobsStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersListBackupJobsStatusResponse(val *ControllersListBackupJobsStatusResponse) *NullableControllersListBackupJobsStatusResponse {
	return &NullableControllersListBackupJobsStatusResponse{value: val, isSet: true}
}

func (v NullableControllersListBackupJobsStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersListBackupJobsStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


