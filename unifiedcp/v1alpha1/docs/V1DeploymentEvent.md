# V1DeploymentEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V1Meta**](V1Meta.md) |  | [optional] 
**Status** | Pointer to [**Deploymenteventsv1Status**](Deploymenteventsv1Status.md) |  | [optional] 

## Methods

### NewV1DeploymentEvent

`func NewV1DeploymentEvent() *V1DeploymentEvent`

NewV1DeploymentEvent instantiates a new V1DeploymentEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeploymentEventWithDefaults

`func NewV1DeploymentEventWithDefaults() *V1DeploymentEvent`

NewV1DeploymentEventWithDefaults instantiates a new V1DeploymentEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V1DeploymentEvent) GetMeta() V1Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V1DeploymentEvent) GetMetaOk() (*V1Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V1DeploymentEvent) SetMeta(v V1Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V1DeploymentEvent) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetStatus

`func (o *V1DeploymentEvent) GetStatus() Deploymenteventsv1Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1DeploymentEvent) GetStatusOk() (*Deploymenteventsv1Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1DeploymentEvent) SetStatus(v Deploymenteventsv1Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1DeploymentEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


