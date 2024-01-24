# V1ListTenantsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenants** | Pointer to [**[]V1Tenant**](V1Tenant.md) |  | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListTenantsResponse

`func NewV1ListTenantsResponse() *V1ListTenantsResponse`

NewV1ListTenantsResponse instantiates a new V1ListTenantsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListTenantsResponseWithDefaults

`func NewV1ListTenantsResponseWithDefaults() *V1ListTenantsResponse`

NewV1ListTenantsResponseWithDefaults instantiates a new V1ListTenantsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenants

`func (o *V1ListTenantsResponse) GetTenants() []V1Tenant`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *V1ListTenantsResponse) GetTenantsOk() (*[]V1Tenant, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *V1ListTenantsResponse) SetTenants(v []V1Tenant)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *V1ListTenantsResponse) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListTenantsResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListTenantsResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListTenantsResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListTenantsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


