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

// checks if the ModelsDeploymentTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsDeploymentTarget{}

// ModelsDeploymentTarget struct for ModelsDeploymentTarget
type ModelsDeploymentTarget struct {
	AccountId *string `json:"account_id,omitempty"`
	Capabilities *ModelsDeploymentTargetCapabilities `json:"capabilities,omitempty"`
	ClusterId *string `json:"cluster_id,omitempty"`
	ConnectionStatus *string `json:"connection_status,omitempty"`
	// CreatedAt is autogenerated on creation
	CreatedAt *string `json:"created_at,omitempty"`
	// ID is auto generated on creation
	Id *string `json:"id,omitempty"`
	LastHealthCheck *string `json:"last_health_check,omitempty"`
	LastOperatorHeartbeat *string `json:"last_operator_heartbeat,omitempty"`
	Metadata *ModelsDeploymentTargetMetadata `json:"metadata,omitempty"`
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	TenantId *string `json:"tenant_id,omitempty"`
	TlsIssuer *string `json:"tls_issuer,omitempty"`
	TlsRequired *bool `json:"tls_required,omitempty"`
	// UpdatedAt is autogenerated on update
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewModelsDeploymentTarget instantiates a new ModelsDeploymentTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsDeploymentTarget() *ModelsDeploymentTarget {
	this := ModelsDeploymentTarget{}
	return &this
}

// NewModelsDeploymentTargetWithDefaults instantiates a new ModelsDeploymentTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsDeploymentTargetWithDefaults() *ModelsDeploymentTarget {
	this := ModelsDeploymentTarget{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ModelsDeploymentTarget) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentTarget) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ModelsDeploymentTarget) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ModelsDeploymentTarget) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *ModelsDeploymentTarget) GetCapabilities() ModelsDeploymentTargetCapabilities {
	if o == nil || IsNil(o.Capabilities) {
		var ret ModelsDeploymentTargetCapabilities
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentTarget) GetCapabilitiesOk() (*ModelsDeploymentTargetCapabilities, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *ModelsDeploymentTarget) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given ModelsDeploymentTargetCapabilities and assigns it to the Capabilities field.
func (o *ModelsDeploymentTarget) SetCapabilities(v ModelsDeploymentTargetCapabilities) {
	o.Capabilities = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *ModelsDeploymentTarget) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentTarget) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *ModelsDeploymentTarget) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *ModelsDeploymentTarget) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *ModelsDeploymentTarget) GetConnectionStatus() string {
	if o == nil || IsNil(o.ConnectionStatus) {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentTarget) GetConnectionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionStatus) {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *ModelsDeploymentTarget) HasConnectionStatus() bool {
	if o != nil && !IsNil(o.ConnectionStatus) {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *ModelsDeploymentTarget) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelsDeploymentTarget) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentTarget) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelsDeploymentTarget) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ModelsDeploymentTarget) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsDeploymentTarget) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentTarget) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsDeploymentTarget) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelsDeploymentTarget) SetId(v string) {
	o.Id = &v
}

// GetLastHealthCheck returns the LastHealthCheck field value if set, zero value otherwise.
func (o *ModelsDeploymentTarget) GetLastHealthCheck() string {
	if o == nil || IsNil(o.LastHealthCheck) {
		var ret string
		return ret
	}
	return *o.LastHealthCheck
}

// GetLastHealthCheckOk returns a tuple with the LastHealthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentTarget) GetLastHealthCheckOk() (*string, bool) {
	if o == nil || IsNil(o.LastHealthCheck) {
		return nil, false
	}
	return o.LastHealthCheck, true
}

// HasLastHealthCheck returns a boolean if a field has been set.
func (o *ModelsDeploymentTarget) HasLastHealthCheck() bool {
	if o != nil && !IsNil(o.LastHealthCheck) {
		return true
	}

	return false
}

// SetLastHealthCheck gets a reference to the given string and assigns it to the LastHealthCheck field.
func (o *ModelsDeploymentTarget) SetLastHealthCheck(v string) {
	o.LastHealthCheck = &v
}

// GetLastOperatorHeartbeat returns the LastOperatorHeartbeat field value if set, zero value otherwise.
func (o *ModelsDeploymentTarget) GetLastOperatorHeartbeat() string {
	if o == nil || IsNil(o.LastOperatorHeartbeat) {
		var ret string
		return ret
	}
	return *o.LastOperatorHeartbeat
}

// GetLastOperatorHeartbeatOk returns a tuple with the LastOperatorHeartbeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentTarget) GetLastOperatorHeartbeatOk() (*string, bool) {
	if o == nil || IsNil(o.LastOperatorHeartbeat) {
		return nil, false
	}
	return o.LastOperatorHeartbeat, true
}

