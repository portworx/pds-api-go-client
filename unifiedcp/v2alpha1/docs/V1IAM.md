# V1IAM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V1Meta**](V1Meta.md) |  | [optional] 
**Config** | Pointer to [**V1Config6**](V1Config6.md) |  | [optional] 

## Methods

### NewV1IAM

`func NewV1IAM() *V1IAM`

NewV1IAM instantiates a new V1IAM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1IAMWithDefaults

`func NewV1IAMWithDefaults() *V1IAM`

NewV1IAMWithDefaults instantiates a new V1IAM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V1IAM) GetMeta() V1Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V1IAM) GetMetaOk() (*V1Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V1IAM) SetMeta(v V1Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V1IAM) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetConfig

`func (o *V1IAM) GetConfig() V1Config6`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V1IAM) GetConfigOk() (*V1Config6, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V1IAM) SetConfig(v V1Config6)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *V1IAM) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


