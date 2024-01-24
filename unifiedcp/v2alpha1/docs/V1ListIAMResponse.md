# V1ListIAMResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iam** | Pointer to [**[]V1IAM**](V1IAM.md) |  | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListIAMResponse

`func NewV1ListIAMResponse() *V1ListIAMResponse`

NewV1ListIAMResponse instantiates a new V1ListIAMResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListIAMResponseWithDefaults

`func NewV1ListIAMResponseWithDefaults() *V1ListIAMResponse`

NewV1ListIAMResponseWithDefaults instantiates a new V1ListIAMResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIam

`func (o *V1ListIAMResponse) GetIam() []V1IAM`

GetIam returns the Iam field if non-nil, zero value otherwise.

### GetIamOk

`func (o *V1ListIAMResponse) GetIamOk() (*[]V1IAM, bool)`

GetIamOk returns a tuple with the Iam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIam

`func (o *V1ListIAMResponse) SetIam(v []V1IAM)`

SetIam sets Iam field to given value.

### HasIam

`func (o *V1ListIAMResponse) HasIam() bool`

HasIam returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListIAMResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListIAMResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListIAMResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListIAMResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


