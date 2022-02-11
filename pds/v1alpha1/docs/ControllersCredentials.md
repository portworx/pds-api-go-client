# ControllersCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Azure** | Pointer to [**ModelsAzureCredentials**](ModelsAzureCredentials.md) |  | [optional] 
**S3** | Pointer to [**ModelsS3Credentials**](ModelsS3Credentials.md) |  | [optional] 
**S3Compatible** | Pointer to [**ModelsS3CompatibleCredentials**](ModelsS3CompatibleCredentials.md) |  | [optional] 

## Methods

### NewControllersCredentials

`func NewControllersCredentials() *ControllersCredentials`

NewControllersCredentials instantiates a new ControllersCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersCredentialsWithDefaults

`func NewControllersCredentialsWithDefaults() *ControllersCredentials`

NewControllersCredentialsWithDefaults instantiates a new ControllersCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzure

`func (o *ControllersCredentials) GetAzure() ModelsAzureCredentials`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *ControllersCredentials) GetAzureOk() (*ModelsAzureCredentials, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *ControllersCredentials) SetAzure(v ModelsAzureCredentials)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *ControllersCredentials) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetS3

`func (o *ControllersCredentials) GetS3() ModelsS3Credentials`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *ControllersCredentials) GetS3Ok() (*ModelsS3Credentials, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *ControllersCredentials) SetS3(v ModelsS3Credentials)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *ControllersCredentials) HasS3() bool`

HasS3 returns a boolean if a field has been set.

### GetS3Compatible

`func (o *ControllersCredentials) GetS3Compatible() ModelsS3CompatibleCredentials`

GetS3Compatible returns the S3Compatible field if non-nil, zero value otherwise.

### GetS3CompatibleOk

`func (o *ControllersCredentials) GetS3CompatibleOk() (*ModelsS3CompatibleCredentials, bool)`

GetS3CompatibleOk returns a tuple with the S3Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Compatible

`func (o *ControllersCredentials) SetS3Compatible(v ModelsS3CompatibleCredentials)`

SetS3Compatible sets S3Compatible field to given value.

### HasS3Compatible

`func (o *ControllersCredentials) HasS3Compatible() bool`

HasS3Compatible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


