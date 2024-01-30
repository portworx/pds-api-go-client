# V1SearchQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataServiceId** | Pointer to **string** | Id of the dataservice. | [optional] 
**Query** | Pointer to **string** | Query string. | [optional] 

## Methods

### NewV1SearchQueryRequest

`func NewV1SearchQueryRequest() *V1SearchQueryRequest`

NewV1SearchQueryRequest instantiates a new V1SearchQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SearchQueryRequestWithDefaults

`func NewV1SearchQueryRequestWithDefaults() *V1SearchQueryRequest`

NewV1SearchQueryRequestWithDefaults instantiates a new V1SearchQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataServiceId

`func (o *V1SearchQueryRequest) GetDataServiceId() string`

GetDataServiceId returns the DataServiceId field if non-nil, zero value otherwise.

### GetDataServiceIdOk

`func (o *V1SearchQueryRequest) GetDataServiceIdOk() (*string, bool)`

GetDataServiceIdOk returns a tuple with the DataServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataServiceId

`func (o *V1SearchQueryRequest) SetDataServiceId(v string)`

SetDataServiceId sets DataServiceId field to given value.

### HasDataServiceId

`func (o *V1SearchQueryRequest) HasDataServiceId() bool`

HasDataServiceId returns a boolean if a field has been set.

### GetQuery

`func (o *V1SearchQueryRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *V1SearchQueryRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *V1SearchQueryRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *V1SearchQueryRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


