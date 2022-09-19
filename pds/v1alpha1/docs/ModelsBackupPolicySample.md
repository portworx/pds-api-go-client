# ModelsBackupPolicySample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the backup policy. Must be unique for the given tenant. | [optional] 
**Schedules** | Pointer to [**[]ModelsBackupSchedule**](ModelsBackupSchedule.md) | A list of the backup schedules. Must be non-empty. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


