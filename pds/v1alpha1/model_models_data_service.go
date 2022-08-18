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

// ModelsDataService struct for ModelsDataService
type ModelsDataService struct {
	ComingSoon *bool `json:"coming_soon,omitempty"`
	// CreatedAt is autogenerated on creation
	CreatedAt *string `json:"created_at,omitempty"`
	HasFullBackup *bool `json:"has_full_backup,omitempty"`
	HasIncrementalBackup *bool `json:"has_incremental_backup,omitempty"`
	// ID is auto generated on creation
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	NodeRestrictions *ModelsNodeRestrictions `json:"node_restrictions,omitempty"`
	// TODO DS-2341 Remove this field once everything is implemented to use the NodeRestrictions.
	NodesLimitations *string `json:"nodes_limitations,omitempty"`
	ShortName *string `json:"short_name,omitempty"`
	// UpdatedAt is autogenerated on update
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewModelsDataService instantiates a new ModelsDataService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsDataService() *ModelsDataService {
	this := ModelsDataService{}
	return &this
}

// NewModelsDataServiceWithDefaults instantiates a new ModelsDataService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsDataServiceWithDefaults() *ModelsDataService {
	this := ModelsDataService{}
	return &this
}

// GetComingSoon returns the ComingSoon field value if set, zero value otherwise.
func (o *ModelsDataService) GetComingSoon() bool {
	if o == nil || o.ComingSoon == nil {
		var ret bool
		return ret
	}
	return *o.ComingSoon
}

// GetComingSoonOk returns a tuple with the ComingSoon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDataService) GetComingSoonOk() (*bool, bool) {
	if o == nil || o.ComingSoon == nil {
		return nil, false
	}
	return o.ComingSoon, true
}

// HasComingSoon returns a boolean if a field has been set.
func (o *ModelsDataService) HasComingSoon() bool {
	if o != nil && o.ComingSoon != nil {
		return true
	}

	return false
}

// SetComingSoon gets a reference to the given bool and assigns it to the ComingSoon field.
func (o *ModelsDataService) SetComingSoon(v bool) {
	o.ComingSoon = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelsDataService) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDataService) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelsDataService) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ModelsDataService) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetHasFullBackup returns the HasFullBackup field value if set, zero value otherwise.
func (o *ModelsDataService) GetHasFullBackup() bool {
	if o == nil || o.HasFullBackup == nil {
		var ret bool
		return ret
	}
	return *o.HasFullBackup
}

// GetHasFullBackupOk returns a tuple with the HasFullBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDataService) GetHasFullBackupOk() (*bool, bool) {
	if o == nil || o.HasFullBackup == nil {
		return nil, false
	}
	return o.HasFullBackup, true
}

// HasHasFullBackup returns a boolean if a field has been set.
func (o *ModelsDataService) HasHasFullBackup() bool {
	if o != nil && o.HasFullBackup != nil {
		return true
	}

	return false
}

// SetHasFullBackup gets a reference to the given bool and assigns it to the HasFullBackup field.
func (o *ModelsDataService) SetHasFullBackup(v bool) {
	o.HasFullBackup = &v
}

// GetHasIncrementalBackup returns the HasIncrementalBackup field value if set, zero value otherwise.
func (o *ModelsDataService) GetHasIncrementalBackup() bool {
	if o == nil || o.HasIncrementalBackup == nil {
		var ret bool
		return ret
	}
	return *o.HasIncrementalBackup
}

// GetHasIncrementalBackupOk returns a tuple with the HasIncrementalBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDataService) GetHasIncrementalBackupOk() (*bool, bool) {
	if o == nil || o.HasIncrementalBackup == nil {
		return nil, false
	}
	return o.HasIncrementalBackup, true
}

// HasHasIncrementalBackup returns a boolean if a field has been set.
func (o *ModelsDataService) HasHasIncrementalBackup() bool {
	if o != nil && o.HasIncrementalBackup != nil {
		return true
	}

	return false
}

