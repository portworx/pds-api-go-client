# V1ListBackupLocationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupLocations** | Pointer to [**[]V1BackupLocation**](V1BackupLocation.md) |  | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListBackupLocationsResponse

`func NewV1ListBackupLocationsResponse() *V1ListBackupLocationsResponse`

NewV1ListBackupLocationsResponse instantiates a new V1ListBackupLocationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListBackupLocationsResponseWithDefaults

`func NewV1ListBackupLocationsResponseWithDefaults() *V1ListBackupLocationsResponse`

NewV1ListBackupLocationsResponseWithDefaults instantiates a new V1ListBackupLocationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupLocations

`func (o *V1ListBackupLocationsResponse) GetBackupLocations() []V1BackupLocation`

GetBackupLocations returns the BackupLocations field if non-nil, zero value otherwise.

### GetBackupLocationsOk

`func (o *V1ListBackupLocationsResponse) GetBackupLocationsOk() (*[]V1BackupLocation, bool)`

GetBackupLocationsOk returns a tuple with the BackupLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLocations

`func (o *V1ListBackupLocationsResponse) SetBackupLocations(v []V1BackupLocation)`

SetBackupLocations sets BackupLocations field to given value.

### HasBackupLocations

`func (o *V1ListBackupLocationsResponse) HasBackupLocations() bool`

HasBackupLocations returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListBackupLocationsResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListBackupLocationsResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListBackupLocationsResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListBackupLocationsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


