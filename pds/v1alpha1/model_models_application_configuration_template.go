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

// checks if the ModelsApplicationConfigurationTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsApplicationConfigurationTemplate{}

// ModelsApplicationConfigurationTemplate struct for ModelsApplicationConfigurationTemplate
type ModelsApplicationConfigurationTemplate struct {
	AccountId *string `json:"account_id,omitempty"`
	ConfigItems []ModelsConfigItem `json:"config_items,omitempty"`
	// CreatedAt is autogenerated on creation
	CreatedAt *string `json:"created_at,omitempty"`
	DataServiceId *string `json:"data_service_id,omitempty"`
	// ID is auto generated on creation
	Id *string `json:"id,omitempty"`
	// Name of the template. Must be unique for a specific data service within the tenant scope.
	Name *string `json:"name,omitempty"`
	TenantId *string `json:"tenant_id,omitempty"`
	// UpdatedAt is autogenerated on update
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewModelsApplicationConfigurationTemplate instantiates a new ModelsApplicationConfigurationTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsApplicationConfigurationTemplate() *ModelsApplicationConfigurationTemplate {
	this := ModelsApplicationConfigurationTemplate{}
	return &this
}

// NewModelsApplicationConfigurationTemplateWithDefaults instantiates a new ModelsApplicationConfigurationTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsApplicationConfigurationTemplateWithDefaults() *ModelsApplicationConfigurationTemplate {
	this := ModelsApplicationConfigurationTemplate{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ModelsApplicationConfigurationTemplate) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsApplicationConfigurationTemplate) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ModelsApplicationConfigurationTemplate) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ModelsApplicationConfigurationTemplate) SetAccountId(v string) {
	o.AccountId = &v
}

// GetConfigItems returns the ConfigItems field value if set, zero value otherwise.
func (o *ModelsApplicationConfigurationTemplate) GetConfigItems() []ModelsConfigItem {
	if o == nil || IsNil(o.ConfigItems) {
		var ret []ModelsConfigItem
		return ret
	}
	return o.ConfigItems
}

// GetConfigItemsOk returns a tuple with the ConfigItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsApplicationConfigurationTemplate) GetConfigItemsOk() ([]ModelsConfigItem, bool) {
	if o == nil || IsNil(o.ConfigItems) {
		return nil, false
	}
	return o.ConfigItems, true
}

// HasConfigItems returns a boolean if a field has been set.
func (o *ModelsApplicationConfigurationTemplate) HasConfigItems() bool {
	if o != nil && !IsNil(o.ConfigItems) {
		return true
	}

	return false
}

// SetConfigItems gets a reference to the given []ModelsConfigItem and assigns it to the ConfigItems field.
func (o *ModelsApplicationConfigurationTemplate) SetConfigItems(v []ModelsConfigItem) {
	o.ConfigItems = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelsApplicationConfigurationTemplate) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsApplicationConfigurationTemplate) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelsApplicationConfigurationTemplate) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ModelsApplicationConfigurationTemplate) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDataServiceId returns the DataServiceId field value if set, zero value otherwise.
func (o *ModelsApplicationConfigurationTemplate) GetDataServiceId() string {
	if o == nil || IsNil(o.DataServiceId) {
		var ret string
		return ret
	}
	return *o.DataServiceId
}

// GetDataServiceIdOk returns a tuple with the DataServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsApplicationConfigurationTemplate) GetDataServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataServiceId) {
		return nil, false
	}
	return o.DataServiceId, true
}

// HasDataServiceId returns a boolean if a field has been set.
func (o *ModelsApplicationConfigurationTemplate) HasDataServiceId() bool {
	if o != nil && !IsNil(o.DataServiceId) {
		return true
	}

	return false
}

// SetDataServiceId gets a reference to the given string and assigns it to the DataServiceId field.
func (o *ModelsApplicationConfigurationTemplate) SetDataServiceId(v string) {
	o.DataServiceId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsApplicationConfigurationTemplate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsApplicationConfigurationTemplate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsApplicationConfigurationTemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelsApplicationConfigurationTemplate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelsApplicationConfigurationTemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsApplicationConfigurationTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelsApplicationConfigurationTemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelsApplicationConfigurationTemplate) SetName(v string) {
	o.Name = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ModelsApplicationConfigurationTemplate) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsApplicationConfigurationTemplate) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ModelsApplicationConfigurationTemplate) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ModelsApplicationConfigurationTemplate) SetTenantId(v string) {
	o.TenantId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelsApplicationConfigurationTemplate) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsApplicationConfigurationTemplate) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelsApplicationConfigurationTemplate) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ModelsApplicationConfigurationTemplate) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ModelsApplicationConfigurationTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsApplicationConfigurationTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.ConfigItems) {
		toSerialize["config_items"] = o.ConfigItems
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DataServiceId) {
		toSerialize["data_service_id"] = o.DataServiceId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenant_id"] = o.TenantId
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableModelsApplicationConfigurationTemplate struct {
	value *ModelsApplicationConfigurationTemplate
	isSet bool
}

func (v NullableModelsApplicationConfigurationTemplate) Get() *ModelsApplicationConfigurationTemplate {
	return v.value
}

func (v *NullableModelsApplicationConfigurationTemplate) Set(val *ModelsApplicationConfigurationTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsApplicationConfigurationTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsApplicationConfigurationTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsApplicationConfigurationTemplate(val *ModelsApplicationConfigurationTemplate) *NullableModelsApplicationConfigurationTemplate {
	return &NullableModelsApplicationConfigurationTemplate{value: val, isSet: true}
}

func (v NullableModelsApplicationConfigurationTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsApplicationConfigurationTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


