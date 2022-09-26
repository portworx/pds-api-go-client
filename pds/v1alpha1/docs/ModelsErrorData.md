# ModelsErrorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | More detailed specification of the error when the &#39;state&#39; is in one of the failed states. | [optional] 
**Details** | Pointer to **string** | More detailed error message possibly containing the root cause. | [optional] 
**Message** | Pointer to **string** | High level human-readable error message determined by the ErrorCode. | [optional] 

## Methods

### NewModelsErrorData

`func NewModelsErrorData() *ModelsErrorData`

NewModelsErrorData instantiates a new ModelsErrorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsErrorDataWithDefaults

`func NewModelsErrorDataWithDefaults() *ModelsErrorData`

NewModelsErrorDataWithDefaults instantiates a new ModelsErrorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ModelsErrorData) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ModelsErrorData) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ModelsErrorData) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ModelsErrorData) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetails

`func (o *ModelsErrorData) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ModelsErrorData) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ModelsErrorData) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ModelsErrorData) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMessage

`func (o *ModelsErrorData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModelsErrorData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModelsErrorData) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ModelsErrorData) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


