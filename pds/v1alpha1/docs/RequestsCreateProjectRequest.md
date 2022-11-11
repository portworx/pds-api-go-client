# RequestsCreateProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the Project. Must be unique within a Tenant. | [optional] 

## Methods

### NewRequestsCreateProjectRequest

`func NewRequestsCreateProjectRequest() *RequestsCreateProjectRequest`

NewRequestsCreateProjectRequest instantiates a new RequestsCreateProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsCreateProjectRequestWithDefaults

`func NewRequestsCreateProjectRequestWithDefaults() *RequestsCreateProjectRequest`

NewRequestsCreateProjectRequestWithDefaults instantiates a new RequestsCreateProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RequestsCreateProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestsCreateProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestsCreateProjectRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequestsCreateProjectRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


