# ControllersPartialCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Azure** | Pointer to [**ControllersPartialAzureCredentials**](ControllersPartialAzureCredentials.md) |  | [optional] 
**S3** | Pointer to [**ControllersPartialS3Credentials**](ControllersPartialS3Credentials.md) |  | [optional] 
**S3Compatible** | Pointer to [**ControllersPartialS3CompatibleCredentials**](ControllersPartialS3CompatibleCredentials.md) |  | [optional] 

## Methods

### NewControllersPartialCredentials

`func NewControllersPartialCredentials() *ControllersPartialCredentials`

NewControllersPartialCredentials instantiates a new ControllersPartialCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersPartialCredentialsWithDefaults

`func NewControllersPartialCredentialsWithDefaults() *ControllersPartialCredentials`

NewControllersPartialCredentialsWithDefaults instantiates a new ControllersPartialCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzure

`func (o *ControllersPartialCredentials) GetAzure() ControllersPartialAzureCredentials`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *ControllersPartialCredentials) GetAzureOk() (*ControllersPartialAzureCredentials, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *ControllersPartialCredentials) SetAzure(v ControllersPartialAzureCredentials)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *ControllersPartialCredentials) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetS3

`func (o *ControllersPartialCredentials) GetS3() ControllersPartialS3Credentials`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *ControllersPartialCredentials) GetS3Ok() (*ControllersPartialS3Credentials, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *ControllersPartialCredentials) SetS3(v ControllersPartialS3Credentials)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *ControllersPartialCredentials) HasS3() bool`

HasS3 returns a boolean if a field has been set.

### GetS3Compatible

`func (o *ControllersPartialCredentials) GetS3Compatible() ControllersPartialS3CompatibleCredentials`

GetS3Compatible returns the S3Compatible field if non-nil, zero value otherwise.

### GetS3CompatibleOk

`func (o *ControllersPartialCredentials) GetS3CompatibleOk() (*ControllersPartialS3CompatibleCredentials, bool)`

GetS3CompatibleOk returns a tuple with the S3Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Compatible

`func (o *ControllersPartialCredentials) SetS3Compatible(v ControllersPartialS3CompatibleCredentials)`

SetS3Compatible sets S3Compatible field to given value.

### HasS3Compatible

`func (o *ControllersPartialCredentials) HasS3Compatible() bool`

HasS3Compatible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


