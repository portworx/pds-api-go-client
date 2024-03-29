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

// RequestsCreateProjectDeploymentRequest struct for RequestsCreateProjectDeploymentRequest
type RequestsCreateProjectDeploymentRequest struct {
	ApplicationConfigurationOverrides *map[string]string `json:"application_configuration_overrides,omitempty"`
	ApplicationConfigurationTemplateId *string `json:"application_configuration_template_id,omitempty"`
	DeploymentTargetId *string `json:"deployment_target_id,omitempty"`
	DnsZone *string `json:"dns_zone,omitempty"`
	ImageId *string `json:"image_id,omitempty"`
	LoadBalancerSourceRanges []string `json:"load_balancer_source_ranges,omitempty"`
	Name *string `json:"name,omitempty"`
	NamespaceId *string `json:"namespace_id,omitempty"`
	NodeCount *int32 `json:"node_count,omitempty"`
	ResourceSettingsTemplateId *string `json:"resource_settings_template_id,omitempty"`
	ScheduledBackup *RequestsDeploymentScheduledBackup `json:"scheduled_backup,omitempty"`
	ServiceType *string `json:"service_type,omitempty"`
	StorageOptionsTemplateId *string `json:"storage_options_template_id,omitempty"`
	TlsEnabled *bool `json:"tls_enabled,omitempty"`
}

// NewRequestsCreateProjectDeploymentRequest instantiates a new RequestsCreateProjectDeploymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestsCreateProjectDeploymentRequest() *RequestsCreateProjectDeploymentRequest {
	this := RequestsCreateProjectDeploymentRequest{}
	return &this
}

// NewRequestsCreateProjectDeploymentRequestWithDefaults instantiates a new RequestsCreateProjectDeploymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestsCreateProjectDeploymentRequestWithDefaults() *RequestsCreateProjectDeploymentRequest {
	this := RequestsCreateProjectDeploymentRequest{}
	return &this
}

// GetApplicationConfigurationOverrides returns the ApplicationConfigurationOverrides field value if set, zero value otherwise.
func (o *RequestsCreateProjectDeploymentRequest) GetApplicationConfigurationOverrides() map[string]string {
	if o == nil || o.ApplicationConfigurationOverrides == nil {
		var ret map[string]string
		return ret
	}
	return *o.ApplicationConfigurationOverrides
}

// GetApplicationConfigurationOverridesOk returns a tuple with the ApplicationConfigurationOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateProjectDeploymentRequest) GetApplicationConfigurationOverridesOk() (*map[string]string, bool) {
	if o == nil || o.ApplicationConfigurationOverrides == nil {
		return nil, false
	}
	return o.ApplicationConfigurationOverrides, true
}

// HasApplicationConfigurationOverrides returns a boolean if a field has been set.
func (o *RequestsCreateProjectDeploymentRequest) HasApplicationConfigurationOverrides() bool {
	if o != nil && o.ApplicationConfigurationOverrides != nil {
		return true
	}

	return false
}

// SetApplicationConfigurationOverrides gets a reference to the given map[string]string and assigns it to the ApplicationConfigurationOverrides field.
func (o *RequestsCreateProjectDeploymentRequest) SetApplicationConfigurationOverrides(v map[string]string) {
	o.ApplicationConfigurationOverrides = &v
}

// GetApplicationConfigurationTemplateId returns the ApplicationConfigurationTemplateId field value if set, zero value otherwise.
func (o *RequestsCreateProjectDeploymentRequest) GetApplicationConfigurationTemplateId() string {
	if o == nil || o.ApplicationConfigurationTemplateId == nil {
		var ret string
		return ret
	}
	return *o.ApplicationConfigurationTemplateId
}

// GetApplicationConfigurationTemplateIdOk returns a tuple with the ApplicationConfigurationTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateProjectDeploymentRequest) GetApplicationConfigurationTemplateIdOk() (*string, bool) {
	if o == nil || o.ApplicationConfigurationTemplateId == nil {
		return nil, false
	}
	return o.ApplicationConfigurationTemplateId, true
}

// HasApplicationConfigurationTemplateId returns a boolean if a field has been set.
func (o *RequestsCreateProjectDeploymentRequest) HasApplicationConfigurationTemplateId() bool {
	if o != nil && o.ApplicationConfigurationTemplateId != nil {
		return true
	}

	return false
}

