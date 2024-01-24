# V1SourceReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | Pointer to **string** | UID of the deployment which was backed up. | [optional] [readonly] 
**BackupId** | Pointer to **string** | UID of the backup. | [optional] 
**BackupLocationId** | Pointer to **string** | UID of the backup location. | [optional] [readonly] 
**CloudsnapId** | Pointer to **string** | UID of the cloud snapshot of the backup volume used for restore. | [optional] [readonly] 

## Methods

### NewV1SourceReferences

`func NewV1SourceReferences() *V1SourceReferences`

NewV1SourceReferences instantiates a new V1SourceReferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SourceReferencesWithDefaults

`func NewV1SourceReferencesWithDefaults() *V1SourceReferences`

NewV1SourceReferencesWithDefaults instantiates a new V1SourceReferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *V1SourceReferences) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *V1SourceReferences) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *V1SourceReferences) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *V1SourceReferences) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetBackupId

`func (o *V1SourceReferences) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *V1SourceReferences) GetBackupIdOk() (*string, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *V1SourceReferences) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.

### HasBackupId

`func (o *V1SourceReferences) HasBackupId() bool`

HasBackupId returns a boolean if a field has been set.

### GetBackupLocationId

`func (o *V1SourceReferences) GetBackupLocationId() string`

GetBackupLocationId returns the BackupLocationId field if non-nil, zero value otherwise.

### GetBackupLocationIdOk

`func (o *V1SourceReferences) GetBackupLocationIdOk() (*string, bool)`

GetBackupLocationIdOk returns a tuple with the BackupLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLocationId

`func (o *V1SourceReferences) SetBackupLocationId(v string)`

SetBackupLocationId sets BackupLocationId field to given value.

### HasBackupLocationId

`func (o *V1SourceReferences) HasBackupLocationId() bool`

HasBackupLocationId returns a boolean if a field has been set.

### GetCloudsnapId

`func (o *V1SourceReferences) GetCloudsnapId() string`

GetCloudsnapId returns the CloudsnapId field if non-nil, zero value otherwise.

### GetCloudsnapIdOk

`func (o *V1SourceReferences) GetCloudsnapIdOk() (*string, bool)`

GetCloudsnapIdOk returns a tuple with the CloudsnapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudsnapId

`func (o *V1SourceReferences) SetCloudsnapId(v string)`

SetCloudsnapId sets CloudsnapId field to given value.

### HasCloudsnapId

`func (o *V1SourceReferences) HasCloudsnapId() bool`

HasCloudsnapId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


