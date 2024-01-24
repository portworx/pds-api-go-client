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

// V1ListAccountsResponse ListAccountsResponse is the response to the ListAccounts API and contains the list of accounts visible to the caller.
type V1ListAccountsResponse struct {
	Accounts []V1Account `json:"accounts,omitempty"`
	Pagination *V1PageBasedPaginationResponse `json:"pagination,omitempty"`
}

// NewV1ListAccountsResponse instantiates a new V1ListAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ListAccountsResponse() *V1ListAccountsResponse {
	this := V1ListAccountsResponse{}
	return &this
}

// NewV1ListAccountsResponseWithDefaults instantiates a new V1ListAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ListAccountsResponseWithDefaults() *V1ListAccountsResponse {
	this := V1ListAccountsResponse{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *V1ListAccountsResponse) GetAccounts() []V1Account {
	if o == nil || o.Accounts == nil {
		var ret []V1Account
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ListAccountsResponse) GetAccountsOk() ([]V1Account, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *V1ListAccountsResponse) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []V1Account and assigns it to the Accounts field.
func (o *V1ListAccountsResponse) SetAccounts(v []V1Account) {
	o.Accounts = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *V1ListAccountsResponse) GetPagination() V1PageBasedPaginationResponse {
	if o == nil || o.Pagination == nil {
		var ret V1PageBasedPaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ListAccountsResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *V1ListAccountsResponse) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given V1PageBasedPaginationResponse and assigns it to the Pagination field.
func (o *V1ListAccountsResponse) SetPagination(v V1PageBasedPaginationResponse) {
	o.Pagination = &v
}

func (o V1ListAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableV1ListAccountsResponse struct {
	value *V1ListAccountsResponse
	isSet bool
}

func (v NullableV1ListAccountsResponse) Get() *V1ListAccountsResponse {
	return v.value
}

func (v *NullableV1ListAccountsResponse) Set(val *V1ListAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ListAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ListAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ListAccountsResponse(val *V1ListAccountsResponse) *NullableV1ListAccountsResponse {
	return &NullableV1ListAccountsResponse{value: val, isSet: true}
}

func (v NullableV1ListAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ListAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


