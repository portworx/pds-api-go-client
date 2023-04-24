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

// checks if the RequestsPutBackupJobRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestsPutBackupJobRequest{}

// RequestsPutBackupJobRequest struct for RequestsPutBackupJobRequest
type RequestsPutBackupJobRequest struct {
	// BackupID which backup created the snapshot (nullable).
	BackupId *string `json:"backup_id,omitempty"`
	// BackupSpecJSON is the specification of the Backup in JSON format at the time the snapshot was taken.
	BackupSpec []int32 `json:"backup_spec,omitempty"`
	// CloudCredentialName credentials to access snapshot.
	CloudCredentialName *string `json:"cloud_credential_name,omitempty"`
	// CloudSnapID snapshot of the backup volume.
	CloudSnapId *string `json:"cloud_snap_id,omitempty"`
	// CompletionStatus of the backup job.
	CompletionStatus *string `json:"completion_status,omitempty"`
	CompletionTime *string `json:"completion_time,omitempty"`
	// BackupSpecJSON is the specification of the Backup in JSON format at the time the snapshot was taken.
	DataServiceSpec []int32 `json:"data_service_spec,omitempty"`
	// DeploymentID which deployment was backed up (nullable).
	DeploymentId *string `json:"deployment_id,omitempty"`
	ErrorCode *string `json:"error_code,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	// FileSize of the CloudSnap image.
	FileSize *int32 `json:"file_size,omitempty"`
	// Name of the BackupJob.
	Name *string `json:"name,omitempty"`
	// ProjectID which created the snapshot (required).
	ProjectId string `json:"project_id"`
	StartTime *string `json:"start_time,omitempty"`
}

// NewRequestsPutBackupJobRequest instantiates a new RequestsPutBackupJobRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestsPutBackupJobRequest(projectId string) *RequestsPutBackupJobRequest {
	this := RequestsPutBackupJobRequest{}
	this.ProjectId = projectId
	return &this
}

// NewRequestsPutBackupJobRequestWithDefaults instantiates a new RequestsPutBackupJobRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestsPutBackupJobRequestWithDefaults() *RequestsPutBackupJobRequest {
	this := RequestsPutBackupJobRequest{}
	return &this
}

// GetBackupId returns the BackupId field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetBackupId() string {
	if o == nil || IsNil(o.BackupId) {
		var ret string
		return ret
	}
	return *o.BackupId
}

// GetBackupIdOk returns a tuple with the BackupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetBackupIdOk() (*string, bool) {
	if o == nil || IsNil(o.BackupId) {
		return nil, false
	}
	return o.BackupId, true
}

// HasBackupId returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasBackupId() bool {
	if o != nil && !IsNil(o.BackupId) {
		return true
	}

	return false
}

// SetBackupId gets a reference to the given string and assigns it to the BackupId field.
func (o *RequestsPutBackupJobRequest) SetBackupId(v string) {
	o.BackupId = &v
}

// GetBackupSpec returns the BackupSpec field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetBackupSpec() []int32 {
	if o == nil || IsNil(o.BackupSpec) {
		var ret []int32
		return ret
	}
	return o.BackupSpec
}

// GetBackupSpecOk returns a tuple with the BackupSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetBackupSpecOk() ([]int32, bool) {
	if o == nil || IsNil(o.BackupSpec) {
		return nil, false
	}
	return o.BackupSpec, true
}

// HasBackupSpec returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasBackupSpec() bool {
	if o != nil && !IsNil(o.BackupSpec) {
		return true
	}

	return false
}

// SetBackupSpec gets a reference to the given []int32 and assigns it to the BackupSpec field.
func (o *RequestsPutBackupJobRequest) SetBackupSpec(v []int32) {
	o.BackupSpec = v
}

// GetCloudCredentialName returns the CloudCredentialName field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetCloudCredentialName() string {
	if o == nil || IsNil(o.CloudCredentialName) {
		var ret string
		return ret
	}
	return *o.CloudCredentialName
}

// GetCloudCredentialNameOk returns a tuple with the CloudCredentialName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetCloudCredentialNameOk() (*string, bool) {
	if o == nil || IsNil(o.CloudCredentialName) {
		return nil, false
	}
	return o.CloudCredentialName, true
}

// HasCloudCredentialName returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasCloudCredentialName() bool {
	if o != nil && !IsNil(o.CloudCredentialName) {
		return true
	}

	return false
}

// SetCloudCredentialName gets a reference to the given string and assigns it to the CloudCredentialName field.
func (o *RequestsPutBackupJobRequest) SetCloudCredentialName(v string) {
	o.CloudCredentialName = &v
}

// GetCloudSnapId returns the CloudSnapId field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetCloudSnapId() string {
	if o == nil || IsNil(o.CloudSnapId) {
		var ret string
		return ret
	}
	return *o.CloudSnapId
}

// GetCloudSnapIdOk returns a tuple with the CloudSnapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetCloudSnapIdOk() (*string, bool) {
	if o == nil || IsNil(o.CloudSnapId) {
		return nil, false
	}
	return o.CloudSnapId, true
}

// HasCloudSnapId returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasCloudSnapId() bool {
	if o != nil && !IsNil(o.CloudSnapId) {
		return true
	}

	return false
}

// SetCloudSnapId gets a reference to the given string and assigns it to the CloudSnapId field.
func (o *RequestsPutBackupJobRequest) SetCloudSnapId(v string) {
	o.CloudSnapId = &v
}

// GetCompletionStatus returns the CompletionStatus field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetCompletionStatus() string {
	if o == nil || IsNil(o.CompletionStatus) {
		var ret string
		return ret
	}
	return *o.CompletionStatus
}

// GetCompletionStatusOk returns a tuple with the CompletionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetCompletionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CompletionStatus) {
		return nil, false
	}
	return o.CompletionStatus, true
}

// HasCompletionStatus returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasCompletionStatus() bool {
	if o != nil && !IsNil(o.CompletionStatus) {
		return true
	}

	return false
}

// SetCompletionStatus gets a reference to the given string and assigns it to the CompletionStatus field.
func (o *RequestsPutBackupJobRequest) SetCompletionStatus(v string) {
	o.CompletionStatus = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetCompletionTime() string {
	if o == nil || IsNil(o.CompletionTime) {
		var ret string
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetCompletionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CompletionTime) {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasCompletionTime() bool {
	if o != nil && !IsNil(o.CompletionTime) {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given string and assigns it to the CompletionTime field.
func (o *RequestsPutBackupJobRequest) SetCompletionTime(v string) {
	o.CompletionTime = &v
}

// GetDataServiceSpec returns the DataServiceSpec field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetDataServiceSpec() []int32 {
	if o == nil || IsNil(o.DataServiceSpec) {
		var ret []int32
		return ret
	}
	return o.DataServiceSpec
}

// GetDataServiceSpecOk returns a tuple with the DataServiceSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetDataServiceSpecOk() ([]int32, bool) {
	if o == nil || IsNil(o.DataServiceSpec) {
		return nil, false
	}
	return o.DataServiceSpec, true
}

// HasDataServiceSpec returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasDataServiceSpec() bool {
	if o != nil && !IsNil(o.DataServiceSpec) {
		return true
	}

	return false
}

// SetDataServiceSpec gets a reference to the given []int32 and assigns it to the DataServiceSpec field.
func (o *RequestsPutBackupJobRequest) SetDataServiceSpec(v []int32) {
	o.DataServiceSpec = v
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetDeploymentId() string {
	if o == nil || IsNil(o.DeploymentId) {
		var ret string
		return ret
	}
	return *o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetDeploymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeploymentId) {
		return nil, false
	}
	return o.DeploymentId, true
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasDeploymentId() bool {
	if o != nil && !IsNil(o.DeploymentId) {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given string and assigns it to the DeploymentId field.
func (o *RequestsPutBackupJobRequest) SetDeploymentId(v string) {
	o.DeploymentId = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *RequestsPutBackupJobRequest) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *RequestsPutBackupJobRequest) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetFileSize() int32 {
	if o == nil || IsNil(o.FileSize) {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetFileSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *RequestsPutBackupJobRequest) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RequestsPutBackupJobRequest) SetName(v string) {
	o.Name = &v
}

// GetProjectId returns the ProjectId field value
func (o *RequestsPutBackupJobRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *RequestsPutBackupJobRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *RequestsPutBackupJobRequest) SetStartTime(v string) {
	o.StartTime = &v
}

func (o RequestsPutBackupJobRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestsPutBackupJobRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackupId) {
		toSerialize["backup_id"] = o.BackupId
	}
	if !IsNil(o.BackupSpec) {
		toSerialize["backup_spec"] = o.BackupSpec
	}
	if !IsNil(o.CloudCredentialName) {
		toSerialize["cloud_credential_name"] = o.CloudCredentialName
	}
	if !IsNil(o.CloudSnapId) {
		toSerialize["cloud_snap_id"] = o.CloudSnapId
	}
	if !IsNil(o.CompletionStatus) {
		toSerialize["completion_status"] = o.CompletionStatus
	}
	if !IsNil(o.CompletionTime) {
		toSerialize["completion_time"] = o.CompletionTime
	}
	if !IsNil(o.DataServiceSpec) {
		toSerialize["data_service_spec"] = o.DataServiceSpec
	}
	if !IsNil(o.DeploymentId) {
		toSerialize["deployment_id"] = o.DeploymentId
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if !IsNil(o.FileSize) {
		toSerialize["file_size"] = o.FileSize
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["project_id"] = o.ProjectId
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	return toSerialize, nil
}

type NullableRequestsPutBackupJobRequest struct {
	value *RequestsPutBackupJobRequest
	isSet bool
}

func (v NullableRequestsPutBackupJobRequest) Get() *RequestsPutBackupJobRequest {
	return v.value
}

func (v *NullableRequestsPutBackupJobRequest) Set(val *RequestsPutBackupJobRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestsPutBackupJobRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestsPutBackupJobRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestsPutBackupJobRequest(val *RequestsPutBackupJobRequest) *NullableRequestsPutBackupJobRequest {
	return &NullableRequestsPutBackupJobRequest{value: val, isSet: true}
}

func (v NullableRequestsPutBackupJobRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestsPutBackupJobRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


