# MetadataForIAM

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

### NewMetadataForIAM

`func NewMetadataForIAM() *MetadataForIAM`

NewMetadataForIAM instantiates a new MetadataForIAM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataForIAMWithDefaults

`func NewMetadataForIAMWithDefaults() *MetadataForIAM`

NewMetadataForIAMWithDefaults instantiates a new MetadataForIAM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetadataForIAM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataForIAM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataForIAM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetadataForIAM) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *MetadataForIAM) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetadataForIAM) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetadataForIAM) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetadataForIAM) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResourceVersion

`func (o *MetadataForIAM) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *MetadataForIAM) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *MetadataForIAM) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *MetadataForIAM) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetCreateTime

`func (o *MetadataForIAM) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *MetadataForIAM) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *MetadataForIAM) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *MetadataForIAM) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *MetadataForIAM) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *MetadataForIAM) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *MetadataForIAM) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *MetadataForIAM) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetLabels

`func (o *MetadataForIAM) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MetadataForIAM) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MetadataForIAM) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *MetadataForIAM) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *MetadataForIAM) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *MetadataForIAM) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *MetadataForIAM) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *MetadataForIAM) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetParentReference

`func (o *MetadataForIAM) GetParentReference() V1Reference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *MetadataForIAM) GetParentReferenceOk() (*V1Reference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *MetadataForIAM) SetParentReference(v V1Reference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *MetadataForIAM) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


