# ControllersUpdateApplicationConfigurationTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigItems** | Pointer to [**[]ModelsConfigItem**](ModelsConfigItem.md) |  | [optional] 
**Name** | Pointer to **string** | See models.ApplicationConfigurationTemplate for more information. | [optional] 

## Methods

### NewControllersUpdateApplicationConfigurationTemplateRequest

`func NewControllersUpdateApplicationConfigurationTemplateRequest() *ControllersUpdateApplicationConfigurationTemplateRequest`

NewControllersUpdateApplicationConfigurationTemplateRequest instantiates a new ControllersUpdateApplicationConfigurationTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersUpdateApplicationConfigurationTemplateRequestWithDefaults

`func NewControllersUpdateApplicationConfigurationTemplateRequestWithDefaults() *ControllersUpdateApplicationConfigurationTemplateRequest`

NewControllersUpdateApplicationConfigurationTemplateRequestWithDefaults instantiates a new ControllersUpdateApplicationConfigurationTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigItems

`func (o *ControllersUpdateApplicationConfigurationTemplateRequest) GetConfigItems() []ModelsConfigItem`

GetConfigItems returns the ConfigItems field if non-nil, zero value otherwise.

### GetConfigItemsOk

`func (o *ControllersUpdateApplicationConfigurationTemplateRequest) GetConfigItemsOk() (*[]ModelsConfigItem, bool)`

GetConfigItemsOk returns a tuple with the ConfigItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigItems

`func (o *ControllersUpdateApplicationConfigurationTemplateRequest) SetConfigItems(v []ModelsConfigItem)`

SetConfigItems sets ConfigItems field to given value.

### HasConfigItems

`func (o *ControllersUpdateApplicationConfigurationTemplateRequest) HasConfigItems() bool`

HasConfigItems returns a boolean if a field has been set.

### GetName

`func (o *ControllersUpdateApplicationConfigurationTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControllersUpdateApplicationConfigurationTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControllersUpdateApplicationConfigurationTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControllersUpdateApplicationConfigurationTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


