# V1Config1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**References** | Pointer to [**V1References2**](V1References2.md) |  | [optional] 
**TlsEnabled** | Pointer to **bool** | Flag to enable TLS for the Data Service. | [optional] 
**DeploymentTopologies** | Pointer to [**[]V1DeploymentTopology**](V1DeploymentTopology.md) | A deployment topology contains a number of nodes that have various attributes as a collective group. | [optional] 

## Methods

### NewV1Config1

`func NewV1Config1() *V1Config1`

NewV1Config1 instantiates a new V1Config1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1Config1WithDefaults

`func NewV1Config1WithDefaults() *V1Config1`

NewV1Config1WithDefaults instantiates a new V1Config1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferences

`func (o *V1Config1) GetReferences() V1References2`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *V1Config1) GetReferencesOk() (*V1References2, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *V1Config1) SetReferences(v V1References2)`

SetReferences sets References field to given value.

### HasReferences

`func (o *V1Config1) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetTlsEnabled

`func (o *V1Config1) GetTlsEnabled() bool`

GetTlsEnabled returns the TlsEnabled field if non-nil, zero value otherwise.

### GetTlsEnabledOk

`func (o *V1Config1) GetTlsEnabledOk() (*bool, bool)`

GetTlsEnabledOk returns a tuple with the TlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnabled

`func (o *V1Config1) SetTlsEnabled(v bool)`

SetTlsEnabled sets TlsEnabled field to given value.

### HasTlsEnabled

`func (o *V1Config1) HasTlsEnabled() bool`

HasTlsEnabled returns a boolean if a field has been set.

### GetDeploymentTopologies

`func (o *V1Config1) GetDeploymentTopologies() []V1DeploymentTopology`

GetDeploymentTopologies returns the DeploymentTopologies field if non-nil, zero value otherwise.

### GetDeploymentTopologiesOk

`func (o *V1Config1) GetDeploymentTopologiesOk() (*[]V1DeploymentTopology, bool)`

GetDeploymentTopologiesOk returns a tuple with the DeploymentTopologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTopologies

`func (o *V1Config1) SetDeploymentTopologies(v []V1DeploymentTopology)`

SetDeploymentTopologies sets DeploymentTopologies field to given value.

### HasDeploymentTopologies

`func (o *V1Config1) HasDeploymentTopologies() bool`

HasDeploymentTopologies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

