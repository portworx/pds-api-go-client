# V1Info1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShortName** | Pointer to **string** | Short name of the data service. | [optional] 
**Enabled** | Pointer to **bool** | Enabled flag suggests if the data service is enabled or not. | [optional] 
**NodesLimitations** | Pointer to **string** | Node limitations. | [optional] 
**NodeRestrictions** | Pointer to [**V1NodeRestrictions**](V1NodeRestrictions.md) |  | [optional] 

## Methods

### NewV1Info1

`func NewV1Info1() *V1Info1`

NewV1Info1 instantiates a new V1Info1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1Info1WithDefaults

`func NewV1Info1WithDefaults() *V1Info1`

NewV1Info1WithDefaults instantiates a new V1Info1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortName

`func (o *V1Info1) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *V1Info1) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *V1Info1) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *V1Info1) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetEnabled

`func (o *V1Info1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *V1Info1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *V1Info1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *V1Info1) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNodesLimitations

`func (o *V1Info1) GetNodesLimitations() string`

GetNodesLimitations returns the NodesLimitations field if non-nil, zero value otherwise.

### GetNodesLimitationsOk

`func (o *V1Info1) GetNodesLimitationsOk() (*string, bool)`

GetNodesLimitationsOk returns a tuple with the NodesLimitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesLimitations

`func (o *V1Info1) SetNodesLimitations(v string)`

SetNodesLimitations sets NodesLimitations field to given value.

### HasNodesLimitations

`func (o *V1Info1) HasNodesLimitations() bool`

HasNodesLimitations returns a boolean if a field has been set.

### GetNodeRestrictions

`func (o *V1Info1) GetNodeRestrictions() V1NodeRestrictions`

GetNodeRestrictions returns the NodeRestrictions field if non-nil, zero value otherwise.

### GetNodeRestrictionsOk

`func (o *V1Info1) GetNodeRestrictionsOk() (*V1NodeRestrictions, bool)`

GetNodeRestrictionsOk returns a tuple with the NodeRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeRestrictions

`func (o *V1Info1) SetNodeRestrictions(v V1NodeRestrictions)`

SetNodeRestrictions sets NodeRestrictions field to given value.

### HasNodeRestrictions

`func (o *V1Info1) HasNodeRestrictions() bool`

HasNodeRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


