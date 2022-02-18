# ModelsBackupSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The only purpose of this is to help UI identify validation failures | [optional] 
**RetentionCount** | Pointer to **int32** |  | [optional] 
**Schedule** | Pointer to **string** | CRON string | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsBackupSchedule

`func NewModelsBackupSchedule() *ModelsBackupSchedule`

NewModelsBackupSchedule instantiates a new ModelsBackupSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsBackupScheduleWithDefaults

`func NewModelsBackupScheduleWithDefaults() *ModelsBackupSchedule`

NewModelsBackupScheduleWithDefaults instantiates a new ModelsBackupSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsBackupSchedule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsBackupSchedule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsBackupSchedule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsBackupSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRetentionCount

`func (o *ModelsBackupSchedule) GetRetentionCount() int32`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *ModelsBackupSchedule) GetRetentionCountOk() (*int32, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *ModelsBackupSchedule) SetRetentionCount(v int32)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *ModelsBackupSchedule) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### GetSchedule

`func (o *ModelsBackupSchedule) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ModelsBackupSchedule) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ModelsBackupSchedule) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ModelsBackupSchedule) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetType

`func (o *ModelsBackupSchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelsBackupSchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelsBackupSchedule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelsBackupSchedule) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


