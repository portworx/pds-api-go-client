# PolicyRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extends** | Pointer to **[]string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPolicyRole

`func NewPolicyRole() *PolicyRole`

NewPolicyRole instantiates a new PolicyRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRoleWithDefaults

`func NewPolicyRoleWithDefaults() *PolicyRole`

NewPolicyRoleWithDefaults instantiates a new PolicyRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtends

`func (o *PolicyRole) GetExtends() []string`

GetExtends returns the Extends field if non-nil, zero value otherwise.

### GetExtendsOk

`func (o *PolicyRole) GetExtendsOk() (*[]string, bool)`

GetExtendsOk returns a tuple with the Extends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtends

`func (o *PolicyRole) SetExtends(v []string)`

SetExtends sets Extends field to given value.

### HasExtends

`func (o *PolicyRole) HasExtends() bool`

HasExtends returns a boolean if a field has been set.

### GetLevel

`func (o *PolicyRole) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *PolicyRole) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *PolicyRole) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *PolicyRole) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetName

`func (o *PolicyRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *PolicyRole) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PolicyRole) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PolicyRole) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PolicyRole) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


