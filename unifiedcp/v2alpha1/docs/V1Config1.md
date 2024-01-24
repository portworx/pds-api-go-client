# V1Config1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**References** | Pointer to [**V1References1**](V1References1.md) |  | [optional] 
**BackupCapability** | Pointer to **string** | BackupCapability of the deployment target when the snapshot was created. | [optional] 

## Methods

### NewV1Config1

`func NewV1Config1() *V1Config1`

NewV1Config1 instantiates a new V1Config1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1Config1WithDefaults

`func NewV1Config1WithDefaults() *V1Config1`

NewV1Config1WithDefaults instantiates a new V1Config1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferences

`func (o *V1Config1) GetReferences() V1References1`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *V1Config1) GetReferencesOk() (*V1References1, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *V1Config1) SetReferences(v V1References1)`

SetReferences sets References field to given value.

### HasReferences

`func (o *V1Config1) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetBackupCapability

`func (o *V1Config1) GetBackupCapability() string`

GetBackupCapability returns the BackupCapability field if non-nil, zero value otherwise.

### GetBackupCapabilityOk

`func (o *V1Config1) GetBackupCapabilityOk() (*string, bool)`

GetBackupCapabilityOk returns a tuple with the BackupCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCapability

`func (o *V1Config1) SetBackupCapability(v string)`

SetBackupCapability sets BackupCapability field to given value.

### HasBackupCapability

`func (o *V1Config1) HasBackupCapability() bool`

HasBackupCapability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


