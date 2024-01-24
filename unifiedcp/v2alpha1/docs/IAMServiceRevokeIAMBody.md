# IAMServiceRevokeIAMBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Iam** | Pointer to [**SpecifiesTheIAMRoleBindingThatNeedsToBeDeletedRevokedForTheUser**](SpecifiesTheIAMRoleBindingThatNeedsToBeDeletedRevokedForTheUser.md) |  | [optional] 

## Methods

### NewIAMServiceRevokeIAMBody

`func NewIAMServiceRevokeIAMBody() *IAMServiceRevokeIAMBody`

NewIAMServiceRevokeIAMBody instantiates a new IAMServiceRevokeIAMBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIAMServiceRevokeIAMBodyWithDefaults

`func NewIAMServiceRevokeIAMBodyWithDefaults() *IAMServiceRevokeIAMBody`

NewIAMServiceRevokeIAMBodyWithDefaults instantiates a new IAMServiceRevokeIAMBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *IAMServiceRevokeIAMBody) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *IAMServiceRevokeIAMBody) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *IAMServiceRevokeIAMBody) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *IAMServiceRevokeIAMBody) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetTenantId

`func (o *IAMServiceRevokeIAMBody) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *IAMServiceRevokeIAMBody) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *IAMServiceRevokeIAMBody) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *IAMServiceRevokeIAMBody) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetProjectId

`func (o *IAMServiceRevokeIAMBody) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *IAMServiceRevokeIAMBody) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *IAMServiceRevokeIAMBody) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *IAMServiceRevokeIAMBody) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetIam

`func (o *IAMServiceRevokeIAMBody) GetIam() SpecifiesTheIAMRoleBindingThatNeedsToBeDeletedRevokedForTheUser`

GetIam returns the Iam field if non-nil, zero value otherwise.

### GetIamOk

`func (o *IAMServiceRevokeIAMBody) GetIamOk() (*SpecifiesTheIAMRoleBindingThatNeedsToBeDeletedRevokedForTheUser, bool)`

GetIamOk returns a tuple with the Iam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIam

`func (o *IAMServiceRevokeIAMBody) SetIam(v SpecifiesTheIAMRoleBindingThatNeedsToBeDeletedRevokedForTheUser)`

SetIam sets Iam field to given value.

### HasIam

`func (o *IAMServiceRevokeIAMBody) HasIam() bool`

HasIam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


