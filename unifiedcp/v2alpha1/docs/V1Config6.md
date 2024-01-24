# V1Config6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorId** | Pointer to **string** |  | [optional] 
**ActorType** | Pointer to **string** |  | [optional] [readonly] 
**AccessPolicy** | Pointer to [**V1AccessPolicy**](V1AccessPolicy.md) |  | [optional] 

## Methods

### NewV1Config6

`func NewV1Config6() *V1Config6`

NewV1Config6 instantiates a new V1Config6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1Config6WithDefaults

`func NewV1Config6WithDefaults() *V1Config6`

NewV1Config6WithDefaults instantiates a new V1Config6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorId

`func (o *V1Config6) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *V1Config6) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *V1Config6) SetActorId(v string)`

SetActorId sets ActorId field to given value.

### HasActorId

`func (o *V1Config6) HasActorId() bool`

HasActorId returns a boolean if a field has been set.

### GetActorType

`func (o *V1Config6) GetActorType() string`

GetActorType returns the ActorType field if non-nil, zero value otherwise.

### GetActorTypeOk

`func (o *V1Config6) GetActorTypeOk() (*string, bool)`

GetActorTypeOk returns a tuple with the ActorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorType

`func (o *V1Config6) SetActorType(v string)`

SetActorType sets ActorType field to given value.

### HasActorType

`func (o *V1Config6) HasActorType() bool`

HasActorType returns a boolean if a field has been set.

### GetAccessPolicy

`func (o *V1Config6) GetAccessPolicy() V1AccessPolicy`

GetAccessPolicy returns the AccessPolicy field if non-nil, zero value otherwise.

### GetAccessPolicyOk

`func (o *V1Config6) GetAccessPolicyOk() (*V1AccessPolicy, bool)`

GetAccessPolicyOk returns a tuple with the AccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicy

`func (o *V1Config6) SetAccessPolicy(v V1AccessPolicy)`

SetAccessPolicy sets AccessPolicy field to given value.

### HasAccessPolicy

`func (o *V1Config6) HasAccessPolicy() bool`

HasAccessPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


