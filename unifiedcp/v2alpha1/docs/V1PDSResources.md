# V1PDSResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployments** | Pointer to **[]string** | Resource ids of deployments. | [optional] 
**Backups** | Pointer to **[]string** | Resource ids of backups. | [optional] 
**Restores** | Pointer to **[]string** | Resource ids of restores. | [optional] 

## Methods

### NewV1PDSResources

`func NewV1PDSResources() *V1PDSResources`

NewV1PDSResources instantiates a new V1PDSResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PDSResourcesWithDefaults

`func NewV1PDSResourcesWithDefaults() *V1PDSResources`

NewV1PDSResourcesWithDefaults instantiates a new V1PDSResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployments

`func (o *V1PDSResources) GetDeployments() []string`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *V1PDSResources) GetDeploymentsOk() (*[]string, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *V1PDSResources) SetDeployments(v []string)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *V1PDSResources) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetBackups

`func (o *V1PDSResources) GetBackups() []string`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *V1PDSResources) GetBackupsOk() (*[]string, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *V1PDSResources) SetBackups(v []string)`

SetBackups sets Backups field to given value.

### HasBackups

`func (o *V1PDSResources) HasBackups() bool`

HasBackups returns a boolean if a field has been set.

### GetRestores

`func (o *V1PDSResources) GetRestores() []string`

GetRestores returns the Restores field if non-nil, zero value otherwise.

### GetRestoresOk

`func (o *V1PDSResources) GetRestoresOk() (*[]string, bool)`

GetRestoresOk returns a tuple with the Restores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestores

`func (o *V1PDSResources) SetRestores(v []string)`

SetRestores sets Restores field to given value.

### HasRestores

`func (o *V1PDSResources) HasRestores() bool`

HasRestores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


