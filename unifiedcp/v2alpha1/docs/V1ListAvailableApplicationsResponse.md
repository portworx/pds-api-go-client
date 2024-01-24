# V1ListAvailableApplicationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]V1Application**](V1Application.md) |  | [optional] 

## Methods

### NewV1ListAvailableApplicationsResponse

`func NewV1ListAvailableApplicationsResponse() *V1ListAvailableApplicationsResponse`

NewV1ListAvailableApplicationsResponse instantiates a new V1ListAvailableApplicationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListAvailableApplicationsResponseWithDefaults

`func NewV1ListAvailableApplicationsResponseWithDefaults() *V1ListAvailableApplicationsResponse`

NewV1ListAvailableApplicationsResponseWithDefaults instantiates a new V1ListAvailableApplicationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *V1ListAvailableApplicationsResponse) GetApplications() []V1Application`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *V1ListAvailableApplicationsResponse) GetApplicationsOk() (*[]V1Application, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *V1ListAvailableApplicationsResponse) SetApplications(v []V1Application)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *V1ListAvailableApplicationsResponse) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


