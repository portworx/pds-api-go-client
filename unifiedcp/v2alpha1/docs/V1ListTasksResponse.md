# V1ListTasksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tasks** | Pointer to [**[]V1Task**](V1Task.md) | List of Task resources. | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListTasksResponse

`func NewV1ListTasksResponse() *V1ListTasksResponse`

NewV1ListTasksResponse instantiates a new V1ListTasksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListTasksResponseWithDefaults

`func NewV1ListTasksResponseWithDefaults() *V1ListTasksResponse`

NewV1ListTasksResponseWithDefaults instantiates a new V1ListTasksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTasks

`func (o *V1ListTasksResponse) GetTasks() []V1Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *V1ListTasksResponse) GetTasksOk() (*[]V1Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *V1ListTasksResponse) SetTasks(v []V1Task)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *V1ListTasksResponse) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListTasksResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListTasksResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListTasksResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListTasksResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


