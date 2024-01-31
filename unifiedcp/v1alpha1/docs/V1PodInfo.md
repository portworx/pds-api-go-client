# V1PodInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The IP of a pod. | [optional] 
**Name** | Pointer to **string** | Name is the Hostname of a pod. | [optional] 
**WorkerNode** | Pointer to **string** | Node that hosts a particular pod. | [optional] 

## Methods

### NewV1PodInfo

`func NewV1PodInfo() *V1PodInfo`

NewV1PodInfo instantiates a new V1PodInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PodInfoWithDefaults

`func NewV1PodInfoWithDefaults() *V1PodInfo`

NewV1PodInfoWithDefaults instantiates a new V1PodInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *V1PodInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V1PodInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V1PodInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V1PodInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetName

`func (o *V1PodInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1PodInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1PodInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1PodInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWorkerNode

`func (o *V1PodInfo) GetWorkerNode() string`

GetWorkerNode returns the WorkerNode field if non-nil, zero value otherwise.

### GetWorkerNodeOk

`func (o *V1PodInfo) GetWorkerNodeOk() (*string, bool)`

GetWorkerNodeOk returns a tuple with the WorkerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerNode

`func (o *V1PodInfo) SetWorkerNode(v string)`

SetWorkerNode sets WorkerNode field to given value.

### HasWorkerNode

`func (o *V1PodInfo) HasWorkerNode() bool`

HasWorkerNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


