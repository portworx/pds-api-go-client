# V1Filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppResource** | Pointer to [**V1ApplicationResourceType**](V1ApplicationResourceType.md) |  | [optional] [default to TYPE_UNSPECIFIED]
**InfraResource** | Pointer to [**V1InfraResourceType**](V1InfraResourceType.md) |  | [optional] [default to TYPE_UNSPECIFIED]
**ResourceId** | Pointer to **string** | ID of the resource for which projects to be listed. | [optional] 

## Methods

### NewV1Filter

`func NewV1Filter() *V1Filter`

NewV1Filter instantiates a new V1Filter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1FilterWithDefaults

`func NewV1FilterWithDefaults() *V1Filter`

NewV1FilterWithDefaults instantiates a new V1Filter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppResource

`func (o *V1Filter) GetAppResource() V1ApplicationResourceType`

GetAppResource returns the AppResource field if non-nil, zero value otherwise.

### GetAppResourceOk

`func (o *V1Filter) GetAppResourceOk() (*V1ApplicationResourceType, bool)`

GetAppResourceOk returns a tuple with the AppResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppResource

`func (o *V1Filter) SetAppResource(v V1ApplicationResourceType)`

SetAppResource sets AppResource field to given value.

### HasAppResource

`func (o *V1Filter) HasAppResource() bool`

HasAppResource returns a boolean if a field has been set.

### GetInfraResource

`func (o *V1Filter) GetInfraResource() V1InfraResourceType`

GetInfraResource returns the InfraResource field if non-nil, zero value otherwise.

### GetInfraResourceOk

`func (o *V1Filter) GetInfraResourceOk() (*V1InfraResourceType, bool)`

GetInfraResourceOk returns a tuple with the InfraResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraResource

`func (o *V1Filter) SetInfraResource(v V1InfraResourceType)`

SetInfraResource sets InfraResource field to given value.

### HasInfraResource

`func (o *V1Filter) HasInfraResource() bool`

HasInfraResource returns a boolean if a field has been set.

### GetResourceId

`func (o *V1Filter) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *V1Filter) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *V1Filter) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *V1Filter) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


