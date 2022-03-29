# ControllersPaginatedBackups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ModelsBackup**](ModelsBackup.md) |  | [optional] 
**Pagination** | Pointer to [**ConstraintPagination**](ConstraintPagination.md) |  | [optional] 

## Methods

### NewControllersPaginatedBackups

`func NewControllersPaginatedBackups() *ControllersPaginatedBackups`

NewControllersPaginatedBackups instantiates a new ControllersPaginatedBackups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersPaginatedBackupsWithDefaults

`func NewControllersPaginatedBackupsWithDefaults() *ControllersPaginatedBackups`

NewControllersPaginatedBackupsWithDefaults instantiates a new ControllersPaginatedBackups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ControllersPaginatedBackups) GetData() []ModelsBackup`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ControllersPaginatedBackups) GetDataOk() (*[]ModelsBackup, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ControllersPaginatedBackups) SetData(v []ModelsBackup)`

SetData sets Data field to given value.

### HasData

`func (o *ControllersPaginatedBackups) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ControllersPaginatedBackups) GetPagination() ConstraintPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ControllersPaginatedBackups) GetPaginationOk() (*ConstraintPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ControllersPaginatedBackups) SetPagination(v ConstraintPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ControllersPaginatedBackups) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


