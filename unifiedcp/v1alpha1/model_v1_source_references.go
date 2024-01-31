/*
public/portworx/pds/tasks/apiv1/tasks.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pdsclient

import (
	"encoding/json"
)

// checks if the V1SourceReferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1SourceReferences{}

// V1SourceReferences SourceReferences for the restore.
type V1SourceReferences struct {
	// UID of the deployment which was backed up.
	DeploymentId *string `json:"deploymentId,omitempty"`
	// UID of the backup.
	BackupId *string `json:"backupId,omitempty"`
	// UID of the backup location.
	BackupLocationId *string `json:"backupLocationId,omitempty"`
	// UID of the cloud snapshot of the backup volume used for restore.
	CloudsnapId *string `json:"cloudsnapId,omitempty"`
}

// NewV1SourceReferences instantiates a new V1SourceReferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1SourceReferences() *V1SourceReferences {
	this := V1SourceReferences{}
	return &this
}

// NewV1SourceReferencesWithDefaults instantiates a new V1SourceReferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1SourceReferencesWithDefaults() *V1SourceReferences {
	this := V1SourceReferences{}
	return &this
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise.
func (o *V1SourceReferences) GetDeploymentId() string {
	if o == nil || IsNil(o.DeploymentId) {
		var ret string
		return ret
	}
	return *o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SourceReferences) GetDeploymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeploymentId) {
		return nil, false
	}
	return o.DeploymentId, true
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *V1SourceReferences) HasDeploymentId() bool {
	if o != nil && !IsNil(o.DeploymentId) {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given string and assigns it to the DeploymentId field.
func (o *V1SourceReferences) SetDeploymentId(v string) {
	o.DeploymentId = &v
}

// GetBackupId returns the BackupId field value if set, zero value otherwise.
func (o *V1SourceReferences) GetBackupId() string {
	if o == nil || IsNil(o.BackupId) {
		var ret string
		return ret
	}
	return *o.BackupId
}

// GetBackupIdOk returns a tuple with the BackupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SourceReferences) GetBackupIdOk() (*string, bool) {
	if o == nil || IsNil(o.BackupId) {
		return nil, false
	}
	return o.BackupId, true
}

// HasBackupId returns a boolean if a field has been set.
func (o *V1SourceReferences) HasBackupId() bool {
	if o != nil && !IsNil(o.BackupId) {
		return true
	}

	return false
}

// SetBackupId gets a reference to the given string and assigns it to the BackupId field.
func (o *V1SourceReferences) SetBackupId(v string) {
	o.BackupId = &v
}

// GetBackupLocationId returns the BackupLocationId field value if set, zero value otherwise.
func (o *V1SourceReferences) GetBackupLocationId() string {
	if o == nil || IsNil(o.BackupLocationId) {
		var ret string
		return ret
	}
	return *o.BackupLocationId
}

// GetBackupLocationIdOk returns a tuple with the BackupLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SourceReferences) GetBackupLocationIdOk() (*string, bool) {
	if o == nil || IsNil(o.BackupLocationId) {
		return nil, false
	}
	return o.BackupLocationId, true
}

// HasBackupLocationId returns a boolean if a field has been set.
func (o *V1SourceReferences) HasBackupLocationId() bool {
	if o != nil && !IsNil(o.BackupLocationId) {
		return true
	}

	return false
}

// SetBackupLocationId gets a reference to the given string and assigns it to the BackupLocationId field.
func (o *V1SourceReferences) SetBackupLocationId(v string) {
	o.BackupLocationId = &v
}

// GetCloudsnapId returns the CloudsnapId field value if set, zero value otherwise.
func (o *V1SourceReferences) GetCloudsnapId() string {
	if o == nil || IsNil(o.CloudsnapId) {
		var ret string
		return ret
	}
	return *o.CloudsnapId
}

// GetCloudsnapIdOk returns a tuple with the CloudsnapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SourceReferences) GetCloudsnapIdOk() (*string, bool) {
	if o == nil || IsNil(o.CloudsnapId) {
		return nil, false
	}
	return o.CloudsnapId, true
}

// HasCloudsnapId returns a boolean if a field has been set.
func (o *V1SourceReferences) HasCloudsnapId() bool {
	if o != nil && !IsNil(o.CloudsnapId) {
		return true
	}

	return false
}

// SetCloudsnapId gets a reference to the given string and assigns it to the CloudsnapId field.
func (o *V1SourceReferences) SetCloudsnapId(v string) {
	o.CloudsnapId = &v
}

func (o V1SourceReferences) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1SourceReferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeploymentId) {
		toSerialize["deploymentId"] = o.DeploymentId
	}
	if !IsNil(o.BackupId) {
		toSerialize["backupId"] = o.BackupId
	}
	if !IsNil(o.BackupLocationId) {
		toSerialize["backupLocationId"] = o.BackupLocationId
	}
	if !IsNil(o.CloudsnapId) {
		toSerialize["cloudsnapId"] = o.CloudsnapId
	}
	return toSerialize, nil
}

type NullableV1SourceReferences struct {
	value *V1SourceReferences
	isSet bool
}

func (v NullableV1SourceReferences) Get() *V1SourceReferences {
	return v.value
}

func (v *NullableV1SourceReferences) Set(val *V1SourceReferences) {
	v.value = val
	v.isSet = true
}

func (v NullableV1SourceReferences) IsSet() bool {
	return v.isSet
}

func (v *NullableV1SourceReferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1SourceReferences(val *V1SourceReferences) *NullableV1SourceReferences {
	return &NullableV1SourceReferences{value: val, isSet: true}
}

func (v NullableV1SourceReferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1SourceReferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


