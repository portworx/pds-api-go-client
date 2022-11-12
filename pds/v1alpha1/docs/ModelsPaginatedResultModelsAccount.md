# ModelsPaginatedResultModelsAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ModelsAccount**](ModelsAccount.md) |  | [optional] 
**Pagination** | Pointer to [**ConstraintPagination**](ConstraintPagination.md) |  | [optional] 

## Methods

### NewModelsPaginatedResultModelsAccount

`func NewModelsPaginatedResultModelsAccount() *ModelsPaginatedResultModelsAccount`

NewModelsPaginatedResultModelsAccount instantiates a new ModelsPaginatedResultModelsAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsPaginatedResultModelsAccountWithDefaults

`func NewModelsPaginatedResultModelsAccountWithDefaults() *ModelsPaginatedResultModelsAccount`

NewModelsPaginatedResultModelsAccountWithDefaults instantiates a new ModelsPaginatedResultModelsAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ModelsPaginatedResultModelsAccount) GetData() []ModelsAccount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelsPaginatedResultModelsAccount) GetDataOk() (*[]ModelsAccount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelsPaginatedResultModelsAccount) SetData(v []ModelsAccount)`

SetData sets Data field to given value.

### HasData

`func (o *ModelsPaginatedResultModelsAccount) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ModelsPaginatedResultModelsAccount) GetPagination() ConstraintPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ModelsPaginatedResultModelsAccount) GetPaginationOk() (*ConstraintPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ModelsPaginatedResultModelsAccount) SetPagination(v ConstraintPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ModelsPaginatedResultModelsAccount) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


