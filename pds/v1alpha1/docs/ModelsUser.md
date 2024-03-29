# ModelsUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is autogenerated on creation | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** | ExternalID identifies a User in context of its Issuer. Therefore, it must be unique in the context of the Issuer. But not necessarily unique across different issuers. In case of PX Central the ExternalID is the &#39;sub&#39; claim from JWT token | [optional] 
**Id** | Pointer to **string** | ID is auto generated on creation | [optional] 
**Issuer** | Pointer to **string** | Issuer identifies the identity provider where the User is originating from. It corresponds to the &#39;iss&#39; claim in the JWT token. Example: \&quot;https://central.portworx.com\&quot; | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is autogenerated on update | [optional] 

## Methods

### NewModelsUser

`func NewModelsUser() *ModelsUser`

NewModelsUser instantiates a new ModelsUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsUserWithDefaults

`func NewModelsUserWithDefaults() *ModelsUser`

NewModelsUserWithDefaults instantiates a new ModelsUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ModelsUser) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsUser) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsUser) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *ModelsUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ModelsUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ModelsUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ModelsUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalId

`func (o *ModelsUser) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ModelsUser) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ModelsUser) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ModelsUser) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *ModelsUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuer

`func (o *ModelsUser) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ModelsUser) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ModelsUser) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ModelsUser) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsUser) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsUser) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsUser) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


