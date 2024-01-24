# V1TargetCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V1Meta**](V1Meta.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to [**PlatformTargetClusterv1Status**](PlatformTargetClusterv1Status.md) |  | [optional] 

## Methods

### NewV1TargetCluster

`func NewV1TargetCluster() *V1TargetCluster`

NewV1TargetCluster instantiates a new V1TargetCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TargetClusterWithDefaults

`func NewV1TargetClusterWithDefaults() *V1TargetCluster`

NewV1TargetClusterWithDefaults instantiates a new V1TargetCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V1TargetCluster) GetMeta() V1Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V1TargetCluster) GetMetaOk() (*V1Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V1TargetCluster) SetMeta(v V1Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V1TargetCluster) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetConfig

`func (o *V1TargetCluster) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V1TargetCluster) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V1TargetCluster) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *V1TargetCluster) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *V1TargetCluster) GetStatus() PlatformTargetClusterv1Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1TargetCluster) GetStatusOk() (*PlatformTargetClusterv1Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1TargetCluster) SetStatus(v PlatformTargetClusterv1Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1TargetCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


