# ModelsBackupTargetState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupTargetId** | Pointer to **string** |  | [optional] 
**DeploymentTargetId** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** | Predefined API error code. | [optional] 
**ErrorDetails** | Pointer to **string** | More detailed error message possibly containing the root cause of the error, not suitable to show in the UI. | [optional] 
**ErrorMessage** | Pointer to **string** | High level human-readable error message determined by the ErrorCode. | [optional] 
**PxCredentialsId** | Pointer to **string** | ID of the credentials in PX cluster. | [optional] 
**PxCredentialsName** | Pointer to **string** | Name of the credentials in PX cluster. This will be used when creating a new backup. | [optional] 
**State** | Pointer to **string** | State of the synchronization of credentials. | [optional] 

## Methods

### NewModelsBackupTargetState

`func NewModelsBackupTargetState() *ModelsBackupTargetState`

NewModelsBackupTargetState instantiates a new ModelsBackupTargetState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsBackupTargetStateWithDefaults

`func NewModelsBackupTargetStateWithDefaults() *ModelsBackupTargetState`

NewModelsBackupTargetStateWithDefaults instantiates a new ModelsBackupTargetState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupTargetId

`func (o *ModelsBackupTargetState) GetBackupTargetId() string`

GetBackupTargetId returns the BackupTargetId field if non-nil, zero value otherwise.

### GetBackupTargetIdOk

`func (o *ModelsBackupTargetState) GetBackupTargetIdOk() (*string, bool)`

GetBackupTargetIdOk returns a tuple with the BackupTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupTargetId

`func (o *ModelsBackupTargetState) SetBackupTargetId(v string)`

SetBackupTargetId sets BackupTargetId field to given value.

### HasBackupTargetId

`func (o *ModelsBackupTargetState) HasBackupTargetId() bool`

HasBackupTargetId returns a boolean if a field has been set.

### GetDeploymentTargetId

`func (o *ModelsBackupTargetState) GetDeploymentTargetId() string`

GetDeploymentTargetId returns the DeploymentTargetId field if non-nil, zero value otherwise.

### GetDeploymentTargetIdOk

`func (o *ModelsBackupTargetState) GetDeploymentTargetIdOk() (*string, bool)`

GetDeploymentTargetIdOk returns a tuple with the DeploymentTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTargetId

`func (o *ModelsBackupTargetState) SetDeploymentTargetId(v string)`

SetDeploymentTargetId sets DeploymentTargetId field to given value.

### HasDeploymentTargetId

`func (o *ModelsBackupTargetState) HasDeploymentTargetId() bool`

HasDeploymentTargetId returns a boolean if a field has been set.

### GetErrorCode

`func (o *ModelsBackupTargetState) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ModelsBackupTargetState) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ModelsBackupTargetState) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ModelsBackupTargetState) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDetails

`func (o *ModelsBackupTargetState) GetErrorDetails() string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *ModelsBackupTargetState) GetErrorDetailsOk() (*string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *ModelsBackupTargetState) SetErrorDetails(v string)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *ModelsBackupTargetState) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ModelsBackupTargetState) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ModelsBackupTargetState) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ModelsBackupTargetState) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ModelsBackupTargetState) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetPxCredentialsId

`func (o *ModelsBackupTargetState) GetPxCredentialsId() string`

GetPxCredentialsId returns the PxCredentialsId field if non-nil, zero value otherwise.

### GetPxCredentialsIdOk

`func (o *ModelsBackupTargetState) GetPxCredentialsIdOk() (*string, bool)`

GetPxCredentialsIdOk returns a tuple with the PxCredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxCredentialsId

`func (o *ModelsBackupTargetState) SetPxCredentialsId(v string)`

SetPxCredentialsId sets PxCredentialsId field to given value.

### HasPxCredentialsId

`func (o *ModelsBackupTargetState) HasPxCredentialsId() bool`

HasPxCredentialsId returns a boolean if a field has been set.

### GetPxCredentialsName

`func (o *ModelsBackupTargetState) GetPxCredentialsName() string`

GetPxCredentialsName returns the PxCredentialsName field if non-nil, zero value otherwise.

### GetPxCredentialsNameOk

`func (o *ModelsBackupTargetState) GetPxCredentialsNameOk() (*string, bool)`

GetPxCredentialsNameOk returns a tuple with the PxCredentialsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxCredentialsName

`func (o *ModelsBackupTargetState) SetPxCredentialsName(v string)`

SetPxCredentialsName sets PxCredentialsName field to given value.

### HasPxCredentialsName

`func (o *ModelsBackupTargetState) HasPxCredentialsName() bool`

HasPxCredentialsName returns a boolean if a field has been set.

### GetState

`func (o *ModelsBackupTargetState) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ModelsBackupTargetState) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ModelsBackupTargetState) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ModelsBackupTargetState) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


