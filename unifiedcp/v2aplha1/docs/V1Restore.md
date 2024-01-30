# V1Restore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V1Meta**](V1Meta.md) |  | [optional] 
**Config** | Pointer to [**V1Config3**](V1Config3.md) |  | [optional] 
**Status** | Pointer to [**Restorev1Status**](Restorev1Status.md) |  | [optional] 

## Methods

### NewV1Restore

`func NewV1Restore() *V1Restore`

NewV1Restore instantiates a new V1Restore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RestoreWithDefaults

`func NewV1RestoreWithDefaults() *V1Restore`

NewV1RestoreWithDefaults instantiates a new V1Restore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V1Restore) GetMeta() V1Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V1Restore) GetMetaOk() (*V1Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V1Restore) SetMeta(v V1Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V1Restore) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetConfig

`func (o *V1Restore) GetConfig() V1Config3`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V1Restore) GetConfigOk() (*V1Config3, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V1Restore) SetConfig(v V1Config3)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *V1Restore) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *V1Restore) GetStatus() Restorev1Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1Restore) GetStatusOk() (*Restorev1Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1Restore) SetStatus(v Restorev1Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1Restore) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


