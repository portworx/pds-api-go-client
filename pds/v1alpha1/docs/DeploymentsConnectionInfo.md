# DeploymentsConnectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterDetails** | Pointer to **map[string]interface{}** |  | [optional] 
**ConnectionDetails** | Pointer to [**DeploymentsConnectionDetails**](DeploymentsConnectionDetails.md) |  | [optional] 
**NotReadyPods** | Pointer to [**[]DeploymentsPodInfo**](DeploymentsPodInfo.md) |  | [optional] 
**Pods** | Pointer to [**[]DeploymentsPodInfo**](DeploymentsPodInfo.md) |  | [optional] 

## Methods

### NewDeploymentsConnectionInfo

`func NewDeploymentsConnectionInfo() *DeploymentsConnectionInfo`

NewDeploymentsConnectionInfo instantiates a new DeploymentsConnectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentsConnectionInfoWithDefaults

`func NewDeploymentsConnectionInfoWithDefaults() *DeploymentsConnectionInfo`

NewDeploymentsConnectionInfoWithDefaults instantiates a new DeploymentsConnectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterDetails

`func (o *DeploymentsConnectionInfo) GetClusterDetails() map[string]interface{}`

GetClusterDetails returns the ClusterDetails field if non-nil, zero value otherwise.

### GetClusterDetailsOk

`func (o *DeploymentsConnectionInfo) GetClusterDetailsOk() (*map[string]interface{}, bool)`

GetClusterDetailsOk returns a tuple with the ClusterDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterDetails

`func (o *DeploymentsConnectionInfo) SetClusterDetails(v map[string]interface{})`

SetClusterDetails sets ClusterDetails field to given value.

### HasClusterDetails

`func (o *DeploymentsConnectionInfo) HasClusterDetails() bool`

HasClusterDetails returns a boolean if a field has been set.

### GetConnectionDetails

`func (o *DeploymentsConnectionInfo) GetConnectionDetails() DeploymentsConnectionDetails`

GetConnectionDetails returns the ConnectionDetails field if non-nil, zero value otherwise.

### GetConnectionDetailsOk

`func (o *DeploymentsConnectionInfo) GetConnectionDetailsOk() (*DeploymentsConnectionDetails, bool)`

GetConnectionDetailsOk returns a tuple with the ConnectionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionDetails

`func (o *DeploymentsConnectionInfo) SetConnectionDetails(v DeploymentsConnectionDetails)`

SetConnectionDetails sets ConnectionDetails field to given value.

### HasConnectionDetails

`func (o *DeploymentsConnectionInfo) HasConnectionDetails() bool`

HasConnectionDetails returns a boolean if a field has been set.

### GetNotReadyPods

`func (o *DeploymentsConnectionInfo) GetNotReadyPods() []DeploymentsPodInfo`

GetNotReadyPods returns the NotReadyPods field if non-nil, zero value otherwise.

### GetNotReadyPodsOk

`func (o *DeploymentsConnectionInfo) GetNotReadyPodsOk() (*[]DeploymentsPodInfo, bool)`

GetNotReadyPodsOk returns a tuple with the NotReadyPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotReadyPods

`func (o *DeploymentsConnectionInfo) SetNotReadyPods(v []DeploymentsPodInfo)`

SetNotReadyPods sets NotReadyPods field to given value.

### HasNotReadyPods

`func (o *DeploymentsConnectionInfo) HasNotReadyPods() bool`

HasNotReadyPods returns a boolean if a field has been set.

### GetPods

`func (o *DeploymentsConnectionInfo) GetPods() []DeploymentsPodInfo`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *DeploymentsConnectionInfo) GetPodsOk() (*[]DeploymentsPodInfo, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *DeploymentsConnectionInfo) SetPods(v []DeploymentsPodInfo)`

SetPods sets Pods field to given value.

### HasPods

`func (o *DeploymentsConnectionInfo) HasPods() bool`

HasPods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


