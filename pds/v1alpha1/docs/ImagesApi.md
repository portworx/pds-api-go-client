# \ImagesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiImagesGet**](ImagesApi.md#ApiImagesGet) | **Get** /api/images | List Images
[**ApiImagesIdGet**](ImagesApi.md#ApiImagesIdGet) | **Get** /api/images/{id} | Get Image
[**ApiVersionsIdImagesGet**](ImagesApi.md#ApiVersionsIdImagesGet) | **Get** /api/versions/{id}/images | List Version&#39;s Images



## ApiImagesGet

> ModelsPaginatedResultModelsImage ApiImagesGet(ctx).DataServiceId(dataServiceId).VersionId(versionId).Latest(latest).SortBy(sortBy).Limit(limit).Continuation(continuation).Execute()

List Images



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
    dataServiceId := "dataServiceId_example" // string | Filter results by data_service_id (optional)
    versionId := "versionId_example" // string | Filter results by version_id (optional)
    latest := true // bool | Only include the latest image for each version_id. (optional)
    sortBy := "sortBy_example" // string | A given Image attribute to sort results by (one of: id, name, created_at). Ignored when latest=true. (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less). Ignored when latest=true. (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows. Ignored when latest=true. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.ApiImagesGet(context.Background()).DataServiceId(dataServiceId).VersionId(versionId).Latest(latest).SortBy(sortBy).Limit(limit).Continuation(continuation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ApiImagesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiImagesGet`: ModelsPaginatedResultModelsImage
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.ApiImagesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiImagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataServiceId** | **string** | Filter results by data_service_id | 
 **versionId** | **string** | Filter results by version_id | 
 **latest** | **bool** | Only include the latest image for each version_id. | 
 **sortBy** | **string** | A given Image attribute to sort results by (one of: id, name, created_at). Ignored when latest&#x3D;true. | 
 **limit** | **string** | Maximum number of rows to return (could be less). Ignored when latest&#x3D;true. | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows. Ignored when latest&#x3D;true. | 

### Return type

[**ModelsPaginatedResultModelsImage**](ModelsPaginatedResultModelsImage.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiImagesIdGet

> ModelsImage ApiImagesIdGet(ctx, id).Execute()

Get Image



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
    id := "id_example" // string | Image ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.ApiImagesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ApiImagesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiImagesIdGet`: ModelsImage
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.ApiImagesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Image ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiImagesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsImage**](ModelsImage.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiVersionsIdImagesGet

> ModelsPaginatedResultModelsImage ApiVersionsIdImagesGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Execute()

List Version's Images



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
    id := "id_example" // string | Version ID (must be valid UUID)
    sortBy := "sortBy_example" // string | A given Image attribute to sort results by (one of: id, name, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by Image id (optional)
    name := "name_example" // string | Filter results by Image's name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.ApiVersionsIdImagesGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ApiVersionsIdImagesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiVersionsIdImagesGet`: ModelsPaginatedResultModelsImage
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.ApiVersionsIdImagesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Version ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiVersionsIdImagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given Image attribute to sort results by (one of: id, name, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by Image id | 
 **name** | **string** | Filter results by Image&#39;s name | 

### Return type

[**ModelsPaginatedResultModelsImage**](ModelsPaginatedResultModelsImage.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

