# V1DeploymentConfigUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V1Meta**](V1Meta.md) |  | [optional] 
**Config** | Pointer to [**Pdsdeploymentconfigupdatev1Config**](Pdsdeploymentconfigupdatev1Config.md) |  | [optional] 
**Status** | Pointer to [**Pdsdeploymentconfigupdatev1Status**](Pdsdeploymentconfigupdatev1Status.md) |  | [optional] 

## Methods

### NewV1DeploymentConfigUpdate

`func NewV1DeploymentConfigUpdate() *V1DeploymentConfigUpdate`

NewV1DeploymentConfigUpdate instantiates a new V1DeploymentConfigUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeploymentConfigUpdateWithDefaults

`func NewV1DeploymentConfigUpdateWithDefaults() *V1DeploymentConfigUpdate`

NewV1DeploymentConfigUpdateWithDefaults instantiates a new V1DeploymentConfigUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V1DeploymentConfigUpdate) GetMeta() V1Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V1DeploymentConfigUpdate) GetMetaOk() (*V1Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V1DeploymentConfigUpdate) SetMeta(v V1Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V1DeploymentConfigUpdate) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetConfig

`func (o *V1DeploymentConfigUpdate) GetConfig() Pdsdeploymentconfigupdatev1Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V1DeploymentConfigUpdate) GetConfigOk() (*Pdsdeploymentconfigupdatev1Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V1DeploymentConfigUpdate) SetConfig(v Pdsdeploymentconfigupdatev1Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *V1DeploymentConfigUpdate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *V1DeploymentConfigUpdate) GetStatus() Pdsdeploymentconfigupdatev1Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1DeploymentConfigUpdate) GetStatusOk() (*Pdsdeploymentconfigupdatev1Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1DeploymentConfigUpdate) SetStatus(v Pdsdeploymentconfigupdatev1Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1DeploymentConfigUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


