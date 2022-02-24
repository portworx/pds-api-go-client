# ControllersCreateProjectDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** |  | [optional] 
**Configuration** | Pointer to **map[string]interface{}** |  | [optional] 
**DataServiceId** | Pointer to **string** |  | [optional] 
**DeploymentTargetId** | Pointer to **string** |  | [optional] 
**DnsZone** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**LoadBalancerSourceRanges** | Pointer to **[]string** |  | [optional] 
**NamespaceId** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int32** |  | [optional] 
**Resources** | Pointer to **map[string]interface{}** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 

## Methods

### NewControllersCreateProjectDeployment

`func NewControllersCreateProjectDeployment() *ControllersCreateProjectDeployment`

NewControllersCreateProjectDeployment instantiates a new ControllersCreateProjectDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersCreateProjectDeploymentWithDefaults

`func NewControllersCreateProjectDeploymentWithDefaults() *ControllersCreateProjectDeployment`

NewControllersCreateProjectDeploymentWithDefaults instantiates a new ControllersCreateProjectDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *ControllersCreateProjectDeployment) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ControllersCreateProjectDeployment) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ControllersCreateProjectDeployment) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *ControllersCreateProjectDeployment) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetConfiguration

`func (o *ControllersCreateProjectDeployment) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ControllersCreateProjectDeployment) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ControllersCreateProjectDeployment) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ControllersCreateProjectDeployment) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDataServiceId

`func (o *ControllersCreateProjectDeployment) GetDataServiceId() string`

GetDataServiceId returns the DataServiceId field if non-nil, zero value otherwise.

### GetDataServiceIdOk

`func (o *ControllersCreateProjectDeployment) GetDataServiceIdOk() (*string, bool)`

GetDataServiceIdOk returns a tuple with the DataServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataServiceId

`func (o *ControllersCreateProjectDeployment) SetDataServiceId(v string)`

SetDataServiceId sets DataServiceId field to given value.

### HasDataServiceId

`func (o *ControllersCreateProjectDeployment) HasDataServiceId() bool`

HasDataServiceId returns a boolean if a field has been set.

### GetDeploymentTargetId

`func (o *ControllersCreateProjectDeployment) GetDeploymentTargetId() string`

GetDeploymentTargetId returns the DeploymentTargetId field if non-nil, zero value otherwise.

### GetDeploymentTargetIdOk

`func (o *ControllersCreateProjectDeployment) GetDeploymentTargetIdOk() (*string, bool)`

GetDeploymentTargetIdOk returns a tuple with the DeploymentTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTargetId

`func (o *ControllersCreateProjectDeployment) SetDeploymentTargetId(v string)`

SetDeploymentTargetId sets DeploymentTargetId field to given value.

### HasDeploymentTargetId

`func (o *ControllersCreateProjectDeployment) HasDeploymentTargetId() bool`

HasDeploymentTargetId returns a boolean if a field has been set.

### GetDnsZone

`func (o *ControllersCreateProjectDeployment) GetDnsZone() string`

GetDnsZone returns the DnsZone field if non-nil, zero value otherwise.

### GetDnsZoneOk

`func (o *ControllersCreateProjectDeployment) GetDnsZoneOk() (*string, bool)`

GetDnsZoneOk returns a tuple with the DnsZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZone

`func (o *ControllersCreateProjectDeployment) SetDnsZone(v string)`

SetDnsZone sets DnsZone field to given value.

### HasDnsZone

`func (o *ControllersCreateProjectDeployment) HasDnsZone() bool`

HasDnsZone returns a boolean if a field has been set.

### GetImageId

`func (o *ControllersCreateProjectDeployment) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ControllersCreateProjectDeployment) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ControllersCreateProjectDeployment) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *ControllersCreateProjectDeployment) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetLoadBalancerSourceRanges

`func (o *ControllersCreateProjectDeployment) GetLoadBalancerSourceRanges() []string`

GetLoadBalancerSourceRanges returns the LoadBalancerSourceRanges field if non-nil, zero value otherwise.

### GetLoadBalancerSourceRangesOk

`func (o *ControllersCreateProjectDeployment) GetLoadBalancerSourceRangesOk() (*[]string, bool)`

GetLoadBalancerSourceRangesOk returns a tuple with the LoadBalancerSourceRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerSourceRanges

`func (o *ControllersCreateProjectDeployment) SetLoadBalancerSourceRanges(v []string)`

SetLoadBalancerSourceRanges sets LoadBalancerSourceRanges field to given value.

### HasLoadBalancerSourceRanges

`func (o *ControllersCreateProjectDeployment) HasLoadBalancerSourceRanges() bool`

HasLoadBalancerSourceRanges returns a boolean if a field has been set.

### GetNamespaceId

`func (o *ControllersCreateProjectDeployment) GetNamespaceId() string`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *ControllersCreateProjectDeployment) GetNamespaceIdOk() (*string, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *ControllersCreateProjectDeployment) SetNamespaceId(v string)`

SetNamespaceId sets NamespaceId field to given value.

### HasNamespaceId

`func (o *ControllersCreateProjectDeployment) HasNamespaceId() bool`

HasNamespaceId returns a boolean if a field has been set.

### GetNodeCount

`func (o *ControllersCreateProjectDeployment) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ControllersCreateProjectDeployment) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ControllersCreateProjectDeployment) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ControllersCreateProjectDeployment) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetResources

`func (o *ControllersCreateProjectDeployment) GetResources() map[string]interface{}`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ControllersCreateProjectDeployment) GetResourcesOk() (*map[string]interface{}, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ControllersCreateProjectDeployment) SetResources(v map[string]interface{})`

SetResources sets Resources field to given value.

### HasResources

`func (o *ControllersCreateProjectDeployment) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetService

`func (o *ControllersCreateProjectDeployment) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ControllersCreateProjectDeployment) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ControllersCreateProjectDeployment) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ControllersCreateProjectDeployment) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceType

`func (o *ControllersCreateProjectDeployment) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ControllersCreateProjectDeployment) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ControllersCreateProjectDeployment) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *ControllersCreateProjectDeployment) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetVersionId

`func (o *ControllersCreateProjectDeployment) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ControllersCreateProjectDeployment) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ControllersCreateProjectDeployment) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *ControllersCreateProjectDeployment) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


