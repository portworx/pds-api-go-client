# V1Config5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to [**V1Provider**](V1Provider.md) |  | [optional] 
**AzureCredentials** | Pointer to [**V1AzureCredentials**](V1AzureCredentials.md) |  | [optional] 
**GoogleCredentials** | Pointer to [**V1GoogleCredentials**](V1GoogleCredentials.md) |  | [optional] 
**S3Credentials** | Pointer to [**V1S3Credentials**](V1S3Credentials.md) |  | [optional] 
**UnstructuredCredentials** | Pointer to [**V1UnstructuredCredentials**](V1UnstructuredCredentials.md) |  | [optional] 

## Methods

### NewV1Config5

`func NewV1Config5() *V1Config5`

NewV1Config5 instantiates a new V1Config5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1Config5WithDefaults

`func NewV1Config5WithDefaults() *V1Config5`

NewV1Config5WithDefaults instantiates a new V1Config5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *V1Config5) GetProvider() V1Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *V1Config5) GetProviderOk() (*V1Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *V1Config5) SetProvider(v V1Provider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *V1Config5) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetAzureCredentials

`func (o *V1Config5) GetAzureCredentials() V1AzureCredentials`

GetAzureCredentials returns the AzureCredentials field if non-nil, zero value otherwise.

### GetAzureCredentialsOk

`func (o *V1Config5) GetAzureCredentialsOk() (*V1AzureCredentials, bool)`

GetAzureCredentialsOk returns a tuple with the AzureCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureCredentials

`func (o *V1Config5) SetAzureCredentials(v V1AzureCredentials)`

SetAzureCredentials sets AzureCredentials field to given value.

### HasAzureCredentials

`func (o *V1Config5) HasAzureCredentials() bool`

HasAzureCredentials returns a boolean if a field has been set.

### GetGoogleCredentials

`func (o *V1Config5) GetGoogleCredentials() V1GoogleCredentials`

GetGoogleCredentials returns the GoogleCredentials field if non-nil, zero value otherwise.

### GetGoogleCredentialsOk

`func (o *V1Config5) GetGoogleCredentialsOk() (*V1GoogleCredentials, bool)`

GetGoogleCredentialsOk returns a tuple with the GoogleCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleCredentials

`func (o *V1Config5) SetGoogleCredentials(v V1GoogleCredentials)`

SetGoogleCredentials sets GoogleCredentials field to given value.

### HasGoogleCredentials

`func (o *V1Config5) HasGoogleCredentials() bool`

HasGoogleCredentials returns a boolean if a field has been set.

### GetS3Credentials

`func (o *V1Config5) GetS3Credentials() V1S3Credentials`

GetS3Credentials returns the S3Credentials field if non-nil, zero value otherwise.

### GetS3CredentialsOk

`func (o *V1Config5) GetS3CredentialsOk() (*V1S3Credentials, bool)`

GetS3CredentialsOk returns a tuple with the S3Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Credentials

`func (o *V1Config5) SetS3Credentials(v V1S3Credentials)`

SetS3Credentials sets S3Credentials field to given value.

### HasS3Credentials

`func (o *V1Config5) HasS3Credentials() bool`

HasS3Credentials returns a boolean if a field has been set.

### GetUnstructuredCredentials

`func (o *V1Config5) GetUnstructuredCredentials() V1UnstructuredCredentials`

GetUnstructuredCredentials returns the UnstructuredCredentials field if non-nil, zero value otherwise.

### GetUnstructuredCredentialsOk

`func (o *V1Config5) GetUnstructuredCredentialsOk() (*V1UnstructuredCredentials, bool)`

GetUnstructuredCredentialsOk returns a tuple with the UnstructuredCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnstructuredCredentials

`func (o *V1Config5) SetUnstructuredCredentials(v V1UnstructuredCredentials)`

SetUnstructuredCredentials sets UnstructuredCredentials field to given value.

### HasUnstructuredCredentials

`func (o *V1Config5) HasUnstructuredCredentials() bool`

HasUnstructuredCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


