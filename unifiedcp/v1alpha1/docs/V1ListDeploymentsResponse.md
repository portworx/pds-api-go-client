# V1ListDeploymentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployments** | Pointer to [**[]V1Deployment**](V1Deployment.md) | List of deployment resources. | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListDeploymentsResponse

`func NewV1ListDeploymentsResponse() *V1ListDeploymentsResponse`

NewV1ListDeploymentsResponse instantiates a new V1ListDeploymentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListDeploymentsResponseWithDefaults

`func NewV1ListDeploymentsResponseWithDefaults() *V1ListDeploymentsResponse`

NewV1ListDeploymentsResponseWithDefaults instantiates a new V1ListDeploymentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployments

`func (o *V1ListDeploymentsResponse) GetDeployments() []V1Deployment`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *V1ListDeploymentsResponse) GetDeploymentsOk() (*[]V1Deployment, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *V1ListDeploymentsResponse) SetDeployments(v []V1Deployment)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *V1ListDeploymentsResponse) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListDeploymentsResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListDeploymentsResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListDeploymentsResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListDeploymentsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


