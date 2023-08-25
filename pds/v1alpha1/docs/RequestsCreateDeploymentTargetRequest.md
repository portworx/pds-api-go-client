# RequestsCreateDeploymentTargetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TeleportAgentVersion** | Pointer to **string** |  | [optional] 
**TlsIssuer** | Pointer to **string** |  | [optional] 
**TlsRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewRequestsCreateDeploymentTargetRequest

`func NewRequestsCreateDeploymentTargetRequest() *RequestsCreateDeploymentTargetRequest`

NewRequestsCreateDeploymentTargetRequest instantiates a new RequestsCreateDeploymentTargetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsCreateDeploymentTargetRequestWithDefaults

`func NewRequestsCreateDeploymentTargetRequestWithDefaults() *RequestsCreateDeploymentTargetRequest`

NewRequestsCreateDeploymentTargetRequestWithDefaults instantiates a new RequestsCreateDeploymentTargetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *RequestsCreateDeploymentTargetRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *RequestsCreateDeploymentTargetRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *RequestsCreateDeploymentTargetRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *RequestsCreateDeploymentTargetRequest) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetName

`func (o *RequestsCreateDeploymentTargetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestsCreateDeploymentTargetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestsCreateDeploymentTargetRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequestsCreateDeploymentTargetRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTeleportAgentVersion

`func (o *RequestsCreateDeploymentTargetRequest) GetTeleportAgentVersion() string`

GetTeleportAgentVersion returns the TeleportAgentVersion field if non-nil, zero value otherwise.

### GetTeleportAgentVersionOk

`func (o *RequestsCreateDeploymentTargetRequest) GetTeleportAgentVersionOk() (*string, bool)`

GetTeleportAgentVersionOk returns a tuple with the TeleportAgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeleportAgentVersion

`func (o *RequestsCreateDeploymentTargetRequest) SetTeleportAgentVersion(v string)`

SetTeleportAgentVersion sets TeleportAgentVersion field to given value.

### HasTeleportAgentVersion

`func (o *RequestsCreateDeploymentTargetRequest) HasTeleportAgentVersion() bool`

HasTeleportAgentVersion returns a boolean if a field has been set.

### GetTlsIssuer

`func (o *RequestsCreateDeploymentTargetRequest) GetTlsIssuer() string`

GetTlsIssuer returns the TlsIssuer field if non-nil, zero value otherwise.

### GetTlsIssuerOk

`func (o *RequestsCreateDeploymentTargetRequest) GetTlsIssuerOk() (*string, bool)`

GetTlsIssuerOk returns a tuple with the TlsIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsIssuer

`func (o *RequestsCreateDeploymentTargetRequest) SetTlsIssuer(v string)`

SetTlsIssuer sets TlsIssuer field to given value.

### HasTlsIssuer

`func (o *RequestsCreateDeploymentTargetRequest) HasTlsIssuer() bool`

HasTlsIssuer returns a boolean if a field has been set.

### GetTlsRequired

`func (o *RequestsCreateDeploymentTargetRequest) GetTlsRequired() bool`

GetTlsRequired returns the TlsRequired field if non-nil, zero value otherwise.

### GetTlsRequiredOk

`func (o *RequestsCreateDeploymentTargetRequest) GetTlsRequiredOk() (*bool, bool)`

GetTlsRequiredOk returns a tuple with the TlsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsRequired

`func (o *RequestsCreateDeploymentTargetRequest) SetTlsRequired(v bool)`

SetTlsRequired sets TlsRequired field to given value.

### HasTlsRequired

`func (o *RequestsCreateDeploymentTargetRequest) HasTlsRequired() bool`

HasTlsRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


