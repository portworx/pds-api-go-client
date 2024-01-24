# DesiredConfigurationOfTheIAM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorType** | Pointer to **string** |  | [optional] [readonly] 
**AccessPolicy** | Pointer to [**V1AccessPolicy**](V1AccessPolicy.md) |  | [optional] 

## Methods

### NewDesiredConfigurationOfTheIAM

`func NewDesiredConfigurationOfTheIAM() *DesiredConfigurationOfTheIAM`

NewDesiredConfigurationOfTheIAM instantiates a new DesiredConfigurationOfTheIAM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesiredConfigurationOfTheIAMWithDefaults

`func NewDesiredConfigurationOfTheIAMWithDefaults() *DesiredConfigurationOfTheIAM`

NewDesiredConfigurationOfTheIAMWithDefaults instantiates a new DesiredConfigurationOfTheIAM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorType

`func (o *DesiredConfigurationOfTheIAM) GetActorType() string`

GetActorType returns the ActorType field if non-nil, zero value otherwise.

### GetActorTypeOk

`func (o *DesiredConfigurationOfTheIAM) GetActorTypeOk() (*string, bool)`

GetActorTypeOk returns a tuple with the ActorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorType

`func (o *DesiredConfigurationOfTheIAM) SetActorType(v string)`

SetActorType sets ActorType field to given value.

### HasActorType

`func (o *DesiredConfigurationOfTheIAM) HasActorType() bool`

HasActorType returns a boolean if a field has been set.

### GetAccessPolicy

`func (o *DesiredConfigurationOfTheIAM) GetAccessPolicy() V1AccessPolicy`

GetAccessPolicy returns the AccessPolicy field if non-nil, zero value otherwise.

### GetAccessPolicyOk

`func (o *DesiredConfigurationOfTheIAM) GetAccessPolicyOk() (*V1AccessPolicy, bool)`

GetAccessPolicyOk returns a tuple with the AccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicy

`func (o *DesiredConfigurationOfTheIAM) SetAccessPolicy(v V1AccessPolicy)`

SetAccessPolicy sets AccessPolicy field to given value.

### HasAccessPolicy

`func (o *DesiredConfigurationOfTheIAM) HasAccessPolicy() bool`

HasAccessPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


