/*
public/portworx/pds/tasks/apiv1/tasks.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pdsclient

import (
	"encoding/json"
	"time"
)

// checks if the MetadataOfTheDeploymentResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataOfTheDeploymentResource{}

// MetadataOfTheDeploymentResource struct for MetadataOfTheDeploymentResource
type MetadataOfTheDeploymentResource struct {
	// Name of the resource.
	Name *string `json:"name,omitempty"`
	// Description of the resource.
	Description *string `json:"description,omitempty"`
	// A string that identifies the version of this object that can be used by clients to determine when objects have changed. This value must be passed unmodified back to the server by the client.
	ResourceVersion *string `json:"resourceVersion,omitempty"`
	// Creation time of the object.
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Update time of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	// Labels to apply to the object.
	Labels *map[string]string `json:"labels,omitempty"`
	// Annotations for the object.
	Annotations *map[string]string `json:"annotations,omitempty"`
	ParentReference *V1Reference `json:"parentReference,omitempty"`
}

// NewMetadataOfTheDeploymentResource instantiates a new MetadataOfTheDeploymentResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataOfTheDeploymentResource() *MetadataOfTheDeploymentResource {
	this := MetadataOfTheDeploymentResource{}
	return &this
}

// NewMetadataOfTheDeploymentResourceWithDefaults instantiates a new MetadataOfTheDeploymentResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataOfTheDeploymentResourceWithDefaults() *MetadataOfTheDeploymentResource {
	this := MetadataOfTheDeploymentResource{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MetadataOfTheDeploymentResource) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataOfTheDeploymentResource) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MetadataOfTheDeploymentResource) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MetadataOfTheDeploymentResource) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MetadataOfTheDeploymentResource) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataOfTheDeploymentResource) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MetadataOfTheDeploymentResource) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MetadataOfTheDeploymentResource) SetDescription(v string) {
	o.Description = &v
}

// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *MetadataOfTheDeploymentResource) GetResourceVersion() string {
	if o == nil || IsNil(o.ResourceVersion) {
		var ret string
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataOfTheDeploymentResource) GetResourceVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceVersion) {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *MetadataOfTheDeploymentResource) HasResourceVersion() bool {
	if o != nil && !IsNil(o.ResourceVersion) {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given string and assigns it to the ResourceVersion field.
func (o *MetadataOfTheDeploymentResource) SetResourceVersion(v string) {
	o.ResourceVersion = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *MetadataOfTheDeploymentResource) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataOfTheDeploymentResource) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *MetadataOfTheDeploymentResource) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *MetadataOfTheDeploymentResource) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *MetadataOfTheDeploymentResource) GetUpdateTime() time.Time {
	if o == nil || IsNil(o.UpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataOfTheDeploymentResource) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *MetadataOfTheDeploymentResource) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *MetadataOfTheDeploymentResource) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *MetadataOfTheDeploymentResource) GetLabels() map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataOfTheDeploymentResource) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *MetadataOfTheDeploymentResource) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *MetadataOfTheDeploymentResource) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *MetadataOfTheDeploymentResource) GetAnnotations() map[string]string {
	if o == nil || IsNil(o.Annotations) {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataOfTheDeploymentResource) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *MetadataOfTheDeploymentResource) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *MetadataOfTheDeploymentResource) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

// GetParentReference returns the ParentReference field value if set, zero value otherwise.
func (o *MetadataOfTheDeploymentResource) GetParentReference() V1Reference {
	if o == nil || IsNil(o.ParentReference) {
		var ret V1Reference
		return ret
	}
	return *o.ParentReference
}

// GetParentReferenceOk returns a tuple with the ParentReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataOfTheDeploymentResource) GetParentReferenceOk() (*V1Reference, bool) {
	if o == nil || IsNil(o.ParentReference) {
		return nil, false
	}
	return o.ParentReference, true
}

// HasParentReference returns a boolean if a field has been set.
func (o *MetadataOfTheDeploymentResource) HasParentReference() bool {
	if o != nil && !IsNil(o.ParentReference) {
		return true
	}

	return false
}

// SetParentReference gets a reference to the given V1Reference and assigns it to the ParentReference field.
func (o *MetadataOfTheDeploymentResource) SetParentReference(v V1Reference) {
	o.ParentReference = &v
}

func (o MetadataOfTheDeploymentResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataOfTheDeploymentResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ResourceVersion) {
		toSerialize["resourceVersion"] = o.ResourceVersion
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !IsNil(o.ParentReference) {
		toSerialize["parentReference"] = o.ParentReference
	}
	return toSerialize, nil
}

type NullableMetadataOfTheDeploymentResource struct {
	value *MetadataOfTheDeploymentResource
	isSet bool
}

func (v NullableMetadataOfTheDeploymentResource) Get() *MetadataOfTheDeploymentResource {
	return v.value
}

func (v *NullableMetadataOfTheDeploymentResource) Set(val *MetadataOfTheDeploymentResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataOfTheDeploymentResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataOfTheDeploymentResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataOfTheDeploymentResource(val *MetadataOfTheDeploymentResource) *NullableMetadataOfTheDeploymentResource {
	return &NullableMetadataOfTheDeploymentResource{value: val, isSet: true}
}

func (v NullableMetadataOfTheDeploymentResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataOfTheDeploymentResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


