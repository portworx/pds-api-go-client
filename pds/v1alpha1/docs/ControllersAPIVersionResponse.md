# ControllersAPIVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**HelmChartVersion** | Pointer to **string** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


