# V1Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | UID of the schedule associated with the backup configuration. | [optional] 
**ResourceVersion** | Pointer to **string** | Resource version of the schedule. | [optional] 

## Methods

### NewV1Schedule

`func NewV1Schedule() *V1Schedule`

NewV1Schedule instantiates a new V1Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ScheduleWithDefaults

`func NewV1ScheduleWithDefaults() *V1Schedule`

NewV1ScheduleWithDefaults instantiates a new V1Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V1Schedule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1Schedule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1Schedule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1Schedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceVersion

`func (o *V1Schedule) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *V1Schedule) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *V1Schedule) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *V1Schedule) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


