# RequestsUpdateDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationConfigurationOverrides** | Pointer to **map[string]string** |  | [optional] 
**ApplicationConfigurationTemplateId** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int32** |  | [optional] 
**ResourceSettingsTemplateId** | Pointer to **string** |  | [optional] 
**ScheduledBackup** | Pointer to [**RequestsUpdateDeploymentScheduledBackup**](RequestsUpdateDeploymentScheduledBackup.md) |  | [optional] 
**TlsEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewRequestsUpdateDeploymentRequest

`func NewRequestsUpdateDeploymentRequest() *RequestsUpdateDeploymentRequest`

NewRequestsUpdateDeploymentRequest instantiates a new RequestsUpdateDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsUpdateDeploymentRequestWithDefaults

`func NewRequestsUpdateDeploymentRequestWithDefaults() *RequestsUpdateDeploymentRequest`

NewRequestsUpdateDeploymentRequestWithDefaults instantiates a new RequestsUpdateDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationConfigurationOverrides

`func (o *RequestsUpdateDeploymentRequest) GetApplicationConfigurationOverrides() map[string]string`

GetApplicationConfigurationOverrides returns the ApplicationConfigurationOverrides field if non-nil, zero value otherwise.

### GetApplicationConfigurationOverridesOk

`func (o *RequestsUpdateDeploymentRequest) GetApplicationConfigurationOverridesOk() (*map[string]string, bool)`

GetApplicationConfigurationOverridesOk returns a tuple with the ApplicationConfigurationOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationConfigurationOverrides

`func (o *RequestsUpdateDeploymentRequest) SetApplicationConfigurationOverrides(v map[string]string)`

SetApplicationConfigurationOverrides sets ApplicationConfigurationOverrides field to given value.

### HasApplicationConfigurationOverrides

`func (o *RequestsUpdateDeploymentRequest) HasApplicationConfigurationOverrides() bool`

HasApplicationConfigurationOverrides returns a boolean if a field has been set.

### GetApplicationConfigurationTemplateId

`func (o *RequestsUpdateDeploymentRequest) GetApplicationConfigurationTemplateId() string`

GetApplicationConfigurationTemplateId returns the ApplicationConfigurationTemplateId field if non-nil, zero value otherwise.

### GetApplicationConfigurationTemplateIdOk

`func (o *RequestsUpdateDeploymentRequest) GetApplicationConfigurationTemplateIdOk() (*string, bool)`

GetApplicationConfigurationTemplateIdOk returns a tuple with the ApplicationConfigurationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationConfigurationTemplateId

`func (o *RequestsUpdateDeploymentRequest) SetApplicationConfigurationTemplateId(v string)`

SetApplicationConfigurationTemplateId sets ApplicationConfigurationTemplateId field to given value.

### HasApplicationConfigurationTemplateId

`func (o *RequestsUpdateDeploymentRequest) HasApplicationConfigurationTemplateId() bool`

HasApplicationConfigurationTemplateId returns a boolean if a field has been set.

### GetImageId

`func (o *RequestsUpdateDeploymentRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *RequestsUpdateDeploymentRequest) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *RequestsUpdateDeploymentRequest) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *RequestsUpdateDeploymentRequest) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetNodeCount

`func (o *RequestsUpdateDeploymentRequest) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *RequestsUpdateDeploymentRequest) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *RequestsUpdateDeploymentRequest) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *RequestsUpdateDeploymentRequest) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetResourceSettingsTemplateId

`func (o *RequestsUpdateDeploymentRequest) GetResourceSettingsTemplateId() string`

GetResourceSettingsTemplateId returns the ResourceSettingsTemplateId field if non-nil, zero value otherwise.

### GetResourceSettingsTemplateIdOk

`func (o *RequestsUpdateDeploymentRequest) GetResourceSettingsTemplateIdOk() (*string, bool)`

GetResourceSettingsTemplateIdOk returns a tuple with the ResourceSettingsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSettingsTemplateId

`func (o *RequestsUpdateDeploymentRequest) SetResourceSettingsTemplateId(v string)`

SetResourceSettingsTemplateId sets ResourceSettingsTemplateId field to given value.

### HasResourceSettingsTemplateId

`func (o *RequestsUpdateDeploymentRequest) HasResourceSettingsTemplateId() bool`

HasResourceSettingsTemplateId returns a boolean if a field has been set.

### GetScheduledBackup

`func (o *RequestsUpdateDeploymentRequest) GetScheduledBackup() RequestsUpdateDeploymentScheduledBackup`

GetScheduledBackup returns the ScheduledBackup field if non-nil, zero value otherwise.

### GetScheduledBackupOk

`func (o *RequestsUpdateDeploymentRequest) GetScheduledBackupOk() (*RequestsUpdateDeploymentScheduledBackup, bool)`

GetScheduledBackupOk returns a tuple with the ScheduledBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledBackup

`func (o *RequestsUpdateDeploymentRequest) SetScheduledBackup(v RequestsUpdateDeploymentScheduledBackup)`

SetScheduledBackup sets ScheduledBackup field to given value.

### HasScheduledBackup

`func (o *RequestsUpdateDeploymentRequest) HasScheduledBackup() bool`

HasScheduledBackup returns a boolean if a field has been set.

### GetTlsEnabled

`func (o *RequestsUpdateDeploymentRequest) GetTlsEnabled() bool`

GetTlsEnabled returns the TlsEnabled field if non-nil, zero value otherwise.

### GetTlsEnabledOk

`func (o *RequestsUpdateDeploymentRequest) GetTlsEnabledOk() (*bool, bool)`

GetTlsEnabledOk returns a tuple with the TlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnabled

`func (o *RequestsUpdateDeploymentRequest) SetTlsEnabled(v bool)`

SetTlsEnabled sets TlsEnabled field to given value.

### HasTlsEnabled

`func (o *RequestsUpdateDeploymentRequest) HasTlsEnabled() bool`

HasTlsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


