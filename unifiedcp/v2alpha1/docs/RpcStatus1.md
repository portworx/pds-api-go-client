# RpcStatus1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**[]ProtobufAny4**](ProtobufAny4.md) |  | [optional] 

## Methods

### NewRpcStatus1

`func NewRpcStatus1() *RpcStatus1`

NewRpcStatus1 instantiates a new RpcStatus1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcStatus1WithDefaults

`func NewRpcStatus1WithDefaults() *RpcStatus1`

NewRpcStatus1WithDefaults instantiates a new RpcStatus1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *RpcStatus1) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RpcStatus1) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RpcStatus1) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *RpcStatus1) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *RpcStatus1) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RpcStatus1) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RpcStatus1) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RpcStatus1) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetails

`func (o *RpcStatus1) GetDetails() []ProtobufAny4`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RpcStatus1) GetDetailsOk() (*[]ProtobufAny4, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RpcStatus1) SetDetails(v []ProtobufAny4)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RpcStatus1) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


