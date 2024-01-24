/*
public/portworx/pds/backupconfig/apiv1/backupconfig.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// V1DeploymentTopology A deployment topology contains a number of nodes that have various attributes as a collective group. This is consumed in a deployment and the functionality is being controlled by the CRUD APIs of deployment.
type V1DeploymentTopology struct {
	// Name of the deployment topology.
	Name *string `json:"name,omitempty"`
	// Description of the deployment topology.
	Description *string `json:"description,omitempty"`
	// Number of replicas of data services.
	Replicas *string `json:"replicas,omitempty"`
	// Service type are standard Kubernetes service types such as clusterIP, NodePort, load balancers, etc.
	ServiceType *string `json:"serviceType,omitempty"`
	// Service name is the name of service as provided by user.
	ServiceName *string `json:"serviceName,omitempty"`
	// Source IP ranges to use for the deployed Load Balancer.
	LoadBalancerSourceRanges []string `json:"loadBalancerSourceRanges,omitempty"`
	ResourceTemplate *V1Template `json:"resourceTemplate,omitempty"`
	ApplicationTemplate *V1Template `json:"applicationTemplate,omitempty"`
	StorageTemplate *V1Template `json:"storageTemplate,omitempty"`
}

// NewV1DeploymentTopology instantiates a new V1DeploymentTopology object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DeploymentTopology() *V1DeploymentTopology {
	this := V1DeploymentTopology{}
	return &this
}

// NewV1DeploymentTopologyWithDefaults instantiates a new V1DeploymentTopology object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DeploymentTopologyWithDefaults() *V1DeploymentTopology {
	this := V1DeploymentTopology{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1DeploymentTopology) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeploymentTopology) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1DeploymentTopology) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1DeploymentTopology) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V1DeploymentTopology) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeploymentTopology) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V1DeploymentTopology) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V1DeploymentTopology) SetDescription(v string) {
	o.Description = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *V1DeploymentTopology) GetReplicas() string {
	if o == nil || o.Replicas == nil {
		var ret string
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeploymentTopology) GetReplicasOk() (*string, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *V1DeploymentTopology) HasReplicas() bool {
	if o != nil && o.Replicas != nil {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given string and assigns it to the Replicas field.
func (o *V1DeploymentTopology) SetReplicas(v string) {
	o.Replicas = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *V1DeploymentTopology) GetServiceType() string {
	if o == nil || o.ServiceType == nil {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeploymentTopology) GetServiceTypeOk() (*string, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *V1DeploymentTopology) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *V1DeploymentTopology) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *V1DeploymentTopology) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeploymentTopology) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *V1DeploymentTopology) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *V1DeploymentTopology) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetLoadBalancerSourceRanges returns the LoadBalancerSourceRanges field value if set, zero value otherwise.
func (o *V1DeploymentTopology) GetLoadBalancerSourceRanges() []string {
	if o == nil || o.LoadBalancerSourceRanges == nil {
		var ret []string
		return ret
	}
	return o.LoadBalancerSourceRanges
}

// GetLoadBalancerSourceRangesOk returns a tuple with the LoadBalancerSourceRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeploymentTopology) GetLoadBalancerSourceRangesOk() ([]string, bool) {
	if o == nil || o.LoadBalancerSourceRanges == nil {
		return nil, false
	}
	return o.LoadBalancerSourceRanges, true
}

// HasLoadBalancerSourceRanges returns a boolean if a field has been set.
func (o *V1DeploymentTopology) HasLoadBalancerSourceRanges() bool {
	if o != nil && o.LoadBalancerSourceRanges != nil {
		return true
	}

	return false
}

// SetLoadBalancerSourceRanges gets a reference to the given []string and assigns it to the LoadBalancerSourceRanges field.
func (o *V1DeploymentTopology) SetLoadBalancerSourceRanges(v []string) {
	o.LoadBalancerSourceRanges = v
}

// GetResourceTemplate returns the ResourceTemplate field value if set, zero value otherwise.
func (o *V1DeploymentTopology) GetResourceTemplate() V1Template {
	if o == nil || o.ResourceTemplate == nil {
		var ret V1Template
		return ret
	}
	return *o.ResourceTemplate
}

// GetResourceTemplateOk returns a tuple with the ResourceTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeploymentTopology) GetResourceTemplateOk() (*V1Template, bool) {
	if o == nil || o.ResourceTemplate == nil {
		return nil, false
	}
	return o.ResourceTemplate, true
}

// HasResourceTemplate returns a boolean if a field has been set.
func (o *V1DeploymentTopology) HasResourceTemplate() bool {
	if o != nil && o.ResourceTemplate != nil {
		return true
	}

	return false
}

// SetResourceTemplate gets a reference to the given V1Template and assigns it to the ResourceTemplate field.
func (o *V1DeploymentTopology) SetResourceTemplate(v V1Template) {
	o.ResourceTemplate = &v
}

// GetApplicationTemplate returns the ApplicationTemplate field value if set, zero value otherwise.
func (o *V1DeploymentTopology) GetApplicationTemplate() V1Template {
	if o == nil || o.ApplicationTemplate == nil {
		var ret V1Template
		return ret
	}
	return *o.ApplicationTemplate
}

// GetApplicationTemplateOk returns a tuple with the ApplicationTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeploymentTopology) GetApplicationTemplateOk() (*V1Template, bool) {
	if o == nil || o.ApplicationTemplate == nil {
		return nil, false
	}
	return o.ApplicationTemplate, true
}

// HasApplicationTemplate returns a boolean if a field has been set.
func (o *V1DeploymentTopology) HasApplicationTemplate() bool {
	if o != nil && o.ApplicationTemplate != nil {
		return true
	}

	return false
}

// SetApplicationTemplate gets a reference to the given V1Template and assigns it to the ApplicationTemplate field.
func (o *V1DeploymentTopology) SetApplicationTemplate(v V1Template) {
	o.ApplicationTemplate = &v
}

// GetStorageTemplate returns the StorageTemplate field value if set, zero value otherwise.
func (o *V1DeploymentTopology) GetStorageTemplate() V1Template {
	if o == nil || o.StorageTemplate == nil {
		var ret V1Template
		return ret
	}
	return *o.StorageTemplate
}

// GetStorageTemplateOk returns a tuple with the StorageTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeploymentTopology) GetStorageTemplateOk() (*V1Template, bool) {
	if o == nil || o.StorageTemplate == nil {
		return nil, false
	}
	return o.StorageTemplate, true
}

// HasStorageTemplate returns a boolean if a field has been set.
func (o *V1DeploymentTopology) HasStorageTemplate() bool {
	if o != nil && o.StorageTemplate != nil {
		return true
	}

	return false
}

// SetStorageTemplate gets a reference to the given V1Template and assigns it to the StorageTemplate field.
func (o *V1DeploymentTopology) SetStorageTemplate(v V1Template) {
	o.StorageTemplate = &v
}

func (o V1DeploymentTopology) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.ServiceType != nil {
		toSerialize["serviceType"] = o.ServiceType
	}
	if o.ServiceName != nil {
		toSerialize["serviceName"] = o.ServiceName
	}
	if o.LoadBalancerSourceRanges != nil {
		toSerialize["loadBalancerSourceRanges"] = o.LoadBalancerSourceRanges
	}
	if o.ResourceTemplate != nil {
		toSerialize["resourceTemplate"] = o.ResourceTemplate
	}
	if o.ApplicationTemplate != nil {
		toSerialize["applicationTemplate"] = o.ApplicationTemplate
	}
	if o.StorageTemplate != nil {
		toSerialize["storageTemplate"] = o.StorageTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableV1DeploymentTopology struct {
	value *V1DeploymentTopology
	isSet bool
}

func (v NullableV1DeploymentTopology) Get() *V1DeploymentTopology {
	return v.value
}

func (v *NullableV1DeploymentTopology) Set(val *V1DeploymentTopology) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DeploymentTopology) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DeploymentTopology) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DeploymentTopology(val *V1DeploymentTopology) *NullableV1DeploymentTopology {
	return &NullableV1DeploymentTopology{value: val, isSet: true}
}

func (v NullableV1DeploymentTopology) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DeploymentTopology) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


