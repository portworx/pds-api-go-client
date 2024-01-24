# V1Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V1Meta**](V1Meta.md) |  | [optional] 
**Config** | Pointer to [**V1Config7**](V1Config7.md) |  | [optional] 
**Status** | Pointer to [**Projectv1Status**](Projectv1Status.md) |  | [optional] 

## Methods

### NewV1Project

`func NewV1Project() *V1Project`

NewV1Project instantiates a new V1Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ProjectWithDefaults

`func NewV1ProjectWithDefaults() *V1Project`

NewV1ProjectWithDefaults instantiates a new V1Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V1Project) GetMeta() V1Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V1Project) GetMetaOk() (*V1Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V1Project) SetMeta(v V1Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V1Project) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetConfig

`func (o *V1Project) GetConfig() V1Config7`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V1Project) GetConfigOk() (*V1Config7, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V1Project) SetConfig(v V1Config7)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *V1Project) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *V1Project) GetStatus() Projectv1Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1Project) GetStatusOk() (*Projectv1Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1Project) SetStatus(v Projectv1Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1Project) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


