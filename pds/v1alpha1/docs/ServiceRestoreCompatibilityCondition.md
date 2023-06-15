# ServiceRestoreCompatibilityCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupJob** | Pointer to **map[string]string** |  | [optional] 
**Image** | Pointer to **map[string]string** |  | [optional] 
**TargetCluster** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewServiceRestoreCompatibilityCondition

`func NewServiceRestoreCompatibilityCondition() *ServiceRestoreCompatibilityCondition`

NewServiceRestoreCompatibilityCondition instantiates a new ServiceRestoreCompatibilityCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRestoreCompatibilityConditionWithDefaults

`func NewServiceRestoreCompatibilityConditionWithDefaults() *ServiceRestoreCompatibilityCondition`

NewServiceRestoreCompatibilityConditionWithDefaults instantiates a new ServiceRestoreCompatibilityCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupJob

`func (o *ServiceRestoreCompatibilityCondition) GetBackupJob() map[string]string`

GetBackupJob returns the BackupJob field if non-nil, zero value otherwise.

### GetBackupJobOk

`func (o *ServiceRestoreCompatibilityCondition) GetBackupJobOk() (*map[string]string, bool)`

GetBackupJobOk returns a tuple with the BackupJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupJob

`func (o *ServiceRestoreCompatibilityCondition) SetBackupJob(v map[string]string)`

SetBackupJob sets BackupJob field to given value.

### HasBackupJob

`func (o *ServiceRestoreCompatibilityCondition) HasBackupJob() bool`

HasBackupJob returns a boolean if a field has been set.

### GetImage

`func (o *ServiceRestoreCompatibilityCondition) GetImage() map[string]string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ServiceRestoreCompatibilityCondition) GetImageOk() (*map[string]string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ServiceRestoreCompatibilityCondition) SetImage(v map[string]string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ServiceRestoreCompatibilityCondition) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetTargetCluster

`func (o *ServiceRestoreCompatibilityCondition) GetTargetCluster() map[string]string`

GetTargetCluster returns the TargetCluster field if non-nil, zero value otherwise.

### GetTargetClusterOk

`func (o *ServiceRestoreCompatibilityCondition) GetTargetClusterOk() (*map[string]string, bool)`

GetTargetClusterOk returns a tuple with the TargetCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCluster

`func (o *ServiceRestoreCompatibilityCondition) SetTargetCluster(v map[string]string)`

SetTargetCluster sets TargetCluster field to given value.

### HasTargetCluster

`func (o *ServiceRestoreCompatibilityCondition) HasTargetCluster() bool`

HasTargetCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


