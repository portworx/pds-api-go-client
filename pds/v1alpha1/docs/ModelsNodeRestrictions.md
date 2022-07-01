# ModelsNodeRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownsizeEnabled** | Pointer to **bool** | Whether downsizing the cluster is enabled during editing. | [optional] 
**ResizeDisabled** | Pointer to **[]int32** | List of all node sizes that are not allowed to change during editing. | [optional] 
**SimpleConfig** | Pointer to **[]int32** | List of all allowed node sizes. | [optional] 

## Methods

### NewModelsNodeRestrictions

`func NewModelsNodeRestrictions() *ModelsNodeRestrictions`

NewModelsNodeRestrictions instantiates a new ModelsNodeRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsNodeRestrictionsWithDefaults

`func NewModelsNodeRestrictionsWithDefaults() *ModelsNodeRestrictions`

NewModelsNodeRestrictionsWithDefaults instantiates a new ModelsNodeRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownsizeEnabled

`func (o *ModelsNodeRestrictions) GetDownsizeEnabled() bool`

GetDownsizeEnabled returns the DownsizeEnabled field if non-nil, zero value otherwise.

### GetDownsizeEnabledOk

`func (o *ModelsNodeRestrictions) GetDownsizeEnabledOk() (*bool, bool)`

GetDownsizeEnabledOk returns a tuple with the DownsizeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownsizeEnabled

`func (o *ModelsNodeRestrictions) SetDownsizeEnabled(v bool)`

SetDownsizeEnabled sets DownsizeEnabled field to given value.

### HasDownsizeEnabled

`func (o *ModelsNodeRestrictions) HasDownsizeEnabled() bool`

HasDownsizeEnabled returns a boolean if a field has been set.

### GetResizeDisabled

`func (o *ModelsNodeRestrictions) GetResizeDisabled() []int32`

GetResizeDisabled returns the ResizeDisabled field if non-nil, zero value otherwise.

### GetResizeDisabledOk

`func (o *ModelsNodeRestrictions) GetResizeDisabledOk() (*[]int32, bool)`

GetResizeDisabledOk returns a tuple with the ResizeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeDisabled

`func (o *ModelsNodeRestrictions) SetResizeDisabled(v []int32)`

SetResizeDisabled sets ResizeDisabled field to given value.

### HasResizeDisabled

`func (o *ModelsNodeRestrictions) HasResizeDisabled() bool`

HasResizeDisabled returns a boolean if a field has been set.

### GetSimpleConfig

`func (o *ModelsNodeRestrictions) GetSimpleConfig() []int32`

GetSimpleConfig returns the SimpleConfig field if non-nil, zero value otherwise.

### GetSimpleConfigOk

`func (o *ModelsNodeRestrictions) GetSimpleConfigOk() (*[]int32, bool)`

GetSimpleConfigOk returns a tuple with the SimpleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleConfig

`func (o *ModelsNodeRestrictions) SetSimpleConfig(v []int32)`

SetSimpleConfig sets SimpleConfig field to given value.

### HasSimpleConfig

`func (o *ModelsNodeRestrictions) HasSimpleConfig() bool`

HasSimpleConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


