# V1EndpointAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | The Hostname of this endpoint +optional | [optional] 
**Ip** | Pointer to **string** | The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24). IPv6 is also accepted but not fully supported on all platforms. Also, certain kubernetes components, like kube-proxy, are not IPv6 ready. TODO: This should allow hostname or IP, See #4447. | [optional] 
**NodeName** | Pointer to **string** | Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node. +optional | [optional] 
**TargetRef** | Pointer to [**V1ObjectReference**](V1ObjectReference.md) |  | [optional] 

## Methods

### NewV1EndpointAddress

`func NewV1EndpointAddress() *V1EndpointAddress`

NewV1EndpointAddress instantiates a new V1EndpointAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EndpointAddressWithDefaults

`func NewV1EndpointAddressWithDefaults() *V1EndpointAddress`

NewV1EndpointAddressWithDefaults instantiates a new V1EndpointAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *V1EndpointAddress) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V1EndpointAddress) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V1EndpointAddress) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V1EndpointAddress) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIp

`func (o *V1EndpointAddress) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V1EndpointAddress) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V1EndpointAddress) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V1EndpointAddress) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNodeName

`func (o *V1EndpointAddress) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *V1EndpointAddress) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *V1EndpointAddress) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *V1EndpointAddress) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetTargetRef

`func (o *V1EndpointAddress) GetTargetRef() V1ObjectReference`

GetTargetRef returns the TargetRef field if non-nil, zero value otherwise.

### GetTargetRefOk

`func (o *V1EndpointAddress) GetTargetRefOk() (*V1ObjectReference, bool)`

GetTargetRefOk returns a tuple with the TargetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRef

`func (o *V1EndpointAddress) SetTargetRef(v V1ObjectReference)`

SetTargetRef sets TargetRef field to given value.

### HasTargetRef

`func (o *V1EndpointAddress) HasTargetRef() bool`

HasTargetRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


