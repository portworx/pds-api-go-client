/*
PDS API

Portworx Data Services API Server

API version: 3.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"encoding/json"
)

// ControllersOIDCTokenResponse struct for ControllersOIDCTokenResponse
type ControllersOIDCTokenResponse struct {
	Claims *AuthClaims `json:"claims,omitempty"`
	ExpiresIn *string `json:"expiresIn,omitempty"`
	IdToken *string `json:"idToken,omitempty"`
	RefreshToken *string `json:"refreshToken,omitempty"`
}

// NewControllersOIDCTokenResponse instantiates a new ControllersOIDCTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersOIDCTokenResponse() *ControllersOIDCTokenResponse {
	this := ControllersOIDCTokenResponse{}
	return &this
}

// NewControllersOIDCTokenResponseWithDefaults instantiates a new ControllersOIDCTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersOIDCTokenResponseWithDefaults() *ControllersOIDCTokenResponse {
	this := ControllersOIDCTokenResponse{}
	return &this
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *ControllersOIDCTokenResponse) GetClaims() AuthClaims {
	if o == nil || o.Claims == nil {
		var ret AuthClaims
		return ret
	}
	return *o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersOIDCTokenResponse) GetClaimsOk() (*AuthClaims, bool) {
	if o == nil || o.Claims == nil {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *ControllersOIDCTokenResponse) HasClaims() bool {
	if o != nil && o.Claims != nil {
		return true
	}

	return false
}

// SetClaims gets a reference to the given AuthClaims and assigns it to the Claims field.
func (o *ControllersOIDCTokenResponse) SetClaims(v AuthClaims) {
	o.Claims = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *ControllersOIDCTokenResponse) GetExpiresIn() string {
	if o == nil || o.ExpiresIn == nil {
		var ret string
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersOIDCTokenResponse) GetExpiresInOk() (*string, bool) {
	if o == nil || o.ExpiresIn == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *ControllersOIDCTokenResponse) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn != nil {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given string and assigns it to the ExpiresIn field.
func (o *ControllersOIDCTokenResponse) SetExpiresIn(v string) {
	o.ExpiresIn = &v
}

// GetIdToken returns the IdToken field value if set, zero value otherwise.
func (o *ControllersOIDCTokenResponse) GetIdToken() string {
	if o == nil || o.IdToken == nil {
		var ret string
		return ret
	}
	return *o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersOIDCTokenResponse) GetIdTokenOk() (*string, bool) {
	if o == nil || o.IdToken == nil {
		return nil, false
	}
	return o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *ControllersOIDCTokenResponse) HasIdToken() bool {
	if o != nil && o.IdToken != nil {
		return true
	}

	return false
}

// SetIdToken gets a reference to the given string and assigns it to the IdToken field.
func (o *ControllersOIDCTokenResponse) SetIdToken(v string) {
	o.IdToken = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *ControllersOIDCTokenResponse) GetRefreshToken() string {
	if o == nil || o.RefreshToken == nil {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersOIDCTokenResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil || o.RefreshToken == nil {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *ControllersOIDCTokenResponse) HasRefreshToken() bool {
	if o != nil && o.RefreshToken != nil {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *ControllersOIDCTokenResponse) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

func (o ControllersOIDCTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Claims != nil {
		toSerialize["claims"] = o.Claims
	}
	if o.ExpiresIn != nil {
		toSerialize["expiresIn"] = o.ExpiresIn
	}
	if o.IdToken != nil {
		toSerialize["idToken"] = o.IdToken
	}
	if o.RefreshToken != nil {
		toSerialize["refreshToken"] = o.RefreshToken
	}
	return json.Marshal(toSerialize)
}

type NullableControllersOIDCTokenResponse struct {
	value *ControllersOIDCTokenResponse
	isSet bool
}

func (v NullableControllersOIDCTokenResponse) Get() *ControllersOIDCTokenResponse {
	return v.value
}

func (v *NullableControllersOIDCTokenResponse) Set(val *ControllersOIDCTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersOIDCTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersOIDCTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersOIDCTokenResponse(val *ControllersOIDCTokenResponse) *NullableControllersOIDCTokenResponse {
	return &NullableControllersOIDCTokenResponse{value: val, isSet: true}
}

func (v NullableControllersOIDCTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersOIDCTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


