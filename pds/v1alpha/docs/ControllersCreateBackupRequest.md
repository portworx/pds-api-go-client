# ControllersCreateBackupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupLevel** | Pointer to **string** |  | [optional] 
**BackupTargetId** | Pointer to **string** |  | [optional] 
**BackupType** | Pointer to **string** |  | [optional] 
**DataServiceId** | Pointer to **string** |  | [optional] 
**DeploymentId** | Pointer to **string** |  | [optional] 
**JobHistoryLimit** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 

## Methods

### NewControllersCreateBackupRequest

`func NewControllersCreateBackupRequest() *ControllersCreateBackupRequest`

NewControllersCreateBackupRequest instantiates a new ControllersCreateBackupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersCreateBackupRequestWithDefaults

`func NewControllersCreateBackupRequestWithDefaults() *ControllersCreateBackupRequest`

NewControllersCreateBackupRequestWithDefaults instantiates a new ControllersCreateBackupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupLevel

`func (o *ControllersCreateBackupRequest) GetBackupLevel() string`

GetBackupLevel returns the BackupLevel field if non-nil, zero value otherwise.

### GetBackupLevelOk

`func (o *ControllersCreateBackupRequest) GetBackupLevelOk() (*string, bool)`

GetBackupLevelOk returns a tuple with the BackupLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLevel

`func (o *ControllersCreateBackupRequest) SetBackupLevel(v string)`

SetBackupLevel sets BackupLevel field to given value.

### HasBackupLevel

`func (o *ControllersCreateBackupRequest) HasBackupLevel() bool`

HasBackupLevel returns a boolean if a field has been set.

### GetBackupTargetId

`func (o *ControllersCreateBackupRequest) GetBackupTargetId() string`

GetBackupTargetId returns the BackupTargetId field if non-nil, zero value otherwise.

### GetBackupTargetIdOk

`func (o *ControllersCreateBackupRequest) GetBackupTargetIdOk() (*string, bool)`

GetBackupTargetIdOk returns a tuple with the BackupTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupTargetId

`func (o *ControllersCreateBackupRequest) SetBackupTargetId(v string)`

SetBackupTargetId sets BackupTargetId field to given value.

### HasBackupTargetId

`func (o *ControllersCreateBackupRequest) HasBackupTargetId() bool`

HasBackupTargetId returns a boolean if a field has been set.

### GetBackupType

`func (o *ControllersCreateBackupRequest) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *ControllersCreateBackupRequest) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *ControllersCreateBackupRequest) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *ControllersCreateBackupRequest) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetDataServiceId

`func (o *ControllersCreateBackupRequest) GetDataServiceId() string`

GetDataServiceId returns the DataServiceId field if non-nil, zero value otherwise.

### GetDataServiceIdOk

`func (o *ControllersCreateBackupRequest) GetDataServiceIdOk() (*string, bool)`

GetDataServiceIdOk returns a tuple with the DataServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataServiceId

`func (o *ControllersCreateBackupRequest) SetDataServiceId(v string)`

SetDataServiceId sets DataServiceId field to given value.

### HasDataServiceId

`func (o *ControllersCreateBackupRequest) HasDataServiceId() bool`

HasDataServiceId returns a boolean if a field has been set.

### GetDeploymentId

`func (o *ControllersCreateBackupRequest) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *ControllersCreateBackupRequest) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *ControllersCreateBackupRequest) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *ControllersCreateBackupRequest) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetJobHistoryLimit

`func (o *ControllersCreateBackupRequest) GetJobHistoryLimit() int32`

GetJobHistoryLimit returns the JobHistoryLimit field if non-nil, zero value otherwise.

### GetJobHistoryLimitOk

`func (o *ControllersCreateBackupRequest) GetJobHistoryLimitOk() (*int32, bool)`

GetJobHistoryLimitOk returns a tuple with the JobHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobHistoryLimit

`func (o *ControllersCreateBackupRequest) SetJobHistoryLimit(v int32)`

SetJobHistoryLimit sets JobHistoryLimit field to given value.

### HasJobHistoryLimit

`func (o *ControllersCreateBackupRequest) HasJobHistoryLimit() bool`

HasJobHistoryLimit returns a boolean if a field has been set.

### GetProjectId

`func (o *ControllersCreateBackupRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ControllersCreateBackupRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ControllersCreateBackupRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ControllersCreateBackupRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSchedule

`func (o *ControllersCreateBackupRequest) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ControllersCreateBackupRequest) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ControllersCreateBackupRequest) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ControllersCreateBackupRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


