# V1Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V1Meta**](V1Meta.md) |  | [optional] 
**Status** | Pointer to [**Taskv1Status**](Taskv1Status.md) |  | [optional] 

## Methods

### NewV1Task

`func NewV1Task() *V1Task`

NewV1Task instantiates a new V1Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TaskWithDefaults

`func NewV1TaskWithDefaults() *V1Task`

NewV1TaskWithDefaults instantiates a new V1Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V1Task) GetMeta() V1Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V1Task) GetMetaOk() (*V1Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V1Task) SetMeta(v V1Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V1Task) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetStatus

`func (o *V1Task) GetStatus() Taskv1Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1Task) GetStatusOk() (*Taskv1Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1Task) SetStatus(v Taskv1Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1Task) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


