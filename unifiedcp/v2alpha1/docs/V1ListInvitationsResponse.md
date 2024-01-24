# V1ListInvitationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invitations** | Pointer to [**[]V1Invitation**](V1Invitation.md) |  | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListInvitationsResponse

`func NewV1ListInvitationsResponse() *V1ListInvitationsResponse`

NewV1ListInvitationsResponse instantiates a new V1ListInvitationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListInvitationsResponseWithDefaults

`func NewV1ListInvitationsResponseWithDefaults() *V1ListInvitationsResponse`

NewV1ListInvitationsResponseWithDefaults instantiates a new V1ListInvitationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvitations

`func (o *V1ListInvitationsResponse) GetInvitations() []V1Invitation`

GetInvitations returns the Invitations field if non-nil, zero value otherwise.

### GetInvitationsOk

`func (o *V1ListInvitationsResponse) GetInvitationsOk() (*[]V1Invitation, bool)`

GetInvitationsOk returns a tuple with the Invitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitations

`func (o *V1ListInvitationsResponse) SetInvitations(v []V1Invitation)`

SetInvitations sets Invitations field to given value.

### HasInvitations

`func (o *V1ListInvitationsResponse) HasInvitations() bool`

HasInvitations returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListInvitationsResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListInvitationsResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListInvitationsResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListInvitationsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


