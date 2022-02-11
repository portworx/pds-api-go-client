# StatusDeploymentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**[]V1EndpointAddress**](V1EndpointAddress.md) |  | [optional] 
**ConnectionStrings** | Pointer to **[]string** |  | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **string** |  | [optional] 
**NotReadyAddresses** | Pointer to [**[]V1EndpointAddress**](V1EndpointAddress.md) |  | [optional] 
**ReadyReplicas** | Pointer to **int32** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 

## Methods

### NewStatusDeploymentStatus

`func NewStatusDeploymentStatus() *StatusDeploymentStatus`

NewStatusDeploymentStatus instantiates a new StatusDeploymentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusDeploymentStatusWithDefaults

`func NewStatusDeploymentStatusWithDefaults() *StatusDeploymentStatus`

NewStatusDeploymentStatusWithDefaults instantiates a new StatusDeploymentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *StatusDeploymentStatus) GetAddresses() []V1EndpointAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *StatusDeploymentStatus) GetAddressesOk() (*[]V1EndpointAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *StatusDeploymentStatus) SetAddresses(v []V1EndpointAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *StatusDeploymentStatus) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetConnectionStrings

`func (o *StatusDeploymentStatus) GetConnectionStrings() []string`

GetConnectionStrings returns the ConnectionStrings field if non-nil, zero value otherwise.

### GetConnectionStringsOk

`func (o *StatusDeploymentStatus) GetConnectionStringsOk() (*[]string, bool)`

GetConnectionStringsOk returns a tuple with the ConnectionStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStrings

`func (o *StatusDeploymentStatus) SetConnectionStrings(v []string)`

SetConnectionStrings sets ConnectionStrings field to given value.

### HasConnectionStrings

`func (o *StatusDeploymentStatus) HasConnectionStrings() bool`

HasConnectionStrings returns a boolean if a field has been set.

### GetHealth

`func (o *StatusDeploymentStatus) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *StatusDeploymentStatus) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *StatusDeploymentStatus) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *StatusDeploymentStatus) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetInitialized

`func (o *StatusDeploymentStatus) GetInitialized() string`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *StatusDeploymentStatus) GetInitializedOk() (*string, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *StatusDeploymentStatus) SetInitialized(v string)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *StatusDeploymentStatus) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetNotReadyAddresses

`func (o *StatusDeploymentStatus) GetNotReadyAddresses() []V1EndpointAddress`

GetNotReadyAddresses returns the NotReadyAddresses field if non-nil, zero value otherwise.

### GetNotReadyAddressesOk

`func (o *StatusDeploymentStatus) GetNotReadyAddressesOk() (*[]V1EndpointAddress, bool)`

GetNotReadyAddressesOk returns a tuple with the NotReadyAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotReadyAddresses

`func (o *StatusDeploymentStatus) SetNotReadyAddresses(v []V1EndpointAddress)`

SetNotReadyAddresses sets NotReadyAddresses field to given value.

### HasNotReadyAddresses

`func (o *StatusDeploymentStatus) HasNotReadyAddresses() bool`

HasNotReadyAddresses returns a boolean if a field has been set.

### GetReadyReplicas

`func (o *StatusDeploymentStatus) GetReadyReplicas() int32`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *StatusDeploymentStatus) GetReadyReplicasOk() (*int32, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *StatusDeploymentStatus) SetReadyReplicas(v int32)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *StatusDeploymentStatus) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.

### GetReplicas

`func (o *StatusDeploymentStatus) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *StatusDeploymentStatus) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *StatusDeploymentStatus) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *StatusDeploymentStatus) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


