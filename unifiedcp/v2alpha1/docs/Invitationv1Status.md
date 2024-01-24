# Invitationv1Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Textual information for the current state of the invitation. | [optional] [readonly] 
**Phase** | Pointer to [**StatusPhase2**](StatusPhase2.md) |  | [optional] [default to PHASE_UNSPECIFIED]

## Methods

### NewInvitationv1Status

`func NewInvitationv1Status() *Invitationv1Status`

NewInvitationv1Status instantiates a new Invitationv1Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationv1StatusWithDefaults

`func NewInvitationv1StatusWithDefaults() *Invitationv1Status`

NewInvitationv1StatusWithDefaults instantiates a new Invitationv1Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *Invitationv1Status) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Invitationv1Status) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Invitationv1Status) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Invitationv1Status) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetPhase

`func (o *Invitationv1Status) GetPhase() StatusPhase2`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *Invitationv1Status) GetPhaseOk() (*StatusPhase2, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *Invitationv1Status) SetPhase(v StatusPhase2)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *Invitationv1Status) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


