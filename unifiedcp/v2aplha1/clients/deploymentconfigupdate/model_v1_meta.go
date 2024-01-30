/*
public/portworx/pds/deploymentconfigupdate/apiv1/deploymentconfigupdate.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package deploymentconfigupdate

import (
	"encoding/json"
	"time"
)

// checks if the V1Meta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1Meta{}

// V1Meta Meta holds general resource metadata.
type V1Meta struct {
	// UID of the resource of the format <resource prefix>-<uuid>.
	Uid *string `json:"uid,omitempty"`
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

// NewV1Meta instantiates a new V1Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Meta() *V1Meta {
	this := V1Meta{}
	return &this
}

// NewV1MetaWithDefaults instantiates a new V1Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1MetaWithDefaults() *V1Meta {
	this := V1Meta{}
	return &this
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *V1Meta) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Meta) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *V1Meta) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *V1Meta) SetUid(v string) {
	o.Uid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1Meta) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Meta) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1Meta) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1Meta) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V1Meta) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Meta) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V1Meta) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V1Meta) SetDescription(v string) {
	o.Description = &v
}

// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *V1Meta) GetResourceVersion() string {
	if o == nil || IsNil(o.ResourceVersion) {
		var ret string
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Meta) GetResourceVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceVersion) {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *V1Meta) HasResourceVersion() bool {
	if o != nil && !IsNil(o.ResourceVersion) {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given string and assigns it to the ResourceVersion field.
func (o *V1Meta) SetResourceVersion(v string) {
	o.ResourceVersion = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *V1Meta) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Meta) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *V1Meta) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *V1Meta) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *V1Meta) GetUpdateTime() time.Time {
	if o == nil || IsNil(o.UpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Meta) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *V1Meta) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *V1Meta) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *V1Meta) GetLabels() map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Meta) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *V1Meta) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *V1Meta) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *V1Meta) GetAnnotations() map[string]string {
	if o == nil || IsNil(o.Annotations) {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Meta) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *V1Meta) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *V1Meta) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

// GetParentReference returns the ParentReference field value if set, zero value otherwise.
func (o *V1Meta) GetParentReference() V1Reference {
	if o == nil || IsNil(o.ParentReference) {
		var ret V1Reference
		return ret
	}
	return *o.ParentReference
}

// GetParentReferenceOk returns a tuple with the ParentReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Meta) GetParentReferenceOk() (*V1Reference, bool) {
	if o == nil || IsNil(o.ParentReference) {
		return nil, false
	}
	return o.ParentReference, true
}

// HasParentReference returns a boolean if a field has been set.
func (o *V1Meta) HasParentReference() bool {
	if o != nil && !IsNil(o.ParentReference) {
		return true
	}

	return false
}

// SetParentReference gets a reference to the given V1Reference and assigns it to the ParentReference field.
func (o *V1Meta) SetParentReference(v V1Reference) {
	o.ParentReference = &v
}

func (o V1Meta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1Meta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
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

type NullableV1Meta struct {
	value *V1Meta
	isSet bool
}

func (v NullableV1Meta) Get() *V1Meta {
	return v.value
}

func (v *NullableV1Meta) Set(val *V1Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Meta) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Meta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Meta(val *V1Meta) *NullableV1Meta {
	return &NullableV1Meta{value: val, isSet: true}
}

func (v NullableV1Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Meta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


