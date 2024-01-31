# V1ListBackupConfigsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupConfigs** | Pointer to [**[]V1BackupConfig**](V1BackupConfig.md) | The list of backup configurations. | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListBackupConfigsResponse

`func NewV1ListBackupConfigsResponse() *V1ListBackupConfigsResponse`

NewV1ListBackupConfigsResponse instantiates a new V1ListBackupConfigsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListBackupConfigsResponseWithDefaults

`func NewV1ListBackupConfigsResponseWithDefaults() *V1ListBackupConfigsResponse`

NewV1ListBackupConfigsResponseWithDefaults instantiates a new V1ListBackupConfigsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupConfigs

`func (o *V1ListBackupConfigsResponse) GetBackupConfigs() []V1BackupConfig`

GetBackupConfigs returns the BackupConfigs field if non-nil, zero value otherwise.

### GetBackupConfigsOk

`func (o *V1ListBackupConfigsResponse) GetBackupConfigsOk() (*[]V1BackupConfig, bool)`

GetBackupConfigsOk returns a tuple with the BackupConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupConfigs

`func (o *V1ListBackupConfigsResponse) SetBackupConfigs(v []V1BackupConfig)`

SetBackupConfigs sets BackupConfigs field to given value.

### HasBackupConfigs

`func (o *V1ListBackupConfigsResponse) HasBackupConfigs() bool`

HasBackupConfigs returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListBackupConfigsResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListBackupConfigsResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListBackupConfigsResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListBackupConfigsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


