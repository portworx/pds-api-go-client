# V1SearchQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataServiceId** | Pointer to **string** | (-- api-linter: core::0158::response-repeated-first-field&#x3D;disabled     aip.dev/not-precedent: We need to do this because of the copilot response. --) (-- api-linter: core::0158::response-plural-first-field&#x3D;disabled     aip.dev/not-precedent: We need to do this because of the copilot response. --) Id of the dataservice. | [optional] 
**Id** | Pointer to **string** | Unique id of response. | [optional] 
**Query** | Pointer to **string** | Query string. | [optional] 
**Response** | Pointer to **string** | Response string. | [optional] 
**ResponseTime** | Pointer to **time.Time** | Response time. | [optional] 

## Methods

### NewV1SearchQueryResponse

`func NewV1SearchQueryResponse() *V1SearchQueryResponse`

NewV1SearchQueryResponse instantiates a new V1SearchQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SearchQueryResponseWithDefaults

`func NewV1SearchQueryResponseWithDefaults() *V1SearchQueryResponse`

NewV1SearchQueryResponseWithDefaults instantiates a new V1SearchQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataServiceId

`func (o *V1SearchQueryResponse) GetDataServiceId() string`

GetDataServiceId returns the DataServiceId field if non-nil, zero value otherwise.

### GetDataServiceIdOk

`func (o *V1SearchQueryResponse) GetDataServiceIdOk() (*string, bool)`

GetDataServiceIdOk returns a tuple with the DataServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataServiceId

`func (o *V1SearchQueryResponse) SetDataServiceId(v string)`

SetDataServiceId sets DataServiceId field to given value.

### HasDataServiceId

`func (o *V1SearchQueryResponse) HasDataServiceId() bool`

HasDataServiceId returns a boolean if a field has been set.

### GetId

`func (o *V1SearchQueryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1SearchQueryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1SearchQueryResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1SearchQueryResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuery

`func (o *V1SearchQueryResponse) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *V1SearchQueryResponse) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *V1SearchQueryResponse) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *V1SearchQueryResponse) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetResponse

`func (o *V1SearchQueryResponse) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *V1SearchQueryResponse) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *V1SearchQueryResponse) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *V1SearchQueryResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetResponseTime

`func (o *V1SearchQueryResponse) GetResponseTime() time.Time`

GetResponseTime returns the ResponseTime field if non-nil, zero value otherwise.

### GetResponseTimeOk

`func (o *V1SearchQueryResponse) GetResponseTimeOk() (*time.Time, bool)`

GetResponseTimeOk returns a tuple with the ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTime

`func (o *V1SearchQueryResponse) SetResponseTime(v time.Time)`

SetResponseTime sets ResponseTime field to given value.

### HasResponseTime

`func (o *V1SearchQueryResponse) HasResponseTime() bool`

HasResponseTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


