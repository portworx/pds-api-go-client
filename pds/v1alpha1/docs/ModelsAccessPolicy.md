# ModelsAccessPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcccountId** | Pointer to **string** |  | [optional] 
**Account** | Pointer to **[]string** |  | [optional] 
**Global** | Pointer to **[]string** |  | [optional] 
**Namespace** | Pointer to [**[]ModelsBinding**](ModelsBinding.md) |  | [optional] 
**Project** | Pointer to [**[]ModelsBinding**](ModelsBinding.md) |  | [optional] 
**Tenant** | Pointer to [**[]ModelsBinding**](ModelsBinding.md) |  | [optional] 

## Methods

### NewModelsAccessPolicy

`func NewModelsAccessPolicy() *ModelsAccessPolicy`

NewModelsAccessPolicy instantiates a new ModelsAccessPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsAccessPolicyWithDefaults

`func NewModelsAccessPolicyWithDefaults() *ModelsAccessPolicy`

NewModelsAccessPolicyWithDefaults instantiates a new ModelsAccessPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcccountId

`func (o *ModelsAccessPolicy) GetAcccountId() string`

GetAcccountId returns the AcccountId field if non-nil, zero value otherwise.

### GetAcccountIdOk

`func (o *ModelsAccessPolicy) GetAcccountIdOk() (*string, bool)`

GetAcccountIdOk returns a tuple with the AcccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcccountId

`func (o *ModelsAccessPolicy) SetAcccountId(v string)`

SetAcccountId sets AcccountId field to given value.

### HasAcccountId

`func (o *ModelsAccessPolicy) HasAcccountId() bool`

HasAcccountId returns a boolean if a field has been set.

### GetAccount

`func (o *ModelsAccessPolicy) GetAccount() []string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ModelsAccessPolicy) GetAccountOk() (*[]string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ModelsAccessPolicy) SetAccount(v []string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ModelsAccessPolicy) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetGlobal

`func (o *ModelsAccessPolicy) GetGlobal() []string`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *ModelsAccessPolicy) GetGlobalOk() (*[]string, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *ModelsAccessPolicy) SetGlobal(v []string)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *ModelsAccessPolicy) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetNamespace

`func (o *ModelsAccessPolicy) GetNamespace() []ModelsBinding`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ModelsAccessPolicy) GetNamespaceOk() (*[]ModelsBinding, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ModelsAccessPolicy) SetNamespace(v []ModelsBinding)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ModelsAccessPolicy) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetProject

`func (o *ModelsAccessPolicy) GetProject() []ModelsBinding`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ModelsAccessPolicy) GetProjectOk() (*[]ModelsBinding, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ModelsAccessPolicy) SetProject(v []ModelsBinding)`

SetProject sets Project field to given value.

### HasProject

`func (o *ModelsAccessPolicy) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetTenant

`func (o *ModelsAccessPolicy) GetTenant() []ModelsBinding`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ModelsAccessPolicy) GetTenantOk() (*[]ModelsBinding, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ModelsAccessPolicy) SetTenant(v []ModelsBinding)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ModelsAccessPolicy) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


