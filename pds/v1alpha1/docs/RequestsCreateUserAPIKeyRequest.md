# RequestsCreateUserAPIKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **string** | Time when the key expires. | [optional] 
**Name** | Pointer to **string** | Name of the UserAPIKey. Must be unique for the given User. | [optional] 

## Methods

### NewRequestsCreateUserAPIKeyRequest

`func NewRequestsCreateUserAPIKeyRequest() *RequestsCreateUserAPIKeyRequest`

NewRequestsCreateUserAPIKeyRequest instantiates a new RequestsCreateUserAPIKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsCreateUserAPIKeyRequestWithDefaults

`func NewRequestsCreateUserAPIKeyRequestWithDefaults() *RequestsCreateUserAPIKeyRequest`

NewRequestsCreateUserAPIKeyRequestWithDefaults instantiates a new RequestsCreateUserAPIKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *RequestsCreateUserAPIKeyRequest) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *RequestsCreateUserAPIKeyRequest) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *RequestsCreateUserAPIKeyRequest) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *RequestsCreateUserAPIKeyRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetName

`func (o *RequestsCreateUserAPIKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestsCreateUserAPIKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestsCreateUserAPIKeyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequestsCreateUserAPIKeyRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


