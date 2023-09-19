/*
PDS API

Portworx Data Services API Server

API version: 121
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"encoding/json"
)

// ModelsBackupJob struct for ModelsBackupJob
type ModelsBackupJob struct {
	// BackupCapability of the deployment target when the snapshot was created
	BackupCapability *string `json:"backup_capability,omitempty"`
	// BackupID of the backup which created the snapshot (required).
	BackupId *string `json:"backup_id,omitempty"`
	// BackupSpec is the specification of the Backup at the time the snapshot was taken.
	BackupSpec map[string]interface{} `json:"backup_spec,omitempty"`
	// CloudCredentialName credentials to access snapshot.
	CloudCredentialName *string `json:"cloud_credential_name,omitempty"`
	// CloudSnapID snapshot of the backup volume.
	CloudSnapId *string `json:"cloud_snap_id,omitempty"`
	// CompletionStatus of the snapshot.
	CompletionStatus *string `json:"completion_status,omitempty"`
	CompletionTime *string `json:"completion_time,omitempty"`
	// CreatedAt is autogenerated on creation
	CreatedAt *string `json:"created_at,omitempty"`
	// DataServiceSpec is the specification of the Deployment at the time the snapshot was taken.
	DataServiceSpec map[string]interface{} `json:"data_service_spec,omitempty"`
	// DeploymentID of the deployment which was backed up (required).
	DeploymentId *string `json:"deployment_id,omitempty"`
	// DeploymentTargetID on which target the backup is created.
	DeploymentTargetId *string `json:"deployment_target_id,omitempty"`
	// ErrorCode if CompletionStatus is \"Failed\"
	ErrorCode *string `json:"error_code,omitempty"`
	// ErrorMessage associated with the ErrorCode
	ErrorMessage *string `json:"error_message,omitempty"`
	// FileSize of the CloudSnap image.
	FileSize *int32 `json:"file_size,omitempty"`
	// ID is auto generated on creation
	Id *string `json:"id,omitempty"`
	// ImageID is the ID of the data service's image that was backed up.
	ImageId *string `json:"image_id,omitempty"`
	// Name of custom resource on deployment target.
	Name *string `json:"name,omitempty"`
	// NamespaceID of the namespace where snapshot was created.
	NamespaceId *string `json:"namespace_id,omitempty"`
	// ProjectID of the project where the snapshot resides (required).
	ProjectId *string `json:"project_id,omitempty"`
	StartTime *string `json:"start_time,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	// UpdatedAt is autogenerated on update
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewModelsBackupJob instantiates a new ModelsBackupJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsBackupJob() *ModelsBackupJob {
	this := ModelsBackupJob{}
	return &this
}

// NewModelsBackupJobWithDefaults instantiates a new ModelsBackupJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsBackupJobWithDefaults() *ModelsBackupJob {
	this := ModelsBackupJob{}
	return &this
}

// GetBackupCapability returns the BackupCapability field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetBackupCapability() string {
	if o == nil || o.BackupCapability == nil {
		var ret string
		return ret
	}
	return *o.BackupCapability
}

// GetBackupCapabilityOk returns a tuple with the BackupCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetBackupCapabilityOk() (*string, bool) {
	if o == nil || o.BackupCapability == nil {
		return nil, false
	}
	return o.BackupCapability, true
}

// HasBackupCapability returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasBackupCapability() bool {
	if o != nil && o.BackupCapability != nil {
		return true
	}

	return false
}

// SetBackupCapability gets a reference to the given string and assigns it to the BackupCapability field.
func (o *ModelsBackupJob) SetBackupCapability(v string) {
	o.BackupCapability = &v
}

// GetBackupId returns the BackupId field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetBackupId() string {
	if o == nil || o.BackupId == nil {
		var ret string
		return ret
	}
	return *o.BackupId
}

// GetBackupIdOk returns a tuple with the BackupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetBackupIdOk() (*string, bool) {
	if o == nil || o.BackupId == nil {
		return nil, false
	}
	return o.BackupId, true
}

// HasBackupId returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasBackupId() bool {
	if o != nil && o.BackupId != nil {
		return true
	}

	return false
}

// SetBackupId gets a reference to the given string and assigns it to the BackupId field.
func (o *ModelsBackupJob) SetBackupId(v string) {
	o.BackupId = &v
}

// GetBackupSpec returns the BackupSpec field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetBackupSpec() map[string]interface{} {
	if o == nil || o.BackupSpec == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.BackupSpec
}

// GetBackupSpecOk returns a tuple with the BackupSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetBackupSpecOk() (map[string]interface{}, bool) {
	if o == nil || o.BackupSpec == nil {
		return nil, false
	}
	return o.BackupSpec, true
}

// HasBackupSpec returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasBackupSpec() bool {
	if o != nil && o.BackupSpec != nil {
		return true
	}

	return false
}

// SetBackupSpec gets a reference to the given map[string]interface{} and assigns it to the BackupSpec field.
func (o *ModelsBackupJob) SetBackupSpec(v map[string]interface{}) {
	o.BackupSpec = v
}

// GetCloudCredentialName returns the CloudCredentialName field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetCloudCredentialName() string {
	if o == nil || o.CloudCredentialName == nil {
		var ret string
		return ret
	}
	return *o.CloudCredentialName
}

// GetCloudCredentialNameOk returns a tuple with the CloudCredentialName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetCloudCredentialNameOk() (*string, bool) {
	if o == nil || o.CloudCredentialName == nil {
		return nil, false
	}
	return o.CloudCredentialName, true
}

// HasCloudCredentialName returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasCloudCredentialName() bool {
	if o != nil && o.CloudCredentialName != nil {
		return true
	}

	return false
}

// SetCloudCredentialName gets a reference to the given string and assigns it to the CloudCredentialName field.
func (o *ModelsBackupJob) SetCloudCredentialName(v string) {
	o.CloudCredentialName = &v
}

// GetCloudSnapId returns the CloudSnapId field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetCloudSnapId() string {
	if o == nil || o.CloudSnapId == nil {
		var ret string
		return ret
	}
	return *o.CloudSnapId
}

// GetCloudSnapIdOk returns a tuple with the CloudSnapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetCloudSnapIdOk() (*string, bool) {
	if o == nil || o.CloudSnapId == nil {
		return nil, false
	}
	return o.CloudSnapId, true
}

// HasCloudSnapId returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasCloudSnapId() bool {
	if o != nil && o.CloudSnapId != nil {
		return true
	}

	return false
}

// SetCloudSnapId gets a reference to the given string and assigns it to the CloudSnapId field.
func (o *ModelsBackupJob) SetCloudSnapId(v string) {
	o.CloudSnapId = &v
}

// GetCompletionStatus returns the CompletionStatus field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetCompletionStatus() string {
	if o == nil || o.CompletionStatus == nil {
		var ret string
		return ret
	}
	return *o.CompletionStatus
}

// GetCompletionStatusOk returns a tuple with the CompletionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetCompletionStatusOk() (*string, bool) {
	if o == nil || o.CompletionStatus == nil {
		return nil, false
	}
	return o.CompletionStatus, true
}

// HasCompletionStatus returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasCompletionStatus() bool {
	if o != nil && o.CompletionStatus != nil {
		return true
	}

	return false
}

// SetCompletionStatus gets a reference to the given string and assigns it to the CompletionStatus field.
func (o *ModelsBackupJob) SetCompletionStatus(v string) {
	o.CompletionStatus = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetCompletionTime() string {
	if o == nil || o.CompletionTime == nil {
		var ret string
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetCompletionTimeOk() (*string, bool) {
	if o == nil || o.CompletionTime == nil {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasCompletionTime() bool {
	if o != nil && o.CompletionTime != nil {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given string and assigns it to the CompletionTime field.
func (o *ModelsBackupJob) SetCompletionTime(v string) {
	o.CompletionTime = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ModelsBackupJob) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDataServiceSpec returns the DataServiceSpec field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetDataServiceSpec() map[string]interface{} {
	if o == nil || o.DataServiceSpec == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DataServiceSpec
}

// GetDataServiceSpecOk returns a tuple with the DataServiceSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetDataServiceSpecOk() (map[string]interface{}, bool) {
	if o == nil || o.DataServiceSpec == nil {
		return nil, false
	}
	return o.DataServiceSpec, true
}

// HasDataServiceSpec returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasDataServiceSpec() bool {
	if o != nil && o.DataServiceSpec != nil {
		return true
	}

	return false
}

// SetDataServiceSpec gets a reference to the given map[string]interface{} and assigns it to the DataServiceSpec field.
func (o *ModelsBackupJob) SetDataServiceSpec(v map[string]interface{}) {
	o.DataServiceSpec = v
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetDeploymentId() string {
	if o == nil || o.DeploymentId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetDeploymentIdOk() (*string, bool) {
	if o == nil || o.DeploymentId == nil {
		return nil, false
	}
	return o.DeploymentId, true
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasDeploymentId() bool {
	if o != nil && o.DeploymentId != nil {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given string and assigns it to the DeploymentId field.
func (o *ModelsBackupJob) SetDeploymentId(v string) {
	o.DeploymentId = &v
}

// GetDeploymentTargetId returns the DeploymentTargetId field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetDeploymentTargetId() string {
	if o == nil || o.DeploymentTargetId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentTargetId
}

// GetDeploymentTargetIdOk returns a tuple with the DeploymentTargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetDeploymentTargetIdOk() (*string, bool) {
	if o == nil || o.DeploymentTargetId == nil {
		return nil, false
	}
	return o.DeploymentTargetId, true
}

// HasDeploymentTargetId returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasDeploymentTargetId() bool {
	if o != nil && o.DeploymentTargetId != nil {
		return true
	}

	return false
}

// SetDeploymentTargetId gets a reference to the given string and assigns it to the DeploymentTargetId field.
func (o *ModelsBackupJob) SetDeploymentTargetId(v string) {
	o.DeploymentTargetId = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *ModelsBackupJob) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ModelsBackupJob) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetFileSize() int32 {
	if o == nil || o.FileSize == nil {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetFileSizeOk() (*int32, bool) {
	if o == nil || o.FileSize == nil {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasFileSize() bool {
	if o != nil && o.FileSize != nil {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *ModelsBackupJob) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelsBackupJob) SetId(v string) {
	o.Id = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetImageId() string {
	if o == nil || o.ImageId == nil {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetImageIdOk() (*string, bool) {
	if o == nil || o.ImageId == nil {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasImageId() bool {
	if o != nil && o.ImageId != nil {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *ModelsBackupJob) SetImageId(v string) {
	o.ImageId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelsBackupJob) SetName(v string) {
	o.Name = &v
}

// GetNamespaceId returns the NamespaceId field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetNamespaceId() string {
	if o == nil || o.NamespaceId == nil {
		var ret string
		return ret
	}
	return *o.NamespaceId
}

// GetNamespaceIdOk returns a tuple with the NamespaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetNamespaceIdOk() (*string, bool) {
	if o == nil || o.NamespaceId == nil {
		return nil, false
	}
	return o.NamespaceId, true
}

// HasNamespaceId returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasNamespaceId() bool {
	if o != nil && o.NamespaceId != nil {
		return true
	}

	return false
}

// SetNamespaceId gets a reference to the given string and assigns it to the NamespaceId field.
func (o *ModelsBackupJob) SetNamespaceId(v string) {
	o.NamespaceId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *ModelsBackupJob) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *ModelsBackupJob) SetStartTime(v string) {
	o.StartTime = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *ModelsBackupJob) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ModelsBackupJob) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ModelsBackupJob) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackupCapability != nil {
		toSerialize["backup_capability"] = o.BackupCapability
	}
	if o.BackupId != nil {
		toSerialize["backup_id"] = o.BackupId
	}
	if o.BackupSpec != nil {
		toSerialize["backup_spec"] = o.BackupSpec
	}
	if o.CloudCredentialName != nil {
		toSerialize["cloud_credential_name"] = o.CloudCredentialName
	}
	if o.CloudSnapId != nil {
		toSerialize["cloud_snap_id"] = o.CloudSnapId
	}
	if o.CompletionStatus != nil {
		toSerialize["completion_status"] = o.CompletionStatus
	}
	if o.CompletionTime != nil {
		toSerialize["completion_time"] = o.CompletionTime
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DataServiceSpec != nil {
		toSerialize["data_service_spec"] = o.DataServiceSpec
	}
	if o.DeploymentId != nil {
		toSerialize["deployment_id"] = o.DeploymentId
	}
	if o.DeploymentTargetId != nil {
		toSerialize["deployment_target_id"] = o.DeploymentTargetId
	}
	if o.ErrorCode != nil {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.FileSize != nil {
		toSerialize["file_size"] = o.FileSize
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ImageId != nil {
		toSerialize["image_id"] = o.ImageId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NamespaceId != nil {
		toSerialize["namespace_id"] = o.NamespaceId
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableModelsBackupJob struct {
	value *ModelsBackupJob
	isSet bool
}

func (v NullableModelsBackupJob) Get() *ModelsBackupJob {
	return v.value
}

func (v *NullableModelsBackupJob) Set(val *ModelsBackupJob) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsBackupJob) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsBackupJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsBackupJob(val *ModelsBackupJob) *NullableModelsBackupJob {
	return &NullableModelsBackupJob{value: val, isSet: true}
}

func (v NullableModelsBackupJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsBackupJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


