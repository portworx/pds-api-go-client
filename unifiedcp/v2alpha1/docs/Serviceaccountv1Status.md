# Serviceaccountv1Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretGenerationCount** | Pointer to **int32** | Represents how many times the service account secret has been rotated. | [optional] [readonly] 
**LastSecretUpdateTime** | Pointer to **time.Time** | When last time the secrets of the service account has been updated. | [optional] [readonly] 

## Methods

### NewServiceaccountv1Status

`func NewServiceaccountv1Status() *Serviceaccountv1Status`

NewServiceaccountv1Status instantiates a new Serviceaccountv1Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceaccountv1StatusWithDefaults

`func NewServiceaccountv1StatusWithDefaults() *Serviceaccountv1Status`

NewServiceaccountv1StatusWithDefaults instantiates a new Serviceaccountv1Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretGenerationCount

`func (o *Serviceaccountv1Status) GetSecretGenerationCount() int32`

GetSecretGenerationCount returns the SecretGenerationCount field if non-nil, zero value otherwise.

### GetSecretGenerationCountOk

`func (o *Serviceaccountv1Status) GetSecretGenerationCountOk() (*int32, bool)`

GetSecretGenerationCountOk returns a tuple with the SecretGenerationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretGenerationCount

`func (o *Serviceaccountv1Status) SetSecretGenerationCount(v int32)`

SetSecretGenerationCount sets SecretGenerationCount field to given value.

### HasSecretGenerationCount

`func (o *Serviceaccountv1Status) HasSecretGenerationCount() bool`

HasSecretGenerationCount returns a boolean if a field has been set.

### GetLastSecretUpdateTime

`func (o *Serviceaccountv1Status) GetLastSecretUpdateTime() time.Time`

GetLastSecretUpdateTime returns the LastSecretUpdateTime field if non-nil, zero value otherwise.

### GetLastSecretUpdateTimeOk

`func (o *Serviceaccountv1Status) GetLastSecretUpdateTimeOk() (*time.Time, bool)`

GetLastSecretUpdateTimeOk returns a tuple with the LastSecretUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSecretUpdateTime

`func (o *Serviceaccountv1Status) SetLastSecretUpdateTime(v time.Time)`

SetLastSecretUpdateTime sets LastSecretUpdateTime field to given value.

### HasLastSecretUpdateTime

`func (o *Serviceaccountv1Status) HasLastSecretUpdateTime() bool`

HasLastSecretUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


