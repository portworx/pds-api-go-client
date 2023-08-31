# RequestsIAMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorId** | **string** |  | 
**Data** | [**ModelsAccessPolicy**](ModelsAccessPolicy.md) |  | 

## Methods

### NewRequestsIAMRequest

`func NewRequestsIAMRequest(actorId string, data ModelsAccessPolicy, ) *RequestsIAMRequest`

NewRequestsIAMRequest instantiates a new RequestsIAMRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsIAMRequestWithDefaults

`func NewRequestsIAMRequestWithDefaults() *RequestsIAMRequest`

NewRequestsIAMRequestWithDefaults instantiates a new RequestsIAMRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorId

`func (o *RequestsIAMRequest) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *RequestsIAMRequest) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *RequestsIAMRequest) SetActorId(v string)`

SetActorId sets ActorId field to given value.


### GetData

`func (o *RequestsIAMRequest) GetData() ModelsAccessPolicy`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RequestsIAMRequest) GetDataOk() (*ModelsAccessPolicy, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RequestsIAMRequest) SetData(v ModelsAccessPolicy)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


