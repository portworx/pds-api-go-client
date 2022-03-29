# ModelsS3Credentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | Access key for the AWS IAM user. | [optional] 
**Endpoint** | Pointer to **string** | Endpoint of S3 storage service. | [optional] 
**SecretKey** | Pointer to **string** | Secret key for the AWS IAM user. | [optional] 

## Methods

### NewModelsS3Credentials

`func NewModelsS3Credentials() *ModelsS3Credentials`

NewModelsS3Credentials instantiates a new ModelsS3Credentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsS3CredentialsWithDefaults

`func NewModelsS3CredentialsWithDefaults() *ModelsS3Credentials`

NewModelsS3CredentialsWithDefaults instantiates a new ModelsS3Credentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *ModelsS3Credentials) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ModelsS3Credentials) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ModelsS3Credentials) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *ModelsS3Credentials) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetEndpoint

`func (o *ModelsS3Credentials) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ModelsS3Credentials) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ModelsS3Credentials) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ModelsS3Credentials) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetSecretKey

`func (o *ModelsS3Credentials) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ModelsS3Credentials) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ModelsS3Credentials) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *ModelsS3Credentials) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


