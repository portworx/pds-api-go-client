# ModelsDNSDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsDetails** | Pointer to [**ModelsAWSDetails**](ModelsAWSDetails.md) |  | [optional] 
**DnsZone** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsDNSDetails

`func NewModelsDNSDetails() *ModelsDNSDetails`

NewModelsDNSDetails instantiates a new ModelsDNSDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsDNSDetailsWithDefaults

`func NewModelsDNSDetailsWithDefaults() *ModelsDNSDetails`

NewModelsDNSDetailsWithDefaults instantiates a new ModelsDNSDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsDetails

`func (o *ModelsDNSDetails) GetAwsDetails() ModelsAWSDetails`

GetAwsDetails returns the AwsDetails field if non-nil, zero value otherwise.

### GetAwsDetailsOk

`func (o *ModelsDNSDetails) GetAwsDetailsOk() (*ModelsAWSDetails, bool)`

GetAwsDetailsOk returns a tuple with the AwsDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsDetails

`func (o *ModelsDNSDetails) SetAwsDetails(v ModelsAWSDetails)`

SetAwsDetails sets AwsDetails field to given value.

### HasAwsDetails

`func (o *ModelsDNSDetails) HasAwsDetails() bool`

HasAwsDetails returns a boolean if a field has been set.

### GetDnsZone

`func (o *ModelsDNSDetails) GetDnsZone() string`

GetDnsZone returns the DnsZone field if non-nil, zero value otherwise.

### GetDnsZoneOk

`func (o *ModelsDNSDetails) GetDnsZoneOk() (*string, bool)`

GetDnsZoneOk returns a tuple with the DnsZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZone

`func (o *ModelsDNSDetails) SetDnsZone(v string)`

SetDnsZone sets DnsZone field to given value.

### HasDnsZone

`func (o *ModelsDNSDetails) HasDnsZone() bool`

HasDnsZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


