# V1Namespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V1Meta**](V1Meta.md) |  | [optional] 
**Status** | Pointer to [**V1NamespaceStatus**](V1NamespaceStatus.md) |  | [optional] 

## Methods

### NewV1Namespace

`func NewV1Namespace() *V1Namespace`

NewV1Namespace instantiates a new V1Namespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NamespaceWithDefaults

`func NewV1NamespaceWithDefaults() *V1Namespace`

NewV1NamespaceWithDefaults instantiates a new V1Namespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V1Namespace) GetMeta() V1Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V1Namespace) GetMetaOk() (*V1Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V1Namespace) SetMeta(v V1Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V1Namespace) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetStatus

`func (o *V1Namespace) GetStatus() V1NamespaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1Namespace) GetStatusOk() (*V1NamespaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1Namespace) SetStatus(v V1NamespaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1Namespace) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


