# AuthClaims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aud** | Pointer to **string** | Audience is the intended recipient of the token. | [optional] 
**Email** | Pointer to **string** | Email of this account. | [optional] 
**Groups** | Pointer to **[]string** | (optional) Groups in which this account is part of. | [optional] 
**Iss** | Pointer to **string** | Issuer is the token issuer. For self-signed token do not prefix with &#x60;https://&#x60;. | [optional] 
**Name** | Pointer to **string** | Name of this account. | [optional] 
**Roles** | Pointer to **[]string** | Roles of this account. | [optional] 
**Sub** | Pointer to **string** | Subject identifier. Unique ID of this account. | [optional] 

## Methods

### NewAuthClaims

`func NewAuthClaims() *AuthClaims`

NewAuthClaims instantiates a new AuthClaims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthClaimsWithDefaults

`func NewAuthClaimsWithDefaults() *AuthClaims`

NewAuthClaimsWithDefaults instantiates a new AuthClaims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAud

`func (o *AuthClaims) GetAud() string`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *AuthClaims) GetAudOk() (*string, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *AuthClaims) SetAud(v string)`

SetAud sets Aud field to given value.

### HasAud

`func (o *AuthClaims) HasAud() bool`

HasAud returns a boolean if a field has been set.

### GetEmail

`func (o *AuthClaims) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthClaims) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthClaims) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuthClaims) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGroups

`func (o *AuthClaims) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AuthClaims) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AuthClaims) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *AuthClaims) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetIss

`func (o *AuthClaims) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *AuthClaims) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *AuthClaims) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *AuthClaims) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetName

`func (o *AuthClaims) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthClaims) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthClaims) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthClaims) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoles

`func (o *AuthClaims) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AuthClaims) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AuthClaims) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AuthClaims) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSub

`func (o *AuthClaims) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *AuthClaims) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *AuthClaims) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *AuthClaims) HasSub() bool`

HasSub returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


