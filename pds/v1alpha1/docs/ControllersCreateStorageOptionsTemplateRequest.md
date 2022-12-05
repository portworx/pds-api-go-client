# ControllersCreateStorageOptionsTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fg** | Pointer to **bool** |  | [optional] 
**Fs** | Pointer to **string** | Filesystem to be laid out. | [optional] 
**Name** | Pointer to **string** | See models.StorageOptionsTemplate for more information. | [optional] 
**Provisioner** | Pointer to **string** | Portworx volume provisioner. Valid values are: \&quot;pxd.portworx.com\&quot; for PX CSI, \&quot;kubernetes.io/portworx-volume\&quot; for PX in-tree or \&quot;auto\&quot; for auto-detect | [optional] 
**Repl** | Pointer to **int32** | Replication factor for the volume. | [optional] 
**Secure** | Pointer to **bool** |  | [optional] 

## Methods

### NewControllersCreateStorageOptionsTemplateRequest

`func NewControllersCreateStorageOptionsTemplateRequest() *ControllersCreateStorageOptionsTemplateRequest`

NewControllersCreateStorageOptionsTemplateRequest instantiates a new ControllersCreateStorageOptionsTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersCreateStorageOptionsTemplateRequestWithDefaults

`func NewControllersCreateStorageOptionsTemplateRequestWithDefaults() *ControllersCreateStorageOptionsTemplateRequest`

NewControllersCreateStorageOptionsTemplateRequestWithDefaults instantiates a new ControllersCreateStorageOptionsTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFg

`func (o *ControllersCreateStorageOptionsTemplateRequest) GetFg() bool`

GetFg returns the Fg field if non-nil, zero value otherwise.

### GetFgOk

`func (o *ControllersCreateStorageOptionsTemplateRequest) GetFgOk() (*bool, bool)`

GetFgOk returns a tuple with the Fg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFg

`func (o *ControllersCreateStorageOptionsTemplateRequest) SetFg(v bool)`

SetFg sets Fg field to given value.

### HasFg

`func (o *ControllersCreateStorageOptionsTemplateRequest) HasFg() bool`

HasFg returns a boolean if a field has been set.

### GetFs

`func (o *ControllersCreateStorageOptionsTemplateRequest) GetFs() string`

GetFs returns the Fs field if non-nil, zero value otherwise.

### GetFsOk

`func (o *ControllersCreateStorageOptionsTemplateRequest) GetFsOk() (*string, bool)`

GetFsOk returns a tuple with the Fs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFs

`func (o *ControllersCreateStorageOptionsTemplateRequest) SetFs(v string)`

SetFs sets Fs field to given value.

### HasFs

`func (o *ControllersCreateStorageOptionsTemplateRequest) HasFs() bool`

HasFs returns a boolean if a field has been set.

### GetName

`func (o *ControllersCreateStorageOptionsTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControllersCreateStorageOptionsTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControllersCreateStorageOptionsTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControllersCreateStorageOptionsTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvisioner

`func (o *ControllersCreateStorageOptionsTemplateRequest) GetProvisioner() string`

GetProvisioner returns the Provisioner field if non-nil, zero value otherwise.

### GetProvisionerOk

`func (o *ControllersCreateStorageOptionsTemplateRequest) GetProvisionerOk() (*string, bool)`

GetProvisionerOk returns a tuple with the Provisioner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioner

`func (o *ControllersCreateStorageOptionsTemplateRequest) SetProvisioner(v string)`

SetProvisioner sets Provisioner field to given value.

### HasProvisioner

`func (o *ControllersCreateStorageOptionsTemplateRequest) HasProvisioner() bool`

HasProvisioner returns a boolean if a field has been set.

### GetRepl

`func (o *ControllersCreateStorageOptionsTemplateRequest) GetRepl() int32`

GetRepl returns the Repl field if non-nil, zero value otherwise.

### GetReplOk

`func (o *ControllersCreateStorageOptionsTemplateRequest) GetReplOk() (*int32, bool)`

GetReplOk returns a tuple with the Repl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepl

`func (o *ControllersCreateStorageOptionsTemplateRequest) SetRepl(v int32)`

SetRepl sets Repl field to given value.

### HasRepl

`func (o *ControllersCreateStorageOptionsTemplateRequest) HasRepl() bool`

HasRepl returns a boolean if a field has been set.

### GetSecure

`func (o *ControllersCreateStorageOptionsTemplateRequest) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *ControllersCreateStorageOptionsTemplateRequest) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *ControllersCreateStorageOptionsTemplateRequest) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *ControllersCreateStorageOptionsTemplateRequest) HasSecure() bool`

HasSecure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


