# V1ListDeploymentConfigUpdatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentConfigUpdates** | Pointer to [**[]V1DeploymentConfigUpdate**](V1DeploymentConfigUpdate.md) |  | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListDeploymentConfigUpdatesResponse

`func NewV1ListDeploymentConfigUpdatesResponse() *V1ListDeploymentConfigUpdatesResponse`

NewV1ListDeploymentConfigUpdatesResponse instantiates a new V1ListDeploymentConfigUpdatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListDeploymentConfigUpdatesResponseWithDefaults

`func NewV1ListDeploymentConfigUpdatesResponseWithDefaults() *V1ListDeploymentConfigUpdatesResponse`

NewV1ListDeploymentConfigUpdatesResponseWithDefaults instantiates a new V1ListDeploymentConfigUpdatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentConfigUpdates

`func (o *V1ListDeploymentConfigUpdatesResponse) GetDeploymentConfigUpdates() []V1DeploymentConfigUpdate`

GetDeploymentConfigUpdates returns the DeploymentConfigUpdates field if non-nil, zero value otherwise.

### GetDeploymentConfigUpdatesOk

`func (o *V1ListDeploymentConfigUpdatesResponse) GetDeploymentConfigUpdatesOk() (*[]V1DeploymentConfigUpdate, bool)`

GetDeploymentConfigUpdatesOk returns a tuple with the DeploymentConfigUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentConfigUpdates

`func (o *V1ListDeploymentConfigUpdatesResponse) SetDeploymentConfigUpdates(v []V1DeploymentConfigUpdate)`

SetDeploymentConfigUpdates sets DeploymentConfigUpdates field to given value.

### HasDeploymentConfigUpdates

`func (o *V1ListDeploymentConfigUpdatesResponse) HasDeploymentConfigUpdates() bool`

HasDeploymentConfigUpdates returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListDeploymentConfigUpdatesResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListDeploymentConfigUpdatesResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListDeploymentConfigUpdatesResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListDeploymentConfigUpdatesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


