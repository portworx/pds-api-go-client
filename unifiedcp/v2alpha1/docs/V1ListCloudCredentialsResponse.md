# V1ListCloudCredentialsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudCredentials** | Pointer to [**[]V1CloudCredential**](V1CloudCredential.md) |  | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListCloudCredentialsResponse

`func NewV1ListCloudCredentialsResponse() *V1ListCloudCredentialsResponse`

NewV1ListCloudCredentialsResponse instantiates a new V1ListCloudCredentialsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListCloudCredentialsResponseWithDefaults

`func NewV1ListCloudCredentialsResponseWithDefaults() *V1ListCloudCredentialsResponse`

NewV1ListCloudCredentialsResponseWithDefaults instantiates a new V1ListCloudCredentialsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudCredentials

`func (o *V1ListCloudCredentialsResponse) GetCloudCredentials() []V1CloudCredential`

GetCloudCredentials returns the CloudCredentials field if non-nil, zero value otherwise.

### GetCloudCredentialsOk

`func (o *V1ListCloudCredentialsResponse) GetCloudCredentialsOk() (*[]V1CloudCredential, bool)`

GetCloudCredentialsOk returns a tuple with the CloudCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentials

`func (o *V1ListCloudCredentialsResponse) SetCloudCredentials(v []V1CloudCredential)`

SetCloudCredentials sets CloudCredentials field to given value.

### HasCloudCredentials

`func (o *V1ListCloudCredentialsResponse) HasCloudCredentials() bool`

HasCloudCredentials returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListCloudCredentialsResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListCloudCredentialsResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListCloudCredentialsResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListCloudCredentialsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


