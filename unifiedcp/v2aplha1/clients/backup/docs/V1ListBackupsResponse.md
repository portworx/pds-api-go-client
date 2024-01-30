# V1ListBackupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backups** | Pointer to [**[]V1Backup**](V1Backup.md) | List of backup resources. | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListBackupsResponse

`func NewV1ListBackupsResponse() *V1ListBackupsResponse`

NewV1ListBackupsResponse instantiates a new V1ListBackupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListBackupsResponseWithDefaults

`func NewV1ListBackupsResponseWithDefaults() *V1ListBackupsResponse`

NewV1ListBackupsResponseWithDefaults instantiates a new V1ListBackupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackups

`func (o *V1ListBackupsResponse) GetBackups() []V1Backup`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *V1ListBackupsResponse) GetBackupsOk() (*[]V1Backup, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *V1ListBackupsResponse) SetBackups(v []V1Backup)`

SetBackups sets Backups field to given value.

### HasBackups

`func (o *V1ListBackupsResponse) HasBackups() bool`

HasBackups returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListBackupsResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListBackupsResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListBackupsResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListBackupsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


