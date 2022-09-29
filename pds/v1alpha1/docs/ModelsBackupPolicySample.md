# ModelsBackupPolicySample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name of the backup policy. Must be unique for the given tenant. | [optional] 
**Schedules** | Pointer to [**[]ModelsBackupSchedule**](ModelsBackupSchedule.md) | A list of the backup schedules. Must be non-empty. | [optional] 
**Updated** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelsBackupPolicySample

`func NewModelsBackupPolicySample() *ModelsBackupPolicySample`

NewModelsBackupPolicySample instantiates a new ModelsBackupPolicySample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsBackupPolicySampleWithDefaults

`func NewModelsBackupPolicySampleWithDefaults() *ModelsBackupPolicySample`

NewModelsBackupPolicySampleWithDefaults instantiates a new ModelsBackupPolicySample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ModelsBackupPolicySample) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsBackupPolicySample) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsBackupPolicySample) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsBackupPolicySample) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetName

`func (o *ModelsBackupPolicySample) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsBackupPolicySample) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsBackupPolicySample) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsBackupPolicySample) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchedules

`func (o *ModelsBackupPolicySample) GetSchedules() []ModelsBackupSchedule`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *ModelsBackupPolicySample) GetSchedulesOk() (*[]ModelsBackupSchedule, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *ModelsBackupPolicySample) SetSchedules(v []ModelsBackupSchedule)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *ModelsBackupPolicySample) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### GetUpdated

`func (o *ModelsBackupPolicySample) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ModelsBackupPolicySample) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ModelsBackupPolicySample) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ModelsBackupPolicySample) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetVersion

`func (o *ModelsBackupPolicySample) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelsBackupPolicySample) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelsBackupPolicySample) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelsBackupPolicySample) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


