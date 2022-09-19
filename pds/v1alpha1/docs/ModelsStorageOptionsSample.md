# ModelsStorageOptionsSample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fg** | Pointer to **bool** | This option enforces volume group policy. If a volume belonging to a group cannot find nodes for its replication sets which don’t have other volumes of same group, the volume creation will fail. | [optional] 
**Fs** | Pointer to **string** | Filesystem to be laid out. | [optional] 
**Name** | Pointer to **string** | Name of the template. Must be unique within the tenant scope. | [optional] 
**Repl** | Pointer to **int32** | Replication factor for the volume. | [optional] 
**Secure** | Pointer to **bool** | Flag to create an encrypted volume. Currently, not supported (should be set to &#x60;false&#x60;). | [optional] 

## Methods

### NewModelsStorageOptionsSample

`func NewModelsStorageOptionsSample() *ModelsStorageOptionsSample`

NewModelsStorageOptionsSample instantiates a new ModelsStorageOptionsSample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsStorageOptionsSampleWithDefaults

`func NewModelsStorageOptionsSampleWithDefaults() *ModelsStorageOptionsSample`

NewModelsStorageOptionsSampleWithDefaults instantiates a new ModelsStorageOptionsSample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFg

`func (o *ModelsStorageOptionsSample) GetFg() bool`

GetFg returns the Fg field if non-nil, zero value otherwise.

### GetFgOk

`func (o *ModelsStorageOptionsSample) GetFgOk() (*bool, bool)`

GetFgOk returns a tuple with the Fg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFg

`func (o *ModelsStorageOptionsSample) SetFg(v bool)`

SetFg sets Fg field to given value.

### HasFg

`func (o *ModelsStorageOptionsSample) HasFg() bool`

HasFg returns a boolean if a field has been set.

### GetFs

`func (o *ModelsStorageOptionsSample) GetFs() string`

GetFs returns the Fs field if non-nil, zero value otherwise.

### GetFsOk

`func (o *ModelsStorageOptionsSample) GetFsOk() (*string, bool)`

GetFsOk returns a tuple with the Fs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFs

`func (o *ModelsStorageOptionsSample) SetFs(v string)`

SetFs sets Fs field to given value.

### HasFs

`func (o *ModelsStorageOptionsSample) HasFs() bool`

HasFs returns a boolean if a field has been set.

### GetName

`func (o *ModelsStorageOptionsSample) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsStorageOptionsSample) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsStorageOptionsSample) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsStorageOptionsSample) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepl

`func (o *ModelsStorageOptionsSample) GetRepl() int32`

GetRepl returns the Repl field if non-nil, zero value otherwise.

### GetReplOk

`func (o *ModelsStorageOptionsSample) GetReplOk() (*int32, bool)`

GetReplOk returns a tuple with the Repl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepl

`func (o *ModelsStorageOptionsSample) SetRepl(v int32)`

SetRepl sets Repl field to given value.

### HasRepl

`func (o *ModelsStorageOptionsSample) HasRepl() bool`

HasRepl returns a boolean if a field has been set.

### GetSecure

`func (o *ModelsStorageOptionsSample) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *ModelsStorageOptionsSample) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *ModelsStorageOptionsSample) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *ModelsStorageOptionsSample) HasSecure() bool`

HasSecure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


