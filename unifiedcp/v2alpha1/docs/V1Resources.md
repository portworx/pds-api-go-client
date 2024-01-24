# V1Resources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to **[]string** | Clusters represents the target k8s clusters. | [optional] 
**Namespaces** | Pointer to **[]string** | Namespaces created in k8s cluster to provide the logical isolation. | [optional] 
**Credentials** | Pointer to **[]string** | Credentials required to connect to a backup target. | [optional] 
**BackupLocations** | Pointer to **[]string** | Backup locations where backups can be placed. | [optional] 
**Templates** | Pointer to **[]string** | Templates can be used by applications to manage its resources. | [optional] 

## Methods

### NewV1Resources

`func NewV1Resources() *V1Resources`

NewV1Resources instantiates a new V1Resources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ResourcesWithDefaults

`func NewV1ResourcesWithDefaults() *V1Resources`

NewV1ResourcesWithDefaults instantiates a new V1Resources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *V1Resources) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *V1Resources) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *V1Resources) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *V1Resources) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetNamespaces

`func (o *V1Resources) GetNamespaces() []string`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *V1Resources) GetNamespacesOk() (*[]string, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *V1Resources) SetNamespaces(v []string)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *V1Resources) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetCredentials

`func (o *V1Resources) GetCredentials() []string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *V1Resources) GetCredentialsOk() (*[]string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *V1Resources) SetCredentials(v []string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *V1Resources) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetBackupLocations

`func (o *V1Resources) GetBackupLocations() []string`

GetBackupLocations returns the BackupLocations field if non-nil, zero value otherwise.

### GetBackupLocationsOk

`func (o *V1Resources) GetBackupLocationsOk() (*[]string, bool)`

GetBackupLocationsOk returns a tuple with the BackupLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLocations

`func (o *V1Resources) SetBackupLocations(v []string)`

SetBackupLocations sets BackupLocations field to given value.

### HasBackupLocations

`func (o *V1Resources) HasBackupLocations() bool`

HasBackupLocations returns a boolean if a field has been set.

### GetTemplates

`func (o *V1Resources) GetTemplates() []string`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *V1Resources) GetTemplatesOk() (*[]string, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *V1Resources) SetTemplates(v []string)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *V1Resources) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


