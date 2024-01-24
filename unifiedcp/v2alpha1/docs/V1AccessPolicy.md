# V1AccessPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalScope** | Pointer to **[]string** |  | [optional] 
**Account** | Pointer to **[]string** |  | [optional] 
**Tenant** | Pointer to [**[]V1RoleBinding**](V1RoleBinding.md) |  | [optional] 
**Project** | Pointer to [**[]V1RoleBinding**](V1RoleBinding.md) |  | [optional] 
**Namespace** | Pointer to [**[]V1RoleBinding**](V1RoleBinding.md) |  | [optional] 

## Methods

### NewV1AccessPolicy

`func NewV1AccessPolicy() *V1AccessPolicy`

NewV1AccessPolicy instantiates a new V1AccessPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AccessPolicyWithDefaults

`func NewV1AccessPolicyWithDefaults() *V1AccessPolicy`

NewV1AccessPolicyWithDefaults instantiates a new V1AccessPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalScope

`func (o *V1AccessPolicy) GetGlobalScope() []string`

GetGlobalScope returns the GlobalScope field if non-nil, zero value otherwise.

### GetGlobalScopeOk

`func (o *V1AccessPolicy) GetGlobalScopeOk() (*[]string, bool)`

GetGlobalScopeOk returns a tuple with the GlobalScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalScope

`func (o *V1AccessPolicy) SetGlobalScope(v []string)`

SetGlobalScope sets GlobalScope field to given value.

### HasGlobalScope

`func (o *V1AccessPolicy) HasGlobalScope() bool`

HasGlobalScope returns a boolean if a field has been set.

### GetAccount

`func (o *V1AccessPolicy) GetAccount() []string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V1AccessPolicy) GetAccountOk() (*[]string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V1AccessPolicy) SetAccount(v []string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V1AccessPolicy) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetTenant

`func (o *V1AccessPolicy) GetTenant() []V1RoleBinding`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *V1AccessPolicy) GetTenantOk() (*[]V1RoleBinding, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *V1AccessPolicy) SetTenant(v []V1RoleBinding)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *V1AccessPolicy) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetProject

`func (o *V1AccessPolicy) GetProject() []V1RoleBinding`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *V1AccessPolicy) GetProjectOk() (*[]V1RoleBinding, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *V1AccessPolicy) SetProject(v []V1RoleBinding)`

SetProject sets Project field to given value.

### HasProject

`func (o *V1AccessPolicy) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetNamespace

`func (o *V1AccessPolicy) GetNamespace() []V1RoleBinding`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1AccessPolicy) GetNamespaceOk() (*[]V1RoleBinding, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1AccessPolicy) SetNamespace(v []V1RoleBinding)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V1AccessPolicy) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


