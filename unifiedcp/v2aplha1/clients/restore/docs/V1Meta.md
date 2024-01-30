# V1Meta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** | UID of the resource of the format &lt;resource prefix&gt;-&lt;uuid&gt;. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**Description** | Pointer to **string** | Description of the resource. | [optional] 
**ResourceVersion** | Pointer to **string** | A string that identifies the version of this object that can be used by clients to determine when objects have changed. This value must be passed unmodified back to the server by the client. | [optional] 
**CreateTime** | Pointer to **time.Time** | Creation time of the object. | [optional] [readonly] 
**UpdateTime** | Pointer to **time.Time** | Update time of the object. | [optional] [readonly] 
**Labels** | Pointer to **map[string]string** | Labels to apply to the object. | [optional] 
**Annotations** | Pointer to **map[string]string** | Annotations for the object. | [optional] 
**ParentReference** | Pointer to [**V1Reference**](V1Reference.md) |  | [optional] 

## Methods

### NewV1Meta

`func NewV1Meta() *V1Meta`

NewV1Meta instantiates a new V1Meta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MetaWithDefaults

`func NewV1MetaWithDefaults() *V1Meta`

NewV1MetaWithDefaults instantiates a new V1Meta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *V1Meta) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *V1Meta) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *V1Meta) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *V1Meta) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *V1Meta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1Meta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1Meta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1Meta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *V1Meta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1Meta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1Meta) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1Meta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResourceVersion

`func (o *V1Meta) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *V1Meta) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *V1Meta) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *V1Meta) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetCreateTime

`func (o *V1Meta) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *V1Meta) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *V1Meta) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *V1Meta) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *V1Meta) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *V1Meta) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *V1Meta) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *V1Meta) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetLabels

`func (o *V1Meta) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *V1Meta) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *V1Meta) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *V1Meta) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *V1Meta) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *V1Meta) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *V1Meta) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *V1Meta) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetParentReference

`func (o *V1Meta) GetParentReference() V1Reference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *V1Meta) GetParentReferenceOk() (*V1Reference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *V1Meta) SetParentReference(v V1Reference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *V1Meta) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


