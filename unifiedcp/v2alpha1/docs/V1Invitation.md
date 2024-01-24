# V1Invitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V1Meta**](V1Meta.md) |  | [optional] 
**Config** | Pointer to [**Platforminvitationv1Config**](Platforminvitationv1Config.md) |  | [optional] 
**Status** | Pointer to [**Invitationv1Status**](Invitationv1Status.md) |  | [optional] 

## Methods

### NewV1Invitation

`func NewV1Invitation() *V1Invitation`

NewV1Invitation instantiates a new V1Invitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InvitationWithDefaults

`func NewV1InvitationWithDefaults() *V1Invitation`

NewV1InvitationWithDefaults instantiates a new V1Invitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V1Invitation) GetMeta() V1Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V1Invitation) GetMetaOk() (*V1Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V1Invitation) SetMeta(v V1Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V1Invitation) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetConfig

`func (o *V1Invitation) GetConfig() Platforminvitationv1Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V1Invitation) GetConfigOk() (*Platforminvitationv1Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V1Invitation) SetConfig(v Platforminvitationv1Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *V1Invitation) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *V1Invitation) GetStatus() Invitationv1Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1Invitation) GetStatusOk() (*Invitationv1Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1Invitation) SetStatus(v Invitationv1Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1Invitation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


