# TargetClusterDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetadataOfTheTargetClusterResourceMeta**](MetadataOfTheTargetClusterResourceMeta.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to [**PlatformTargetClusterv1Status**](PlatformTargetClusterv1Status.md) |  | [optional] 

## Methods

### NewTargetClusterDetails

`func NewTargetClusterDetails() *TargetClusterDetails`

NewTargetClusterDetails instantiates a new TargetClusterDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetClusterDetailsWithDefaults

`func NewTargetClusterDetailsWithDefaults() *TargetClusterDetails`

NewTargetClusterDetailsWithDefaults instantiates a new TargetClusterDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *TargetClusterDetails) GetMeta() MetadataOfTheTargetClusterResourceMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TargetClusterDetails) GetMetaOk() (*MetadataOfTheTargetClusterResourceMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TargetClusterDetails) SetMeta(v MetadataOfTheTargetClusterResourceMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TargetClusterDetails) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetConfig

`func (o *TargetClusterDetails) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *TargetClusterDetails) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *TargetClusterDetails) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *TargetClusterDetails) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *TargetClusterDetails) GetStatus() PlatformTargetClusterv1Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TargetClusterDetails) GetStatusOk() (*PlatformTargetClusterv1Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TargetClusterDetails) SetStatus(v PlatformTargetClusterv1Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TargetClusterDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


