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

// checks if the ModelsStorageOptionsSample type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsStorageOptionsSample{}

// ModelsStorageOptionsSample struct for ModelsStorageOptionsSample
type ModelsStorageOptionsSample struct {
	Created *string `json:"created,omitempty"`
	// This option enforces volume group policy. If a volume belonging to a group cannot find nodes for its replication sets which don’t have other volumes of same group, the volume creation will fail.
	Fg *bool `json:"fg,omitempty"`
	// Filesystem to be laid out.
	Fs *string `json:"fs,omitempty"`
	// Name of the template. Must be unique within the tenant scope.
	Name *string `json:"name,omitempty"`
	// Portworx volume provisioner. Valid values are: \"pxd.portworx.com\" for PX CSI, \"kubernetes.io/portworx-volume\" for PX in-tree or \"auto\" for auto-detect
	Provisioner *string `json:"provisioner,omitempty"`
	// Replication factor for the volume.
	Repl *int32 `json:"repl,omitempty"`
	// Flag to create an encrypted volume. Currently, not supported (should be set to `false`).
	Secure *bool `json:"secure,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewModelsStorageOptionsSample instantiates a new ModelsStorageOptionsSample object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsStorageOptionsSample() *ModelsStorageOptionsSample {
	this := ModelsStorageOptionsSample{}
	return &this
}

// NewModelsStorageOptionsSampleWithDefaults instantiates a new ModelsStorageOptionsSample object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsStorageOptionsSampleWithDefaults() *ModelsStorageOptionsSample {
	this := ModelsStorageOptionsSample{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ModelsStorageOptionsSample) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsStorageOptionsSample) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ModelsStorageOptionsSample) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *ModelsStorageOptionsSample) SetCreated(v string) {
	o.Created = &v
}

// GetFg returns the Fg field value if set, zero value otherwise.
func (o *ModelsStorageOptionsSample) GetFg() bool {
	if o == nil || IsNil(o.Fg) {
		var ret bool
		return ret
	}
	return *o.Fg
}

// GetFgOk returns a tuple with the Fg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsStorageOptionsSample) GetFgOk() (*bool, bool) {
	if o == nil || IsNil(o.Fg) {
		return nil, false
	}
	return o.Fg, true
}

// HasFg returns a boolean if a field has been set.
func (o *ModelsStorageOptionsSample) HasFg() bool {
	if o != nil && !IsNil(o.Fg) {
		return true
	}

	return false
}

// SetFg gets a reference to the given bool and assigns it to the Fg field.
func (o *ModelsStorageOptionsSample) SetFg(v bool) {
	o.Fg = &v
}

// GetFs returns the Fs field value if set, zero value otherwise.
func (o *ModelsStorageOptionsSample) GetFs() string {
	if o == nil || IsNil(o.Fs) {
		var ret string
		return ret
	}
	return *o.Fs
}

// GetFsOk returns a tuple with the Fs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsStorageOptionsSample) GetFsOk() (*string, bool) {
	if o == nil || IsNil(o.Fs) {
		return nil, false
	}
	return o.Fs, true
}

// HasFs returns a boolean if a field has been set.
func (o *ModelsStorageOptionsSample) HasFs() bool {
	if o != nil && !IsNil(o.Fs) {
		return true
	}

	return false
}

// SetFs gets a reference to the given string and assigns it to the Fs field.
func (o *ModelsStorageOptionsSample) SetFs(v string) {
	o.Fs = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelsStorageOptionsSample) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsStorageOptionsSample) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelsStorageOptionsSample) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelsStorageOptionsSample) SetName(v string) {
	o.Name = &v
}

// GetProvisioner returns the Provisioner field value if set, zero value otherwise.
func (o *ModelsStorageOptionsSample) GetProvisioner() string {
	if o == nil || IsNil(o.Provisioner) {
		var ret string
		return ret
	}
	return *o.Provisioner
}

// GetProvisionerOk returns a tuple with the Provisioner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsStorageOptionsSample) GetProvisionerOk() (*string, bool) {
	if o == nil || IsNil(o.Provisioner) {
		return nil, false
	}
	return o.Provisioner, true
}

// HasProvisioner returns a boolean if a field has been set.
func (o *ModelsStorageOptionsSample) HasProvisioner() bool {
	if o != nil && !IsNil(o.Provisioner) {
		return true
	}

	return false
}

// SetProvisioner gets a reference to the given string and assigns it to the Provisioner field.
func (o *ModelsStorageOptionsSample) SetProvisioner(v string) {
	o.Provisioner = &v
}

// GetRepl returns the Repl field value if set, zero value otherwise.
func (o *ModelsStorageOptionsSample) GetRepl() int32 {
	if o == nil || IsNil(o.Repl) {
		var ret int32
		return ret
	}
	return *o.Repl
}

// GetReplOk returns a tuple with the Repl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsStorageOptionsSample) GetReplOk() (*int32, bool) {
	if o == nil || IsNil(o.Repl) {
		return nil, false
	}
	return o.Repl, true
}

// HasRepl returns a boolean if a field has been set.
func (o *ModelsStorageOptionsSample) HasRepl() bool {
	if o != nil && !IsNil(o.Repl) {
		return true
	}

	return false
}

// SetRepl gets a reference to the given int32 and assigns it to the Repl field.
func (o *ModelsStorageOptionsSample) SetRepl(v int32) {
	o.Repl = &v
}

// GetSecure returns the Secure field value if set, zero value otherwise.
func (o *ModelsStorageOptionsSample) GetSecure() bool {
	if o == nil || IsNil(o.Secure) {
		var ret bool
		return ret
	}
	return *o.Secure
}

// GetSecureOk returns a tuple with the Secure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsStorageOptionsSample) GetSecureOk() (*bool, bool) {
	if o == nil || IsNil(o.Secure) {
		return nil, false
	}
	return o.Secure, true
}

// HasSecure returns a boolean if a field has been set.
func (o *ModelsStorageOptionsSample) HasSecure() bool {
	if o != nil && !IsNil(o.Secure) {
		return true
	}

	return false
}

// SetSecure gets a reference to the given bool and assigns it to the Secure field.
func (o *ModelsStorageOptionsSample) SetSecure(v bool) {
	o.Secure = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ModelsStorageOptionsSample) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsStorageOptionsSample) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ModelsStorageOptionsSample) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *ModelsStorageOptionsSample) SetUpdated(v string) {
	o.Updated = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ModelsStorageOptionsSample) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsStorageOptionsSample) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ModelsStorageOptionsSample) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ModelsStorageOptionsSample) SetVersion(v int32) {
	o.Version = &v
}

func (o ModelsStorageOptionsSample) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsStorageOptionsSample) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Fg) {
		toSerialize["fg"] = o.Fg
	}
	if !IsNil(o.Fs) {
		toSerialize["fs"] = o.Fs
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Provisioner) {
		toSerialize["provisioner"] = o.Provisioner
	}
	if !IsNil(o.Repl) {
		toSerialize["repl"] = o.Repl
	}
	if !IsNil(o.Secure) {
		toSerialize["secure"] = o.Secure
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableModelsStorageOptionsSample struct {
	value *ModelsStorageOptionsSample
	isSet bool
}

func (v NullableModelsStorageOptionsSample) Get() *ModelsStorageOptionsSample {
	return v.value
}

func (v *NullableModelsStorageOptionsSample) Set(val *ModelsStorageOptionsSample) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsStorageOptionsSample) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsStorageOptionsSample) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsStorageOptionsSample(val *ModelsStorageOptionsSample) *NullableModelsStorageOptionsSample {
	return &NullableModelsStorageOptionsSample{value: val, isSet: true}
}

func (v NullableModelsStorageOptionsSample) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsStorageOptionsSample) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


