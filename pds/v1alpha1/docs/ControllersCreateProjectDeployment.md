# ControllersCreateProjectDeployment

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
**NodeCount** | Pointer to **int32** | See models.Deployment for more information. | [optional] 
**ResourceSettingsTemplateId** | Pointer to **string** |  | [optional] 
**ScheduledBackup** | Pointer to [**ControllersCreateDeploymentScheduledBackup**](ControllersCreateDeploymentScheduledBackup.md) |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**StorageOptionsTemplateId** | Pointer to **string** |  | [optional] 

## Methods

### NewControllersCreateProjectDeployment

`func NewControllersCreateProjectDeployment() *ControllersCreateProjectDeployment`

NewControllersCreateProjectDeployment instantiates a new ControllersCreateProjectDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersCreateProjectDeploymentWithDefaults

`func NewControllersCreateProjectDeploymentWithDefaults() *ControllersCreateProjectDeployment`

NewControllersCreateProjectDeploymentWithDefaults instantiates a new ControllersCreateProjectDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationConfigurationOverrides

`func (o *ControllersCreateProjectDeployment) GetApplicationConfigurationOverrides() map[string]string`

GetApplicationConfigurationOverrides returns the ApplicationConfigurationOverrides field if non-nil, zero value otherwise.

### GetApplicationConfigurationOverridesOk

`func (o *ControllersCreateProjectDeployment) GetApplicationConfigurationOverridesOk() (*map[string]string, bool)`

GetApplicationConfigurationOverridesOk returns a tuple with the ApplicationConfigurationOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationConfigurationOverrides

`func (o *ControllersCreateProjectDeployment) SetApplicationConfigurationOverrides(v map[string]string)`

SetApplicationConfigurationOverrides sets ApplicationConfigurationOverrides field to given value.

### HasApplicationConfigurationOverrides

`func (o *ControllersCreateProjectDeployment) HasApplicationConfigurationOverrides() bool`

HasApplicationConfigurationOverrides returns a boolean if a field has been set.

### GetApplicationConfigurationTemplateId

`func (o *ControllersCreateProjectDeployment) GetApplicationConfigurationTemplateId() string`

GetApplicationConfigurationTemplateId returns the ApplicationConfigurationTemplateId field if non-nil, zero value otherwise.

### GetApplicationConfigurationTemplateIdOk

`func (o *ControllersCreateProjectDeployment) GetApplicationConfigurationTemplateIdOk() (*string, bool)`

GetApplicationConfigurationTemplateIdOk returns a tuple with the ApplicationConfigurationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationConfigurationTemplateId

`func (o *ControllersCreateProjectDeployment) SetApplicationConfigurationTemplateId(v string)`

SetApplicationConfigurationTemplateId sets ApplicationConfigurationTemplateId field to given value.

### HasApplicationConfigurationTemplateId

`func (o *ControllersCreateProjectDeployment) HasApplicationConfigurationTemplateId() bool`

HasApplicationConfigurationTemplateId returns a boolean if a field has been set.

### GetDeploymentTargetId

`func (o *ControllersCreateProjectDeployment) GetDeploymentTargetId() string`

GetDeploymentTargetId returns the DeploymentTargetId field if non-nil, zero value otherwise.

### GetDeploymentTargetIdOk

`func (o *ControllersCreateProjectDeployment) GetDeploymentTargetIdOk() (*string, bool)`

GetDeploymentTargetIdOk returns a tuple with the DeploymentTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTargetId

`func (o *ControllersCreateProjectDeployment) SetDeploymentTargetId(v string)`

SetDeploymentTargetId sets DeploymentTargetId field to given value.

### HasDeploymentTargetId

`func (o *ControllersCreateProjectDeployment) HasDeploymentTargetId() bool`

HasDeploymentTargetId returns a boolean if a field has been set.

### GetDnsZone

`func (o *ControllersCreateProjectDeployment) GetDnsZone() string`

GetDnsZone returns the DnsZone field if non-nil, zero value otherwise.

### GetDnsZoneOk

`func (o *ControllersCreateProjectDeployment) GetDnsZoneOk() (*string, bool)`

GetDnsZoneOk returns a tuple with the DnsZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZone

`func (o *ControllersCreateProjectDeployment) SetDnsZone(v string)`

SetDnsZone sets DnsZone field to given value.

### HasDnsZone

`func (o *ControllersCreateProjectDeployment) HasDnsZone() bool`

HasDnsZone returns a boolean if a field has been set.

### GetImageId

