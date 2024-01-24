# V1ListServiceAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccounts** | Pointer to [**[]V1ServiceAccount**](V1ServiceAccount.md) |  | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListServiceAccountResponse

`func NewV1ListServiceAccountResponse() *V1ListServiceAccountResponse`

NewV1ListServiceAccountResponse instantiates a new V1ListServiceAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListServiceAccountResponseWithDefaults

`func NewV1ListServiceAccountResponseWithDefaults() *V1ListServiceAccountResponse`

NewV1ListServiceAccountResponseWithDefaults instantiates a new V1ListServiceAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccounts

`func (o *V1ListServiceAccountResponse) GetServiceAccounts() []V1ServiceAccount`

GetServiceAccounts returns the ServiceAccounts field if non-nil, zero value otherwise.

### GetServiceAccountsOk

`func (o *V1ListServiceAccountResponse) GetServiceAccountsOk() (*[]V1ServiceAccount, bool)`

GetServiceAccountsOk returns a tuple with the ServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccounts

`func (o *V1ListServiceAccountResponse) SetServiceAccounts(v []V1ServiceAccount)`

SetServiceAccounts sets ServiceAccounts field to given value.

### HasServiceAccounts

`func (o *V1ListServiceAccountResponse) HasServiceAccounts() bool`

HasServiceAccounts returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListServiceAccountResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListServiceAccountResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListServiceAccountResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListServiceAccountResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


