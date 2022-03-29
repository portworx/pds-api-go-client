# ModelsConfigItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeployTime** | Pointer to **bool** | If true, user must fill the value for this configuration parameter when creating a new deployment. | [optional] 
**Key** | Pointer to **string** | Configuration key, transfers to the image&#39;s environment variable. | [optional] 
**Value** | Pointer to **string** | Configuration value. | [optional] 

## Methods

### NewModelsConfigItem

`func NewModelsConfigItem() *ModelsConfigItem`

NewModelsConfigItem instantiates a new ModelsConfigItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsConfigItemWithDefaults

`func NewModelsConfigItemWithDefaults() *ModelsConfigItem`

NewModelsConfigItemWithDefaults instantiates a new ModelsConfigItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployTime

`func (o *ModelsConfigItem) GetDeployTime() bool`

GetDeployTime returns the DeployTime field if non-nil, zero value otherwise.

### GetDeployTimeOk

`func (o *ModelsConfigItem) GetDeployTimeOk() (*bool, bool)`

GetDeployTimeOk returns a tuple with the DeployTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployTime

`func (o *ModelsConfigItem) SetDeployTime(v bool)`

SetDeployTime sets DeployTime field to given value.

### HasDeployTime

`func (o *ModelsConfigItem) HasDeployTime() bool`

HasDeployTime returns a boolean if a field has been set.

### GetKey

`func (o *ModelsConfigItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ModelsConfigItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ModelsConfigItem) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ModelsConfigItem) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *ModelsConfigItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ModelsConfigItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ModelsConfigItem) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ModelsConfigItem) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


