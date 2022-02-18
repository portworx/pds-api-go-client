# ControllersUpdateBackupPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Schedules** | Pointer to [**[]ModelsBackupSchedule**](ModelsBackupSchedule.md) |  | [optional] 

## Methods

### NewControllersUpdateBackupPolicyRequest

`func NewControllersUpdateBackupPolicyRequest() *ControllersUpdateBackupPolicyRequest`

NewControllersUpdateBackupPolicyRequest instantiates a new ControllersUpdateBackupPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersUpdateBackupPolicyRequestWithDefaults

`func NewControllersUpdateBackupPolicyRequestWithDefaults() *ControllersUpdateBackupPolicyRequest`

NewControllersUpdateBackupPolicyRequestWithDefaults instantiates a new ControllersUpdateBackupPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ControllersUpdateBackupPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControllersUpdateBackupPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControllersUpdateBackupPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControllersUpdateBackupPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchedules

`func (o *ControllersUpdateBackupPolicyRequest) GetSchedules() []ModelsBackupSchedule`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *ControllersUpdateBackupPolicyRequest) GetSchedulesOk() (*[]ModelsBackupSchedule, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *ControllersUpdateBackupPolicyRequest) SetSchedules(v []ModelsBackupSchedule)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *ControllersUpdateBackupPolicyRequest) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


