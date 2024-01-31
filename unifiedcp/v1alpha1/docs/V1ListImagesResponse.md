# V1ListImagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | Pointer to [**[]V1Image**](V1Image.md) | List of images. | [optional] 
**Pagination** | Pointer to [**V1PageBasedPaginationResponse**](V1PageBasedPaginationResponse.md) |  | [optional] 

## Methods

### NewV1ListImagesResponse

`func NewV1ListImagesResponse() *V1ListImagesResponse`

NewV1ListImagesResponse instantiates a new V1ListImagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListImagesResponseWithDefaults

`func NewV1ListImagesResponseWithDefaults() *V1ListImagesResponse`

NewV1ListImagesResponseWithDefaults instantiates a new V1ListImagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *V1ListImagesResponse) GetImages() []V1Image`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *V1ListImagesResponse) GetImagesOk() (*[]V1Image, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *V1ListImagesResponse) SetImages(v []V1Image)`

SetImages sets Images field to given value.

### HasImages

`func (o *V1ListImagesResponse) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetPagination

`func (o *V1ListImagesResponse) GetPagination() V1PageBasedPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1ListImagesResponse) GetPaginationOk() (*V1PageBasedPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1ListImagesResponse) SetPagination(v V1PageBasedPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1ListImagesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


