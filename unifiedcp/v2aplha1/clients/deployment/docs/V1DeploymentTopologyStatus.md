# V1DeploymentTopologyStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | Pointer to [**V1DeploymentTopologyStatusHealth**](V1DeploymentTopologyStatusHealth.md) |  | [optional] [default to UNKNOWN]
**Phase** | Pointer to [**V1DeploymentTopologyStatusPhase**](V1DeploymentTopologyStatusPhase.md) |  | [optional] [default to PHASE_UNSPECIFIED]
**ReadyReplicas** | Pointer to **string** | Number of replicas reported by Target Cluster that are up and running. | [optional] 
**ConnectionInfo** | Pointer to [**V1ConnectionInfo**](V1ConnectionInfo.md) |  | [optional] 

## Methods

### NewV1DeploymentTopologyStatus

`func NewV1DeploymentTopologyStatus() *V1DeploymentTopologyStatus`

NewV1DeploymentTopologyStatus instantiates a new V1DeploymentTopologyStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeploymentTopologyStatusWithDefaults

`func NewV1DeploymentTopologyStatusWithDefaults() *V1DeploymentTopologyStatus`

NewV1DeploymentTopologyStatusWithDefaults instantiates a new V1DeploymentTopologyStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *V1DeploymentTopologyStatus) GetHealth() V1DeploymentTopologyStatusHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *V1DeploymentTopologyStatus) GetHealthOk() (*V1DeploymentTopologyStatusHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *V1DeploymentTopologyStatus) SetHealth(v V1DeploymentTopologyStatusHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *V1DeploymentTopologyStatus) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetPhase

`func (o *V1DeploymentTopologyStatus) GetPhase() V1DeploymentTopologyStatusPhase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *V1DeploymentTopologyStatus) GetPhaseOk() (*V1DeploymentTopologyStatusPhase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *V1DeploymentTopologyStatus) SetPhase(v V1DeploymentTopologyStatusPhase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *V1DeploymentTopologyStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetReadyReplicas

`func (o *V1DeploymentTopologyStatus) GetReadyReplicas() string`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *V1DeploymentTopologyStatus) GetReadyReplicasOk() (*string, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *V1DeploymentTopologyStatus) SetReadyReplicas(v string)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *V1DeploymentTopologyStatus) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.

### GetConnectionInfo

`func (o *V1DeploymentTopologyStatus) GetConnectionInfo() V1ConnectionInfo`

GetConnectionInfo returns the ConnectionInfo field if non-nil, zero value otherwise.

### GetConnectionInfoOk

`func (o *V1DeploymentTopologyStatus) GetConnectionInfoOk() (*V1ConnectionInfo, bool)`

GetConnectionInfoOk returns a tuple with the ConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionInfo

`func (o *V1DeploymentTopologyStatus) SetConnectionInfo(v V1ConnectionInfo)`

SetConnectionInfo sets ConnectionInfo field to given value.

### HasConnectionInfo

`func (o *V1DeploymentTopologyStatus) HasConnectionInfo() bool`

HasConnectionInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


