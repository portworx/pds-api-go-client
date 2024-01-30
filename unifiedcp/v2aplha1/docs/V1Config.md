# V1Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**References** | Pointer to [**V1References1**](V1References1.md) |  | [optional] 
**JobHistoryLimit** | Pointer to **int32** | Job History Limit is a number of retained backup jobs. Must be greater than or equal to 1. | [optional] 
**Schedule** | Pointer to [**V1Schedule**](V1Schedule.md) |  | [optional] 
**Suspend** | Pointer to **bool** | Suspend flag is used to suspend a scheduled backup from creating new backup jobs. | [optional] 
**BackupType** | Pointer to [**ConfigBackupType**](ConfigBackupType.md) |  | [optional] [default to CONFIGBACKUPTYPE_BACKUP_TYPE_UNSPECIFIED]
**BackupLevel** | Pointer to [**ConfigBackupLevel**](ConfigBackupLevel.md) |  | [optional] [default to CONFIGBACKUPLEVEL_BACKUP_LEVEL_UNSPECIFIED]
**ReclaimPolicy** | Pointer to [**ConfigReclaimPolicyType**](ConfigReclaimPolicyType.md) |  | [optional] [default to CONFIGRECLAIMPOLICYTYPE_RECLAIM_POLICY_TYPE_UNSPECIFIED]

## Methods

### NewV1Config

`func NewV1Config() *V1Config`

NewV1Config instantiates a new V1Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ConfigWithDefaults

`func NewV1ConfigWithDefaults() *V1Config`

NewV1ConfigWithDefaults instantiates a new V1Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferences

`func (o *V1Config) GetReferences() V1References1`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *V1Config) GetReferencesOk() (*V1References1, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *V1Config) SetReferences(v V1References1)`

SetReferences sets References field to given value.

### HasReferences

`func (o *V1Config) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetJobHistoryLimit

`func (o *V1Config) GetJobHistoryLimit() int32`

GetJobHistoryLimit returns the JobHistoryLimit field if non-nil, zero value otherwise.

### GetJobHistoryLimitOk

`func (o *V1Config) GetJobHistoryLimitOk() (*int32, bool)`

GetJobHistoryLimitOk returns a tuple with the JobHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobHistoryLimit

`func (o *V1Config) SetJobHistoryLimit(v int32)`

SetJobHistoryLimit sets JobHistoryLimit field to given value.

### HasJobHistoryLimit

`func (o *V1Config) HasJobHistoryLimit() bool`

HasJobHistoryLimit returns a boolean if a field has been set.

### GetSchedule

`func (o *V1Config) GetSchedule() V1Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *V1Config) GetScheduleOk() (*V1Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *V1Config) SetSchedule(v V1Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *V1Config) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetSuspend

`func (o *V1Config) GetSuspend() bool`

GetSuspend returns the Suspend field if non-nil, zero value otherwise.

### GetSuspendOk

`func (o *V1Config) GetSuspendOk() (*bool, bool)`

GetSuspendOk returns a tuple with the Suspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspend

`func (o *V1Config) SetSuspend(v bool)`

SetSuspend sets Suspend field to given value.

### HasSuspend

`func (o *V1Config) HasSuspend() bool`

HasSuspend returns a boolean if a field has been set.

### GetBackupType

`func (o *V1Config) GetBackupType() ConfigBackupType`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *V1Config) GetBackupTypeOk() (*ConfigBackupType, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *V1Config) SetBackupType(v ConfigBackupType)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *V1Config) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetBackupLevel

`func (o *V1Config) GetBackupLevel() ConfigBackupLevel`

GetBackupLevel returns the BackupLevel field if non-nil, zero value otherwise.

### GetBackupLevelOk

`func (o *V1Config) GetBackupLevelOk() (*ConfigBackupLevel, bool)`

GetBackupLevelOk returns a tuple with the BackupLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLevel

`func (o *V1Config) SetBackupLevel(v ConfigBackupLevel)`

SetBackupLevel sets BackupLevel field to given value.

### HasBackupLevel

`func (o *V1Config) HasBackupLevel() bool`

HasBackupLevel returns a boolean if a field has been set.

### GetReclaimPolicy

`func (o *V1Config) GetReclaimPolicy() ConfigReclaimPolicyType`

GetReclaimPolicy returns the ReclaimPolicy field if non-nil, zero value otherwise.

### GetReclaimPolicyOk

`func (o *V1Config) GetReclaimPolicyOk() (*ConfigReclaimPolicyType, bool)`

GetReclaimPolicyOk returns a tuple with the ReclaimPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimPolicy

`func (o *V1Config) SetReclaimPolicy(v ConfigReclaimPolicyType)`

SetReclaimPolicy sets ReclaimPolicy field to given value.

### HasReclaimPolicy

`func (o *V1Config) HasReclaimPolicy() bool`

HasReclaimPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


