# RequestsPatchDeploymentTargetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**TlsIssuer** | Pointer to **string** |  | [optional] 
**TlsRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewRequestsPatchDeploymentTargetRequest

`func NewRequestsPatchDeploymentTargetRequest() *RequestsPatchDeploymentTargetRequest`

NewRequestsPatchDeploymentTargetRequest instantiates a new RequestsPatchDeploymentTargetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsPatchDeploymentTargetRequestWithDefaults

`func NewRequestsPatchDeploymentTargetRequestWithDefaults() *RequestsPatchDeploymentTargetRequest`

NewRequestsPatchDeploymentTargetRequestWithDefaults instantiates a new RequestsPatchDeploymentTargetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RequestsPatchDeploymentTargetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestsPatchDeploymentTargetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestsPatchDeploymentTargetRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequestsPatchDeploymentTargetRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTlsIssuer

`func (o *RequestsPatchDeploymentTargetRequest) GetTlsIssuer() string`

GetTlsIssuer returns the TlsIssuer field if non-nil, zero value otherwise.

### GetTlsIssuerOk

`func (o *RequestsPatchDeploymentTargetRequest) GetTlsIssuerOk() (*string, bool)`

GetTlsIssuerOk returns a tuple with the TlsIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsIssuer

`func (o *RequestsPatchDeploymentTargetRequest) SetTlsIssuer(v string)`

SetTlsIssuer sets TlsIssuer field to given value.

### HasTlsIssuer

`func (o *RequestsPatchDeploymentTargetRequest) HasTlsIssuer() bool`

HasTlsIssuer returns a boolean if a field has been set.

### GetTlsRequired

`func (o *RequestsPatchDeploymentTargetRequest) GetTlsRequired() bool`

GetTlsRequired returns the TlsRequired field if non-nil, zero value otherwise.

### GetTlsRequiredOk

`func (o *RequestsPatchDeploymentTargetRequest) GetTlsRequiredOk() (*bool, bool)`

GetTlsRequiredOk returns a tuple with the TlsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsRequired

`func (o *RequestsPatchDeploymentTargetRequest) SetTlsRequired(v bool)`

SetTlsRequired sets TlsRequired field to given value.

### HasTlsRequired

`func (o *RequestsPatchDeploymentTargetRequest) HasTlsRequired() bool`

HasTlsRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


