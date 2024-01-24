# MetadataOfTheTargetClusterResourceMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**Description** | Pointer to **string** | Description of the resource. | [optional] 
**ResourceVersion** | Pointer to **string** | A string that identifies the version of this object that can be used by clients to determine when objects have changed. This value must be passed unmodified back to the server by the client. | [optional] 
**CreateTime** | Pointer to **time.Time** | Creation time of the object. | [optional] [readonly] 
**UpdateTime** | Pointer to **time.Time** | Update time of the object. | [optional] [readonly] 
**Labels** | Pointer to **map[string]string** | Labels to apply to the object. | [optional] 
**Annotations** | Pointer to **map[string]string** | Annotations for the object. | [optional] 
**ParentReference** | Pointer to [**V1Reference**](V1Reference.md) |  | [optional] 

## Methods

### NewMetadataOfTheTargetClusterResourceMeta

`func NewMetadataOfTheTargetClusterResourceMeta() *MetadataOfTheTargetClusterResourceMeta`

NewMetadataOfTheTargetClusterResourceMeta instantiates a new MetadataOfTheTargetClusterResourceMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataOfTheTargetClusterResourceMetaWithDefaults

`func NewMetadataOfTheTargetClusterResourceMetaWithDefaults() *MetadataOfTheTargetClusterResourceMeta`

NewMetadataOfTheTargetClusterResourceMetaWithDefaults instantiates a new MetadataOfTheTargetClusterResourceMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetadataOfTheTargetClusterResourceMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataOfTheTargetClusterResourceMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataOfTheTargetClusterResourceMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetadataOfTheTargetClusterResourceMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *MetadataOfTheTargetClusterResourceMeta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetadataOfTheTargetClusterResourceMeta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetadataOfTheTargetClusterResourceMeta) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetadataOfTheTargetClusterResourceMeta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResourceVersion

`func (o *MetadataOfTheTargetClusterResourceMeta) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *MetadataOfTheTargetClusterResourceMeta) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *MetadataOfTheTargetClusterResourceMeta) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *MetadataOfTheTargetClusterResourceMeta) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetCreateTime

`func (o *MetadataOfTheTargetClusterResourceMeta) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *MetadataOfTheTargetClusterResourceMeta) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *MetadataOfTheTargetClusterResourceMeta) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *MetadataOfTheTargetClusterResourceMeta) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *MetadataOfTheTargetClusterResourceMeta) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *MetadataOfTheTargetClusterResourceMeta) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *MetadataOfTheTargetClusterResourceMeta) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *MetadataOfTheTargetClusterResourceMeta) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetLabels

`func (o *MetadataOfTheTargetClusterResourceMeta) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MetadataOfTheTargetClusterResourceMeta) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MetadataOfTheTargetClusterResourceMeta) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *MetadataOfTheTargetClusterResourceMeta) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *MetadataOfTheTargetClusterResourceMeta) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *MetadataOfTheTargetClusterResourceMeta) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *MetadataOfTheTargetClusterResourceMeta) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *MetadataOfTheTargetClusterResourceMeta) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetParentReference

`func (o *MetadataOfTheTargetClusterResourceMeta) GetParentReference() V1Reference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *MetadataOfTheTargetClusterResourceMeta) GetParentReferenceOk() (*V1Reference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *MetadataOfTheTargetClusterResourceMeta) SetParentReference(v V1Reference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *MetadataOfTheTargetClusterResourceMeta) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


