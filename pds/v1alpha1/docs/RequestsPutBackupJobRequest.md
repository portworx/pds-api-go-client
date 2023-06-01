# RequestsPutBackupJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupCapability** | Pointer to **string** |  | [optional] 
**BackupId** | Pointer to **string** |  | [optional] 
**BackupSpec** | Pointer to **map[string]interface{}** |  | [optional] 
**CloudCredentialName** | Pointer to **string** |  | [optional] 
**CloudSnapId** | Pointer to **string** |  | [optional] 
**CompletionStatus** | Pointer to **string** | CompletionStatus of the snapshot. | [optional] 
**CompletionTime** | Pointer to **string** |  | [optional] 
**DataServiceSpec** | Pointer to **map[string]interface{}** |  | [optional] 
**DeploymentId** | Pointer to **string** |  | [optional] 
**DeploymentTargetId** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** | ErrorCode if CompletionStatus is \&quot;Failed\&quot; | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**FileSize** | Pointer to **int32** |  | [optional] 
**ImageName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**ProjectId** | **string** |  | 
**StartTime** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 

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

### GetBackupCapability

`func (o *RequestsPutBackupJobRequest) GetBackupCapability() string`

GetBackupCapability returns the BackupCapability field if non-nil, zero value otherwise.

### GetBackupCapabilityOk

`func (o *RequestsPutBackupJobRequest) GetBackupCapabilityOk() (*string, bool)`

GetBackupCapabilityOk returns a tuple with the BackupCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCapability

`func (o *RequestsPutBackupJobRequest) SetBackupCapability(v string)`

SetBackupCapability sets BackupCapability field to given value.

### HasBackupCapability

`func (o *RequestsPutBackupJobRequest) HasBackupCapability() bool`

HasBackupCapability returns a boolean if a field has been set.

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

`func (o *RequestsPutBackupJobRequest) GetBackupSpec() map[string]interface{}`

GetBackupSpec returns the BackupSpec field if non-nil, zero value otherwise.

### GetBackupSpecOk

`func (o *RequestsPutBackupJobRequest) GetBackupSpecOk() (*map[string]interface{}, bool)`

GetBackupSpecOk returns a tuple with the BackupSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSpec

`func (o *RequestsPutBackupJobRequest) SetBackupSpec(v map[string]interface{})`

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

`func (o *RequestsPutBackupJobRequest) GetDataServiceSpec() map[string]interface{}`

GetDataServiceSpec returns the DataServiceSpec field if non-nil, zero value otherwise.

### GetDataServiceSpecOk

`func (o *RequestsPutBackupJobRequest) GetDataServiceSpecOk() (*map[string]interface{}, bool)`

GetDataServiceSpecOk returns a tuple with the DataServiceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataServiceSpec

`func (o *RequestsPutBackupJobRequest) SetDataServiceSpec(v map[string]interface{})`

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

### GetDeploymentTargetId

`func (o *RequestsPutBackupJobRequest) GetDeploymentTargetId() string`

GetDeploymentTargetId returns the DeploymentTargetId field if non-nil, zero value otherwise.

### GetDeploymentTargetIdOk

`func (o *RequestsPutBackupJobRequest) GetDeploymentTargetIdOk() (*string, bool)`

GetDeploymentTargetIdOk returns a tuple with the DeploymentTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTargetId

`func (o *RequestsPutBackupJobRequest) SetDeploymentTargetId(v string)`

SetDeploymentTargetId sets DeploymentTargetId field to given value.

### HasDeploymentTargetId

`func (o *RequestsPutBackupJobRequest) HasDeploymentTargetId() bool`

HasDeploymentTargetId returns a boolean if a field has been set.

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

### GetImageName

`func (o *RequestsPutBackupJobRequest) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *RequestsPutBackupJobRequest) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *RequestsPutBackupJobRequest) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *RequestsPutBackupJobRequest) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

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

### GetNamespace

`func (o *RequestsPutBackupJobRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *RequestsPutBackupJobRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *RequestsPutBackupJobRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *RequestsPutBackupJobRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

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

### GetTimestamp

`func (o *RequestsPutBackupJobRequest) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RequestsPutBackupJobRequest) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RequestsPutBackupJobRequest) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RequestsPutBackupJobRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


