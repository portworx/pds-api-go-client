# RequestsPatchUserAPIKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether the UserAPIKey is enabled or disabled. | [optional] 
**Name** | Pointer to **string** | The name of the UserAPIKey. | [optional] 

## Methods

### NewRequestsPatchUserAPIKeyRequest

`func NewRequestsPatchUserAPIKeyRequest() *RequestsPatchUserAPIKeyRequest`

NewRequestsPatchUserAPIKeyRequest instantiates a new RequestsPatchUserAPIKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsPatchUserAPIKeyRequestWithDefaults

`func NewRequestsPatchUserAPIKeyRequestWithDefaults() *RequestsPatchUserAPIKeyRequest`

NewRequestsPatchUserAPIKeyRequestWithDefaults instantiates a new RequestsPatchUserAPIKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *RequestsPatchUserAPIKeyRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RequestsPatchUserAPIKeyRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RequestsPatchUserAPIKeyRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RequestsPatchUserAPIKeyRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *RequestsPatchUserAPIKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestsPatchUserAPIKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestsPatchUserAPIKeyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequestsPatchUserAPIKeyRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


