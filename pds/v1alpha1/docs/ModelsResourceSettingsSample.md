# ModelsResourceSettingsSample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuLimit** | Pointer to **string** | Maximal amount of CPU cores the deployment will have access to inside k8s. | [optional] 
**CpuRequest** | Pointer to **string** | Minimal amount of CPU cores the deployment will get reserved inside k8s. | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**DataServiceId** | Pointer to **string** |  | [optional] 
**MemoryLimit** | Pointer to **string** | Maximal amount of RAM the deployment will have access to inside k8s. | [optional] 
**MemoryRequest** | Pointer to **string** | Minimal amount of RAM the deployment will get reserved inside k8s. | [optional] 
**Name** | Pointer to **string** | Name of the template. Must be unique for a specific data service within the tenant scope. | [optional] 
**StorageRequest** | Pointer to **string** | Amount of disk space the deployment will get reserved inside k8s. | [optional] 
**Updated** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelsResourceSettingsSample

`func NewModelsResourceSettingsSample() *ModelsResourceSettingsSample`

NewModelsResourceSettingsSample instantiates a new ModelsResourceSettingsSample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsResourceSettingsSampleWithDefaults

`func NewModelsResourceSettingsSampleWithDefaults() *ModelsResourceSettingsSample`

NewModelsResourceSettingsSampleWithDefaults instantiates a new ModelsResourceSettingsSample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuLimit

`func (o *ModelsResourceSettingsSample) GetCpuLimit() string`

GetCpuLimit returns the CpuLimit field if non-nil, zero value otherwise.

### GetCpuLimitOk

`func (o *ModelsResourceSettingsSample) GetCpuLimitOk() (*string, bool)`

GetCpuLimitOk returns a tuple with the CpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimit

`func (o *ModelsResourceSettingsSample) SetCpuLimit(v string)`

SetCpuLimit sets CpuLimit field to given value.

### HasCpuLimit

`func (o *ModelsResourceSettingsSample) HasCpuLimit() bool`

HasCpuLimit returns a boolean if a field has been set.

### GetCpuRequest

`func (o *ModelsResourceSettingsSample) GetCpuRequest() string`

GetCpuRequest returns the CpuRequest field if non-nil, zero value otherwise.

### GetCpuRequestOk

`func (o *ModelsResourceSettingsSample) GetCpuRequestOk() (*string, bool)`

GetCpuRequestOk returns a tuple with the CpuRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuRequest

`func (o *ModelsResourceSettingsSample) SetCpuRequest(v string)`

SetCpuRequest sets CpuRequest field to given value.

### HasCpuRequest

`func (o *ModelsResourceSettingsSample) HasCpuRequest() bool`

HasCpuRequest returns a boolean if a field has been set.

### GetCreated

`func (o *ModelsResourceSettingsSample) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelsResourceSettingsSample) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelsResourceSettingsSample) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelsResourceSettingsSample) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDataServiceId

`func (o *ModelsResourceSettingsSample) GetDataServiceId() string`

GetDataServiceId returns the DataServiceId field if non-nil, zero value otherwise.

### GetDataServiceIdOk

`func (o *ModelsResourceSettingsSample) GetDataServiceIdOk() (*string, bool)`

GetDataServiceIdOk returns a tuple with the DataServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataServiceId

`func (o *ModelsResourceSettingsSample) SetDataServiceId(v string)`

SetDataServiceId sets DataServiceId field to given value.

### HasDataServiceId

`func (o *ModelsResourceSettingsSample) HasDataServiceId() bool`

HasDataServiceId returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *ModelsResourceSettingsSample) GetMemoryLimit() string`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *ModelsResourceSettingsSample) GetMemoryLimitOk() (*string, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *ModelsResourceSettingsSample) SetMemoryLimit(v string)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *ModelsResourceSettingsSample) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetMemoryRequest

`func (o *ModelsResourceSettingsSample) GetMemoryRequest() string`

GetMemoryRequest returns the MemoryRequest field if non-nil, zero value otherwise.

### GetMemoryRequestOk

`func (o *ModelsResourceSettingsSample) GetMemoryRequestOk() (*string, bool)`

GetMemoryRequestOk returns a tuple with the MemoryRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequest

`func (o *ModelsResourceSettingsSample) SetMemoryRequest(v string)`

SetMemoryRequest sets MemoryRequest field to given value.

### HasMemoryRequest

`func (o *ModelsResourceSettingsSample) HasMemoryRequest() bool`

HasMemoryRequest returns a boolean if a field has been set.

### GetName

`func (o *ModelsResourceSettingsSample) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsResourceSettingsSample) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsResourceSettingsSample) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsResourceSettingsSample) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStorageRequest

`func (o *ModelsResourceSettingsSample) GetStorageRequest() string`

GetStorageRequest returns the StorageRequest field if non-nil, zero value otherwise.

### GetStorageRequestOk

`func (o *ModelsResourceSettingsSample) GetStorageRequestOk() (*string, bool)`

GetStorageRequestOk returns a tuple with the StorageRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRequest

`func (o *ModelsResourceSettingsSample) SetStorageRequest(v string)`

SetStorageRequest sets StorageRequest field to given value.

### HasStorageRequest

`func (o *ModelsResourceSettingsSample) HasStorageRequest() bool`

HasStorageRequest returns a boolean if a field has been set.

### GetUpdated

`func (o *ModelsResourceSettingsSample) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ModelsResourceSettingsSample) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ModelsResourceSettingsSample) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ModelsResourceSettingsSample) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetVersion

`func (o *ModelsResourceSettingsSample) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelsResourceSettingsSample) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelsResourceSettingsSample) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelsResourceSettingsSample) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