// SetApplicationConfigurationTemplateId gets a reference to the given string and assigns it to the ApplicationConfigurationTemplateId field.
func (o *RequestsCreateProjectDeploymentRequest) SetApplicationConfigurationTemplateId(v string) {
	o.ApplicationConfigurationTemplateId = &v
}

// GetDeploymentTargetId returns the DeploymentTargetId field value if set, zero value otherwise.
func (o *RequestsCreateProjectDeploymentRequest) GetDeploymentTargetId() string {
	if o == nil || o.DeploymentTargetId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentTargetId
}

// GetDeploymentTargetIdOk returns a tuple with the DeploymentTargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateProjectDeploymentRequest) GetDeploymentTargetIdOk() (*string, bool) {
	if o == nil || o.DeploymentTargetId == nil {
		return nil, false
	}
	return o.DeploymentTargetId, true
}

// HasDeploymentTargetId returns a boolean if a field has been set.
func (o *RequestsCreateProjectDeploymentRequest) HasDeploymentTargetId() bool {
	if o != nil && o.DeploymentTargetId != nil {
		return true
	}

	return false
}

// SetDeploymentTargetId gets a reference to the given string and assigns it to the DeploymentTargetId field.
func (o *RequestsCreateProjectDeploymentRequest) SetDeploymentTargetId(v string) {
	o.DeploymentTargetId = &v
}

// GetDnsZone returns the DnsZone field value if set, zero value otherwise.
func (o *RequestsCreateProjectDeploymentRequest) GetDnsZone() string {
	if o == nil || o.DnsZone == nil {
		var ret string
		return ret
	}
	return *o.DnsZone
}

// GetDnsZoneOk returns a tuple with the DnsZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateProjectDeploymentRequest) GetDnsZoneOk() (*string, bool) {
	if o == nil || o.DnsZone == nil {
		return nil, false
	}
	return o.DnsZone, true
}

// HasDnsZone returns a boolean if a field has been set.
func (o *RequestsCreateProjectDeploymentRequest) HasDnsZone() bool {
	if o != nil && o.DnsZone != nil {
		return true
	}

	return false
}

