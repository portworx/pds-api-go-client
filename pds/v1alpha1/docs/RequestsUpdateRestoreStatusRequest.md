# RequestsUpdateRestoreStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionStatus** | Pointer to **string** | Status of the restore process on the Target Cluster | [optional] 
**CompletionTime** | Pointer to **string** | CompletionTime of the restore process. | [optional] 
**ErrorCode** | Pointer to **string** | Error code of the restore from Target Cluster | [optional] 
**StartTime** | Pointer to **string** | StartTime of the restore process. | [optional] 

## Methods

### NewRequestsUpdateRestoreStatusRequest

`func NewRequestsUpdateRestoreStatusRequest() *RequestsUpdateRestoreStatusRequest`

NewRequestsUpdateRestoreStatusRequest instantiates a new RequestsUpdateRestoreStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsUpdateRestoreStatusRequestWithDefaults

`func NewRequestsUpdateRestoreStatusRequestWithDefaults() *RequestsUpdateRestoreStatusRequest`

NewRequestsUpdateRestoreStatusRequestWithDefaults instantiates a new RequestsUpdateRestoreStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionStatus

`func (o *RequestsUpdateRestoreStatusRequest) GetCompletionStatus() string`

GetCompletionStatus returns the CompletionStatus field if non-nil, zero value otherwise.

### GetCompletionStatusOk

`func (o *RequestsUpdateRestoreStatusRequest) GetCompletionStatusOk() (*string, bool)`

GetCompletionStatusOk returns a tuple with the CompletionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionStatus

`func (o *RequestsUpdateRestoreStatusRequest) SetCompletionStatus(v string)`

SetCompletionStatus sets CompletionStatus field to given value.

### HasCompletionStatus

`func (o *RequestsUpdateRestoreStatusRequest) HasCompletionStatus() bool`

HasCompletionStatus returns a boolean if a field has been set.

### GetCompletionTime

`func (o *RequestsUpdateRestoreStatusRequest) GetCompletionTime() string`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *RequestsUpdateRestoreStatusRequest) GetCompletionTimeOk() (*string, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *RequestsUpdateRestoreStatusRequest) SetCompletionTime(v string)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *RequestsUpdateRestoreStatusRequest) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetErrorCode

`func (o *RequestsUpdateRestoreStatusRequest) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *RequestsUpdateRestoreStatusRequest) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *RequestsUpdateRestoreStatusRequest) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *RequestsUpdateRestoreStatusRequest) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetStartTime

`func (o *RequestsUpdateRestoreStatusRequest) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RequestsUpdateRestoreStatusRequest) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RequestsUpdateRestoreStatusRequest) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RequestsUpdateRestoreStatusRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


