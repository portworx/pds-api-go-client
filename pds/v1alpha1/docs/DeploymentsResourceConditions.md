# DeploymentsResourceConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]DeploymentsCondition**](DeploymentsCondition.md) |  | [optional] 
**Resource** | Pointer to [**V1TypedLocalObjectReference**](V1TypedLocalObjectReference.md) |  | [optional] 

## Methods

### NewDeploymentsResourceConditions

`func NewDeploymentsResourceConditions() *DeploymentsResourceConditions`

NewDeploymentsResourceConditions instantiates a new DeploymentsResourceConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentsResourceConditionsWithDefaults

`func NewDeploymentsResourceConditionsWithDefaults() *DeploymentsResourceConditions`

NewDeploymentsResourceConditionsWithDefaults instantiates a new DeploymentsResourceConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *DeploymentsResourceConditions) GetConditions() []DeploymentsCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *DeploymentsResourceConditions) GetConditionsOk() (*[]DeploymentsCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *DeploymentsResourceConditions) SetConditions(v []DeploymentsCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *DeploymentsResourceConditions) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetResource

`func (o *DeploymentsResourceConditions) GetResource() V1TypedLocalObjectReference`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DeploymentsResourceConditions) GetResourceOk() (*V1TypedLocalObjectReference, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DeploymentsResourceConditions) SetResource(v V1TypedLocalObjectReference)`

SetResource sets Resource field to given value.

### HasResource

`func (o *DeploymentsResourceConditions) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


