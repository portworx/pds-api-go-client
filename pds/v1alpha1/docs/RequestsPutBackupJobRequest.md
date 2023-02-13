# RequestsPutBackupJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupId** | Pointer to **string** | BackupID which backup created the snapshot (nullable). | [optional] 
**BackupSpec** | Pointer to **[]int32** | BackupSpecJSON is the specification of the Backup in JSON format at the time the snapshot was taken. | [optional] 
**CloudCredentialName** | Pointer to **string** | CloudCredentialName credentials to access snapshot. | [optional] 
**CloudSnapId** | Pointer to **string** | CloudSnapID snapshot of the backup volume. | [optional] 
**CompletionStatus** | Pointer to **string** | CompletionStatus of the backup job. | [optional] 
**CompletionTime** | Pointer to **string** |  | [optional] 
**DataServiceSpec** | Pointer to **[]int32** | BackupSpecJSON is the specification of the Backup in JSON format at the time the snapshot was taken. | [optional] 
**DeploymentId** | Pointer to **string** | DeploymentID which deployment was backed up (nullable). | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**FileSize** | Pointer to **int32** | FileSize of the CloudSnap image. | [optional] 
**Name** | Pointer to **string** | Name of the BackupJob. | [optional] 
**ProjectId** | **string** | ProjectID which created the snapshot (required). | 
**StartTime** | Pointer to **string** |  | [optional] 

## Methods

### NewRequestsPutBackupJobRequest

`func NewRequestsPutBackupJobRequest(projectId string, ) *RequestsPutBackupJobRequest`

NewRequestsPutBackupJobRequest instantiates a new RequestsPutBackupJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsPutBackupJobRequestWithDefaults

`func NewRequestsPutBackupJobRequestWithDefaults() *RequestsPutBackupJobRequest`

NewRequestsPutBackupJobRequestWithDefaults instantiates a new RequestsPutBackupJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupId

`func (o *RequestsPutBackupJobRequest) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *RequestsPutBackupJobRequest) GetBackupIdOk() (*string, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *RequestsPutBackupJobRequest) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.

### HasBackupId

`func (o *RequestsPutBackupJobRequest) HasBackupId() bool`

HasBackupId returns a boolean if a field has been set.

### GetBackupSpec

`func (o *RequestsPutBackupJobRequest) GetBackupSpec() []int32`

GetBackupSpec returns the BackupSpec field if non-nil, zero value otherwise.

### GetBackupSpecOk

`func (o *RequestsPutBackupJobRequest) GetBackupSpecOk() (*[]int32, bool)`

GetBackupSpecOk returns a tuple with the BackupSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSpec

`func (o *RequestsPutBackupJobRequest) SetBackupSpec(v []int32)`

SetBackupSpec sets BackupSpec field to given value.

### HasBackupSpec

`func (o *RequestsPutBackupJobRequest) HasBackupSpec() bool`

HasBackupSpec returns a boolean if a field has been set.

### GetCloudCredentialName

`func (o *RequestsPutBackupJobRequest) GetCloudCredentialName() string`

GetCloudCredentialName returns the CloudCredentialName field if non-nil, zero value otherwise.

### GetCloudCredentialNameOk

`func (o *RequestsPutBackupJobRequest) GetCloudCredentialNameOk() (*string, bool)`

GetCloudCredentialNameOk returns a tuple with the CloudCredentialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialName

`func (o *RequestsPutBackupJobRequest) SetCloudCredentialName(v string)`

SetCloudCredentialName sets CloudCredentialName field to given value.

### HasCloudCredentialName

`func (o *RequestsPutBackupJobRequest) HasCloudCredentialName() bool`

HasCloudCredentialName returns a boolean if a field has been set.

### GetCloudSnapId

`func (o *RequestsPutBackupJobRequest) GetCloudSnapId() string`

GetCloudSnapId returns the CloudSnapId field if non-nil, zero value otherwise.

### GetCloudSnapIdOk

`func (o *RequestsPutBackupJobRequest) GetCloudSnapIdOk() (*string, bool)`

GetCloudSnapIdOk returns a tuple with the CloudSnapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSnapId

`func (o *RequestsPutBackupJobRequest) SetCloudSnapId(v string)`

SetCloudSnapId sets CloudSnapId field to given value.

### HasCloudSnapId

`func (o *RequestsPutBackupJobRequest) HasCloudSnapId() bool`

HasCloudSnapId returns a boolean if a field has been set.

### GetCompletionStatus

`func (o *RequestsPutBackupJobRequest) GetCompletionStatus() string`

GetCompletionStatus returns the CompletionStatus field if non-nil, zero value otherwise.

### GetCompletionStatusOk

`func (o *RequestsPutBackupJobRequest) GetCompletionStatusOk() (*string, bool)`

GetCompletionStatusOk returns a tuple with the CompletionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionStatus

`func (o *RequestsPutBackupJobRequest) SetCompletionStatus(v string)`

SetCompletionStatus sets CompletionStatus field to given value.

### HasCompletionStatus

`func (o *RequestsPutBackupJobRequest) HasCompletionStatus() bool`

HasCompletionStatus returns a boolean if a field has been set.

### GetCompletionTime

`func (o *RequestsPutBackupJobRequest) GetCompletionTime() string`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *RequestsPutBackupJobRequest) GetCompletionTimeOk() (*string, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *RequestsPutBackupJobRequest) SetCompletionTime(v string)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *RequestsPutBackupJobRequest) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetDataServiceSpec

`func (o *RequestsPutBackupJobRequest) GetDataServiceSpec() []int32`

GetDataServiceSpec returns the DataServiceSpec field if non-nil, zero value otherwise.

### GetDataServiceSpecOk

`func (o *RequestsPutBackupJobRequest) GetDataServiceSpecOk() (*[]int32, bool)`

GetDataServiceSpecOk returns a tuple with the DataServiceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataServiceSpec

`func (o *RequestsPutBackupJobRequest) SetDataServiceSpec(v []int32)`

SetDataServiceSpec sets DataServiceSpec field to given value.

### HasDataServiceSpec

`func (o *RequestsPutBackupJobRequest) HasDataServiceSpec() bool`

HasDataServiceSpec returns a boolean if a field has been set.

### GetDeploymentId

`func (o *RequestsPutBackupJobRequest) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *RequestsPutBackupJobRequest) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *RequestsPutBackupJobRequest) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *RequestsPutBackupJobRequest) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetErrorCode

`func (o *RequestsPutBackupJobRequest) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *RequestsPutBackupJobRequest) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *RequestsPutBackupJobRequest) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *RequestsPutBackupJobRequest) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *RequestsPutBackupJobRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *RequestsPutBackupJobRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *RequestsPutBackupJobRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *RequestsPutBackupJobRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetFileSize

`func (o *RequestsPutBackupJobRequest) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *RequestsPutBackupJobRequest) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *RequestsPutBackupJobRequest) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *RequestsPutBackupJobRequest) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetName

`func (o *RequestsPutBackupJobRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestsPutBackupJobRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestsPutBackupJobRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequestsPutBackupJobRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *RequestsPutBackupJobRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RequestsPutBackupJobRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RequestsPutBackupJobRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetStartTime

`func (o *RequestsPutBackupJobRequest) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RequestsPutBackupJobRequest) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RequestsPutBackupJobRequest) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RequestsPutBackupJobRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


