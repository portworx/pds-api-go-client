# \VersionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDataServicesIdVersionsGet**](VersionsApi.md#ApiDataServicesIdVersionsGet) | **Get** /api/data-services/{id}/versions | List Data Service&#39;s Versions
[**ApiVersionsIdGet**](VersionsApi.md#ApiVersionsIdGet) | **Get** /api/versions/{id} | Get Version
[**ApiVersionsIdLatestPatchPost**](VersionsApi.md#ApiVersionsIdLatestPatchPost) | **Post** /api/versions/{id}/latest-patch | Get latest patch version.
[**ApiVersionsIdListCompatiblePost**](VersionsApi.md#ApiVersionsIdListCompatiblePost) | **Post** /api/versions/{id}/list-compatible | List compatible versions.



## ApiDataServicesIdVersionsGet

> ControllersPaginatedVersions ApiDataServicesIdVersionsGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Enabled(enabled).Execute()

List Data Service's Versions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Data Service ID (must be valid UUID)
    sortBy := "sortBy_example" // string | A given Version attribute to sort results by (one of: id, name, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by Version id (optional)
    name := "name_example" // string | Filter results by Version's name (optional)
    enabled := true // bool | Filter results by Version's enabled parameter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsApi.ApiDataServicesIdVersionsGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Enabled(enabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.ApiDataServicesIdVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDataServicesIdVersionsGet`: ControllersPaginatedVersions
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.ApiDataServicesIdVersionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Data Service ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDataServicesIdVersionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given Version attribute to sort results by (one of: id, name, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by Version id | 
 **name** | **string** | Filter results by Version&#39;s name | 
 **enabled** | **bool** | Filter results by Version&#39;s enabled parameter | 

### Return type

[**ControllersPaginatedVersions**](ControllersPaginatedVersions.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiVersionsIdGet

> ModelsVersion ApiVersionsIdGet(ctx, id).Execute()

Get Version



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Version ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsApi.ApiVersionsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.ApiVersionsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiVersionsIdGet`: ModelsVersion
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.ApiVersionsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Version ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiVersionsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsVersion**](ModelsVersion.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiVersionsIdLatestPatchPost

> ModelsVersion ApiVersionsIdLatestPatchPost(ctx, id).Execute()

Get latest patch version.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Version ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsApi.ApiVersionsIdLatestPatchPost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.ApiVersionsIdLatestPatchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiVersionsIdLatestPatchPost`: ModelsVersion
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.ApiVersionsIdLatestPatchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Version ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiVersionsIdLatestPatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsVersion**](ModelsVersion.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiVersionsIdListCompatiblePost

> []ModelsVersion ApiVersionsIdListCompatiblePost(ctx, id).Latest(latest).Execute()

List compatible versions.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Version ID (must be valid UUID)
    latest := "latest_example" // string | List latest patch versions of compatible major/minor versions for the given version (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsApi.ApiVersionsIdListCompatiblePost(context.Background(), id).Latest(latest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.ApiVersionsIdListCompatiblePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiVersionsIdListCompatiblePost`: []ModelsVersion
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.ApiVersionsIdListCompatiblePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Version ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiVersionsIdListCompatiblePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **latest** | **string** | List latest patch versions of compatible major/minor versions for the given version | 

### Return type

[**[]ModelsVersion**](ModelsVersion.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

