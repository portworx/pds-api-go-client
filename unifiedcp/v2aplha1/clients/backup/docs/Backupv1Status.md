# Backupv1Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudSnapId** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**CompletionTime** | Pointer to **time.Time** |  | [optional] 
**Phase** | Pointer to [**StatusPhase**](StatusPhase.md) |  | [optional] [default to PHASE_UNSPECIFIED]
**ErrorCode** | Pointer to **string** | ErrorCode if CompletionStatus is \&quot;Failed\&quot;. | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**FileSize** | Pointer to **string** | FileSize of the CloudSnap image. | [optional] 

## Methods

### NewBackupv1Status

`func NewBackupv1Status() *Backupv1Status`

NewBackupv1Status instantiates a new Backupv1Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupv1StatusWithDefaults

`func NewBackupv1StatusWithDefaults() *Backupv1Status`

NewBackupv1StatusWithDefaults instantiates a new Backupv1Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudSnapId

`func (o *Backupv1Status) GetCloudSnapId() string`

GetCloudSnapId returns the CloudSnapId field if non-nil, zero value otherwise.

### GetCloudSnapIdOk

`func (o *Backupv1Status) GetCloudSnapIdOk() (*string, bool)`

GetCloudSnapIdOk returns a tuple with the CloudSnapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSnapId

`func (o *Backupv1Status) SetCloudSnapId(v string)`

SetCloudSnapId sets CloudSnapId field to given value.

### HasCloudSnapId

`func (o *Backupv1Status) HasCloudSnapId() bool`

HasCloudSnapId returns a boolean if a field has been set.

### GetStartTime

`func (o *Backupv1Status) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Backupv1Status) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Backupv1Status) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Backupv1Status) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetCompletionTime

`func (o *Backupv1Status) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *Backupv1Status) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *Backupv1Status) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *Backupv1Status) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetPhase

`func (o *Backupv1Status) GetPhase() StatusPhase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *Backupv1Status) GetPhaseOk() (*StatusPhase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *Backupv1Status) SetPhase(v StatusPhase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *Backupv1Status) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetErrorCode

`func (o *Backupv1Status) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Backupv1Status) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Backupv1Status) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *Backupv1Status) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Backupv1Status) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Backupv1Status) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Backupv1Status) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Backupv1Status) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetFileSize

`func (o *Backupv1Status) GetFileSize() string`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *Backupv1Status) GetFileSizeOk() (*string, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *Backupv1Status) SetFileSize(v string)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *Backupv1Status) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


