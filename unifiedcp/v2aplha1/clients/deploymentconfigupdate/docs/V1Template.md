# V1Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | UID of the Template. | [optional] 
**ResourceVersion** | Pointer to **string** | Resource version of the template. | [optional] [readonly] 
**Values** | Pointer to [**map[string]ProtobufAny**](ProtobufAny.md) | Values required for template. | [optional] 

## Methods

### NewV1Template

`func NewV1Template() *V1Template`

NewV1Template instantiates a new V1Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TemplateWithDefaults

`func NewV1TemplateWithDefaults() *V1Template`

NewV1TemplateWithDefaults instantiates a new V1Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V1Template) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1Template) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1Template) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1Template) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceVersion

`func (o *V1Template) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *V1Template) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *V1Template) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *V1Template) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetValues

`func (o *V1Template) GetValues() map[string]ProtobufAny`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *V1Template) GetValuesOk() (*map[string]ProtobufAny, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *V1Template) SetValues(v map[string]ProtobufAny)`

SetValues sets Values field to given value.

### HasValues

`func (o *V1Template) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


