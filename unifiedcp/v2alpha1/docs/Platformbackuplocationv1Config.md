# Platformbackuplocationv1Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to [**V1Provider**](V1Provider.md) |  | [optional] 
**CloudCredentialId** | Pointer to **string** |  | [optional] 
**AzureStorage** | Pointer to [**V1AzureBlobStorage**](V1AzureBlobStorage.md) |  | [optional] 
**GoogleStorage** | Pointer to [**V1GoogleCloudStorage**](V1GoogleCloudStorage.md) |  | [optional] 
**S3Storage** | Pointer to [**V1S3ObjectStorage**](V1S3ObjectStorage.md) |  | [optional] 

## Methods

### NewPlatformbackuplocationv1Config

`func NewPlatformbackuplocationv1Config() *Platformbackuplocationv1Config`

NewPlatformbackuplocationv1Config instantiates a new Platformbackuplocationv1Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformbackuplocationv1ConfigWithDefaults

`func NewPlatformbackuplocationv1ConfigWithDefaults() *Platformbackuplocationv1Config`

NewPlatformbackuplocationv1ConfigWithDefaults instantiates a new Platformbackuplocationv1Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *Platformbackuplocationv1Config) GetProvider() V1Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Platformbackuplocationv1Config) GetProviderOk() (*V1Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Platformbackuplocationv1Config) SetProvider(v V1Provider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Platformbackuplocationv1Config) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetCloudCredentialId

`func (o *Platformbackuplocationv1Config) GetCloudCredentialId() string`

GetCloudCredentialId returns the CloudCredentialId field if non-nil, zero value otherwise.

### GetCloudCredentialIdOk

`func (o *Platformbackuplocationv1Config) GetCloudCredentialIdOk() (*string, bool)`

GetCloudCredentialIdOk returns a tuple with the CloudCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialId

`func (o *Platformbackuplocationv1Config) SetCloudCredentialId(v string)`

SetCloudCredentialId sets CloudCredentialId field to given value.

### HasCloudCredentialId

`func (o *Platformbackuplocationv1Config) HasCloudCredentialId() bool`

HasCloudCredentialId returns a boolean if a field has been set.

### GetAzureStorage

`func (o *Platformbackuplocationv1Config) GetAzureStorage() V1AzureBlobStorage`

GetAzureStorage returns the AzureStorage field if non-nil, zero value otherwise.

### GetAzureStorageOk

`func (o *Platformbackuplocationv1Config) GetAzureStorageOk() (*V1AzureBlobStorage, bool)`

GetAzureStorageOk returns a tuple with the AzureStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureStorage

`func (o *Platformbackuplocationv1Config) SetAzureStorage(v V1AzureBlobStorage)`

SetAzureStorage sets AzureStorage field to given value.

### HasAzureStorage

`func (o *Platformbackuplocationv1Config) HasAzureStorage() bool`

HasAzureStorage returns a boolean if a field has been set.

### GetGoogleStorage

`func (o *Platformbackuplocationv1Config) GetGoogleStorage() V1GoogleCloudStorage`

GetGoogleStorage returns the GoogleStorage field if non-nil, zero value otherwise.

### GetGoogleStorageOk

`func (o *Platformbackuplocationv1Config) GetGoogleStorageOk() (*V1GoogleCloudStorage, bool)`

GetGoogleStorageOk returns a tuple with the GoogleStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleStorage

`func (o *Platformbackuplocationv1Config) SetGoogleStorage(v V1GoogleCloudStorage)`

SetGoogleStorage sets GoogleStorage field to given value.

### HasGoogleStorage

`func (o *Platformbackuplocationv1Config) HasGoogleStorage() bool`

HasGoogleStorage returns a boolean if a field has been set.

### GetS3Storage

`func (o *Platformbackuplocationv1Config) GetS3Storage() V1S3ObjectStorage`

GetS3Storage returns the S3Storage field if non-nil, zero value otherwise.

### GetS3StorageOk

`func (o *Platformbackuplocationv1Config) GetS3StorageOk() (*V1S3ObjectStorage, bool)`

GetS3StorageOk returns a tuple with the S3Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Storage

`func (o *Platformbackuplocationv1Config) SetS3Storage(v V1S3ObjectStorage)`

SetS3Storage sets S3Storage field to given value.

### HasS3Storage

`func (o *Platformbackuplocationv1Config) HasS3Storage() bool`

HasS3Storage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


