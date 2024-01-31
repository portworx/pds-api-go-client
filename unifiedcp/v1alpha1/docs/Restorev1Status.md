# Restorev1Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**ErrorCode** | Pointer to [**V1ErrorCode**](V1ErrorCode.md) |  | [optional] [default to V1ERRORCODE_ERROR_CODE_UNSPECIFIED]
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to [**V1Phase**](V1Phase.md) |  | [optional] [default to V1PHASE_PHASE_UNSPECIFIED]

## Methods

### NewRestorev1Status

`func NewRestorev1Status() *Restorev1Status`

NewRestorev1Status instantiates a new Restorev1Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestorev1StatusWithDefaults

`func NewRestorev1StatusWithDefaults() *Restorev1Status`

NewRestorev1StatusWithDefaults instantiates a new Restorev1Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartedAt

`func (o *Restorev1Status) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Restorev1Status) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Restorev1Status) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Restorev1Status) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *Restorev1Status) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Restorev1Status) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Restorev1Status) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *Restorev1Status) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetErrorCode

`func (o *Restorev1Status) GetErrorCode() V1ErrorCode`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Restorev1Status) GetErrorCodeOk() (*V1ErrorCode, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Restorev1Status) SetErrorCode(v V1ErrorCode)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *Restorev1Status) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Restorev1Status) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Restorev1Status) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Restorev1Status) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Restorev1Status) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetPhase

`func (o *Restorev1Status) GetPhase() V1Phase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *Restorev1Status) GetPhaseOk() (*V1Phase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *Restorev1Status) SetPhase(v V1Phase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *Restorev1Status) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


