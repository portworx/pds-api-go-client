/*
PDS API

Portworx Data Services API Server

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"encoding/json"
)

// checks if the ControllersCreateTenantBackupTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControllersCreateTenantBackupTarget{}

// ControllersCreateTenantBackupTarget struct for ControllersCreateTenantBackupTarget
type ControllersCreateTenantBackupTarget struct {
	BackupCredentialsId *string `json:"backup_credentials_id,omitempty"`
	// Bucket name for S3 or S3 compatible. Container name for Azure.
	Bucket *string `json:"bucket,omitempty"`
	// Name of the backup target. Must be unique for the given tenant.
	Name *string `json:"name,omitempty"`
	// Region of the bucket. Required for S3. Otherwise must be empty.
	Region *string `json:"region,omitempty"`
	// Type of the backup target. Must match the used backup credentials.
	Type *string `json:"type,omitempty"`
}

// NewControllersCreateTenantBackupTarget instantiates a new ControllersCreateTenantBackupTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersCreateTenantBackupTarget() *ControllersCreateTenantBackupTarget {
	this := ControllersCreateTenantBackupTarget{}
	return &this
}

// NewControllersCreateTenantBackupTargetWithDefaults instantiates a new ControllersCreateTenantBackupTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersCreateTenantBackupTargetWithDefaults() *ControllersCreateTenantBackupTarget {
	this := ControllersCreateTenantBackupTarget{}
	return &this
}

// GetBackupCredentialsId returns the BackupCredentialsId field value if set, zero value otherwise.
func (o *ControllersCreateTenantBackupTarget) GetBackupCredentialsId() string {
	if o == nil || IsNil(o.BackupCredentialsId) {
		var ret string
		return ret
	}
	return *o.BackupCredentialsId
}

// GetBackupCredentialsIdOk returns a tuple with the BackupCredentialsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateTenantBackupTarget) GetBackupCredentialsIdOk() (*string, bool) {
	if o == nil || IsNil(o.BackupCredentialsId) {
		return nil, false
	}
	return o.BackupCredentialsId, true
}

// HasBackupCredentialsId returns a boolean if a field has been set.
func (o *ControllersCreateTenantBackupTarget) HasBackupCredentialsId() bool {
	if o != nil && !IsNil(o.BackupCredentialsId) {
		return true
	}

	return false
}

// SetBackupCredentialsId gets a reference to the given string and assigns it to the BackupCredentialsId field.
func (o *ControllersCreateTenantBackupTarget) SetBackupCredentialsId(v string) {
	o.BackupCredentialsId = &v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *ControllersCreateTenantBackupTarget) GetBucket() string {
	if o == nil || IsNil(o.Bucket) {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateTenantBackupTarget) GetBucketOk() (*string, bool) {
	if o == nil || IsNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *ControllersCreateTenantBackupTarget) HasBucket() bool {
	if o != nil && !IsNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *ControllersCreateTenantBackupTarget) SetBucket(v string) {
	o.Bucket = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ControllersCreateTenantBackupTarget) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateTenantBackupTarget) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ControllersCreateTenantBackupTarget) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ControllersCreateTenantBackupTarget) SetName(v string) {
	o.Name = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ControllersCreateTenantBackupTarget) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateTenantBackupTarget) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ControllersCreateTenantBackupTarget) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ControllersCreateTenantBackupTarget) SetRegion(v string) {
	o.Region = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ControllersCreateTenantBackupTarget) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateTenantBackupTarget) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ControllersCreateTenantBackupTarget) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ControllersCreateTenantBackupTarget) SetType(v string) {
	o.Type = &v
}

func (o ControllersCreateTenantBackupTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControllersCreateTenantBackupTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackupCredentialsId) {
		toSerialize["backup_credentials_id"] = o.BackupCredentialsId
	}
	if !IsNil(o.Bucket) {
		toSerialize["bucket"] = o.Bucket
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableControllersCreateTenantBackupTarget struct {
	value *ControllersCreateTenantBackupTarget
	isSet bool
}

func (v NullableControllersCreateTenantBackupTarget) Get() *ControllersCreateTenantBackupTarget {
	return v.value
}

func (v *NullableControllersCreateTenantBackupTarget) Set(val *ControllersCreateTenantBackupTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersCreateTenantBackupTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersCreateTenantBackupTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersCreateTenantBackupTarget(val *ControllersCreateTenantBackupTarget) *NullableControllersCreateTenantBackupTarget {
	return &NullableControllersCreateTenantBackupTarget{value: val, isSet: true}
}

func (v NullableControllersCreateTenantBackupTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersCreateTenantBackupTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


