# DeploymentsResourceEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]DeploymentsResourceEvent**](DeploymentsResourceEvent.md) |  | [optional] 
**Resource** | Pointer to [**V1TypedLocalObjectReference**](V1TypedLocalObjectReference.md) |  | [optional] 

## Methods

### NewDeploymentsResourceEvents

`func NewDeploymentsResourceEvents() *DeploymentsResourceEvents`

NewDeploymentsResourceEvents instantiates a new DeploymentsResourceEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentsResourceEventsWithDefaults

`func NewDeploymentsResourceEventsWithDefaults() *DeploymentsResourceEvents`

NewDeploymentsResourceEventsWithDefaults instantiates a new DeploymentsResourceEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *DeploymentsResourceEvents) GetEvents() []DeploymentsResourceEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *DeploymentsResourceEvents) GetEventsOk() (*[]DeploymentsResourceEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *DeploymentsResourceEvents) SetEvents(v []DeploymentsResourceEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *DeploymentsResourceEvents) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetResource

`func (o *DeploymentsResourceEvents) GetResource() V1TypedLocalObjectReference`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DeploymentsResourceEvents) GetResourceOk() (*V1TypedLocalObjectReference, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DeploymentsResourceEvents) SetResource(v V1TypedLocalObjectReference)`

SetResource sets Resource field to given value.

### HasResource

`func (o *DeploymentsResourceEvents) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


