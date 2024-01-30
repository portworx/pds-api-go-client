# V1DeploymentMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the deployment. | [optional] 
**CustomResourceName** | Pointer to **string** | Custom Resource Name is the kubernetes resource name for the deployment meta data. | [optional] [readonly] 
**DeploymentTargetName** | Pointer to **string** | Deployment Target Name associated with the backup configuration. | [optional] 
**NamespaceName** | Pointer to **string** | Namespace name to which the backup configuration is associated. | [optional] 
**TlsEnabled** | Pointer to **bool** | Flag to check whether Transport Layer Security support is enabled or not. | [optional] 

## Methods

### NewV1DeploymentMetaData

`func NewV1DeploymentMetaData() *V1DeploymentMetaData`

NewV1DeploymentMetaData instantiates a new V1DeploymentMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeploymentMetaDataWithDefaults

`func NewV1DeploymentMetaDataWithDefaults() *V1DeploymentMetaData`

NewV1DeploymentMetaDataWithDefaults instantiates a new V1DeploymentMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1DeploymentMetaData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1DeploymentMetaData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1DeploymentMetaData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1DeploymentMetaData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCustomResourceName

`func (o *V1DeploymentMetaData) GetCustomResourceName() string`

GetCustomResourceName returns the CustomResourceName field if non-nil, zero value otherwise.

### GetCustomResourceNameOk

`func (o *V1DeploymentMetaData) GetCustomResourceNameOk() (*string, bool)`

GetCustomResourceNameOk returns a tuple with the CustomResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomResourceName

`func (o *V1DeploymentMetaData) SetCustomResourceName(v string)`

SetCustomResourceName sets CustomResourceName field to given value.

### HasCustomResourceName

`func (o *V1DeploymentMetaData) HasCustomResourceName() bool`

HasCustomResourceName returns a boolean if a field has been set.

### GetDeploymentTargetName

`func (o *V1DeploymentMetaData) GetDeploymentTargetName() string`

GetDeploymentTargetName returns the DeploymentTargetName field if non-nil, zero value otherwise.

### GetDeploymentTargetNameOk

`func (o *V1DeploymentMetaData) GetDeploymentTargetNameOk() (*string, bool)`

GetDeploymentTargetNameOk returns a tuple with the DeploymentTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTargetName

`func (o *V1DeploymentMetaData) SetDeploymentTargetName(v string)`

SetDeploymentTargetName sets DeploymentTargetName field to given value.

### HasDeploymentTargetName

`func (o *V1DeploymentMetaData) HasDeploymentTargetName() bool`

HasDeploymentTargetName returns a boolean if a field has been set.

### GetNamespaceName

`func (o *V1DeploymentMetaData) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *V1DeploymentMetaData) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *V1DeploymentMetaData) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.

### HasNamespaceName

`func (o *V1DeploymentMetaData) HasNamespaceName() bool`

HasNamespaceName returns a boolean if a field has been set.

### GetTlsEnabled

`func (o *V1DeploymentMetaData) GetTlsEnabled() bool`

GetTlsEnabled returns the TlsEnabled field if non-nil, zero value otherwise.

### GetTlsEnabledOk

`func (o *V1DeploymentMetaData) GetTlsEnabledOk() (*bool, bool)`

GetTlsEnabledOk returns a tuple with the TlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnabled

`func (o *V1DeploymentMetaData) SetTlsEnabled(v bool)`

SetTlsEnabled sets TlsEnabled field to given value.

### HasTlsEnabled

`func (o *V1DeploymentMetaData) HasTlsEnabled() bool`

HasTlsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


