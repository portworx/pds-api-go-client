# DeploymentsPodInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The IP of this pod. | [optional] 
**Name** | Pointer to **string** | Name is the Hostname of this pod. | [optional] 
**WorkerNode** | Pointer to **string** | Node hosting this pod. | [optional] 

## Methods

### NewDeploymentsPodInfo

`func NewDeploymentsPodInfo() *DeploymentsPodInfo`

NewDeploymentsPodInfo instantiates a new DeploymentsPodInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentsPodInfoWithDefaults

`func NewDeploymentsPodInfoWithDefaults() *DeploymentsPodInfo`

NewDeploymentsPodInfoWithDefaults instantiates a new DeploymentsPodInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *DeploymentsPodInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DeploymentsPodInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DeploymentsPodInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *DeploymentsPodInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetName

`func (o *DeploymentsPodInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentsPodInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentsPodInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentsPodInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWorkerNode

`func (o *DeploymentsPodInfo) GetWorkerNode() string`

GetWorkerNode returns the WorkerNode field if non-nil, zero value otherwise.

### GetWorkerNodeOk

`func (o *DeploymentsPodInfo) GetWorkerNodeOk() (*string, bool)`

GetWorkerNodeOk returns a tuple with the WorkerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerNode

`func (o *DeploymentsPodInfo) SetWorkerNode(v string)`

SetWorkerNode sets WorkerNode field to given value.

### HasWorkerNode

`func (o *DeploymentsPodInfo) HasWorkerNode() bool`

HasWorkerNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


