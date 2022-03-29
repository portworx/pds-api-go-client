# ControllersCreateBackupCredentialsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**ControllersCredentials**](ControllersCredentials.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the backup credentials. Must be unique for the given tenant. | [optional] 

## Methods

### NewControllersCreateBackupCredentialsRequest

`func NewControllersCreateBackupCredentialsRequest() *ControllersCreateBackupCredentialsRequest`

NewControllersCreateBackupCredentialsRequest instantiates a new ControllersCreateBackupCredentialsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersCreateBackupCredentialsRequestWithDefaults

`func NewControllersCreateBackupCredentialsRequestWithDefaults() *ControllersCreateBackupCredentialsRequest`

NewControllersCreateBackupCredentialsRequestWithDefaults instantiates a new ControllersCreateBackupCredentialsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ControllersCreateBackupCredentialsRequest) GetCredentials() ControllersCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ControllersCreateBackupCredentialsRequest) GetCredentialsOk() (*ControllersCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ControllersCreateBackupCredentialsRequest) SetCredentials(v ControllersCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ControllersCreateBackupCredentialsRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *ControllersCreateBackupCredentialsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControllersCreateBackupCredentialsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControllersCreateBackupCredentialsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControllersCreateBackupCredentialsRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