// SetDnsZone gets a reference to the given string and assigns it to the DnsZone field.
func (o *RequestsCreateProjectDeploymentRequest) SetDnsZone(v string) {
	o.DnsZone = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *RequestsCreateProjectDeploymentRequest) GetImageId() string {
	if o == nil || o.ImageId == nil {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateProjectDeploymentRequest) GetImageIdOk() (*string, bool) {
	if o == nil || o.ImageId == nil {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *RequestsCreateProjectDeploymentRequest) HasImageId() bool {
	if o != nil && o.ImageId != nil {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *RequestsCreateProjectDeploymentRequest) SetImageId(v string) {
	o.ImageId = &v
}

// GetLoadBalancerSourceRanges returns the LoadBalancerSourceRanges field value if set, zero value otherwise.
func (o *RequestsCreateProjectDeploymentRequest) GetLoadBalancerSourceRanges() []string {
	if o == nil || o.LoadBalancerSourceRanges == nil {
		var ret []string
		return ret
	}
	return o.LoadBalancerSourceRanges
}

// GetLoadBalancerSourceRangesOk returns a tuple with the LoadBalancerSourceRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateProjectDeploymentRequest) GetLoadBalancerSourceRangesOk() ([]string, bool) {
	if o == nil || o.LoadBalancerSourceRanges == nil {
		return nil, false
	}
	return o.LoadBalancerSourceRanges, true
}

// HasLoadBalancerSourceRanges returns a boolean if a field has been set.
func (o *RequestsCreateProjectDeploymentRequest) HasLoadBalancerSourceRanges() bool {
	if o != nil && o.LoadBalancerSourceRanges != nil {
		return true
	}

	return false
}

// SetLoadBalancerSourceRanges gets a reference to the given []string and assigns it to the LoadBalancerSourceRanges field.
func (o *RequestsCreateProjectDeploymentRequest) SetLoadBalancerSourceRanges(v []string) {
	o.LoadBalancerSourceRanges = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RequestsCreateProjectDeploymentRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateProjectDeploymentRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RequestsCreateProjectDeploymentRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RequestsCreateProjectDeploymentRequest) SetName(v string) {
	o.Name = &v
}

// GetNamespaceId returns the NamespaceId field value if set, zero value otherwise.
func (o *RequestsCreateProjectDeploymentRequest) GetNamespaceId() string {
	if o == nil || o.NamespaceId == nil {
		var ret string
		return ret
	}
	return *o.NamespaceId
}

// GetNamespaceIdOk returns a tuple with the NamespaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateProjectDeploymentRequest) GetNamespaceIdOk() (*string, bool) {
	if o == nil || o.NamespaceId == nil {
		return nil, false
	}
	return o.NamespaceId, true
}

// HasNamespaceId returns a boolean if a field has been set.
func (o *RequestsCreateProjectDeploymentRequest) HasNamespaceId() bool {
	if o != nil && o.NamespaceId != nil {
		return true
	}

	return false
}

// SetNamespaceId gets a reference to the given string and assigns it to the NamespaceId field.
func (o *RequestsCreateProjectDeploymentRequest) SetNamespaceId(v string) {
	o.NamespaceId = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *RequestsCreateProjectDeploymentRequest) GetNodeCount() int32 {
	if o == nil || o.NodeCount == nil {
		var ret int32
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateProjectDeploymentRequest) GetNodeCountOk() (*int32, bool) {
	if o == nil || o.NodeCount == nil {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *RequestsCreateProjectDeploymentRequest) HasNodeCount() bool {
	if o != nil && o.NodeCount != nil {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int32 and assigns it to the NodeCount field.
func (o *RequestsCreateProjectDeploymentRequest) SetNodeCount(v int32) {
	o.NodeCount = &v
}

// GetResourceSettingsTemplateId returns the ResourceSettingsTemplateId field value if set, zero value otherwise.
func (o *RequestsCreateProjectDeploymentRequest) GetResourceSettingsTemplateId() string {
	if o == nil || o.ResourceSettingsTemplateId == nil {
		var ret string
		return ret
	}
	return *o.ResourceSettingsTemplateId
}

// GetResourceSettingsTemplateIdOk returns a tuple with the ResourceSettingsTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateProjectDeploymentRequest) GetResourceSettingsTemplateIdOk() (*string, bool) {
	if o == nil || o.ResourceSettingsTemplateId == nil {
		return nil, false
	}
	return o.ResourceSettingsTemplateId, true
}

// HasResourceSettingsTemplateId returns a boolean if a field has been set.
func (o *RequestsCreateProjectDeploymentRequest) HasResourceSettingsTemplateId() bool {
	if o != nil && o.ResourceSettingsTemplateId != nil {
		return true
	}

	return false
}

// SetResourceSettingsTemplateId gets a reference to the given string and assigns it to the ResourceSettingsTemplateId field.
func (o *RequestsCreateProjectDeploymentRequest) SetResourceSettingsTemplateId(v string) {
	o.ResourceSettingsTemplateId = &v
}

// GetScheduledBackup returns the ScheduledBackup field value if set, zero value otherwise.
func (o *RequestsCreateProjectDeploymentRequest) GetScheduledBackup() RequestsDeploymentScheduledBackup {
	if o == nil || o.ScheduledBackup == nil {
		var ret RequestsDeploymentScheduledBackup
		return ret
	}
	return *o.ScheduledBackup
}

// GetScheduledBackupOk returns a tuple with the ScheduledBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateProjectDeploymentRequest) GetScheduledBackupOk() (*RequestsDeploymentScheduledBackup, bool) {
	if o == nil || o.ScheduledBackup == nil {
		return nil, false
	}
	return o.ScheduledBackup, true
}

// HasScheduledBackup returns a boolean if a field has been set.
func (o *RequestsCreateProjectDeploymentRequest) HasScheduledBackup() bool {
	if o != nil && o.ScheduledBackup != nil {
		return true
	}

	return false
}

// SetScheduledBackup gets a reference to the given RequestsDeploymentScheduledBackup and assigns it to the ScheduledBackup field.
func (o *RequestsCreateProjectDeploymentRequest) SetScheduledBackup(v RequestsDeploymentScheduledBackup) {
	o.ScheduledBackup = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *RequestsCreateProjectDeploymentRequest) GetServiceType() string {
	if o == nil || o.ServiceType == nil {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateProjectDeploymentRequest) GetServiceTypeOk() (*string, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *RequestsCreateProjectDeploymentRequest) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *RequestsCreateProjectDeploymentRequest) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetStorageOptionsTemplateId returns the StorageOptionsTemplateId field value if set, zero value otherwise.
func (o *RequestsCreateProjectDeploymentRequest) GetStorageOptionsTemplateId() string {
	if o == nil || o.StorageOptionsTemplateId == nil {
		var ret string
		return ret
	}
	return *o.StorageOptionsTemplateId
}

// GetStorageOptionsTemplateIdOk returns a tuple with the StorageOptionsTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateProjectDeploymentRequest) GetStorageOptionsTemplateIdOk() (*string, bool) {
	if o == nil || o.StorageOptionsTemplateId == nil {
		return nil, false
	}
	return o.StorageOptionsTemplateId, true
}

// HasStorageOptionsTemplateId returns a boolean if a field has been set.
func (o *RequestsCreateProjectDeploymentRequest) HasStorageOptionsTemplateId() bool {
	if o != nil && o.StorageOptionsTemplateId != nil {
		return true
	}

	return false
}

// SetStorageOptionsTemplateId gets a reference to the given string and assigns it to the StorageOptionsTemplateId field.
func (o *RequestsCreateProjectDeploymentRequest) SetStorageOptionsTemplateId(v string) {
	o.StorageOptionsTemplateId = &v
}

// GetTlsEnabled returns the TlsEnabled field value if set, zero value otherwise.
func (o *RequestsCreateProjectDeploymentRequest) GetTlsEnabled() bool {
	if o == nil || o.TlsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TlsEnabled
}

// GetTlsEnabledOk returns a tuple with the TlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateProjectDeploymentRequest) GetTlsEnabledOk() (*bool, bool) {
	if o == nil || o.TlsEnabled == nil {
		return nil, false
	}
	return o.TlsEnabled, true
}

// HasTlsEnabled returns a boolean if a field has been set.
func (o *RequestsCreateProjectDeploymentRequest) HasTlsEnabled() bool {
	if o != nil && o.TlsEnabled != nil {
		return true
	}

	return false
}

// SetTlsEnabled gets a reference to the given bool and assigns it to the TlsEnabled field.
func (o *RequestsCreateProjectDeploymentRequest) SetTlsEnabled(v bool) {
	o.TlsEnabled = &v
}

func (o RequestsCreateProjectDeploymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationConfigurationOverrides != nil {
		toSerialize["application_configuration_overrides"] = o.ApplicationConfigurationOverrides
	}
	if o.ApplicationConfigurationTemplateId != nil {
		toSerialize["application_configuration_template_id"] = o.ApplicationConfigurationTemplateId
	}
	if o.DeploymentTargetId != nil {
		toSerialize["deployment_target_id"] = o.DeploymentTargetId
	}
	if o.DnsZone != nil {
		toSerialize["dns_zone"] = o.DnsZone
	}
	if o.ImageId != nil {
		toSerialize["image_id"] = o.ImageId
	}
	if o.LoadBalancerSourceRanges != nil {
		toSerialize["load_balancer_source_ranges"] = o.LoadBalancerSourceRanges
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NamespaceId != nil {
		toSerialize["namespace_id"] = o.NamespaceId
	}
	if o.NodeCount != nil {
		toSerialize["node_count"] = o.NodeCount
	}
	if o.ResourceSettingsTemplateId != nil {
		toSerialize["resource_settings_template_id"] = o.ResourceSettingsTemplateId
	}
	if o.ScheduledBackup != nil {
		toSerialize["scheduled_backup"] = o.ScheduledBackup
	}
	if o.ServiceType != nil {
		toSerialize["service_type"] = o.ServiceType
	}
	if o.StorageOptionsTemplateId != nil {
		toSerialize["storage_options_template_id"] = o.StorageOptionsTemplateId
	}
	if o.TlsEnabled != nil {
		toSerialize["tls_enabled"] = o.TlsEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableRequestsCreateProjectDeploymentRequest struct {
	value *RequestsCreateProjectDeploymentRequest
	isSet bool
}

func (v NullableRequestsCreateProjectDeploymentRequest) Get() *RequestsCreateProjectDeploymentRequest {
	return v.value
}

func (v *NullableRequestsCreateProjectDeploymentRequest) Set(val *RequestsCreateProjectDeploymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestsCreateProjectDeploymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestsCreateProjectDeploymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestsCreateProjectDeploymentRequest(val *RequestsCreateProjectDeploymentRequest) *NullableRequestsCreateProjectDeploymentRequest {
	return &NullableRequestsCreateProjectDeploymentRequest{value: val, isSet: true}
}

func (v NullableRequestsCreateProjectDeploymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestsCreateProjectDeploymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


