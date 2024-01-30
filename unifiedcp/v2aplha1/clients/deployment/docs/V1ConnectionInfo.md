# V1ConnectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pods** | Pointer to [**[]V1PodInfo**](V1PodInfo.md) | Ready pods. | [optional] 
**NotReadyPods** | Pointer to [**[]V1PodInfo**](V1PodInfo.md) | Pods that are not ready. | [optional] 
**ConnectionDetails** | Pointer to [**V1ConnectionDetails**](V1ConnectionDetails.md) |  | [optional] 
**ClusterDetails** | Pointer to [**map[string]ProtobufAny**](ProtobufAny.md) | Stores details about the cluster. | [optional] 

## Methods

### NewV1ConnectionInfo

`func NewV1ConnectionInfo() *V1ConnectionInfo`

NewV1ConnectionInfo instantiates a new V1ConnectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ConnectionInfoWithDefaults

`func NewV1ConnectionInfoWithDefaults() *V1ConnectionInfo`

NewV1ConnectionInfoWithDefaults instantiates a new V1ConnectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPods

`func (o *V1ConnectionInfo) GetPods() []V1PodInfo`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *V1ConnectionInfo) GetPodsOk() (*[]V1PodInfo, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *V1ConnectionInfo) SetPods(v []V1PodInfo)`

SetPods sets Pods field to given value.

### HasPods

`func (o *V1ConnectionInfo) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetNotReadyPods

`func (o *V1ConnectionInfo) GetNotReadyPods() []V1PodInfo`

GetNotReadyPods returns the NotReadyPods field if non-nil, zero value otherwise.

### GetNotReadyPodsOk

`func (o *V1ConnectionInfo) GetNotReadyPodsOk() (*[]V1PodInfo, bool)`

GetNotReadyPodsOk returns a tuple with the NotReadyPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotReadyPods

`func (o *V1ConnectionInfo) SetNotReadyPods(v []V1PodInfo)`

SetNotReadyPods sets NotReadyPods field to given value.

### HasNotReadyPods

`func (o *V1ConnectionInfo) HasNotReadyPods() bool`

HasNotReadyPods returns a boolean if a field has been set.

### GetConnectionDetails

`func (o *V1ConnectionInfo) GetConnectionDetails() V1ConnectionDetails`

GetConnectionDetails returns the ConnectionDetails field if non-nil, zero value otherwise.

### GetConnectionDetailsOk

`func (o *V1ConnectionInfo) GetConnectionDetailsOk() (*V1ConnectionDetails, bool)`

GetConnectionDetailsOk returns a tuple with the ConnectionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionDetails

`func (o *V1ConnectionInfo) SetConnectionDetails(v V1ConnectionDetails)`

SetConnectionDetails sets ConnectionDetails field to given value.

### HasConnectionDetails

`func (o *V1ConnectionInfo) HasConnectionDetails() bool`

HasConnectionDetails returns a boolean if a field has been set.

### GetClusterDetails

`func (o *V1ConnectionInfo) GetClusterDetails() map[string]ProtobufAny`

GetClusterDetails returns the ClusterDetails field if non-nil, zero value otherwise.

### GetClusterDetailsOk

`func (o *V1ConnectionInfo) GetClusterDetailsOk() (*map[string]ProtobufAny, bool)`

GetClusterDetailsOk returns a tuple with the ClusterDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterDetails

`func (o *V1ConnectionInfo) SetClusterDetails(v map[string]ProtobufAny)`

SetClusterDetails sets ClusterDetails field to given value.

### HasClusterDetails

`func (o *V1ConnectionInfo) HasClusterDetails() bool`

HasClusterDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


