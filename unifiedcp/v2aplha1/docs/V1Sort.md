# V1Sort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SortBy** | Pointer to [**SortByField**](SortByField.md) |  | [optional] [default to SORTBYFIELD_FIELD_UNSPECIFIED]
**SortOrder** | Pointer to [**SortOrderValue**](SortOrderValue.md) |  | [optional] [default to SORTORDERVALUE_VALUE_UNSPECIFIED]

## Methods

### NewV1Sort

`func NewV1Sort() *V1Sort`

NewV1Sort instantiates a new V1Sort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SortWithDefaults

`func NewV1SortWithDefaults() *V1Sort`

NewV1SortWithDefaults instantiates a new V1Sort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSortBy

`func (o *V1Sort) GetSortBy() SortByField`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *V1Sort) GetSortByOk() (*SortByField, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *V1Sort) SetSortBy(v SortByField)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *V1Sort) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetSortOrder

`func (o *V1Sort) GetSortOrder() SortOrderValue`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *V1Sort) GetSortOrderOk() (*SortOrderValue, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *V1Sort) SetSortOrder(v SortOrderValue)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *V1Sort) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


