# RequestsDeploymentEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | Pointer to **string** | DeploymentID of the event. | [optional] 
**Events** | Pointer to [**[]RequestsK8sEvent**](RequestsK8sEvent.md) | List of k8s events for the deployment | [optional] 
**Namespace** | Pointer to **string** | Namespace in which the deployment events were generated. | [optional] 

## Methods

### NewRequestsDeploymentEvent

`func NewRequestsDeploymentEvent() *RequestsDeploymentEvent`

NewRequestsDeploymentEvent instantiates a new RequestsDeploymentEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsDeploymentEventWithDefaults

`func NewRequestsDeploymentEventWithDefaults() *RequestsDeploymentEvent`

NewRequestsDeploymentEventWithDefaults instantiates a new RequestsDeploymentEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *RequestsDeploymentEvent) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *RequestsDeploymentEvent) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *RequestsDeploymentEvent) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *RequestsDeploymentEvent) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetEvents

`func (o *RequestsDeploymentEvent) GetEvents() []RequestsK8sEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RequestsDeploymentEvent) GetEventsOk() (*[]RequestsK8sEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RequestsDeploymentEvent) SetEvents(v []RequestsK8sEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *RequestsDeploymentEvent) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetNamespace

`func (o *RequestsDeploymentEvent) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *RequestsDeploymentEvent) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *RequestsDeploymentEvent) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *RequestsDeploymentEvent) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


