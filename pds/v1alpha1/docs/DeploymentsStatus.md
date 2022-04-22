# DeploymentsStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **string** |  | [optional] 
**ReadyReplicas** | Pointer to **int32** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**Resources** | Pointer to [**[]DeploymentsResourceConditions**](DeploymentsResourceConditions.md) |  | [optional] 

## Methods

### NewDeploymentsStatus

`func NewDeploymentsStatus() *DeploymentsStatus`

NewDeploymentsStatus instantiates a new DeploymentsStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentsStatusWithDefaults

`func NewDeploymentsStatusWithDefaults() *DeploymentsStatus`

NewDeploymentsStatusWithDefaults instantiates a new DeploymentsStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *DeploymentsStatus) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *DeploymentsStatus) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *DeploymentsStatus) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *DeploymentsStatus) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetInitialized

`func (o *DeploymentsStatus) GetInitialized() string`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *DeploymentsStatus) GetInitializedOk() (*string, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *DeploymentsStatus) SetInitialized(v string)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *DeploymentsStatus) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetReadyReplicas

`func (o *DeploymentsStatus) GetReadyReplicas() int32`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *DeploymentsStatus) GetReadyReplicasOk() (*int32, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *DeploymentsStatus) SetReadyReplicas(v int32)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *DeploymentsStatus) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.

### GetReplicas

`func (o *DeploymentsStatus) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *DeploymentsStatus) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *DeploymentsStatus) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *DeploymentsStatus) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetResources

`func (o *DeploymentsStatus) GetResources() []DeploymentsResourceConditions`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DeploymentsStatus) GetResourcesOk() (*[]DeploymentsResourceConditions, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DeploymentsStatus) SetResources(v []DeploymentsResourceConditions)`

SetResources sets Resources field to given value.

### HasResources

`func (o *DeploymentsStatus) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


