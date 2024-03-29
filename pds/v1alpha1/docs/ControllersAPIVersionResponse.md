# ControllersAPIVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **map[string]string** |  | [optional] 
**HelmChartVersion** | Pointer to **string** |  | [optional] 
**PdsBuildNumber** | Pointer to **int32** |  | [optional] 

## Methods

### NewControllersAPIVersionResponse

`func NewControllersAPIVersionResponse() *ControllersAPIVersionResponse`

NewControllersAPIVersionResponse instantiates a new ControllersAPIVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersAPIVersionResponseWithDefaults

`func NewControllersAPIVersionResponseWithDefaults() *ControllersAPIVersionResponse`

NewControllersAPIVersionResponseWithDefaults instantiates a new ControllersAPIVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ControllersAPIVersionResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ControllersAPIVersionResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ControllersAPIVersionResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ControllersAPIVersionResponse) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetFeatures

`func (o *ControllersAPIVersionResponse) GetFeatures() map[string]string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ControllersAPIVersionResponse) GetFeaturesOk() (*map[string]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ControllersAPIVersionResponse) SetFeatures(v map[string]string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ControllersAPIVersionResponse) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetHelmChartVersion

`func (o *ControllersAPIVersionResponse) GetHelmChartVersion() string`

GetHelmChartVersion returns the HelmChartVersion field if non-nil, zero value otherwise.

### GetHelmChartVersionOk

`func (o *ControllersAPIVersionResponse) GetHelmChartVersionOk() (*string, bool)`

GetHelmChartVersionOk returns a tuple with the HelmChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmChartVersion

`func (o *ControllersAPIVersionResponse) SetHelmChartVersion(v string)`

SetHelmChartVersion sets HelmChartVersion field to given value.

### HasHelmChartVersion

`func (o *ControllersAPIVersionResponse) HasHelmChartVersion() bool`

HasHelmChartVersion returns a boolean if a field has been set.

### GetPdsBuildNumber

`func (o *ControllersAPIVersionResponse) GetPdsBuildNumber() int32`

GetPdsBuildNumber returns the PdsBuildNumber field if non-nil, zero value otherwise.

### GetPdsBuildNumberOk

`func (o *ControllersAPIVersionResponse) GetPdsBuildNumberOk() (*int32, bool)`

GetPdsBuildNumberOk returns a tuple with the PdsBuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdsBuildNumber

`func (o *ControllersAPIVersionResponse) SetPdsBuildNumber(v int32)`

SetPdsBuildNumber sets PdsBuildNumber field to given value.

### HasPdsBuildNumber

`func (o *ControllersAPIVersionResponse) HasPdsBuildNumber() bool`

HasPdsBuildNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


