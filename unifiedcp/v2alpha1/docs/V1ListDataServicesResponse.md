# V1ListDataServicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataServices** | Pointer to [**[]V1DataService**](V1DataService.md) | List of data service resources. | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListDataServicesResponse

`func NewV1ListDataServicesResponse() *V1ListDataServicesResponse`

NewV1ListDataServicesResponse instantiates a new V1ListDataServicesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListDataServicesResponseWithDefaults

`func NewV1ListDataServicesResponseWithDefaults() *V1ListDataServicesResponse`

NewV1ListDataServicesResponseWithDefaults instantiates a new V1ListDataServicesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataServices

`func (o *V1ListDataServicesResponse) GetDataServices() []V1DataService`

GetDataServices returns the DataServices field if non-nil, zero value otherwise.

### GetDataServicesOk

`func (o *V1ListDataServicesResponse) GetDataServicesOk() (*[]V1DataService, bool)`

GetDataServicesOk returns a tuple with the DataServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataServices

`func (o *V1ListDataServicesResponse) SetDataServices(v []V1DataService)`

SetDataServices sets DataServices field to given value.

### HasDataServices

`func (o *V1ListDataServicesResponse) HasDataServices() bool`

HasDataServices returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListDataServicesResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListDataServicesResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListDataServicesResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListDataServicesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


