# ControllersPartialS3Credentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | Access key for the AWS IAM user. | [optional] 
**Endpoint** | Pointer to **string** | Endpoint of S3 storage service. | [optional] 

## Methods

### NewControllersPartialS3Credentials

`func NewControllersPartialS3Credentials() *ControllersPartialS3Credentials`

NewControllersPartialS3Credentials instantiates a new ControllersPartialS3Credentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersPartialS3CredentialsWithDefaults

`func NewControllersPartialS3CredentialsWithDefaults() *ControllersPartialS3Credentials`

NewControllersPartialS3CredentialsWithDefaults instantiates a new ControllersPartialS3Credentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *ControllersPartialS3Credentials) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ControllersPartialS3Credentials) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ControllersPartialS3Credentials) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *ControllersPartialS3Credentials) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetEndpoint

`func (o *ControllersPartialS3Credentials) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ControllersPartialS3Credentials) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ControllersPartialS3Credentials) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ControllersPartialS3Credentials) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


