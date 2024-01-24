# V1ListAccountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]V1Account**](V1Account.md) |  | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListAccountsResponse

`func NewV1ListAccountsResponse() *V1ListAccountsResponse`

NewV1ListAccountsResponse instantiates a new V1ListAccountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListAccountsResponseWithDefaults

`func NewV1ListAccountsResponseWithDefaults() *V1ListAccountsResponse`

NewV1ListAccountsResponseWithDefaults instantiates a new V1ListAccountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *V1ListAccountsResponse) GetAccounts() []V1Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V1ListAccountsResponse) GetAccountsOk() (*[]V1Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V1ListAccountsResponse) SetAccounts(v []V1Account)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V1ListAccountsResponse) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListAccountsResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListAccountsResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListAccountsResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListAccountsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


