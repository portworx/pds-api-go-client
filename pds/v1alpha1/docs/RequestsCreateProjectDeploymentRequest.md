# RequestsCreateProjectDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationConfigurationOverrides** | Pointer to **map[string]string** |  | [optional] 
**ApplicationConfigurationTemplateId** | Pointer to **string** |  | [optional] 
**DeploymentTargetId** | Pointer to **string** |  | [optional] 
**DnsZone** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**LoadBalancerSourceRanges** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NamespaceId** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int32** |  | [optional] 
**ResourceSettingsTemplateId** | Pointer to **string** |  | [optional] 
**ScheduledBackup** | Pointer to [**RequestsDeploymentScheduledBackup**](RequestsDeploymentScheduledBackup.md) |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**StorageOptionsTemplateId** | Pointer to **string** |  | [optional] 
**TlsEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewRequestsCreateProjectDeploymentRequest

`func NewRequestsCreateProjectDeploymentRequest() *RequestsCreateProjectDeploymentRequest`

NewRequestsCreateProjectDeploymentRequest instantiates a new RequestsCreateProjectDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsCreateProjectDeploymentRequestWithDefaults

`func NewRequestsCreateProjectDeploymentRequestWithDefaults() *RequestsCreateProjectDeploymentRequest`

NewRequestsCreateProjectDeploymentRequestWithDefaults instantiates a new RequestsCreateProjectDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationConfigurationOverrides

`func (o *RequestsCreateProjectDeploymentRequest) GetApplicationConfigurationOverrides() map[string]string`

GetApplicationConfigurationOverrides returns the ApplicationConfigurationOverrides field if non-nil, zero value otherwise.

### GetApplicationConfigurationOverridesOk

`func (o *RequestsCreateProjectDeploymentRequest) GetApplicationConfigurationOverridesOk() (*map[string]string, bool)`

GetApplicationConfigurationOverridesOk returns a tuple with the ApplicationConfigurationOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationConfigurationOverrides

`func (o *RequestsCreateProjectDeploymentRequest) SetApplicationConfigurationOverrides(v map[string]string)`

SetApplicationConfigurationOverrides sets ApplicationConfigurationOverrides field to given value.

### HasApplicationConfigurationOverrides

`func (o *RequestsCreateProjectDeploymentRequest) HasApplicationConfigurationOverrides() bool`

HasApplicationConfigurationOverrides returns a boolean if a field has been set.

### GetApplicationConfigurationTemplateId

`func (o *RequestsCreateProjectDeploymentRequest) GetApplicationConfigurationTemplateId() string`

GetApplicationConfigurationTemplateId returns the ApplicationConfigurationTemplateId field if non-nil, zero value otherwise.

### GetApplicationConfigurationTemplateIdOk

`func (o *RequestsCreateProjectDeploymentRequest) GetApplicationConfigurationTemplateIdOk() (*string, bool)`

GetApplicationConfigurationTemplateIdOk returns a tuple with the ApplicationConfigurationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationConfigurationTemplateId

`func (o *RequestsCreateProjectDeploymentRequest) SetApplicationConfigurationTemplateId(v string)`

SetApplicationConfigurationTemplateId sets ApplicationConfigurationTemplateId field to given value.

### HasApplicationConfigurationTemplateId

`func (o *RequestsCreateProjectDeploymentRequest) HasApplicationConfigurationTemplateId() bool`

HasApplicationConfigurationTemplateId returns a boolean if a field has been set.

### GetDeploymentTargetId

`func (o *RequestsCreateProjectDeploymentRequest) GetDeploymentTargetId() string`

GetDeploymentTargetId returns the DeploymentTargetId field if non-nil, zero value otherwise.

### GetDeploymentTargetIdOk

`func (o *RequestsCreateProjectDeploymentRequest) GetDeploymentTargetIdOk() (*string, bool)`

GetDeploymentTargetIdOk returns a tuple with the DeploymentTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTargetId

`func (o *RequestsCreateProjectDeploymentRequest) SetDeploymentTargetId(v string)`

SetDeploymentTargetId sets DeploymentTargetId field to given value.

### HasDeploymentTargetId

`func (o *RequestsCreateProjectDeploymentRequest) HasDeploymentTargetId() bool`

HasDeploymentTargetId returns a boolean if a field has been set.

### GetDnsZone

`func (o *RequestsCreateProjectDeploymentRequest) GetDnsZone() string`

GetDnsZone returns the DnsZone field if non-nil, zero value otherwise.

### GetDnsZoneOk

`func (o *RequestsCreateProjectDeploymentRequest) GetDnsZoneOk() (*string, bool)`

GetDnsZoneOk returns a tuple with the DnsZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZone

`func (o *RequestsCreateProjectDeploymentRequest) SetDnsZone(v string)`

SetDnsZone sets DnsZone field to given value.

### HasDnsZone

`func (o *RequestsCreateProjectDeploymentRequest) HasDnsZone() bool`

HasDnsZone returns a boolean if a field has been set.

### GetImageId

`func (o *RequestsCreateProjectDeploymentRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *RequestsCreateProjectDeploymentRequest) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *RequestsCreateProjectDeploymentRequest) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *RequestsCreateProjectDeploymentRequest) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetLoadBalancerSourceRanges

`func (o *RequestsCreateProjectDeploymentRequest) GetLoadBalancerSourceRanges() []string`

GetLoadBalancerSourceRanges returns the LoadBalancerSourceRanges field if non-nil, zero value otherwise.

### GetLoadBalancerSourceRangesOk

`func (o *RequestsCreateProjectDeploymentRequest) GetLoadBalancerSourceRangesOk() (*[]string, bool)`

GetLoadBalancerSourceRangesOk returns a tuple with the LoadBalancerSourceRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerSourceRanges

`func (o *RequestsCreateProjectDeploymentRequest) SetLoadBalancerSourceRanges(v []string)`

SetLoadBalancerSourceRanges sets LoadBalancerSourceRanges field to given value.

### HasLoadBalancerSourceRanges

`func (o *RequestsCreateProjectDeploymentRequest) HasLoadBalancerSourceRanges() bool`

HasLoadBalancerSourceRanges returns a boolean if a field has been set.

### GetName

`func (o *RequestsCreateProjectDeploymentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestsCreateProjectDeploymentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestsCreateProjectDeploymentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequestsCreateProjectDeploymentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespaceId

`func (o *RequestsCreateProjectDeploymentRequest) GetNamespaceId() string`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *RequestsCreateProjectDeploymentRequest) GetNamespaceIdOk() (*string, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *RequestsCreateProjectDeploymentRequest) SetNamespaceId(v string)`

SetNamespaceId sets NamespaceId field to given value.

### HasNamespaceId

`func (o *RequestsCreateProjectDeploymentRequest) HasNamespaceId() bool`

HasNamespaceId returns a boolean if a field has been set.

### GetNodeCount

`func (o *RequestsCreateProjectDeploymentRequest) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *RequestsCreateProjectDeploymentRequest) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *RequestsCreateProjectDeploymentRequest) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *RequestsCreateProjectDeploymentRequest) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetResourceSettingsTemplateId

`func (o *RequestsCreateProjectDeploymentRequest) GetResourceSettingsTemplateId() string`

GetResourceSettingsTemplateId returns the ResourceSettingsTemplateId field if non-nil, zero value otherwise.

### GetResourceSettingsTemplateIdOk

`func (o *RequestsCreateProjectDeploymentRequest) GetResourceSettingsTemplateIdOk() (*string, bool)`

GetResourceSettingsTemplateIdOk returns a tuple with the ResourceSettingsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSettingsTemplateId

`func (o *RequestsCreateProjectDeploymentRequest) SetResourceSettingsTemplateId(v string)`

SetResourceSettingsTemplateId sets ResourceSettingsTemplateId field to given value.

### HasResourceSettingsTemplateId

`func (o *RequestsCreateProjectDeploymentRequest) HasResourceSettingsTemplateId() bool`

HasResourceSettingsTemplateId returns a boolean if a field has been set.

### GetScheduledBackup

`func (o *RequestsCreateProjectDeploymentRequest) GetScheduledBackup() RequestsDeploymentScheduledBackup`

GetScheduledBackup returns the ScheduledBackup field if non-nil, zero value otherwise.

### GetScheduledBackupOk

`func (o *RequestsCreateProjectDeploymentRequest) GetScheduledBackupOk() (*RequestsDeploymentScheduledBackup, bool)`

GetScheduledBackupOk returns a tuple with the ScheduledBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledBackup

`func (o *RequestsCreateProjectDeploymentRequest) SetScheduledBackup(v RequestsDeploymentScheduledBackup)`

SetScheduledBackup sets ScheduledBackup field to given value.

### HasScheduledBackup

`func (o *RequestsCreateProjectDeploymentRequest) HasScheduledBackup() bool`

HasScheduledBackup returns a boolean if a field has been set.

### GetServiceType

`func (o *RequestsCreateProjectDeploymentRequest) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *RequestsCreateProjectDeploymentRequest) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *RequestsCreateProjectDeploymentRequest) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *RequestsCreateProjectDeploymentRequest) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetStorageOptionsTemplateId

`func (o *RequestsCreateProjectDeploymentRequest) GetStorageOptionsTemplateId() string`

GetStorageOptionsTemplateId returns the StorageOptionsTemplateId field if non-nil, zero value otherwise.

### GetStorageOptionsTemplateIdOk

`func (o *RequestsCreateProjectDeploymentRequest) GetStorageOptionsTemplateIdOk() (*string, bool)`

GetStorageOptionsTemplateIdOk returns a tuple with the StorageOptionsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageOptionsTemplateId

`func (o *RequestsCreateProjectDeploymentRequest) SetStorageOptionsTemplateId(v string)`

SetStorageOptionsTemplateId sets StorageOptionsTemplateId field to given value.

### HasStorageOptionsTemplateId

`func (o *RequestsCreateProjectDeploymentRequest) HasStorageOptionsTemplateId() bool`

HasStorageOptionsTemplateId returns a boolean if a field has been set.

### GetTlsEnabled

`func (o *RequestsCreateProjectDeploymentRequest) GetTlsEnabled() bool`

GetTlsEnabled returns the TlsEnabled field if non-nil, zero value otherwise.

### GetTlsEnabledOk

`func (o *RequestsCreateProjectDeploymentRequest) GetTlsEnabledOk() (*bool, bool)`

GetTlsEnabledOk returns a tuple with the TlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnabled

`func (o *RequestsCreateProjectDeploymentRequest) SetTlsEnabled(v bool)`

SetTlsEnabled sets TlsEnabled field to given value.

### HasTlsEnabled

`func (o *RequestsCreateProjectDeploymentRequest) HasTlsEnabled() bool`

HasTlsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


