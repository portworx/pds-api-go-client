# V1ListDeploymentEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentEvents** | Pointer to [**[]V1DeploymentEvent**](V1DeploymentEvent.md) | List of Event resources. | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListDeploymentEventsResponse

`func NewV1ListDeploymentEventsResponse() *V1ListDeploymentEventsResponse`

NewV1ListDeploymentEventsResponse instantiates a new V1ListDeploymentEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListDeploymentEventsResponseWithDefaults

`func NewV1ListDeploymentEventsResponseWithDefaults() *V1ListDeploymentEventsResponse`

NewV1ListDeploymentEventsResponseWithDefaults instantiates a new V1ListDeploymentEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentEvents

`func (o *V1ListDeploymentEventsResponse) GetDeploymentEvents() []V1DeploymentEvent`

GetDeploymentEvents returns the DeploymentEvents field if non-nil, zero value otherwise.

### GetDeploymentEventsOk

`func (o *V1ListDeploymentEventsResponse) GetDeploymentEventsOk() (*[]V1DeploymentEvent, bool)`

GetDeploymentEventsOk returns a tuple with the DeploymentEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEvents

`func (o *V1ListDeploymentEventsResponse) SetDeploymentEvents(v []V1DeploymentEvent)`

SetDeploymentEvents sets DeploymentEvents field to given value.

### HasDeploymentEvents

`func (o *V1ListDeploymentEventsResponse) HasDeploymentEvents() bool`

HasDeploymentEvents returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListDeploymentEventsResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListDeploymentEventsResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListDeploymentEventsResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListDeploymentEventsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


