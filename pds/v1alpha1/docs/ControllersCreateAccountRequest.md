# ControllersCreateAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsDetails** | Pointer to [**ModelsDNSDetails**](ModelsDNSDetails.md) |  | [optional] 
**MaasDetails** | Pointer to [**ModelsMAASDetails**](ModelsMAASDetails.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the account. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


