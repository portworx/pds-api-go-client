# V1References

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | Pointer to **string** | UID of the deployment to which the backup configuration belong. | [optional] [readonly] 
**BackupLocationId** | Pointer to **string** | UID of the backup target to which the backup configuration belong. | [optional] 
**DataServiceId** | Pointer to **string** | UID of the data service to which the backup configuration belong. | [optional] [readonly] 

## Methods

### NewV1References

`func NewV1References() *V1References`

NewV1References instantiates a new V1References object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ReferencesWithDefaults

`func NewV1ReferencesWithDefaults() *V1References`

NewV1ReferencesWithDefaults instantiates a new V1References object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *V1References) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *V1References) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *V1References) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *V1References) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetBackupLocationId

`func (o *V1References) GetBackupLocationId() string`

GetBackupLocationId returns the BackupLocationId field if non-nil, zero value otherwise.

### GetBackupLocationIdOk

`func (o *V1References) GetBackupLocationIdOk() (*string, bool)`

GetBackupLocationIdOk returns a tuple with the BackupLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLocationId

`func (o *V1References) SetBackupLocationId(v string)`

SetBackupLocationId sets BackupLocationId field to given value.

### HasBackupLocationId

`func (o *V1References) HasBackupLocationId() bool`

HasBackupLocationId returns a boolean if a field has been set.

### GetDataServiceId

`func (o *V1References) GetDataServiceId() string`

GetDataServiceId returns the DataServiceId field if non-nil, zero value otherwise.

### GetDataServiceIdOk

`func (o *V1References) GetDataServiceIdOk() (*string, bool)`

GetDataServiceIdOk returns a tuple with the DataServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataServiceId

`func (o *V1References) SetDataServiceId(v string)`

SetDataServiceId sets DataServiceId field to given value.

### HasDataServiceId

`func (o *V1References) HasDataServiceId() bool`

HasDataServiceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


