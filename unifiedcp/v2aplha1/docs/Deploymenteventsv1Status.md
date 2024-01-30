# Deploymenteventsv1Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action involved in the event. | [optional] 
**Count** | Pointer to **string** | No. of times the event has been generated. | [optional] 
**Message** | Pointer to **string** | Message related to the event. | [optional] 
**Reason** | Pointer to **string** | Reason for the event. | [optional] 
**ResourceKind** | Pointer to **string** | Resource Kind. | [optional] 
**ResourceName** | Pointer to **string** | Resource Name. | [optional] 
**TimestampTime** | Pointer to **time.Time** | Timestamp of the event. | [optional] 
**Type** | Pointer to [**V1EventType**](V1EventType.md) |  | [optional] [default to V1EVENTTYPE_EVENT_TYPE_UNSPECIFIED]

## Methods

### NewDeploymenteventsv1Status

`func NewDeploymenteventsv1Status() *Deploymenteventsv1Status`

NewDeploymenteventsv1Status instantiates a new Deploymenteventsv1Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymenteventsv1StatusWithDefaults

`func NewDeploymenteventsv1StatusWithDefaults() *Deploymenteventsv1Status`

NewDeploymenteventsv1StatusWithDefaults instantiates a new Deploymenteventsv1Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Deploymenteventsv1Status) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Deploymenteventsv1Status) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Deploymenteventsv1Status) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Deploymenteventsv1Status) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCount

`func (o *Deploymenteventsv1Status) GetCount() string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Deploymenteventsv1Status) GetCountOk() (*string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Deploymenteventsv1Status) SetCount(v string)`

SetCount sets Count field to given value.

### HasCount

`func (o *Deploymenteventsv1Status) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMessage

`func (o *Deploymenteventsv1Status) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Deploymenteventsv1Status) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Deploymenteventsv1Status) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Deploymenteventsv1Status) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *Deploymenteventsv1Status) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Deploymenteventsv1Status) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Deploymenteventsv1Status) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Deploymenteventsv1Status) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetResourceKind

`func (o *Deploymenteventsv1Status) GetResourceKind() string`

GetResourceKind returns the ResourceKind field if non-nil, zero value otherwise.

### GetResourceKindOk

`func (o *Deploymenteventsv1Status) GetResourceKindOk() (*string, bool)`

GetResourceKindOk returns a tuple with the ResourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKind

`func (o *Deploymenteventsv1Status) SetResourceKind(v string)`

SetResourceKind sets ResourceKind field to given value.

### HasResourceKind

`func (o *Deploymenteventsv1Status) HasResourceKind() bool`

HasResourceKind returns a boolean if a field has been set.

### GetResourceName

`func (o *Deploymenteventsv1Status) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *Deploymenteventsv1Status) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *Deploymenteventsv1Status) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *Deploymenteventsv1Status) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetTimestampTime

`func (o *Deploymenteventsv1Status) GetTimestampTime() time.Time`

GetTimestampTime returns the TimestampTime field if non-nil, zero value otherwise.

### GetTimestampTimeOk

`func (o *Deploymenteventsv1Status) GetTimestampTimeOk() (*time.Time, bool)`

GetTimestampTimeOk returns a tuple with the TimestampTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampTime

`func (o *Deploymenteventsv1Status) SetTimestampTime(v time.Time)`

SetTimestampTime sets TimestampTime field to given value.

### HasTimestampTime

`func (o *Deploymenteventsv1Status) HasTimestampTime() bool`

HasTimestampTime returns a boolean if a field has been set.

### GetType

`func (o *Deploymenteventsv1Status) GetType() V1EventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Deploymenteventsv1Status) GetTypeOk() (*V1EventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Deploymenteventsv1Status) SetType(v V1EventType)`

SetType sets Type field to given value.

### HasType

`func (o *Deploymenteventsv1Status) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


