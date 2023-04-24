# \VersionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCompatibleVersionsGet**](VersionsApi.md#ApiCompatibleVersionsGet) | **Get** /api/compatible-versions | List compatible versions
[**ApiDataServicesIdVersionsGet**](VersionsApi.md#ApiDataServicesIdVersionsGet) | **Get** /api/data-services/{id}/versions | List Data Service&#39;s Versions
[**ApiVersionsIdGet**](VersionsApi.md#ApiVersionsIdGet) | **Get** /api/versions/{id} | Get Version



## ApiCompatibleVersionsGet

> ControllersCompatibleVersionsResponse ApiCompatibleVersionsGet(ctx).DataServiceId(dataServiceId).VersionId(versionId).Execute()

List compatible versions



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsApi.ApiCompatibleVersionsGet(context.Background()).DataServiceId(dataServiceId).VersionId(versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.ApiCompatibleVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiCompatibleVersionsGet`: ControllersCompatibleVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.ApiCompatibleVersionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiCompatibleVersionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataServiceId** | **string** | Filter results by data_service_id | 
 **versionId** | **string** | Filter results by version_id | 

### Return type

[**ControllersCompatibleVersionsResponse**](ControllersCompatibleVersionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDataServicesIdVersionsGet

> ModelsPaginatedResultModelsVersion ApiDataServicesIdVersionsGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Enabled(enabled).Execute()

List Data Service's Versions



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
    // response from `ApiDataServicesIdVersionsGet`: ModelsPaginatedResultModelsVersion
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

[**ModelsPaginatedResultModelsVersion**](ModelsPaginatedResultModelsVersion.md)

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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

