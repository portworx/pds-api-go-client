# \ImageServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImageServiceGetImage**](ImageServiceAPI.md#ImageServiceGetImage) | **Get** /pds/v1/catalog/images/{id} | GetImage API returns the information about the specified image.
[**ImageServiceListImages**](ImageServiceAPI.md#ImageServiceListImages) | **Get** /pds/v1/catalog/images | ListImages API lists the data service images.



## ImageServiceGetImage

> V1Image ImageServiceGetImage(ctx, id).Execute()

GetImage API returns the information about the specified image.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | UID of the image.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageServiceAPI.ImageServiceGetImage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageServiceAPI.ImageServiceGetImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImageServiceGetImage`: V1Image
	fmt.Fprintf(os.Stdout, "Response from `ImageServiceAPI.ImageServiceGetImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UID of the image. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageServiceGetImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1Image**](V1Image.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageServiceListImages

> V1ListImagesResponse ImageServiceListImages(ctx).DataServiceId(dataServiceId).VersionId(versionId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()

ListImages API lists the data service images.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dataServiceId := "dataServiceId_example" // string | UID of the account. (optional)
	versionId := "versionId_example" // string | UID of the tenant. (optional)
	paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
	paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageServiceAPI.ImageServiceListImages(context.Background()).DataServiceId(dataServiceId).VersionId(versionId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageServiceAPI.ImageServiceListImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImageServiceListImages`: V1ListImagesResponse
	fmt.Fprintf(os.Stdout, "Response from `ImageServiceAPI.ImageServiceListImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImageServiceListImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataServiceId** | **string** | UID of the account. | 
 **versionId** | **string** | UID of the tenant. | 
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 

### Return type

[**V1ListImagesResponse**](V1ListImagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

