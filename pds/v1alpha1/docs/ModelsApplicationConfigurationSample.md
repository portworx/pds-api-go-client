# ModelsApplicationConfigurationSample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigItems** | Pointer to [**[]ModelsConfigItem**](ModelsConfigItem.md) |  | [optional] 
**DataServiceId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name of the template. Must be unique for a specific data service within the tenant scope. | [optional] 

## Methods

### NewModelsApplicationConfigurationSample

`func NewModelsApplicationConfigurationSample() *ModelsApplicationConfigurationSample`

NewModelsApplicationConfigurationSample instantiates a new ModelsApplicationConfigurationSample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsApplicationConfigurationSampleWithDefaults

`func NewModelsApplicationConfigurationSampleWithDefaults() *ModelsApplicationConfigurationSample`

NewModelsApplicationConfigurationSampleWithDefaults instantiates a new ModelsApplicationConfigurationSample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigItems

`func (o *ModelsApplicationConfigurationSample) GetConfigItems() []ModelsConfigItem`

GetConfigItems returns the ConfigItems field if non-nil, zero value otherwise.

### GetConfigItemsOk

`func (o *ModelsApplicationConfigurationSample) GetConfigItemsOk() (*[]ModelsConfigItem, bool)`

GetConfigItemsOk returns a tuple with the ConfigItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigItems

`func (o *ModelsApplicationConfigurationSample) SetConfigItems(v []ModelsConfigItem)`

SetConfigItems sets ConfigItems field to given value.

### HasConfigItems

`func (o *ModelsApplicationConfigurationSample) HasConfigItems() bool`

HasConfigItems returns a boolean if a field has been set.

### GetDataServiceId

`func (o *ModelsApplicationConfigurationSample) GetDataServiceId() string`

GetDataServiceId returns the DataServiceId field if non-nil, zero value otherwise.

### GetDataServiceIdOk

`func (o *ModelsApplicationConfigurationSample) GetDataServiceIdOk() (*string, bool)`

GetDataServiceIdOk returns a tuple with the DataServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataServiceId

`func (o *ModelsApplicationConfigurationSample) SetDataServiceId(v string)`

SetDataServiceId sets DataServiceId field to given value.

### HasDataServiceId

`func (o *ModelsApplicationConfigurationSample) HasDataServiceId() bool`

HasDataServiceId returns a boolean if a field has been set.

### GetName

`func (o *ModelsApplicationConfigurationSample) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsApplicationConfigurationSample) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsApplicationConfigurationSample) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsApplicationConfigurationSample) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


