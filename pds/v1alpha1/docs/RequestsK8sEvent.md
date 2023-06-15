# RequestsK8sEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action that was taken/failed regarding the given object. | [optional] 
**Message** | Pointer to **string** | Message is a human-readable description of the status of this operation. | [optional] 
**Name** | Pointer to **string** | Name of the Event resource in target cluster. | [optional] 
**Reason** | Pointer to **string** | Reason is a short, machine understandable string that gives the reason for the transition into the object&#39;s current status. | [optional] 
**ResourceKind** | Pointer to **string** | ResourceKind is the type of resource being referenced | [optional] 
**ResourceName** | Pointer to **string** | ResourceName is the name of resource being referenced | [optional] 
**Timestamp** | Pointer to **string** | Timestamp informs about when did the event occur most recently. | [optional] 
**Type** | Pointer to **string** | Type of this event , new types could be added in the future. | [optional] 

## Methods

### NewRequestsK8sEvent

`func NewRequestsK8sEvent() *RequestsK8sEvent`

NewRequestsK8sEvent instantiates a new RequestsK8sEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsK8sEventWithDefaults

`func NewRequestsK8sEventWithDefaults() *RequestsK8sEvent`

NewRequestsK8sEventWithDefaults instantiates a new RequestsK8sEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RequestsK8sEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RequestsK8sEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RequestsK8sEvent) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RequestsK8sEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetMessage

`func (o *RequestsK8sEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RequestsK8sEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RequestsK8sEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RequestsK8sEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *RequestsK8sEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestsK8sEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestsK8sEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequestsK8sEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReason

`func (o *RequestsK8sEvent) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RequestsK8sEvent) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RequestsK8sEvent) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RequestsK8sEvent) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetResourceKind

`func (o *RequestsK8sEvent) GetResourceKind() string`

GetResourceKind returns the ResourceKind field if non-nil, zero value otherwise.

### GetResourceKindOk

`func (o *RequestsK8sEvent) GetResourceKindOk() (*string, bool)`

GetResourceKindOk returns a tuple with the ResourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKind

`func (o *RequestsK8sEvent) SetResourceKind(v string)`

SetResourceKind sets ResourceKind field to given value.

### HasResourceKind

`func (o *RequestsK8sEvent) HasResourceKind() bool`

HasResourceKind returns a boolean if a field has been set.

### GetResourceName

`func (o *RequestsK8sEvent) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *RequestsK8sEvent) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *RequestsK8sEvent) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *RequestsK8sEvent) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetTimestamp

`func (o *RequestsK8sEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RequestsK8sEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RequestsK8sEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RequestsK8sEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *RequestsK8sEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestsK8sEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestsK8sEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RequestsK8sEvent) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


