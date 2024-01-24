# V1WhoAmIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Accounts** | Pointer to [**[]V1Account1**](V1Account1.md) |  | [optional] 

## Methods

### NewV1WhoAmIResponse

`func NewV1WhoAmIResponse() *V1WhoAmIResponse`

NewV1WhoAmIResponse instantiates a new V1WhoAmIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1WhoAmIResponseWithDefaults

`func NewV1WhoAmIResponseWithDefaults() *V1WhoAmIResponse`

NewV1WhoAmIResponseWithDefaults instantiates a new V1WhoAmIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V1WhoAmIResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1WhoAmIResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1WhoAmIResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1WhoAmIResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *V1WhoAmIResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *V1WhoAmIResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *V1WhoAmIResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *V1WhoAmIResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAccounts

`func (o *V1WhoAmIResponse) GetAccounts() []V1Account1`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V1WhoAmIResponse) GetAccountsOk() (*[]V1Account1, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V1WhoAmIResponse) SetAccounts(v []V1Account1)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V1WhoAmIResponse) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


