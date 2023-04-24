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

// checks if the ControllersAPIVersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControllersAPIVersionResponse{}

// ControllersAPIVersionResponse struct for ControllersAPIVersionResponse
type ControllersAPIVersionResponse struct {
	ApiVersion *string `json:"api_version,omitempty"`
	Features *map[string]string `json:"features,omitempty"`
	HelmChartVersion *string `json:"helm_chart_version,omitempty"`
	PdsBuildNumber *int32 `json:"pds_build_number,omitempty"`
}

// NewControllersAPIVersionResponse instantiates a new ControllersAPIVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersAPIVersionResponse() *ControllersAPIVersionResponse {
	this := ControllersAPIVersionResponse{}
	return &this
}

// NewControllersAPIVersionResponseWithDefaults instantiates a new ControllersAPIVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersAPIVersionResponseWithDefaults() *ControllersAPIVersionResponse {
	this := ControllersAPIVersionResponse{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ControllersAPIVersionResponse) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersAPIVersionResponse) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ControllersAPIVersionResponse) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ControllersAPIVersionResponse) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ControllersAPIVersionResponse) GetFeatures() map[string]string {
	if o == nil || IsNil(o.Features) {
		var ret map[string]string
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersAPIVersionResponse) GetFeaturesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ControllersAPIVersionResponse) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given map[string]string and assigns it to the Features field.
func (o *ControllersAPIVersionResponse) SetFeatures(v map[string]string) {
	o.Features = &v
}

// GetHelmChartVersion returns the HelmChartVersion field value if set, zero value otherwise.
func (o *ControllersAPIVersionResponse) GetHelmChartVersion() string {
	if o == nil || IsNil(o.HelmChartVersion) {
		var ret string
		return ret
	}
	return *o.HelmChartVersion
}

// GetHelmChartVersionOk returns a tuple with the HelmChartVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersAPIVersionResponse) GetHelmChartVersionOk() (*string, bool) {
	if o == nil || IsNil(o.HelmChartVersion) {
		return nil, false
	}
	return o.HelmChartVersion, true
}

// HasHelmChartVersion returns a boolean if a field has been set.
func (o *ControllersAPIVersionResponse) HasHelmChartVersion() bool {
	if o != nil && !IsNil(o.HelmChartVersion) {
		return true
	}

	return false
}

// SetHelmChartVersion gets a reference to the given string and assigns it to the HelmChartVersion field.
func (o *ControllersAPIVersionResponse) SetHelmChartVersion(v string) {
	o.HelmChartVersion = &v
}

// GetPdsBuildNumber returns the PdsBuildNumber field value if set, zero value otherwise.
func (o *ControllersAPIVersionResponse) GetPdsBuildNumber() int32 {
	if o == nil || IsNil(o.PdsBuildNumber) {
		var ret int32
		return ret
	}
	return *o.PdsBuildNumber
}

// GetPdsBuildNumberOk returns a tuple with the PdsBuildNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersAPIVersionResponse) GetPdsBuildNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PdsBuildNumber) {
		return nil, false
	}
	return o.PdsBuildNumber, true
}

// HasPdsBuildNumber returns a boolean if a field has been set.
func (o *ControllersAPIVersionResponse) HasPdsBuildNumber() bool {
	if o != nil && !IsNil(o.PdsBuildNumber) {
		return true
	}

	return false
}

// SetPdsBuildNumber gets a reference to the given int32 and assigns it to the PdsBuildNumber field.
func (o *ControllersAPIVersionResponse) SetPdsBuildNumber(v int32) {
	o.PdsBuildNumber = &v
}

func (o ControllersAPIVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControllersAPIVersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiVersion) {
		toSerialize["api_version"] = o.ApiVersion
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.HelmChartVersion) {
		toSerialize["helm_chart_version"] = o.HelmChartVersion
	}
	if !IsNil(o.PdsBuildNumber) {
		toSerialize["pds_build_number"] = o.PdsBuildNumber
	}
	return toSerialize, nil
}

type NullableControllersAPIVersionResponse struct {
	value *ControllersAPIVersionResponse
	isSet bool
}

func (v NullableControllersAPIVersionResponse) Get() *ControllersAPIVersionResponse {
	return v.value
}

func (v *NullableControllersAPIVersionResponse) Set(val *ControllersAPIVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersAPIVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersAPIVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersAPIVersionResponse(val *ControllersAPIVersionResponse) *NullableControllersAPIVersionResponse {
	return &NullableControllersAPIVersionResponse{value: val, isSet: true}
}

func (v NullableControllersAPIVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersAPIVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


