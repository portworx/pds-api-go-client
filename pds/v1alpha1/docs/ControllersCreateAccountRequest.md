# ControllersCreateAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaasFeaturesEnabled** | Pointer to **bool** | Whether BAAS should be enabled for this account. | [optional] 
**DnsDetails** | Pointer to [**ModelsDNSDetails**](ModelsDNSDetails.md) |  | [optional] 
**MaasDetails** | Pointer to [**ModelsMAASDetails**](ModelsMAASDetails.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the account. | [optional] 
**PdsFeaturesEnabled** | Pointer to **bool** | Whether PDS should be enabled for this account. | [optional] 
**PxoneFeaturesEnabled** | Pointer to **bool** | Whether PX-ONE should be enabled for this account. | [optional] 
**Subdomain** | Pointer to **string** | (optional) Account subdomain name. | [optional] 

## Methods

### NewControllersCreateAccountRequest

`func NewControllersCreateAccountRequest() *ControllersCreateAccountRequest`

NewControllersCreateAccountRequest instantiates a new ControllersCreateAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersCreateAccountRequestWithDefaults

`func NewControllersCreateAccountRequestWithDefaults() *ControllersCreateAccountRequest`

NewControllersCreateAccountRequestWithDefaults instantiates a new ControllersCreateAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaasFeaturesEnabled

`func (o *ControllersCreateAccountRequest) GetBaasFeaturesEnabled() bool`

GetBaasFeaturesEnabled returns the BaasFeaturesEnabled field if non-nil, zero value otherwise.

### GetBaasFeaturesEnabledOk

`func (o *ControllersCreateAccountRequest) GetBaasFeaturesEnabledOk() (*bool, bool)`

GetBaasFeaturesEnabledOk returns a tuple with the BaasFeaturesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaasFeaturesEnabled

`func (o *ControllersCreateAccountRequest) SetBaasFeaturesEnabled(v bool)`

SetBaasFeaturesEnabled sets BaasFeaturesEnabled field to given value.

### HasBaasFeaturesEnabled

`func (o *ControllersCreateAccountRequest) HasBaasFeaturesEnabled() bool`

HasBaasFeaturesEnabled returns a boolean if a field has been set.

### GetDnsDetails

`func (o *ControllersCreateAccountRequest) GetDnsDetails() ModelsDNSDetails`

GetDnsDetails returns the DnsDetails field if non-nil, zero value otherwise.

### GetDnsDetailsOk

`func (o *ControllersCreateAccountRequest) GetDnsDetailsOk() (*ModelsDNSDetails, bool)`

GetDnsDetailsOk returns a tuple with the DnsDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDetails

`func (o *ControllersCreateAccountRequest) SetDnsDetails(v ModelsDNSDetails)`

SetDnsDetails sets DnsDetails field to given value.

### HasDnsDetails

`func (o *ControllersCreateAccountRequest) HasDnsDetails() bool`

HasDnsDetails returns a boolean if a field has been set.

### GetMaasDetails

`func (o *ControllersCreateAccountRequest) GetMaasDetails() ModelsMAASDetails`

GetMaasDetails returns the MaasDetails field if non-nil, zero value otherwise.

### GetMaasDetailsOk

`func (o *ControllersCreateAccountRequest) GetMaasDetailsOk() (*ModelsMAASDetails, bool)`

GetMaasDetailsOk returns a tuple with the MaasDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaasDetails

`func (o *ControllersCreateAccountRequest) SetMaasDetails(v ModelsMAASDetails)`

SetMaasDetails sets MaasDetails field to given value.

### HasMaasDetails

`func (o *ControllersCreateAccountRequest) HasMaasDetails() bool`

HasMaasDetails returns a boolean if a field has been set.

### GetName

`func (o *ControllersCreateAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControllersCreateAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControllersCreateAccountRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControllersCreateAccountRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPdsFeaturesEnabled

`func (o *ControllersCreateAccountRequest) GetPdsFeaturesEnabled() bool`

GetPdsFeaturesEnabled returns the PdsFeaturesEnabled field if non-nil, zero value otherwise.

### GetPdsFeaturesEnabledOk

`func (o *ControllersCreateAccountRequest) GetPdsFeaturesEnabledOk() (*bool, bool)`

GetPdsFeaturesEnabledOk returns a tuple with the PdsFeaturesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdsFeaturesEnabled

`func (o *ControllersCreateAccountRequest) SetPdsFeaturesEnabled(v bool)`

SetPdsFeaturesEnabled sets PdsFeaturesEnabled field to given value.

### HasPdsFeaturesEnabled

`func (o *ControllersCreateAccountRequest) HasPdsFeaturesEnabled() bool`

HasPdsFeaturesEnabled returns a boolean if a field has been set.

### GetPxoneFeaturesEnabled

`func (o *ControllersCreateAccountRequest) GetPxoneFeaturesEnabled() bool`

GetPxoneFeaturesEnabled returns the PxoneFeaturesEnabled field if non-nil, zero value otherwise.

### GetPxoneFeaturesEnabledOk

`func (o *ControllersCreateAccountRequest) GetPxoneFeaturesEnabledOk() (*bool, bool)`

GetPxoneFeaturesEnabledOk returns a tuple with the PxoneFeaturesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxoneFeaturesEnabled

`func (o *ControllersCreateAccountRequest) SetPxoneFeaturesEnabled(v bool)`

SetPxoneFeaturesEnabled sets PxoneFeaturesEnabled field to given value.

### HasPxoneFeaturesEnabled

`func (o *ControllersCreateAccountRequest) HasPxoneFeaturesEnabled() bool`

HasPxoneFeaturesEnabled returns a boolean if a field has been set.

### GetSubdomain

`func (o *ControllersCreateAccountRequest) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ControllersCreateAccountRequest) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ControllersCreateAccountRequest) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ControllersCreateAccountRequest) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


