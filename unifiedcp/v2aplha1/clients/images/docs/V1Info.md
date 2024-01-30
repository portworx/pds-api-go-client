# V1Info

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**References** | Pointer to [**V1References**](V1References.md) |  | [optional] 
**Registry** | Pointer to **string** | Image registry where the image is stored. | [optional] 
**Namespace** | Pointer to **string** | Image registry namespace where the image is stored. | [optional] 
**Tag** | Pointer to **string** | Tag associated with the image. | [optional] 
**Build** | Pointer to **string** | Build version of the image. | [optional] 
**TlsSupport** | Pointer to **bool** | Flag indicating if TLS is supported for a data service using this image. | [optional] 
**Capabilities** | Pointer to **map[string]string** | Capabilities associated with this image. | [optional] 
**AdditionalImages** | Pointer to **map[string]string** | Additional images associated with this data service image. | [optional] 

## Methods

### NewV1Info

`func NewV1Info() *V1Info`

NewV1Info instantiates a new V1Info object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InfoWithDefaults

`func NewV1InfoWithDefaults() *V1Info`

NewV1InfoWithDefaults instantiates a new V1Info object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferences

`func (o *V1Info) GetReferences() V1References`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *V1Info) GetReferencesOk() (*V1References, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *V1Info) SetReferences(v V1References)`

SetReferences sets References field to given value.

### HasReferences

`func (o *V1Info) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRegistry

`func (o *V1Info) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *V1Info) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *V1Info) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *V1Info) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetNamespace

`func (o *V1Info) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1Info) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1Info) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V1Info) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTag

`func (o *V1Info) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *V1Info) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *V1Info) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *V1Info) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetBuild

`func (o *V1Info) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *V1Info) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *V1Info) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *V1Info) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetTlsSupport

`func (o *V1Info) GetTlsSupport() bool`

GetTlsSupport returns the TlsSupport field if non-nil, zero value otherwise.

### GetTlsSupportOk

`func (o *V1Info) GetTlsSupportOk() (*bool, bool)`

GetTlsSupportOk returns a tuple with the TlsSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsSupport

`func (o *V1Info) SetTlsSupport(v bool)`

SetTlsSupport sets TlsSupport field to given value.

### HasTlsSupport

`func (o *V1Info) HasTlsSupport() bool`

HasTlsSupport returns a boolean if a field has been set.

### GetCapabilities

`func (o *V1Info) GetCapabilities() map[string]string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *V1Info) GetCapabilitiesOk() (*map[string]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *V1Info) SetCapabilities(v map[string]string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *V1Info) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetAdditionalImages

`func (o *V1Info) GetAdditionalImages() map[string]string`

GetAdditionalImages returns the AdditionalImages field if non-nil, zero value otherwise.

### GetAdditionalImagesOk

`func (o *V1Info) GetAdditionalImagesOk() (*map[string]string, bool)`

GetAdditionalImagesOk returns a tuple with the AdditionalImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalImages

`func (o *V1Info) SetAdditionalImages(v map[string]string)`

SetAdditionalImages sets AdditionalImages field to given value.

### HasAdditionalImages

`func (o *V1Info) HasAdditionalImages() bool`

HasAdditionalImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