// HasLastOperatorHeartbeat returns a boolean if a field has been set.
func (o *ModelsDeploymentTarget) HasLastOperatorHeartbeat() bool {
	if o != nil && !IsNil(o.LastOperatorHeartbeat) {
		return true
	}

	return false
}

// SetLastOperatorHeartbeat gets a reference to the given string and assigns it to the LastOperatorHeartbeat field.
func (o *ModelsDeploymentTarget) SetLastOperatorHeartbeat(v string) {
	o.LastOperatorHeartbeat = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ModelsDeploymentTarget) GetMetadata() ModelsDeploymentTargetMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ModelsDeploymentTargetMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentTarget) GetMetadataOk() (*ModelsDeploymentTargetMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ModelsDeploymentTarget) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ModelsDeploymentTargetMetadata and assigns it to the Metadata field.
func (o *ModelsDeploymentTarget) SetMetadata(v ModelsDeploymentTargetMetadata) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelsDeploymentTarget) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentTarget) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelsDeploymentTarget) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelsDeploymentTarget) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ModelsDeploymentTarget) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentTarget) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelsDeploymentTarget) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ModelsDeploymentTarget) SetStatus(v string) {
	o.Status = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ModelsDeploymentTarget) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentTarget) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ModelsDeploymentTarget) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ModelsDeploymentTarget) SetTenantId(v string) {
	o.TenantId = &v
}

// GetTlsIssuer returns the TlsIssuer field value if set, zero value otherwise.
func (o *ModelsDeploymentTarget) GetTlsIssuer() string {
	if o == nil || IsNil(o.TlsIssuer) {
		var ret string
		return ret
	}
	return *o.TlsIssuer
}

// GetTlsIssuerOk returns a tuple with the TlsIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentTarget) GetTlsIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.TlsIssuer) {
		return nil, false
	}
	return o.TlsIssuer, true
}

// HasTlsIssuer returns a boolean if a field has been set.
func (o *ModelsDeploymentTarget) HasTlsIssuer() bool {
	if o != nil && !IsNil(o.TlsIssuer) {
		return true
	}

	return false
}

// SetTlsIssuer gets a reference to the given string and assigns it to the TlsIssuer field.
func (o *ModelsDeploymentTarget) SetTlsIssuer(v string) {
	o.TlsIssuer = &v
}

// GetTlsRequired returns the TlsRequired field value if set, zero value otherwise.
func (o *ModelsDeploymentTarget) GetTlsRequired() bool {
	if o == nil || IsNil(o.TlsRequired) {
		var ret bool
		return ret
	}
	return *o.TlsRequired
}

// GetTlsRequiredOk returns a tuple with the TlsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentTarget) GetTlsRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.TlsRequired) {
		return nil, false
	}
	return o.TlsRequired, true
}

// HasTlsRequired returns a boolean if a field has been set.
func (o *ModelsDeploymentTarget) HasTlsRequired() bool {
	if o != nil && !IsNil(o.TlsRequired) {
		return true
	}

	return false
}

// SetTlsRequired gets a reference to the given bool and assigns it to the TlsRequired field.
func (o *ModelsDeploymentTarget) SetTlsRequired(v bool) {
	o.TlsRequired = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelsDeploymentTarget) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentTarget) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelsDeploymentTarget) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ModelsDeploymentTarget) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ModelsDeploymentTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsDeploymentTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.ClusterId) {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if !IsNil(o.ConnectionStatus) {
		toSerialize["connection_status"] = o.ConnectionStatus
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastHealthCheck) {
		toSerialize["last_health_check"] = o.LastHealthCheck
	}
	if !IsNil(o.LastOperatorHeartbeat) {
		toSerialize["last_operator_heartbeat"] = o.LastOperatorHeartbeat
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenant_id"] = o.TenantId
	}
	if !IsNil(o.TlsIssuer) {
		toSerialize["tls_issuer"] = o.TlsIssuer
	}
	if !IsNil(o.TlsRequired) {
		toSerialize["tls_required"] = o.TlsRequired
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableModelsDeploymentTarget struct {
	value *ModelsDeploymentTarget
	isSet bool
}

func (v NullableModelsDeploymentTarget) Get() *ModelsDeploymentTarget {
	return v.value
}

func (v *NullableModelsDeploymentTarget) Set(val *ModelsDeploymentTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsDeploymentTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsDeploymentTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsDeploymentTarget(val *ModelsDeploymentTarget) *NullableModelsDeploymentTarget {
	return &NullableModelsDeploymentTarget{value: val, isSet: true}
}

func (v NullableModelsDeploymentTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsDeploymentTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


