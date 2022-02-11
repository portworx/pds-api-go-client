# ControllersOIDCTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claims** | Pointer to [**AuthClaims**](AuthClaims.md) |  | [optional] 
**ExpiresIn** | Pointer to **string** |  | [optional] 
**IdToken** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 

## Methods

### NewControllersOIDCTokenResponse

`func NewControllersOIDCTokenResponse() *ControllersOIDCTokenResponse`

NewControllersOIDCTokenResponse instantiates a new ControllersOIDCTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersOIDCTokenResponseWithDefaults

`func NewControllersOIDCTokenResponseWithDefaults() *ControllersOIDCTokenResponse`

NewControllersOIDCTokenResponseWithDefaults instantiates a new ControllersOIDCTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaims

`func (o *ControllersOIDCTokenResponse) GetClaims() AuthClaims`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *ControllersOIDCTokenResponse) GetClaimsOk() (*AuthClaims, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *ControllersOIDCTokenResponse) SetClaims(v AuthClaims)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *ControllersOIDCTokenResponse) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetExpiresIn

`func (o *ControllersOIDCTokenResponse) GetExpiresIn() string`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *ControllersOIDCTokenResponse) GetExpiresInOk() (*string, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *ControllersOIDCTokenResponse) SetExpiresIn(v string)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *ControllersOIDCTokenResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetIdToken

`func (o *ControllersOIDCTokenResponse) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *ControllersOIDCTokenResponse) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *ControllersOIDCTokenResponse) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.

### HasIdToken

`func (o *ControllersOIDCTokenResponse) HasIdToken() bool`

HasIdToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *ControllersOIDCTokenResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *ControllersOIDCTokenResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *ControllersOIDCTokenResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *ControllersOIDCTokenResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


