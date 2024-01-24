# V1NamespaceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | Pointer to [**NamespacePhasePhase**](NamespacePhasePhase.md) |  | [optional] [default to PHASE_UNSPECIFIED]

## Methods

### NewV1NamespaceStatus

`func NewV1NamespaceStatus() *V1NamespaceStatus`

NewV1NamespaceStatus instantiates a new V1NamespaceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NamespaceStatusWithDefaults

`func NewV1NamespaceStatusWithDefaults() *V1NamespaceStatus`

NewV1NamespaceStatusWithDefaults instantiates a new V1NamespaceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *V1NamespaceStatus) GetPhase() NamespacePhasePhase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *V1NamespaceStatus) GetPhaseOk() (*NamespacePhasePhase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *V1NamespaceStatus) SetPhase(v NamespacePhasePhase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *V1NamespaceStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


