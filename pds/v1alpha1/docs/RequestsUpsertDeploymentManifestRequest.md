# RequestsUpsertDeploymentManifestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Manifest** | Pointer to **map[string]interface{}** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 

## Methods

### NewRequestsUpsertDeploymentManifestRequest

`func NewRequestsUpsertDeploymentManifestRequest() *RequestsUpsertDeploymentManifestRequest`

NewRequestsUpsertDeploymentManifestRequest instantiates a new RequestsUpsertDeploymentManifestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsUpsertDeploymentManifestRequestWithDefaults

`func NewRequestsUpsertDeploymentManifestRequestWithDefaults() *RequestsUpsertDeploymentManifestRequest`

NewRequestsUpsertDeploymentManifestRequestWithDefaults instantiates a new RequestsUpsertDeploymentManifestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManifest

`func (o *RequestsUpsertDeploymentManifestRequest) GetManifest() map[string]interface{}`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *RequestsUpsertDeploymentManifestRequest) GetManifestOk() (*map[string]interface{}, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *RequestsUpsertDeploymentManifestRequest) SetManifest(v map[string]interface{})`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *RequestsUpsertDeploymentManifestRequest) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### GetTimestamp

`func (o *RequestsUpsertDeploymentManifestRequest) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RequestsUpsertDeploymentManifestRequest) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RequestsUpsertDeploymentManifestRequest) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RequestsUpsertDeploymentManifestRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


