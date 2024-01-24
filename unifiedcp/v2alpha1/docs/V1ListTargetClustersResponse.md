# V1ListTargetClustersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]V1TargetCluster**](V1TargetCluster.md) |  | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListTargetClustersResponse

`func NewV1ListTargetClustersResponse() *V1ListTargetClustersResponse`

NewV1ListTargetClustersResponse instantiates a new V1ListTargetClustersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListTargetClustersResponseWithDefaults

`func NewV1ListTargetClustersResponseWithDefaults() *V1ListTargetClustersResponse`

NewV1ListTargetClustersResponseWithDefaults instantiates a new V1ListTargetClustersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *V1ListTargetClustersResponse) GetClusters() []V1TargetCluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *V1ListTargetClustersResponse) GetClustersOk() (*[]V1TargetCluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *V1ListTargetClustersResponse) SetClusters(v []V1TargetCluster)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *V1ListTargetClustersResponse) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListTargetClustersResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListTargetClustersResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListTargetClustersResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListTargetClustersResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


