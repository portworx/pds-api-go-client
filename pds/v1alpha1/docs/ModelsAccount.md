# ModelsAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaasFeaturesEnabled** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is autogenerated on creation | [optional] 
**Eula** | Pointer to [**ModelsEULADetails**](ModelsEULADetails.md) |  | [optional] 
**GlobalConfig** | Pointer to [**ModelsAccountGlobalConfig**](ModelsAccountGlobalConfig.md) |  | [optional] 
**Id** | Pointer to **string** | ID is auto generated on creation | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PdsFeaturesEnabled** | Pointer to **bool** |  | [optional] 
**PxoneFeaturesEnabled** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is autogenerated on update | [optional] 

## Methods

### NewModelsAccount

`func NewModelsAccount() *ModelsAccount`

NewModelsAccount instantiates a new ModelsAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsAccountWithDefaults

`func NewModelsAccountWithDefaults() *ModelsAccount`

NewModelsAccountWithDefaults instantiates a new ModelsAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaasFeaturesEnabled

`func (o *ModelsAccount) GetBaasFeaturesEnabled() bool`

GetBaasFeaturesEnabled returns the BaasFeaturesEnabled field if non-nil, zero value otherwise.

### GetBaasFeaturesEnabledOk

`func (o *ModelsAccount) GetBaasFeaturesEnabledOk() (*bool, bool)`

GetBaasFeaturesEnabledOk returns a tuple with the BaasFeaturesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaasFeaturesEnabled

`func (o *ModelsAccount) SetBaasFeaturesEnabled(v bool)`

SetBaasFeaturesEnabled sets BaasFeaturesEnabled field to given value.

### HasBaasFeaturesEnabled

`func (o *ModelsAccount) HasBaasFeaturesEnabled() bool`

HasBaasFeaturesEnabled returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsAccount) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsAccount) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsAccount) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEula

`func (o *ModelsAccount) GetEula() ModelsEULADetails`

GetEula returns the Eula field if non-nil, zero value otherwise.

### GetEulaOk

`func (o *ModelsAccount) GetEulaOk() (*ModelsEULADetails, bool)`

GetEulaOk returns a tuple with the Eula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEula

`func (o *ModelsAccount) SetEula(v ModelsEULADetails)`

SetEula sets Eula field to given value.

### HasEula

`func (o *ModelsAccount) HasEula() bool`

HasEula returns a boolean if a field has been set.

### GetGlobalConfig

`func (o *ModelsAccount) GetGlobalConfig() ModelsAccountGlobalConfig`

GetGlobalConfig returns the GlobalConfig field if non-nil, zero value otherwise.

### GetGlobalConfigOk

`func (o *ModelsAccount) GetGlobalConfigOk() (*ModelsAccountGlobalConfig, bool)`

GetGlobalConfigOk returns a tuple with the GlobalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalConfig

`func (o *ModelsAccount) SetGlobalConfig(v ModelsAccountGlobalConfig)`

SetGlobalConfig sets GlobalConfig field to given value.

### HasGlobalConfig

`func (o *ModelsAccount) HasGlobalConfig() bool`

HasGlobalConfig returns a boolean if a field has been set.

### GetId

`func (o *ModelsAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ModelsAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPdsFeaturesEnabled

`func (o *ModelsAccount) GetPdsFeaturesEnabled() bool`

GetPdsFeaturesEnabled returns the PdsFeaturesEnabled field if non-nil, zero value otherwise.

### GetPdsFeaturesEnabledOk

`func (o *ModelsAccount) GetPdsFeaturesEnabledOk() (*bool, bool)`

GetPdsFeaturesEnabledOk returns a tuple with the PdsFeaturesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdsFeaturesEnabled

`func (o *ModelsAccount) SetPdsFeaturesEnabled(v bool)`

SetPdsFeaturesEnabled sets PdsFeaturesEnabled field to given value.

### HasPdsFeaturesEnabled

`func (o *ModelsAccount) HasPdsFeaturesEnabled() bool`

HasPdsFeaturesEnabled returns a boolean if a field has been set.

### GetPxoneFeaturesEnabled

`func (o *ModelsAccount) GetPxoneFeaturesEnabled() bool`

GetPxoneFeaturesEnabled returns the PxoneFeaturesEnabled field if non-nil, zero value otherwise.

### GetPxoneFeaturesEnabledOk

`func (o *ModelsAccount) GetPxoneFeaturesEnabledOk() (*bool, bool)`

GetPxoneFeaturesEnabledOk returns a tuple with the PxoneFeaturesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxoneFeaturesEnabled

`func (o *ModelsAccount) SetPxoneFeaturesEnabled(v bool)`

SetPxoneFeaturesEnabled sets PxoneFeaturesEnabled field to given value.

### HasPxoneFeaturesEnabled

`func (o *ModelsAccount) HasPxoneFeaturesEnabled() bool`

HasPxoneFeaturesEnabled returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsAccount) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsAccount) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsAccount) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsAccount) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


