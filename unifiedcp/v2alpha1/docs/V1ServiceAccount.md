# V1ServiceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V1Meta**](V1Meta.md) |  | [optional] 
**Config** | Pointer to [**V1Config8**](V1Config8.md) |  | [optional] 
**Status** | Pointer to [**Serviceaccountv1Status**](Serviceaccountv1Status.md) |  | [optional] 

## Methods

### NewV1ServiceAccount

`func NewV1ServiceAccount() *V1ServiceAccount`

NewV1ServiceAccount instantiates a new V1ServiceAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ServiceAccountWithDefaults

`func NewV1ServiceAccountWithDefaults() *V1ServiceAccount`

NewV1ServiceAccountWithDefaults instantiates a new V1ServiceAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V1ServiceAccount) GetMeta() V1Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V1ServiceAccount) GetMetaOk() (*V1Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V1ServiceAccount) SetMeta(v V1Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V1ServiceAccount) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetConfig

`func (o *V1ServiceAccount) GetConfig() V1Config8`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V1ServiceAccount) GetConfigOk() (*V1Config8, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V1ServiceAccount) SetConfig(v V1Config8)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *V1ServiceAccount) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *V1ServiceAccount) GetStatus() Serviceaccountv1Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1ServiceAccount) GetStatusOk() (*Serviceaccountv1Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1ServiceAccount) SetStatus(v Serviceaccountv1Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1ServiceAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


