# V1References3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetClusterId** | Pointer to **string** | UID of the target cluster in which Data Service will be deployed. | [optional] 
**ImageId** | Pointer to **string** | UID of the image to be used for the Data Service Deployment. | [optional] 
**ProjectId** | Pointer to **string** | UID of the project to which DataService Deployment associated. | [optional] 
**RestoreId** | Pointer to **string** | UID of the restore id for the Deployment. | [optional] [readonly] 

## Methods

### NewV1References3

`func NewV1References3() *V1References3`

NewV1References3 instantiates a new V1References3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1References3WithDefaults

`func NewV1References3WithDefaults() *V1References3`

NewV1References3WithDefaults instantiates a new V1References3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetClusterId

`func (o *V1References3) GetTargetClusterId() string`

GetTargetClusterId returns the TargetClusterId field if non-nil, zero value otherwise.

### GetTargetClusterIdOk

`func (o *V1References3) GetTargetClusterIdOk() (*string, bool)`

GetTargetClusterIdOk returns a tuple with the TargetClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterId

`func (o *V1References3) SetTargetClusterId(v string)`

SetTargetClusterId sets TargetClusterId field to given value.

### HasTargetClusterId

`func (o *V1References3) HasTargetClusterId() bool`

HasTargetClusterId returns a boolean if a field has been set.

### GetImageId

`func (o *V1References3) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *V1References3) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *V1References3) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *V1References3) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetProjectId

`func (o *V1References3) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *V1References3) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *V1References3) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *V1References3) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRestoreId

`func (o *V1References3) GetRestoreId() string`

GetRestoreId returns the RestoreId field if non-nil, zero value otherwise.

### GetRestoreIdOk

`func (o *V1References3) GetRestoreIdOk() (*string, bool)`

GetRestoreIdOk returns a tuple with the RestoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreId

`func (o *V1References3) SetRestoreId(v string)`

SetRestoreId sets RestoreId field to given value.

### HasRestoreId

`func (o *V1References3) HasRestoreId() bool`

HasRestoreId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


