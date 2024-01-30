# V1NodeRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedSizes** | Pointer to **[]string** | List of all allowed node sizes. | [optional] 
**DownsizeEnabled** | Pointer to **bool** | Flag to determine whether downsizing the cluster is enabled during editing. | [optional] 
**ResizeDisabledSizes** | Pointer to **[]string** | List of all node sizes that are not allowed to change during editing. | [optional] 

## Methods

### NewV1NodeRestrictions

`func NewV1NodeRestrictions() *V1NodeRestrictions`

NewV1NodeRestrictions instantiates a new V1NodeRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NodeRestrictionsWithDefaults

`func NewV1NodeRestrictionsWithDefaults() *V1NodeRestrictions`

NewV1NodeRestrictionsWithDefaults instantiates a new V1NodeRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedSizes

`func (o *V1NodeRestrictions) GetAllowedSizes() []string`

GetAllowedSizes returns the AllowedSizes field if non-nil, zero value otherwise.

### GetAllowedSizesOk

`func (o *V1NodeRestrictions) GetAllowedSizesOk() (*[]string, bool)`

GetAllowedSizesOk returns a tuple with the AllowedSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSizes

`func (o *V1NodeRestrictions) SetAllowedSizes(v []string)`

SetAllowedSizes sets AllowedSizes field to given value.

### HasAllowedSizes

`func (o *V1NodeRestrictions) HasAllowedSizes() bool`

HasAllowedSizes returns a boolean if a field has been set.

### GetDownsizeEnabled

`func (o *V1NodeRestrictions) GetDownsizeEnabled() bool`

GetDownsizeEnabled returns the DownsizeEnabled field if non-nil, zero value otherwise.

### GetDownsizeEnabledOk

`func (o *V1NodeRestrictions) GetDownsizeEnabledOk() (*bool, bool)`

GetDownsizeEnabledOk returns a tuple with the DownsizeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownsizeEnabled

`func (o *V1NodeRestrictions) SetDownsizeEnabled(v bool)`

SetDownsizeEnabled sets DownsizeEnabled field to given value.

### HasDownsizeEnabled

`func (o *V1NodeRestrictions) HasDownsizeEnabled() bool`

HasDownsizeEnabled returns a boolean if a field has been set.

### GetResizeDisabledSizes

`func (o *V1NodeRestrictions) GetResizeDisabledSizes() []string`

GetResizeDisabledSizes returns the ResizeDisabledSizes field if non-nil, zero value otherwise.

### GetResizeDisabledSizesOk

`func (o *V1NodeRestrictions) GetResizeDisabledSizesOk() (*[]string, bool)`

GetResizeDisabledSizesOk returns a tuple with the ResizeDisabledSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeDisabledSizes

`func (o *V1NodeRestrictions) SetResizeDisabledSizes(v []string)`

SetResizeDisabledSizes sets ResizeDisabledSizes field to given value.

### HasResizeDisabledSizes

`func (o *V1NodeRestrictions) HasResizeDisabledSizes() bool`

HasResizeDisabledSizes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


