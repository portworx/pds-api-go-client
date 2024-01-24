# V1Provider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to [**V1ProviderType**](V1ProviderType.md) |  | [optional] [default to TYPE_UNSPECIFIED]

## Methods

### NewV1Provider

`func NewV1Provider() *V1Provider`

NewV1Provider instantiates a new V1Provider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ProviderWithDefaults

`func NewV1ProviderWithDefaults() *V1Provider`

NewV1ProviderWithDefaults instantiates a new V1Provider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *V1Provider) GetCloudProvider() V1ProviderType`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *V1Provider) GetCloudProviderOk() (*V1ProviderType, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *V1Provider) SetCloudProvider(v V1ProviderType)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *V1Provider) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


