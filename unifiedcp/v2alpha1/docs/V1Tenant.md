# V1Tenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V1Meta**](V1Meta.md) |  | [optional] 
**Status** | Pointer to [**Tenantv1Status**](Tenantv1Status.md) |  | [optional] 

## Methods

### NewV1Tenant

`func NewV1Tenant() *V1Tenant`

NewV1Tenant instantiates a new V1Tenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TenantWithDefaults

`func NewV1TenantWithDefaults() *V1Tenant`

NewV1TenantWithDefaults instantiates a new V1Tenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V1Tenant) GetMeta() V1Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V1Tenant) GetMetaOk() (*V1Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V1Tenant) SetMeta(v V1Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V1Tenant) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetStatus

`func (o *V1Tenant) GetStatus() Tenantv1Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1Tenant) GetStatusOk() (*Tenantv1Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1Tenant) SetStatus(v Tenantv1Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1Tenant) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


