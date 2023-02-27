# ControllersUpdateGlobalConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TlsPreviewEnabled** | Pointer to **string** | Specify if the TLS Preview feature should be enabled for this account. | [optional] 
**VersionAvailability** | Pointer to **string** | Specify what data service versions are available for deployment for this account. | [optional] 
**VersionUpdatability** | Pointer to **string** | Specify what data service versions are updatable for this account. | [optional] 

## Methods

### NewControllersUpdateGlobalConfigRequest

`func NewControllersUpdateGlobalConfigRequest() *ControllersUpdateGlobalConfigRequest`

NewControllersUpdateGlobalConfigRequest instantiates a new ControllersUpdateGlobalConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersUpdateGlobalConfigRequestWithDefaults

`func NewControllersUpdateGlobalConfigRequestWithDefaults() *ControllersUpdateGlobalConfigRequest`

NewControllersUpdateGlobalConfigRequestWithDefaults instantiates a new ControllersUpdateGlobalConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTlsPreviewEnabled

`func (o *ControllersUpdateGlobalConfigRequest) GetTlsPreviewEnabled() string`

GetTlsPreviewEnabled returns the TlsPreviewEnabled field if non-nil, zero value otherwise.

### GetTlsPreviewEnabledOk

`func (o *ControllersUpdateGlobalConfigRequest) GetTlsPreviewEnabledOk() (*string, bool)`

GetTlsPreviewEnabledOk returns a tuple with the TlsPreviewEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsPreviewEnabled

`func (o *ControllersUpdateGlobalConfigRequest) SetTlsPreviewEnabled(v string)`

SetTlsPreviewEnabled sets TlsPreviewEnabled field to given value.

### HasTlsPreviewEnabled

`func (o *ControllersUpdateGlobalConfigRequest) HasTlsPreviewEnabled() bool`

HasTlsPreviewEnabled returns a boolean if a field has been set.

### GetVersionAvailability

`func (o *ControllersUpdateGlobalConfigRequest) GetVersionAvailability() string`

GetVersionAvailability returns the VersionAvailability field if non-nil, zero value otherwise.

### GetVersionAvailabilityOk

`func (o *ControllersUpdateGlobalConfigRequest) GetVersionAvailabilityOk() (*string, bool)`

GetVersionAvailabilityOk returns a tuple with the VersionAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionAvailability

`func (o *ControllersUpdateGlobalConfigRequest) SetVersionAvailability(v string)`

SetVersionAvailability sets VersionAvailability field to given value.

### HasVersionAvailability

`func (o *ControllersUpdateGlobalConfigRequest) HasVersionAvailability() bool`

HasVersionAvailability returns a boolean if a field has been set.

### GetVersionUpdatability

`func (o *ControllersUpdateGlobalConfigRequest) GetVersionUpdatability() string`

GetVersionUpdatability returns the VersionUpdatability field if non-nil, zero value otherwise.

### GetVersionUpdatabilityOk

`func (o *ControllersUpdateGlobalConfigRequest) GetVersionUpdatabilityOk() (*string, bool)`

GetVersionUpdatabilityOk returns a tuple with the VersionUpdatability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionUpdatability

`func (o *ControllersUpdateGlobalConfigRequest) SetVersionUpdatability(v string)`

SetVersionUpdatability sets VersionUpdatability field to given value.

### HasVersionUpdatability

`func (o *ControllersUpdateGlobalConfigRequest) HasVersionUpdatability() bool`

HasVersionUpdatability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


