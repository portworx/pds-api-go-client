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

// ModelsStorageOptionsSample struct for ModelsStorageOptionsSample
type ModelsStorageOptionsSample struct {
	// This option enforces volume group policy. If a volume belonging to a group cannot find nodes for its replication sets which don’t have other volumes of same group, the volume creation will fail.
	Fg *bool `json:"fg,omitempty"`
	// Filesystem to be laid out.
	Fs *string `json:"fs,omitempty"`
	// Name of the template. Must be unique within the tenant scope.
	Name *string `json:"name,omitempty"`
	// Replication factor for the volume.
	Repl *int32 `json:"repl,omitempty"`
	// Flag to create an encrypted volume. Currently, not supported (should be set to `false`).
	Secure *bool `json:"secure,omitempty"`
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

// GetFg returns the Fg field value if set, zero value otherwise.
func (o *ModelsStorageOptionsSample) GetFg() bool {
	if o == nil || o.Fg == nil {
		var ret bool
		return ret
	}
	return *o.Fg
}

// GetFgOk returns a tuple with the Fg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsStorageOptionsSample) GetFgOk() (*bool, bool) {
	if o == nil || o.Fg == nil {
		return nil, false
	}
	return o.Fg, true
}

// HasFg returns a boolean if a field has been set.
func (o *ModelsStorageOptionsSample) HasFg() bool {
	if o != nil && o.Fg != nil {
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
	if o == nil || o.Fs == nil {
		var ret string
		return ret
	}
	return *o.Fs
}

// GetFsOk returns a tuple with the Fs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsStorageOptionsSample) GetFsOk() (*string, bool) {
	if o == nil || o.Fs == nil {
		return nil, false
	}
	return o.Fs, true
}

// HasFs returns a boolean if a field has been set.
func (o *ModelsStorageOptionsSample) HasFs() bool {
	if o != nil && o.Fs != nil {
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
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsStorageOptionsSample) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelsStorageOptionsSample) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelsStorageOptionsSample) SetName(v string) {
	o.Name = &v
}

// GetRepl returns the Repl field value if set, zero value otherwise.
func (o *ModelsStorageOptionsSample) GetRepl() int32 {
	if o == nil || o.Repl == nil {
		var ret int32
		return ret
	}
	return *o.Repl
}

// GetReplOk returns a tuple with the Repl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsStorageOptionsSample) GetReplOk() (*int32, bool) {
	if o == nil || o.Repl == nil {
		return nil, false
	}
	return o.Repl, true
}

// HasRepl returns a boolean if a field has been set.
func (o *ModelsStorageOptionsSample) HasRepl() bool {
	if o != nil && o.Repl != nil {
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
	if o == nil || o.Secure == nil {
		var ret bool
		return ret
	}
	return *o.Secure
}

// GetSecureOk returns a tuple with the Secure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsStorageOptionsSample) GetSecureOk() (*bool, bool) {
	if o == nil || o.Secure == nil {
		return nil, false
	}
	return o.Secure, true
}

// HasSecure returns a boolean if a field has been set.
func (o *ModelsStorageOptionsSample) HasSecure() bool {
	if o != nil && o.Secure != nil {
		return true
	}

	return false
}

// SetSecure gets a reference to the given bool and assigns it to the Secure field.
func (o *ModelsStorageOptionsSample) SetSecure(v bool) {
	o.Secure = &v
}

func (o ModelsStorageOptionsSample) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fg != nil {
		toSerialize["fg"] = o.Fg
	}
	if o.Fs != nil {
		toSerialize["fs"] = o.Fs
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Repl != nil {
		toSerialize["repl"] = o.Repl
	}
	if o.Secure != nil {
		toSerialize["secure"] = o.Secure
	}
	return json.Marshal(toSerialize)
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


