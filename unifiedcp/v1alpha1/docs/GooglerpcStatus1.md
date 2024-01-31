# GooglerpcStatus1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**[]ProtobufAny2**](ProtobufAny2.md) |  | [optional] 

## Methods

### NewGooglerpcStatus1

`func NewGooglerpcStatus1() *GooglerpcStatus1`

NewGooglerpcStatus1 instantiates a new GooglerpcStatus1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGooglerpcStatus1WithDefaults

`func NewGooglerpcStatus1WithDefaults() *GooglerpcStatus1`

NewGooglerpcStatus1WithDefaults instantiates a new GooglerpcStatus1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GooglerpcStatus1) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GooglerpcStatus1) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GooglerpcStatus1) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GooglerpcStatus1) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *GooglerpcStatus1) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GooglerpcStatus1) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GooglerpcStatus1) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GooglerpcStatus1) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetails

`func (o *GooglerpcStatus1) GetDetails() []ProtobufAny2`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GooglerpcStatus1) GetDetailsOk() (*[]ProtobufAny2, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GooglerpcStatus1) SetDetails(v []ProtobufAny2)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GooglerpcStatus1) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


