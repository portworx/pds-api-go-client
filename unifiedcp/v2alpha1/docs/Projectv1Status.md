# Projectv1Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Textual information for the current state of the project. | [optional] 
**Phase** | Pointer to [**V1PhaseType**](V1PhaseType.md) |  | [optional] [default to TYPE_UNSPECIFIED]

## Methods

### NewProjectv1Status

`func NewProjectv1Status() *Projectv1Status`

NewProjectv1Status instantiates a new Projectv1Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectv1StatusWithDefaults

`func NewProjectv1StatusWithDefaults() *Projectv1Status`

NewProjectv1StatusWithDefaults instantiates a new Projectv1Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *Projectv1Status) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Projectv1Status) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Projectv1Status) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Projectv1Status) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetPhase

`func (o *Projectv1Status) GetPhase() V1PhaseType`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *Projectv1Status) GetPhaseOk() (*V1PhaseType, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *Projectv1Status) SetPhase(v V1PhaseType)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *Projectv1Status) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


