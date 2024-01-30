# V1CompatibleVersions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataServiceId** | Pointer to **string** |  | [optional] 
**DataServiceName** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**VersionName** | Pointer to **string** |  | [optional] 
**LatestCompatiblePatch** | Pointer to [**V1DataServiceVersion**](V1DataServiceVersion.md) |  | [optional] 
**LatestCompatibleVersions** | Pointer to [**[]V1DataServiceVersion**](V1DataServiceVersion.md) | LatestCompatibleVersions contains the latest patch versions of compatible major/minor versions. | [optional] 
**CompatibleVersions** | Pointer to [**[]V1DataServiceVersion**](V1DataServiceVersion.md) | CompatibleVersions contains all compatible versions. | [optional] 

## Methods

### NewV1CompatibleVersions

`func NewV1CompatibleVersions() *V1CompatibleVersions`

NewV1CompatibleVersions instantiates a new V1CompatibleVersions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1CompatibleVersionsWithDefaults

`func NewV1CompatibleVersionsWithDefaults() *V1CompatibleVersions`

NewV1CompatibleVersionsWithDefaults instantiates a new V1CompatibleVersions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataServiceId

`func (o *V1CompatibleVersions) GetDataServiceId() string`

GetDataServiceId returns the DataServiceId field if non-nil, zero value otherwise.

### GetDataServiceIdOk

`func (o *V1CompatibleVersions) GetDataServiceIdOk() (*string, bool)`

GetDataServiceIdOk returns a tuple with the DataServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataServiceId

`func (o *V1CompatibleVersions) SetDataServiceId(v string)`

SetDataServiceId sets DataServiceId field to given value.

### HasDataServiceId

`func (o *V1CompatibleVersions) HasDataServiceId() bool`

HasDataServiceId returns a boolean if a field has been set.

### GetDataServiceName

`func (o *V1CompatibleVersions) GetDataServiceName() string`

GetDataServiceName returns the DataServiceName field if non-nil, zero value otherwise.

### GetDataServiceNameOk

`func (o *V1CompatibleVersions) GetDataServiceNameOk() (*string, bool)`

GetDataServiceNameOk returns a tuple with the DataServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataServiceName

`func (o *V1CompatibleVersions) SetDataServiceName(v string)`

SetDataServiceName sets DataServiceName field to given value.

### HasDataServiceName

`func (o *V1CompatibleVersions) HasDataServiceName() bool`

HasDataServiceName returns a boolean if a field has been set.

### GetVersionId

`func (o *V1CompatibleVersions) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *V1CompatibleVersions) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *V1CompatibleVersions) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *V1CompatibleVersions) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetVersionName

`func (o *V1CompatibleVersions) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *V1CompatibleVersions) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *V1CompatibleVersions) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *V1CompatibleVersions) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.

### GetLatestCompatiblePatch

`func (o *V1CompatibleVersions) GetLatestCompatiblePatch() V1DataServiceVersion`

GetLatestCompatiblePatch returns the LatestCompatiblePatch field if non-nil, zero value otherwise.

### GetLatestCompatiblePatchOk

`func (o *V1CompatibleVersions) GetLatestCompatiblePatchOk() (*V1DataServiceVersion, bool)`

GetLatestCompatiblePatchOk returns a tuple with the LatestCompatiblePatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestCompatiblePatch

`func (o *V1CompatibleVersions) SetLatestCompatiblePatch(v V1DataServiceVersion)`

SetLatestCompatiblePatch sets LatestCompatiblePatch field to given value.

### HasLatestCompatiblePatch

`func (o *V1CompatibleVersions) HasLatestCompatiblePatch() bool`

HasLatestCompatiblePatch returns a boolean if a field has been set.

### GetLatestCompatibleVersions

`func (o *V1CompatibleVersions) GetLatestCompatibleVersions() []V1DataServiceVersion`

GetLatestCompatibleVersions returns the LatestCompatibleVersions field if non-nil, zero value otherwise.

### GetLatestCompatibleVersionsOk

`func (o *V1CompatibleVersions) GetLatestCompatibleVersionsOk() (*[]V1DataServiceVersion, bool)`

GetLatestCompatibleVersionsOk returns a tuple with the LatestCompatibleVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestCompatibleVersions

`func (o *V1CompatibleVersions) SetLatestCompatibleVersions(v []V1DataServiceVersion)`

SetLatestCompatibleVersions sets LatestCompatibleVersions field to given value.

### HasLatestCompatibleVersions

`func (o *V1CompatibleVersions) HasLatestCompatibleVersions() bool`

HasLatestCompatibleVersions returns a boolean if a field has been set.

### GetCompatibleVersions

`func (o *V1CompatibleVersions) GetCompatibleVersions() []V1DataServiceVersion`

GetCompatibleVersions returns the CompatibleVersions field if non-nil, zero value otherwise.

### GetCompatibleVersionsOk

`func (o *V1CompatibleVersions) GetCompatibleVersionsOk() (*[]V1DataServiceVersion, bool)`

GetCompatibleVersionsOk returns a tuple with the CompatibleVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleVersions

`func (o *V1CompatibleVersions) SetCompatibleVersions(v []V1DataServiceVersion)`

SetCompatibleVersions sets CompatibleVersions field to given value.

### HasCompatibleVersions

`func (o *V1CompatibleVersions) HasCompatibleVersions() bool`

HasCompatibleVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


