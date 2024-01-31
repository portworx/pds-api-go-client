# V1Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V1Meta**](V1Meta.md) |  | [optional] 
**Info** | Pointer to [**V1Info**](V1Info.md) |  | [optional] 

## Methods

### NewV1Image

`func NewV1Image() *V1Image`

NewV1Image instantiates a new V1Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageWithDefaults

`func NewV1ImageWithDefaults() *V1Image`

NewV1ImageWithDefaults instantiates a new V1Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V1Image) GetMeta() V1Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V1Image) GetMetaOk() (*V1Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V1Image) SetMeta(v V1Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V1Image) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetInfo

`func (o *V1Image) GetInfo() V1Info`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *V1Image) GetInfoOk() (*V1Info, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *V1Image) SetInfo(v V1Info)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *V1Image) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


