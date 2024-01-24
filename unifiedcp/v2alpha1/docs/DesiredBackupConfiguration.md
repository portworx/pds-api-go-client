# DesiredBackupConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetadataOfTheBackupConfiguration**](MetadataOfTheBackupConfiguration.md) |  | [optional] 
**Config** | Pointer to [**V1Config**](V1Config.md) |  | [optional] 
**Status** | Pointer to [**Backupconfigv1Status**](Backupconfigv1Status.md) |  | [optional] 

## Methods

### NewDesiredBackupConfiguration

`func NewDesiredBackupConfiguration() *DesiredBackupConfiguration`

NewDesiredBackupConfiguration instantiates a new DesiredBackupConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesiredBackupConfigurationWithDefaults

`func NewDesiredBackupConfigurationWithDefaults() *DesiredBackupConfiguration`

NewDesiredBackupConfigurationWithDefaults instantiates a new DesiredBackupConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *DesiredBackupConfiguration) GetMeta() MetadataOfTheBackupConfiguration`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DesiredBackupConfiguration) GetMetaOk() (*MetadataOfTheBackupConfiguration, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DesiredBackupConfiguration) SetMeta(v MetadataOfTheBackupConfiguration)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DesiredBackupConfiguration) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetConfig

`func (o *DesiredBackupConfiguration) GetConfig() V1Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DesiredBackupConfiguration) GetConfigOk() (*V1Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DesiredBackupConfiguration) SetConfig(v V1Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DesiredBackupConfiguration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *DesiredBackupConfiguration) GetStatus() Backupconfigv1Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DesiredBackupConfiguration) GetStatusOk() (*Backupconfigv1Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DesiredBackupConfiguration) SetStatus(v Backupconfigv1Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DesiredBackupConfiguration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


