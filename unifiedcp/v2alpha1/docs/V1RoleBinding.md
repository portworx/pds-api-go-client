# V1RoleBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleName** | Pointer to **string** | Role name represents the role for a tenant/project/namespace. | [optional] 
**ResourceIds** | Pointer to **[]string** | Resource IDs represent the IDs bounded for the given role. | [optional] 

## Methods

### NewV1RoleBinding

`func NewV1RoleBinding() *V1RoleBinding`

NewV1RoleBinding instantiates a new V1RoleBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RoleBindingWithDefaults

`func NewV1RoleBindingWithDefaults() *V1RoleBinding`

NewV1RoleBindingWithDefaults instantiates a new V1RoleBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleName

`func (o *V1RoleBinding) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *V1RoleBinding) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *V1RoleBinding) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *V1RoleBinding) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetResourceIds

`func (o *V1RoleBinding) GetResourceIds() []string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *V1RoleBinding) GetResourceIdsOk() (*[]string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *V1RoleBinding) SetResourceIds(v []string)`

SetResourceIds sets ResourceIds field to given value.

### HasResourceIds

`func (o *V1RoleBinding) HasResourceIds() bool`

HasResourceIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


