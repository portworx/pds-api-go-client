# V1Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KubeServerVersion** | Pointer to **string** |  | [optional] 
**KubePlatform** | Pointer to [**V1KubePlatformType**](V1KubePlatformType.md) |  | [optional] [default to TYPE_UNSPECIFIED]

## Methods

### NewV1Metadata

`func NewV1Metadata() *V1Metadata`

NewV1Metadata instantiates a new V1Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MetadataWithDefaults

`func NewV1MetadataWithDefaults() *V1Metadata`

NewV1MetadataWithDefaults instantiates a new V1Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubeServerVersion

`func (o *V1Metadata) GetKubeServerVersion() string`

GetKubeServerVersion returns the KubeServerVersion field if non-nil, zero value otherwise.

### GetKubeServerVersionOk

`func (o *V1Metadata) GetKubeServerVersionOk() (*string, bool)`

GetKubeServerVersionOk returns a tuple with the KubeServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeServerVersion

`func (o *V1Metadata) SetKubeServerVersion(v string)`

SetKubeServerVersion sets KubeServerVersion field to given value.

### HasKubeServerVersion

`func (o *V1Metadata) HasKubeServerVersion() bool`

HasKubeServerVersion returns a boolean if a field has been set.

### GetKubePlatform

`func (o *V1Metadata) GetKubePlatform() V1KubePlatformType`

GetKubePlatform returns the KubePlatform field if non-nil, zero value otherwise.

### GetKubePlatformOk

`func (o *V1Metadata) GetKubePlatformOk() (*V1KubePlatformType, bool)`

GetKubePlatformOk returns a tuple with the KubePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubePlatform

`func (o *V1Metadata) SetKubePlatform(v V1KubePlatformType)`

SetKubePlatform sets KubePlatform field to given value.

### HasKubePlatform

`func (o *V1Metadata) HasKubePlatform() bool`

HasKubePlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


