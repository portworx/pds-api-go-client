# V1ListDataServiceVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataServiceVersions** | Pointer to [**[]V1DataServiceVersion**](V1DataServiceVersion.md) |  | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListDataServiceVersionsResponse

`func NewV1ListDataServiceVersionsResponse() *V1ListDataServiceVersionsResponse`

NewV1ListDataServiceVersionsResponse instantiates a new V1ListDataServiceVersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListDataServiceVersionsResponseWithDefaults

`func NewV1ListDataServiceVersionsResponseWithDefaults() *V1ListDataServiceVersionsResponse`

NewV1ListDataServiceVersionsResponseWithDefaults instantiates a new V1ListDataServiceVersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataServiceVersions

`func (o *V1ListDataServiceVersionsResponse) GetDataServiceVersions() []V1DataServiceVersion`

GetDataServiceVersions returns the DataServiceVersions field if non-nil, zero value otherwise.

### GetDataServiceVersionsOk

`func (o *V1ListDataServiceVersionsResponse) GetDataServiceVersionsOk() (*[]V1DataServiceVersion, bool)`

GetDataServiceVersionsOk returns a tuple with the DataServiceVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataServiceVersions

`func (o *V1ListDataServiceVersionsResponse) SetDataServiceVersions(v []V1DataServiceVersion)`

SetDataServiceVersions sets DataServiceVersions field to given value.

### HasDataServiceVersions

`func (o *V1ListDataServiceVersionsResponse) HasDataServiceVersions() bool`

HasDataServiceVersions returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListDataServiceVersionsResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListDataServiceVersionsResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListDataServiceVersionsResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListDataServiceVersionsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


