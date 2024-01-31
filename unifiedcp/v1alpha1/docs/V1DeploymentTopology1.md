# V1DeploymentTopology1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the deployment topology. | [optional] 
**Description** | Pointer to **string** | Description of the deployment topology. | [optional] 
**Replicas** | Pointer to **string** | Number of replicas of data services. | [optional] 
**ServiceType** | Pointer to **string** | Service type are standard Kubernetes service types such as clusterIP, NodePort, load balancers, etc. | [optional] 
**ServiceName** | Pointer to **string** | Service name is the name of service as provided by user. | [optional] 
**LoadBalancerSourceRanges** | Pointer to **[]string** | Source IP ranges to use for the deployed Load Balancer. | [optional] 
**ResourceTemplate** | Pointer to [**V1Template**](V1Template.md) |  | [optional] 
**ApplicationTemplate** | Pointer to [**V1Template**](V1Template.md) |  | [optional] 
**StorageTemplate** | Pointer to [**V1Template**](V1Template.md) |  | [optional] 

## Methods

### NewV1DeploymentTopology1

`func NewV1DeploymentTopology1() *V1DeploymentTopology1`

NewV1DeploymentTopology1 instantiates a new V1DeploymentTopology1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeploymentTopology1WithDefaults

`func NewV1DeploymentTopology1WithDefaults() *V1DeploymentTopology1`

NewV1DeploymentTopology1WithDefaults instantiates a new V1DeploymentTopology1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1DeploymentTopology1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1DeploymentTopology1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1DeploymentTopology1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1DeploymentTopology1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *V1DeploymentTopology1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1DeploymentTopology1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1DeploymentTopology1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1DeploymentTopology1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReplicas

`func (o *V1DeploymentTopology1) GetReplicas() string`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *V1DeploymentTopology1) GetReplicasOk() (*string, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *V1DeploymentTopology1) SetReplicas(v string)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *V1DeploymentTopology1) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetServiceType

`func (o *V1DeploymentTopology1) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *V1DeploymentTopology1) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *V1DeploymentTopology1) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *V1DeploymentTopology1) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetServiceName

`func (o *V1DeploymentTopology1) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *V1DeploymentTopology1) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *V1DeploymentTopology1) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *V1DeploymentTopology1) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetLoadBalancerSourceRanges

`func (o *V1DeploymentTopology1) GetLoadBalancerSourceRanges() []string`

GetLoadBalancerSourceRanges returns the LoadBalancerSourceRanges field if non-nil, zero value otherwise.

### GetLoadBalancerSourceRangesOk

`func (o *V1DeploymentTopology1) GetLoadBalancerSourceRangesOk() (*[]string, bool)`

GetLoadBalancerSourceRangesOk returns a tuple with the LoadBalancerSourceRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerSourceRanges

`func (o *V1DeploymentTopology1) SetLoadBalancerSourceRanges(v []string)`

SetLoadBalancerSourceRanges sets LoadBalancerSourceRanges field to given value.

### HasLoadBalancerSourceRanges

`func (o *V1DeploymentTopology1) HasLoadBalancerSourceRanges() bool`

HasLoadBalancerSourceRanges returns a boolean if a field has been set.

### GetResourceTemplate

`func (o *V1DeploymentTopology1) GetResourceTemplate() V1Template`

GetResourceTemplate returns the ResourceTemplate field if non-nil, zero value otherwise.

### GetResourceTemplateOk

`func (o *V1DeploymentTopology1) GetResourceTemplateOk() (*V1Template, bool)`

GetResourceTemplateOk returns a tuple with the ResourceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTemplate

`func (o *V1DeploymentTopology1) SetResourceTemplate(v V1Template)`

SetResourceTemplate sets ResourceTemplate field to given value.

### HasResourceTemplate

`func (o *V1DeploymentTopology1) HasResourceTemplate() bool`

HasResourceTemplate returns a boolean if a field has been set.

### GetApplicationTemplate

`func (o *V1DeploymentTopology1) GetApplicationTemplate() V1Template`

GetApplicationTemplate returns the ApplicationTemplate field if non-nil, zero value otherwise.

### GetApplicationTemplateOk

`func (o *V1DeploymentTopology1) GetApplicationTemplateOk() (*V1Template, bool)`

GetApplicationTemplateOk returns a tuple with the ApplicationTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationTemplate

`func (o *V1DeploymentTopology1) SetApplicationTemplate(v V1Template)`

SetApplicationTemplate sets ApplicationTemplate field to given value.

### HasApplicationTemplate

`func (o *V1DeploymentTopology1) HasApplicationTemplate() bool`

HasApplicationTemplate returns a boolean if a field has been set.

### GetStorageTemplate

`func (o *V1DeploymentTopology1) GetStorageTemplate() V1Template`

GetStorageTemplate returns the StorageTemplate field if non-nil, zero value otherwise.

### GetStorageTemplateOk

`func (o *V1DeploymentTopology1) GetStorageTemplateOk() (*V1Template, bool)`

GetStorageTemplateOk returns a tuple with the StorageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTemplate

`func (o *V1DeploymentTopology1) SetStorageTemplate(v V1Template)`

SetStorageTemplate sets StorageTemplate field to given value.

### HasStorageTemplate

`func (o *V1DeploymentTopology1) HasStorageTemplate() bool`

HasStorageTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


