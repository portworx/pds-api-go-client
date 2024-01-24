# V1Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V1Meta**](V1Meta.md) |  | [optional] 
**Config** | Pointer to [**V1Config4**](V1Config4.md) |  | [optional] 

## Methods

### NewV1Account

`func NewV1Account() *V1Account`

NewV1Account instantiates a new V1Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AccountWithDefaults

`func NewV1AccountWithDefaults() *V1Account`

NewV1AccountWithDefaults instantiates a new V1Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V1Account) GetMeta() V1Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V1Account) GetMetaOk() (*V1Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V1Account) SetMeta(v V1Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V1Account) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetConfig

`func (o *V1Account) GetConfig() V1Config4`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V1Account) GetConfigOk() (*V1Config4, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V1Account) SetConfig(v V1Config4)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *V1Account) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


