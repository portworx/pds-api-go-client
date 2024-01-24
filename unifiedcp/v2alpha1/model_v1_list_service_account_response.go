/*
public/portworx/pds/backupconfig/apiv1/backupconfig.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// V1ListServiceAccountResponse Response of requested list of service accounts.
type V1ListServiceAccountResponse struct {
	ServiceAccounts []V1ServiceAccount `json:"serviceAccounts,omitempty"`
	Pagination *V1PageBasedPaginationResponse `json:"pagination,omitempty"`
}

// NewV1ListServiceAccountResponse instantiates a new V1ListServiceAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ListServiceAccountResponse() *V1ListServiceAccountResponse {
	this := V1ListServiceAccountResponse{}
	return &this
}

// NewV1ListServiceAccountResponseWithDefaults instantiates a new V1ListServiceAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ListServiceAccountResponseWithDefaults() *V1ListServiceAccountResponse {
	this := V1ListServiceAccountResponse{}
	return &this
}

// GetServiceAccounts returns the ServiceAccounts field value if set, zero value otherwise.
func (o *V1ListServiceAccountResponse) GetServiceAccounts() []V1ServiceAccount {
	if o == nil || o.ServiceAccounts == nil {
		var ret []V1ServiceAccount
		return ret
	}
	return o.ServiceAccounts
}

// GetServiceAccountsOk returns a tuple with the ServiceAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ListServiceAccountResponse) GetServiceAccountsOk() ([]V1ServiceAccount, bool) {
	if o == nil || o.ServiceAccounts == nil {
		return nil, false
	}
	return o.ServiceAccounts, true
}

// HasServiceAccounts returns a boolean if a field has been set.
func (o *V1ListServiceAccountResponse) HasServiceAccounts() bool {
	if o != nil && o.ServiceAccounts != nil {
		return true
	}

	return false
}

// SetServiceAccounts gets a reference to the given []V1ServiceAccount and assigns it to the ServiceAccounts field.
func (o *V1ListServiceAccountResponse) SetServiceAccounts(v []V1ServiceAccount) {
	o.ServiceAccounts = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *V1ListServiceAccountResponse) GetPagination() V1PageBasedPaginationResponse {
	if o == nil || o.Pagination == nil {
		var ret V1PageBasedPaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ListServiceAccountResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *V1ListServiceAccountResponse) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given V1PageBasedPaginationResponse and assigns it to the Pagination field.
func (o *V1ListServiceAccountResponse) SetPagination(v V1PageBasedPaginationResponse) {
	o.Pagination = &v
}

func (o V1ListServiceAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceAccounts != nil {
		toSerialize["serviceAccounts"] = o.ServiceAccounts
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableV1ListServiceAccountResponse struct {
	value *V1ListServiceAccountResponse
	isSet bool
}

func (v NullableV1ListServiceAccountResponse) Get() *V1ListServiceAccountResponse {
	return v.value
}

func (v *NullableV1ListServiceAccountResponse) Set(val *V1ListServiceAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ListServiceAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ListServiceAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ListServiceAccountResponse(val *V1ListServiceAccountResponse) *NullableV1ListServiceAccountResponse {
	return &NullableV1ListServiceAccountResponse{value: val, isSet: true}
}

func (v NullableV1ListServiceAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ListServiceAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


