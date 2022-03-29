# ControllersUpdateDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationConfigurationOverrides** | Pointer to **map[string]string** |  | [optional] 
**ApplicationConfigurationTemplateId** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int32** | See models.Deployment for more information. | [optional] 
**ResourceSettingsTemplateId** | Pointer to **string** |  | [optional] 
**ScheduledBackup** | Pointer to [**ControllersUpdateDeploymentScheduledBackup**](ControllersUpdateDeploymentScheduledBackup.md) |  | [optional] 

## Methods

### NewControllersUpdateDeploymentRequest

`func NewControllersUpdateDeploymentRequest() *ControllersUpdateDeploymentRequest`

NewControllersUpdateDeploymentRequest instantiates a new ControllersUpdateDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersUpdateDeploymentRequestWithDefaults

`func NewControllersUpdateDeploymentRequestWithDefaults() *ControllersUpdateDeploymentRequest`

NewControllersUpdateDeploymentRequestWithDefaults instantiates a new ControllersUpdateDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationConfigurationOverrides

`func (o *ControllersUpdateDeploymentRequest) GetApplicationConfigurationOverrides() map[string]string`

GetApplicationConfigurationOverrides returns the ApplicationConfigurationOverrides field if non-nil, zero value otherwise.

### GetApplicationConfigurationOverridesOk

`func (o *ControllersUpdateDeploymentRequest) GetApplicationConfigurationOverridesOk() (*map[string]string, bool)`

GetApplicationConfigurationOverridesOk returns a tuple with the ApplicationConfigurationOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationConfigurationOverrides

`func (o *ControllersUpdateDeploymentRequest) SetApplicationConfigurationOverrides(v map[string]string)`

SetApplicationConfigurationOverrides sets ApplicationConfigurationOverrides field to given value.

### HasApplicationConfigurationOverrides

`func (o *ControllersUpdateDeploymentRequest) HasApplicationConfigurationOverrides() bool`

HasApplicationConfigurationOverrides returns a boolean if a field has been set.

### GetApplicationConfigurationTemplateId

`func (o *ControllersUpdateDeploymentRequest) GetApplicationConfigurationTemplateId() string`

GetApplicationConfigurationTemplateId returns the ApplicationConfigurationTemplateId field if non-nil, zero value otherwise.

### GetApplicationConfigurationTemplateIdOk

`func (o *ControllersUpdateDeploymentRequest) GetApplicationConfigurationTemplateIdOk() (*string, bool)`

GetApplicationConfigurationTemplateIdOk returns a tuple with the ApplicationConfigurationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationConfigurationTemplateId

`func (o *ControllersUpdateDeploymentRequest) SetApplicationConfigurationTemplateId(v string)`

SetApplicationConfigurationTemplateId sets ApplicationConfigurationTemplateId field to given value.

### HasApplicationConfigurationTemplateId

`func (o *ControllersUpdateDeploymentRequest) HasApplicationConfigurationTemplateId() bool`

HasApplicationConfigurationTemplateId returns a boolean if a field has been set.

### GetImageId

`func (o *ControllersUpdateDeploymentRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ControllersUpdateDeploymentRequest) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ControllersUpdateDeploymentRequest) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *ControllersUpdateDeploymentRequest) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetNodeCount

`func (o *ControllersUpdateDeploymentRequest) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ControllersUpdateDeploymentRequest) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ControllersUpdateDeploymentRequest) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ControllersUpdateDeploymentRequest) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetResourceSettingsTemplateId

`func (o *ControllersUpdateDeploymentRequest) GetResourceSettingsTemplateId() string`

GetResourceSettingsTemplateId returns the ResourceSettingsTemplateId field if non-nil, zero value otherwise.

### GetResourceSettingsTemplateIdOk

`func (o *ControllersUpdateDeploymentRequest) GetResourceSettingsTemplateIdOk() (*string, bool)`

GetResourceSettingsTemplateIdOk returns a tuple with the ResourceSettingsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSettingsTemplateId

`func (o *ControllersUpdateDeploymentRequest) SetResourceSettingsTemplateId(v string)`

SetResourceSettingsTemplateId sets ResourceSettingsTemplateId field to given value.

### HasResourceSettingsTemplateId

`func (o *ControllersUpdateDeploymentRequest) HasResourceSettingsTemplateId() bool`

HasResourceSettingsTemplateId returns a boolean if a field has been set.

### GetScheduledBackup

`func (o *ControllersUpdateDeploymentRequest) GetScheduledBackup() ControllersUpdateDeploymentScheduledBackup`

GetScheduledBackup returns the ScheduledBackup field if non-nil, zero value otherwise.

### GetScheduledBackupOk

`func (o *ControllersUpdateDeploymentRequest) GetScheduledBackupOk() (*ControllersUpdateDeploymentScheduledBackup, bool)`

GetScheduledBackupOk returns a tuple with the ScheduledBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledBackup

`func (o *ControllersUpdateDeploymentRequest) SetScheduledBackup(v ControllersUpdateDeploymentScheduledBackup)`

SetScheduledBackup sets ScheduledBackup field to given value.

### HasScheduledBackup

`func (o *ControllersUpdateDeploymentRequest) HasScheduledBackup() bool`

HasScheduledBackup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


