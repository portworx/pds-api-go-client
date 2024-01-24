# Tenantv1Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Textual information for the current state of the tenant. | [optional] 
**Phase** | Pointer to [**V1PhaseType1**](V1PhaseType1.md) |  | [optional] [default to TYPE_UNSPECIFIED]

## Methods

### NewTenantv1Status

`func NewTenantv1Status() *Tenantv1Status`

NewTenantv1Status instantiates a new Tenantv1Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantv1StatusWithDefaults

`func NewTenantv1StatusWithDefaults() *Tenantv1Status`

NewTenantv1StatusWithDefaults instantiates a new Tenantv1Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *Tenantv1Status) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Tenantv1Status) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Tenantv1Status) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Tenantv1Status) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetPhase

`func (o *Tenantv1Status) GetPhase() V1PhaseType1`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *Tenantv1Status) GetPhaseOk() (*V1PhaseType1, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *Tenantv1Status) SetPhase(v V1PhaseType1)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *Tenantv1Status) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


