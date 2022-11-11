# RequestsCreateTenantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the Tenant. Must be unique within an Account. | [optional] 

## Methods

### NewRequestsCreateTenantRequest

`func NewRequestsCreateTenantRequest() *RequestsCreateTenantRequest`

NewRequestsCreateTenantRequest instantiates a new RequestsCreateTenantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsCreateTenantRequestWithDefaults

`func NewRequestsCreateTenantRequestWithDefaults() *RequestsCreateTenantRequest`

NewRequestsCreateTenantRequestWithDefaults instantiates a new RequestsCreateTenantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RequestsCreateTenantRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestsCreateTenantRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestsCreateTenantRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequestsCreateTenantRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