// SetHasIncrementalBackup gets a reference to the given bool and assigns it to the HasIncrementalBackup field.
func (o *ModelsDataService) SetHasIncrementalBackup(v bool) {
	o.HasIncrementalBackup = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsDataService) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDataService) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsDataService) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelsDataService) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelsDataService) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDataService) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelsDataService) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelsDataService) SetName(v string) {
	o.Name = &v
}

// GetNodeRestrictions returns the NodeRestrictions field value if set, zero value otherwise.
func (o *ModelsDataService) GetNodeRestrictions() ModelsNodeRestrictions {
	if o == nil || o.NodeRestrictions == nil {
		var ret ModelsNodeRestrictions
		return ret
	}
	return *o.NodeRestrictions
}

// GetNodeRestrictionsOk returns a tuple with the NodeRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDataService) GetNodeRestrictionsOk() (*ModelsNodeRestrictions, bool) {
	if o == nil || o.NodeRestrictions == nil {
		return nil, false
	}
	return o.NodeRestrictions, true
}

// HasNodeRestrictions returns a boolean if a field has been set.
func (o *ModelsDataService) HasNodeRestrictions() bool {
	if o != nil && o.NodeRestrictions != nil {
		return true
	}

	return false
}

// SetNodeRestrictions gets a reference to the given ModelsNodeRestrictions and assigns it to the NodeRestrictions field.
func (o *ModelsDataService) SetNodeRestrictions(v ModelsNodeRestrictions) {
	o.NodeRestrictions = &v
}

// GetNodesLimitations returns the NodesLimitations field value if set, zero value otherwise.
func (o *ModelsDataService) GetNodesLimitations() string {
	if o == nil || o.NodesLimitations == nil {
		var ret string
		return ret
	}
	return *o.NodesLimitations
}

// GetNodesLimitationsOk returns a tuple with the NodesLimitations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDataService) GetNodesLimitationsOk() (*string, bool) {
	if o == nil || o.NodesLimitations == nil {
		return nil, false
	}
	return o.NodesLimitations, true
}

// HasNodesLimitations returns a boolean if a field has been set.
func (o *ModelsDataService) HasNodesLimitations() bool {
	if o != nil && o.NodesLimitations != nil {
		return true
	}

	return false
}

// SetNodesLimitations gets a reference to the given string and assigns it to the NodesLimitations field.
func (o *ModelsDataService) SetNodesLimitations(v string) {
	o.NodesLimitations = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *ModelsDataService) GetShortName() string {
	if o == nil || o.ShortName == nil {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDataService) GetShortNameOk() (*string, bool) {
	if o == nil || o.ShortName == nil {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *ModelsDataService) HasShortName() bool {
	if o != nil && o.ShortName != nil {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *ModelsDataService) SetShortName(v string) {
	o.ShortName = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelsDataService) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDataService) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelsDataService) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ModelsDataService) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ModelsDataService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ComingSoon != nil {
		toSerialize["coming_soon"] = o.ComingSoon
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.HasFullBackup != nil {
		toSerialize["has_full_backup"] = o.HasFullBackup
	}
	if o.HasIncrementalBackup != nil {
		toSerialize["has_incremental_backup"] = o.HasIncrementalBackup
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NodeRestrictions != nil {
		toSerialize["node_restrictions"] = o.NodeRestrictions
	}
	if o.NodesLimitations != nil {
		toSerialize["nodes_limitations"] = o.NodesLimitations
	}
	if o.ShortName != nil {
		toSerialize["short_name"] = o.ShortName
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableModelsDataService struct {
	value *ModelsDataService
	isSet bool
}

func (v NullableModelsDataService) Get() *ModelsDataService {
	return v.value
}

func (v *NullableModelsDataService) Set(val *ModelsDataService) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsDataService) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsDataService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsDataService(val *ModelsDataService) *NullableModelsDataService {
	return &NullableModelsDataService{value: val, isSet: true}
}

func (v NullableModelsDataService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsDataService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


