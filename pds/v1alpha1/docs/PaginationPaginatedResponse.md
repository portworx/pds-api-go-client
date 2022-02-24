# PaginationPaginatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Members** | Pointer to **map[string]interface{}** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PerPage** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] 

## Methods

### NewPaginationPaginatedResponse

`func NewPaginationPaginatedResponse() *PaginationPaginatedResponse`

NewPaginationPaginatedResponse instantiates a new PaginationPaginatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationPaginatedResponseWithDefaults

`func NewPaginationPaginatedResponseWithDefaults() *PaginationPaginatedResponse`

NewPaginationPaginatedResponseWithDefaults instantiates a new PaginationPaginatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginationPaginatedResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginationPaginatedResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginationPaginatedResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginationPaginatedResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMembers

`func (o *PaginationPaginatedResponse) GetMembers() map[string]interface{}`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *PaginationPaginatedResponse) GetMembersOk() (*map[string]interface{}, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *PaginationPaginatedResponse) SetMembers(v map[string]interface{})`

SetMembers sets Members field to given value.

### HasMembers

`func (o *PaginationPaginatedResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetPage

`func (o *PaginationPaginatedResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PaginationPaginatedResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PaginationPaginatedResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PaginationPaginatedResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *PaginationPaginatedResponse) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *PaginationPaginatedResponse) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *PaginationPaginatedResponse) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *PaginationPaginatedResponse) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetTotal

`func (o *PaginationPaginatedResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PaginationPaginatedResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PaginationPaginatedResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PaginationPaginatedResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalPages

`func (o *PaginationPaginatedResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PaginationPaginatedResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PaginationPaginatedResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PaginationPaginatedResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


