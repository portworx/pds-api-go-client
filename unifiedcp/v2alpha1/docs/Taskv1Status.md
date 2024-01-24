# Taskv1Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalSteps** | Pointer to **string** | Total number of steps involved in the task. | [optional] 
**CurrentStep** | Pointer to **string** | Current step of the running task. | [optional] 
**Logs** | Pointer to **[]string** | List of log strings associated with the task. | [optional] 
**AssociatedResource** | Pointer to [**V1Reference**](V1Reference.md) |  | [optional] 
**State** | Pointer to [**StatusState**](StatusState.md) |  | [optional] [default to STATE_UNSPECIFIED]

## Methods

### NewTaskv1Status

`func NewTaskv1Status() *Taskv1Status`

NewTaskv1Status instantiates a new Taskv1Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskv1StatusWithDefaults

`func NewTaskv1StatusWithDefaults() *Taskv1Status`

NewTaskv1StatusWithDefaults instantiates a new Taskv1Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalSteps

`func (o *Taskv1Status) GetTotalSteps() string`

GetTotalSteps returns the TotalSteps field if non-nil, zero value otherwise.

### GetTotalStepsOk

`func (o *Taskv1Status) GetTotalStepsOk() (*string, bool)`

GetTotalStepsOk returns a tuple with the TotalSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSteps

`func (o *Taskv1Status) SetTotalSteps(v string)`

SetTotalSteps sets TotalSteps field to given value.

### HasTotalSteps

`func (o *Taskv1Status) HasTotalSteps() bool`

HasTotalSteps returns a boolean if a field has been set.

### GetCurrentStep

`func (o *Taskv1Status) GetCurrentStep() string`

GetCurrentStep returns the CurrentStep field if non-nil, zero value otherwise.

### GetCurrentStepOk

`func (o *Taskv1Status) GetCurrentStepOk() (*string, bool)`

GetCurrentStepOk returns a tuple with the CurrentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStep

`func (o *Taskv1Status) SetCurrentStep(v string)`

SetCurrentStep sets CurrentStep field to given value.

### HasCurrentStep

`func (o *Taskv1Status) HasCurrentStep() bool`

HasCurrentStep returns a boolean if a field has been set.

### GetLogs

`func (o *Taskv1Status) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *Taskv1Status) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *Taskv1Status) SetLogs(v []string)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *Taskv1Status) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetAssociatedResource

`func (o *Taskv1Status) GetAssociatedResource() V1Reference`

GetAssociatedResource returns the AssociatedResource field if non-nil, zero value otherwise.

### GetAssociatedResourceOk

`func (o *Taskv1Status) GetAssociatedResourceOk() (*V1Reference, bool)`

GetAssociatedResourceOk returns a tuple with the AssociatedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedResource

`func (o *Taskv1Status) SetAssociatedResource(v V1Reference)`

SetAssociatedResource sets AssociatedResource field to given value.

### HasAssociatedResource

`func (o *Taskv1Status) HasAssociatedResource() bool`

HasAssociatedResource returns a boolean if a field has been set.

### GetState

`func (o *Taskv1Status) GetState() StatusState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Taskv1Status) GetStateOk() (*StatusState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Taskv1Status) SetState(v StatusState)`

SetState sets State field to given value.

### HasState

`func (o *Taskv1Status) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


