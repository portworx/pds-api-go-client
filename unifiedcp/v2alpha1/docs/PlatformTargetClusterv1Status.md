# PlatformTargetClusterv1Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**V1Metadata**](V1Metadata.md) |  | [optional] 
**Phase** | Pointer to [**V1TargetClusterPhasePhase**](V1TargetClusterPhasePhase.md) |  | [optional] [default to PHASE_UNSPECIFIED]
**LastStatusUpdateTime** | Pointer to **time.Time** |  | [optional] 
**PlatformAgent** | Pointer to [**V1ApplicationPhasePhase**](V1ApplicationPhasePhase.md) |  | [optional] [default to PHASE_UNSPECIFIED]
**Applications** | Pointer to [**map[string]V1ApplicationPhasePhase**](V1ApplicationPhasePhase.md) |  | [optional] 

## Methods

### NewPlatformTargetClusterv1Status

`func NewPlatformTargetClusterv1Status() *PlatformTargetClusterv1Status`

NewPlatformTargetClusterv1Status instantiates a new PlatformTargetClusterv1Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformTargetClusterv1StatusWithDefaults

`func NewPlatformTargetClusterv1StatusWithDefaults() *PlatformTargetClusterv1Status`

NewPlatformTargetClusterv1StatusWithDefaults instantiates a new PlatformTargetClusterv1Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *PlatformTargetClusterv1Status) GetMetadata() V1Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PlatformTargetClusterv1Status) GetMetadataOk() (*V1Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PlatformTargetClusterv1Status) SetMetadata(v V1Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PlatformTargetClusterv1Status) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPhase

`func (o *PlatformTargetClusterv1Status) GetPhase() V1TargetClusterPhasePhase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *PlatformTargetClusterv1Status) GetPhaseOk() (*V1TargetClusterPhasePhase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *PlatformTargetClusterv1Status) SetPhase(v V1TargetClusterPhasePhase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *PlatformTargetClusterv1Status) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetLastStatusUpdateTime

`func (o *PlatformTargetClusterv1Status) GetLastStatusUpdateTime() time.Time`

GetLastStatusUpdateTime returns the LastStatusUpdateTime field if non-nil, zero value otherwise.

### GetLastStatusUpdateTimeOk

`func (o *PlatformTargetClusterv1Status) GetLastStatusUpdateTimeOk() (*time.Time, bool)`

GetLastStatusUpdateTimeOk returns a tuple with the LastStatusUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusUpdateTime

`func (o *PlatformTargetClusterv1Status) SetLastStatusUpdateTime(v time.Time)`

SetLastStatusUpdateTime sets LastStatusUpdateTime field to given value.

### HasLastStatusUpdateTime

`func (o *PlatformTargetClusterv1Status) HasLastStatusUpdateTime() bool`

HasLastStatusUpdateTime returns a boolean if a field has been set.

### GetPlatformAgent

`func (o *PlatformTargetClusterv1Status) GetPlatformAgent() V1ApplicationPhasePhase`

GetPlatformAgent returns the PlatformAgent field if non-nil, zero value otherwise.

### GetPlatformAgentOk

`func (o *PlatformTargetClusterv1Status) GetPlatformAgentOk() (*V1ApplicationPhasePhase, bool)`

GetPlatformAgentOk returns a tuple with the PlatformAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAgent

`func (o *PlatformTargetClusterv1Status) SetPlatformAgent(v V1ApplicationPhasePhase)`

SetPlatformAgent sets PlatformAgent field to given value.

### HasPlatformAgent

`func (o *PlatformTargetClusterv1Status) HasPlatformAgent() bool`

HasPlatformAgent returns a boolean if a field has been set.

### GetApplications

`func (o *PlatformTargetClusterv1Status) GetApplications() map[string]V1ApplicationPhasePhase`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *PlatformTargetClusterv1Status) GetApplicationsOk() (*map[string]V1ApplicationPhasePhase, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *PlatformTargetClusterv1Status) SetApplications(v map[string]V1ApplicationPhasePhase)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *PlatformTargetClusterv1Status) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