`func (o *ControllersCreateProjectDeployment) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ControllersCreateProjectDeployment) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ControllersCreateProjectDeployment) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *ControllersCreateProjectDeployment) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetLoadBalancerSourceRanges

`func (o *ControllersCreateProjectDeployment) GetLoadBalancerSourceRanges() []string`

GetLoadBalancerSourceRanges returns the LoadBalancerSourceRanges field if non-nil, zero value otherwise.

### GetLoadBalancerSourceRangesOk

`func (o *ControllersCreateProjectDeployment) GetLoadBalancerSourceRangesOk() (*[]string, bool)`

GetLoadBalancerSourceRangesOk returns a tuple with the LoadBalancerSourceRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerSourceRanges

`func (o *ControllersCreateProjectDeployment) SetLoadBalancerSourceRanges(v []string)`

SetLoadBalancerSourceRanges sets LoadBalancerSourceRanges field to given value.

### HasLoadBalancerSourceRanges

`func (o *ControllersCreateProjectDeployment) HasLoadBalancerSourceRanges() bool`

HasLoadBalancerSourceRanges returns a boolean if a field has been set.

### GetName

`func (o *ControllersCreateProjectDeployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControllersCreateProjectDeployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControllersCreateProjectDeployment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControllersCreateProjectDeployment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespaceId

`func (o *ControllersCreateProjectDeployment) GetNamespaceId() string`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *ControllersCreateProjectDeployment) GetNamespaceIdOk() (*string, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *ControllersCreateProjectDeployment) SetNamespaceId(v string)`

SetNamespaceId sets NamespaceId field to given value.

### HasNamespaceId

`func (o *ControllersCreateProjectDeployment) HasNamespaceId() bool`

HasNamespaceId returns a boolean if a field has been set.

### GetNodeCount

`func (o *ControllersCreateProjectDeployment) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ControllersCreateProjectDeployment) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ControllersCreateProjectDeployment) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ControllersCreateProjectDeployment) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetResourceSettingsTemplateId

`func (o *ControllersCreateProjectDeployment) GetResourceSettingsTemplateId() string`

GetResourceSettingsTemplateId returns the ResourceSettingsTemplateId field if non-nil, zero value otherwise.

### GetResourceSettingsTemplateIdOk

`func (o *ControllersCreateProjectDeployment) GetResourceSettingsTemplateIdOk() (*string, bool)`

GetResourceSettingsTemplateIdOk returns a tuple with the ResourceSettingsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSettingsTemplateId

`func (o *ControllersCreateProjectDeployment) SetResourceSettingsTemplateId(v string)`

SetResourceSettingsTemplateId sets ResourceSettingsTemplateId field to given value.

### HasResourceSettingsTemplateId

`func (o *ControllersCreateProjectDeployment) HasResourceSettingsTemplateId() bool`

HasResourceSettingsTemplateId returns a boolean if a field has been set.

### GetScheduledBackup

`func (o *ControllersCreateProjectDeployment) GetScheduledBackup() ControllersCreateDeploymentScheduledBackup`

GetScheduledBackup returns the ScheduledBackup field if non-nil, zero value otherwise.

### GetScheduledBackupOk

`func (o *ControllersCreateProjectDeployment) GetScheduledBackupOk() (*ControllersCreateDeploymentScheduledBackup, bool)`

GetScheduledBackupOk returns a tuple with the ScheduledBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledBackup

`func (o *ControllersCreateProjectDeployment) SetScheduledBackup(v ControllersCreateDeploymentScheduledBackup)`

SetScheduledBackup sets ScheduledBackup field to given value.

### HasScheduledBackup

`func (o *ControllersCreateProjectDeployment) HasScheduledBackup() bool`

HasScheduledBackup returns a boolean if a field has been set.

### GetServiceType

`func (o *ControllersCreateProjectDeployment) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ControllersCreateProjectDeployment) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ControllersCreateProjectDeployment) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *ControllersCreateProjectDeployment) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetStorageOptionsTemplateId

`func (o *ControllersCreateProjectDeployment) GetStorageOptionsTemplateId() string`

GetStorageOptionsTemplateId returns the StorageOptionsTemplateId field if non-nil, zero value otherwise.

### GetStorageOptionsTemplateIdOk

`func (o *ControllersCreateProjectDeployment) GetStorageOptionsTemplateIdOk() (*string, bool)`

GetStorageOptionsTemplateIdOk returns a tuple with the StorageOptionsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageOptionsTemplateId

`func (o *ControllersCreateProjectDeployment) SetStorageOptionsTemplateId(v string)`

SetStorageOptionsTemplateId sets StorageOptionsTemplateId field to given value.

### HasStorageOptionsTemplateId

`func (o *ControllersCreateProjectDeployment) HasStorageOptionsTemplateId() bool`

HasStorageOptionsTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


