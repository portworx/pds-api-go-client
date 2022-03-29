# ControllersCreateTenantBackupTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupCredentialsId** | Pointer to **string** |  | [optional] 
**Bucket** | Pointer to **string** | Bucket name for S3 or S3 compatible. Container name for Azure. | [optional] 
**Name** | Pointer to **string** | Name of the backup target. Must be unique for the given tenant. | [optional] 
**Region** | Pointer to **string** | Region of the bucket. Required for S3. Otherwise must be empty. | [optional] 
**Type** | Pointer to **string** | Type of the backup target. Must match the used backup credentials. | [optional] 

## Methods

### NewControllersCreateTenantBackupTarget

`func NewControllersCreateTenantBackupTarget() *ControllersCreateTenantBackupTarget`

NewControllersCreateTenantBackupTarget instantiates a new ControllersCreateTenantBackupTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersCreateTenantBackupTargetWithDefaults

`func NewControllersCreateTenantBackupTargetWithDefaults() *ControllersCreateTenantBackupTarget`

NewControllersCreateTenantBackupTargetWithDefaults instantiates a new ControllersCreateTenantBackupTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupCredentialsId

`func (o *ControllersCreateTenantBackupTarget) GetBackupCredentialsId() string`

GetBackupCredentialsId returns the BackupCredentialsId field if non-nil, zero value otherwise.

### GetBackupCredentialsIdOk

`func (o *ControllersCreateTenantBackupTarget) GetBackupCredentialsIdOk() (*string, bool)`

GetBackupCredentialsIdOk returns a tuple with the BackupCredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCredentialsId

`func (o *ControllersCreateTenantBackupTarget) SetBackupCredentialsId(v string)`

SetBackupCredentialsId sets BackupCredentialsId field to given value.

### HasBackupCredentialsId

`func (o *ControllersCreateTenantBackupTarget) HasBackupCredentialsId() bool`

HasBackupCredentialsId returns a boolean if a field has been set.

### GetBucket

`func (o *ControllersCreateTenantBackupTarget) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ControllersCreateTenantBackupTarget) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ControllersCreateTenantBackupTarget) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *ControllersCreateTenantBackupTarget) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetName

`func (o *ControllersCreateTenantBackupTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControllersCreateTenantBackupTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControllersCreateTenantBackupTarget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControllersCreateTenantBackupTarget) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *ControllersCreateTenantBackupTarget) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ControllersCreateTenantBackupTarget) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ControllersCreateTenantBackupTarget) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ControllersCreateTenantBackupTarget) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetType

`func (o *ControllersCreateTenantBackupTarget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ControllersCreateTenantBackupTarget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ControllersCreateTenantBackupTarget) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ControllersCreateTenantBackupTarget) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


