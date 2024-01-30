# V1Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceReferences** | Pointer to [**V1SourceReferences**](V1SourceReferences.md) |  | [optional] 
**DestinationReferences** | Pointer to [**V1DestinationReferences**](V1DestinationReferences.md) |  | [optional] 
**CustomResourceName** | Pointer to **string** | K8s resource name for restore, built from [\&quot;restore-\&quot; + name + short-id]. | [optional] [readonly] 

## Methods

### NewV1Config

`func NewV1Config() *V1Config`

NewV1Config instantiates a new V1Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ConfigWithDefaults

`func NewV1ConfigWithDefaults() *V1Config`

NewV1ConfigWithDefaults instantiates a new V1Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceReferences

`func (o *V1Config) GetSourceReferences() V1SourceReferences`

GetSourceReferences returns the SourceReferences field if non-nil, zero value otherwise.

### GetSourceReferencesOk

`func (o *V1Config) GetSourceReferencesOk() (*V1SourceReferences, bool)`

GetSourceReferencesOk returns a tuple with the SourceReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceReferences

`func (o *V1Config) SetSourceReferences(v V1SourceReferences)`

SetSourceReferences sets SourceReferences field to given value.

### HasSourceReferences

`func (o *V1Config) HasSourceReferences() bool`

HasSourceReferences returns a boolean if a field has been set.

### GetDestinationReferences

`func (o *V1Config) GetDestinationReferences() V1DestinationReferences`

GetDestinationReferences returns the DestinationReferences field if non-nil, zero value otherwise.

### GetDestinationReferencesOk

`func (o *V1Config) GetDestinationReferencesOk() (*V1DestinationReferences, bool)`

GetDestinationReferencesOk returns a tuple with the DestinationReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationReferences

`func (o *V1Config) SetDestinationReferences(v V1DestinationReferences)`

SetDestinationReferences sets DestinationReferences field to given value.

### HasDestinationReferences

`func (o *V1Config) HasDestinationReferences() bool`

HasDestinationReferences returns a boolean if a field has been set.

### GetCustomResourceName

`func (o *V1Config) GetCustomResourceName() string`

GetCustomResourceName returns the CustomResourceName field if non-nil, zero value otherwise.

### GetCustomResourceNameOk

`func (o *V1Config) GetCustomResourceNameOk() (*string, bool)`

GetCustomResourceNameOk returns a tuple with the CustomResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomResourceName

`func (o *V1Config) SetCustomResourceName(v string)`

SetCustomResourceName sets CustomResourceName field to given value.

### HasCustomResourceName

`func (o *V1Config) HasCustomResourceName() bool`

HasCustomResourceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


