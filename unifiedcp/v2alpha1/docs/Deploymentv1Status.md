# Deploymentv1Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | Pointer to [**V1StatusHealth**](V1StatusHealth.md) |  | [optional] [default to HEALTH_UNSPECIFIED]
**Phase** | Pointer to [**V1StatusPhase**](V1StatusPhase.md) |  | [optional] [default to PHASE_UNSPECIFIED]
**ConnectionInfo** | Pointer to [**map[string]ProtobufAny2**](ProtobufAny2.md) | ConnectionDetails urls, ports, credentials, etc for connecting to the data service. | [optional] 
**Initialized** | Pointer to **string** | Initialize used to control startup scripts. | [optional] 
**DeploymentTopologyStatus** | Pointer to [**[]V1DeploymentTopologyStatus**](V1DeploymentTopologyStatus.md) | Status of the deployment topology. | [optional] 

## Methods

### NewDeploymentv1Status

`func NewDeploymentv1Status() *Deploymentv1Status`

NewDeploymentv1Status instantiates a new Deploymentv1Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentv1StatusWithDefaults

`func NewDeploymentv1StatusWithDefaults() *Deploymentv1Status`

NewDeploymentv1StatusWithDefaults instantiates a new Deploymentv1Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *Deploymentv1Status) GetHealth() V1StatusHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *Deploymentv1Status) GetHealthOk() (*V1StatusHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *Deploymentv1Status) SetHealth(v V1StatusHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *Deploymentv1Status) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetPhase

`func (o *Deploymentv1Status) GetPhase() V1StatusPhase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *Deploymentv1Status) GetPhaseOk() (*V1StatusPhase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *Deploymentv1Status) SetPhase(v V1StatusPhase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *Deploymentv1Status) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetConnectionInfo

`func (o *Deploymentv1Status) GetConnectionInfo() map[string]ProtobufAny2`

GetConnectionInfo returns the ConnectionInfo field if non-nil, zero value otherwise.

### GetConnectionInfoOk

`func (o *Deploymentv1Status) GetConnectionInfoOk() (*map[string]ProtobufAny2, bool)`

GetConnectionInfoOk returns a tuple with the ConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionInfo

`func (o *Deploymentv1Status) SetConnectionInfo(v map[string]ProtobufAny2)`

SetConnectionInfo sets ConnectionInfo field to given value.

### HasConnectionInfo

`func (o *Deploymentv1Status) HasConnectionInfo() bool`

HasConnectionInfo returns a boolean if a field has been set.

### GetInitialized

`func (o *Deploymentv1Status) GetInitialized() string`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *Deploymentv1Status) GetInitializedOk() (*string, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *Deploymentv1Status) SetInitialized(v string)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *Deploymentv1Status) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetDeploymentTopologyStatus

`func (o *Deploymentv1Status) GetDeploymentTopologyStatus() []V1DeploymentTopologyStatus`

GetDeploymentTopologyStatus returns the DeploymentTopologyStatus field if non-nil, zero value otherwise.

### GetDeploymentTopologyStatusOk

`func (o *Deploymentv1Status) GetDeploymentTopologyStatusOk() (*[]V1DeploymentTopologyStatus, bool)`

GetDeploymentTopologyStatusOk returns a tuple with the DeploymentTopologyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTopologyStatus

`func (o *Deploymentv1Status) SetDeploymentTopologyStatus(v []V1DeploymentTopologyStatus)`

SetDeploymentTopologyStatus sets DeploymentTopologyStatus field to given value.

### HasDeploymentTopologyStatus

`func (o *Deploymentv1Status) HasDeploymentTopologyStatus() bool`

HasDeploymentTopologyStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


