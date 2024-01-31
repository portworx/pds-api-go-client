# V1ListRestoresResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Restores** | Pointer to [**[]V1Restore**](V1Restore.md) |  | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListRestoresResponse

`func NewV1ListRestoresResponse() *V1ListRestoresResponse`

NewV1ListRestoresResponse instantiates a new V1ListRestoresResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListRestoresResponseWithDefaults

`func NewV1ListRestoresResponseWithDefaults() *V1ListRestoresResponse`

NewV1ListRestoresResponseWithDefaults instantiates a new V1ListRestoresResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestores

`func (o *V1ListRestoresResponse) GetRestores() []V1Restore`

GetRestores returns the Restores field if non-nil, zero value otherwise.

### GetRestoresOk

`func (o *V1ListRestoresResponse) GetRestoresOk() (*[]V1Restore, bool)`

GetRestoresOk returns a tuple with the Restores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestores

`func (o *V1ListRestoresResponse) SetRestores(v []V1Restore)`

SetRestores sets Restores field to given value.

### HasRestores

`func (o *V1ListRestoresResponse) HasRestores() bool`

HasRestores returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListRestoresResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListRestoresResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListRestoresResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListRestoresResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


