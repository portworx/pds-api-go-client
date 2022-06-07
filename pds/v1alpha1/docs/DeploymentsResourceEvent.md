# DeploymentsResourceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action that was taken/failed regarding to the given object. | [optional] 
**Message** | Pointer to **string** | Message is a human-readable description of the status of this operation. | [optional] 
**Name** | Pointer to **string** | Name of the Event resource in target cluster. | [optional] 
**Reason** | Pointer to **string** | Reason is a short, machine understandable string that gives the reason for the transition into the object&#39;s current status. | [optional] 
**Timestamp** | Pointer to **string** | Timestamp informs about when did the event occur most recently. | [optional] 
**Type** | Pointer to **string** | Type of this event , new types could be added in the future. | [optional] 

## Methods

### NewDeploymentsResourceEvent

`func NewDeploymentsResourceEvent() *DeploymentsResourceEvent`

NewDeploymentsResourceEvent instantiates a new DeploymentsResourceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentsResourceEventWithDefaults

`func NewDeploymentsResourceEventWithDefaults() *DeploymentsResourceEvent`

NewDeploymentsResourceEventWithDefaults instantiates a new DeploymentsResourceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DeploymentsResourceEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeploymentsResourceEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeploymentsResourceEvent) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DeploymentsResourceEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetMessage

`func (o *DeploymentsResourceEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeploymentsResourceEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeploymentsResourceEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeploymentsResourceEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *DeploymentsResourceEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentsResourceEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentsResourceEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentsResourceEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReason

`func (o *DeploymentsResourceEvent) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DeploymentsResourceEvent) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DeploymentsResourceEvent) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DeploymentsResourceEvent) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetTimestamp

`func (o *DeploymentsResourceEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DeploymentsResourceEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DeploymentsResourceEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DeploymentsResourceEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *DeploymentsResourceEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentsResourceEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentsResourceEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeploymentsResourceEvent) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


