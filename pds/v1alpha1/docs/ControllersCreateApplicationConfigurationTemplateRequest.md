# ControllersCreateApplicationConfigurationTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigItems** | Pointer to [**[]ModelsConfigItem**](ModelsConfigItem.md) |  | [optional] 
**DataServiceId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | See models.ApplicationConfigurationTemplate for more information. | [optional] 

## Methods

### NewControllersCreateApplicationConfigurationTemplateRequest

`func NewControllersCreateApplicationConfigurationTemplateRequest() *ControllersCreateApplicationConfigurationTemplateRequest`

NewControllersCreateApplicationConfigurationTemplateRequest instantiates a new ControllersCreateApplicationConfigurationTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersCreateApplicationConfigurationTemplateRequestWithDefaults

`func NewControllersCreateApplicationConfigurationTemplateRequestWithDefaults() *ControllersCreateApplicationConfigurationTemplateRequest`

NewControllersCreateApplicationConfigurationTemplateRequestWithDefaults instantiates a new ControllersCreateApplicationConfigurationTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigItems

`func (o *ControllersCreateApplicationConfigurationTemplateRequest) GetConfigItems() []ModelsConfigItem`

GetConfigItems returns the ConfigItems field if non-nil, zero value otherwise.

### GetConfigItemsOk

`func (o *ControllersCreateApplicationConfigurationTemplateRequest) GetConfigItemsOk() (*[]ModelsConfigItem, bool)`

GetConfigItemsOk returns a tuple with the ConfigItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigItems

`func (o *ControllersCreateApplicationConfigurationTemplateRequest) SetConfigItems(v []ModelsConfigItem)`

SetConfigItems sets ConfigItems field to given value.

### HasConfigItems

`func (o *ControllersCreateApplicationConfigurationTemplateRequest) HasConfigItems() bool`

HasConfigItems returns a boolean if a field has been set.

### GetDataServiceId

`func (o *ControllersCreateApplicationConfigurationTemplateRequest) GetDataServiceId() string`

GetDataServiceId returns the DataServiceId field if non-nil, zero value otherwise.

### GetDataServiceIdOk

`func (o *ControllersCreateApplicationConfigurationTemplateRequest) GetDataServiceIdOk() (*string, bool)`

GetDataServiceIdOk returns a tuple with the DataServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataServiceId

`func (o *ControllersCreateApplicationConfigurationTemplateRequest) SetDataServiceId(v string)`

SetDataServiceId sets DataServiceId field to given value.

### HasDataServiceId

`func (o *ControllersCreateApplicationConfigurationTemplateRequest) HasDataServiceId() bool`

HasDataServiceId returns a boolean if a field has been set.

### GetName

`func (o *ControllersCreateApplicationConfigurationTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControllersCreateApplicationConfigurationTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControllersCreateApplicationConfigurationTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControllersCreateApplicationConfigurationTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


