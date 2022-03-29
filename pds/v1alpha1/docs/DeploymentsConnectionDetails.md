# DeploymentsConnectionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to **[]string** | Nodes of the data service. | [optional] 
**Ports** | Pointer to **map[string]int32** | Ports provided by the data service (name and number). | [optional] 

## Methods

### NewDeploymentsConnectionDetails

`func NewDeploymentsConnectionDetails() *DeploymentsConnectionDetails`

NewDeploymentsConnectionDetails instantiates a new DeploymentsConnectionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentsConnectionDetailsWithDefaults

`func NewDeploymentsConnectionDetailsWithDefaults() *DeploymentsConnectionDetails`

NewDeploymentsConnectionDetailsWithDefaults instantiates a new DeploymentsConnectionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *DeploymentsConnectionDetails) GetNodes() []string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *DeploymentsConnectionDetails) GetNodesOk() (*[]string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *DeploymentsConnectionDetails) SetNodes(v []string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *DeploymentsConnectionDetails) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetPorts

`func (o *DeploymentsConnectionDetails) GetPorts() map[string]int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *DeploymentsConnectionDetails) GetPortsOk() (*map[string]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *DeploymentsConnectionDetails) SetPorts(v map[string]int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *DeploymentsConnectionDetails) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


