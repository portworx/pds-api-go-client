# RequestsServiceIdentityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** |  | 
**Name** | **string** |  | 

## Methods

### NewRequestsServiceIdentityRequest

`func NewRequestsServiceIdentityRequest(enabled bool, name string, ) *RequestsServiceIdentityRequest`

NewRequestsServiceIdentityRequest instantiates a new RequestsServiceIdentityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsServiceIdentityRequestWithDefaults

`func NewRequestsServiceIdentityRequestWithDefaults() *RequestsServiceIdentityRequest`

NewRequestsServiceIdentityRequestWithDefaults instantiates a new RequestsServiceIdentityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RequestsServiceIdentityRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestsServiceIdentityRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestsServiceIdentityRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestsServiceIdentityRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *RequestsServiceIdentityRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RequestsServiceIdentityRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RequestsServiceIdentityRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetName

`func (o *RequestsServiceIdentityRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestsServiceIdentityRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestsServiceIdentityRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


