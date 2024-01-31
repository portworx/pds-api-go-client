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

// checks if the GooglerpcStatus1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GooglerpcStatus1{}

// GooglerpcStatus1 struct for GooglerpcStatus1
type GooglerpcStatus1 struct {
	Code *int32 `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Details []ProtobufAny2 `json:"details,omitempty"`
}

// NewGooglerpcStatus1 instantiates a new GooglerpcStatus1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGooglerpcStatus1() *GooglerpcStatus1 {
	this := GooglerpcStatus1{}
	return &this
}

// NewGooglerpcStatus1WithDefaults instantiates a new GooglerpcStatus1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGooglerpcStatus1WithDefaults() *GooglerpcStatus1 {
	this := GooglerpcStatus1{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GooglerpcStatus1) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GooglerpcStatus1) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GooglerpcStatus1) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *GooglerpcStatus1) SetCode(v int32) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GooglerpcStatus1) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GooglerpcStatus1) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GooglerpcStatus1) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GooglerpcStatus1) SetMessage(v string) {
	o.Message = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *GooglerpcStatus1) GetDetails() []ProtobufAny2 {
	if o == nil || IsNil(o.Details) {
		var ret []ProtobufAny2
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GooglerpcStatus1) GetDetailsOk() ([]ProtobufAny2, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *GooglerpcStatus1) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []ProtobufAny2 and assigns it to the Details field.
func (o *GooglerpcStatus1) SetDetails(v []ProtobufAny2) {
	o.Details = v
}

func (o GooglerpcStatus1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GooglerpcStatus1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableGooglerpcStatus1 struct {
	value *GooglerpcStatus1
	isSet bool
}

func (v NullableGooglerpcStatus1) Get() *GooglerpcStatus1 {
	return v.value
}

func (v *NullableGooglerpcStatus1) Set(val *GooglerpcStatus1) {
	v.value = val
	v.isSet = true
}

func (v NullableGooglerpcStatus1) IsSet() bool {
	return v.isSet
}

func (v *NullableGooglerpcStatus1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGooglerpcStatus1(val *GooglerpcStatus1) *NullableGooglerpcStatus1 {
	return &NullableGooglerpcStatus1{value: val, isSet: true}
}

func (v NullableGooglerpcStatus1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGooglerpcStatus1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


