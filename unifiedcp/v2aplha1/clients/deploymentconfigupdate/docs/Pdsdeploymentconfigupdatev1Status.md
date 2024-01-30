# Pdsdeploymentconfigupdatev1Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **string** | Error Code is a short string that represents the error. | [optional] 
**ErrorMessage** | Pointer to **string** | Error Message is a description of the error. | [optional] 
**RetryCount** | Pointer to **int32** |  | [optional] 
**Phase** | Pointer to [**Pdsdeploymentconfigupdatev1StatusPhase**](Pdsdeploymentconfigupdatev1StatusPhase.md) |  | [optional] [default to PHASE_UNSPECIFIED]

## Methods

### NewPdsdeploymentconfigupdatev1Status

`func NewPdsdeploymentconfigupdatev1Status() *Pdsdeploymentconfigupdatev1Status`

NewPdsdeploymentconfigupdatev1Status instantiates a new Pdsdeploymentconfigupdatev1Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPdsdeploymentconfigupdatev1StatusWithDefaults

`func NewPdsdeploymentconfigupdatev1StatusWithDefaults() *Pdsdeploymentconfigupdatev1Status`

NewPdsdeploymentconfigupdatev1StatusWithDefaults instantiates a new Pdsdeploymentconfigupdatev1Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *Pdsdeploymentconfigupdatev1Status) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Pdsdeploymentconfigupdatev1Status) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Pdsdeploymentconfigupdatev1Status) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *Pdsdeploymentconfigupdatev1Status) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Pdsdeploymentconfigupdatev1Status) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Pdsdeploymentconfigupdatev1Status) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Pdsdeploymentconfigupdatev1Status) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Pdsdeploymentconfigupdatev1Status) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetRetryCount

`func (o *Pdsdeploymentconfigupdatev1Status) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *Pdsdeploymentconfigupdatev1Status) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *Pdsdeploymentconfigupdatev1Status) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *Pdsdeploymentconfigupdatev1Status) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetPhase

`func (o *Pdsdeploymentconfigupdatev1Status) GetPhase() Pdsdeploymentconfigupdatev1StatusPhase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *Pdsdeploymentconfigupdatev1Status) GetPhaseOk() (*Pdsdeploymentconfigupdatev1StatusPhase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *Pdsdeploymentconfigupdatev1Status) SetPhase(v Pdsdeploymentconfigupdatev1StatusPhase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *Pdsdeploymentconfigupdatev1Status) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


