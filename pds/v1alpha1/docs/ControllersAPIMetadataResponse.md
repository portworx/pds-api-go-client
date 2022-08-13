# ControllersAPIMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | Pointer to **map[string]string** |  | [optional] 
**HelmChartVersion** | Pointer to **string** |  | [optional] 
**PdsBuildNumber** | Pointer to **int32** |  | [optional] 

## Methods

### NewControllersAPIMetadataResponse

`func NewControllersAPIMetadataResponse() *ControllersAPIMetadataResponse`

NewControllersAPIMetadataResponse instantiates a new ControllersAPIMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersAPIMetadataResponseWithDefaults

`func NewControllersAPIMetadataResponseWithDefaults() *ControllersAPIMetadataResponse`

NewControllersAPIMetadataResponseWithDefaults instantiates a new ControllersAPIMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *ControllersAPIMetadataResponse) GetFeatures() map[string]string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ControllersAPIMetadataResponse) GetFeaturesOk() (*map[string]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ControllersAPIMetadataResponse) SetFeatures(v map[string]string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ControllersAPIMetadataResponse) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetHelmChartVersion

`func (o *ControllersAPIMetadataResponse) GetHelmChartVersion() string`

GetHelmChartVersion returns the HelmChartVersion field if non-nil, zero value otherwise.

### GetHelmChartVersionOk

`func (o *ControllersAPIMetadataResponse) GetHelmChartVersionOk() (*string, bool)`

GetHelmChartVersionOk returns a tuple with the HelmChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmChartVersion

`func (o *ControllersAPIMetadataResponse) SetHelmChartVersion(v string)`

SetHelmChartVersion sets HelmChartVersion field to given value.

### HasHelmChartVersion

`func (o *ControllersAPIMetadataResponse) HasHelmChartVersion() bool`

HasHelmChartVersion returns a boolean if a field has been set.

### GetPdsBuildNumber

`func (o *ControllersAPIMetadataResponse) GetPdsBuildNumber() int32`

GetPdsBuildNumber returns the PdsBuildNumber field if non-nil, zero value otherwise.

### GetPdsBuildNumberOk

`func (o *ControllersAPIMetadataResponse) GetPdsBuildNumberOk() (*int32, bool)`

GetPdsBuildNumberOk returns a tuple with the PdsBuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdsBuildNumber

`func (o *ControllersAPIMetadataResponse) SetPdsBuildNumber(v int32)`

SetPdsBuildNumber sets PdsBuildNumber field to given value.

### HasPdsBuildNumber

`func (o *ControllersAPIMetadataResponse) HasPdsBuildNumber() bool`

HasPdsBuildNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


