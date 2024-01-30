# V1ListCompatibleDataServiceVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompatibleDataServiceVersions** | Pointer to [**[]V1CompatibleVersions**](V1CompatibleVersions.md) |  | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListCompatibleDataServiceVersionsResponse

`func NewV1ListCompatibleDataServiceVersionsResponse() *V1ListCompatibleDataServiceVersionsResponse`

NewV1ListCompatibleDataServiceVersionsResponse instantiates a new V1ListCompatibleDataServiceVersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListCompatibleDataServiceVersionsResponseWithDefaults

`func NewV1ListCompatibleDataServiceVersionsResponseWithDefaults() *V1ListCompatibleDataServiceVersionsResponse`

NewV1ListCompatibleDataServiceVersionsResponseWithDefaults instantiates a new V1ListCompatibleDataServiceVersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatibleDataServiceVersions

`func (o *V1ListCompatibleDataServiceVersionsResponse) GetCompatibleDataServiceVersions() []V1CompatibleVersions`

GetCompatibleDataServiceVersions returns the CompatibleDataServiceVersions field if non-nil, zero value otherwise.

### GetCompatibleDataServiceVersionsOk

`func (o *V1ListCompatibleDataServiceVersionsResponse) GetCompatibleDataServiceVersionsOk() (*[]V1CompatibleVersions, bool)`

GetCompatibleDataServiceVersionsOk returns a tuple with the CompatibleDataServiceVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleDataServiceVersions

`func (o *V1ListCompatibleDataServiceVersionsResponse) SetCompatibleDataServiceVersions(v []V1CompatibleVersions)`

SetCompatibleDataServiceVersions sets CompatibleDataServiceVersions field to given value.

### HasCompatibleDataServiceVersions

`func (o *V1ListCompatibleDataServiceVersionsResponse) HasCompatibleDataServiceVersions() bool`

HasCompatibleDataServiceVersions returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListCompatibleDataServiceVersionsResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListCompatibleDataServiceVersionsResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListCompatibleDataServiceVersionsResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListCompatibleDataServiceVersionsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


