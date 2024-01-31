# V1DestinationReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetClusterId** | Pointer to **string** |  | [optional] 
**DeploymentId** | Pointer to **string** | UID of the deployment created by the restore. | [optional] [readonly] 

## Methods

### NewV1DestinationReferences

`func NewV1DestinationReferences() *V1DestinationReferences`

NewV1DestinationReferences instantiates a new V1DestinationReferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DestinationReferencesWithDefaults

`func NewV1DestinationReferencesWithDefaults() *V1DestinationReferences`

NewV1DestinationReferencesWithDefaults instantiates a new V1DestinationReferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetClusterId

`func (o *V1DestinationReferences) GetTargetClusterId() string`

GetTargetClusterId returns the TargetClusterId field if non-nil, zero value otherwise.

### GetTargetClusterIdOk

`func (o *V1DestinationReferences) GetTargetClusterIdOk() (*string, bool)`

GetTargetClusterIdOk returns a tuple with the TargetClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterId

`func (o *V1DestinationReferences) SetTargetClusterId(v string)`

SetTargetClusterId sets TargetClusterId field to given value.

### HasTargetClusterId

`func (o *V1DestinationReferences) HasTargetClusterId() bool`

HasTargetClusterId returns a boolean if a field has been set.

### GetDeploymentId

`func (o *V1DestinationReferences) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *V1DestinationReferences) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *V1DestinationReferences) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *V1DestinationReferences) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


