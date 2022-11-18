# \DataServicesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDataServicesGet**](DataServicesApi.md#ApiDataServicesGet) | **Get** /api/data-services | List Data Services
[**ApiDataServicesIdGet**](DataServicesApi.md#ApiDataServicesIdGet) | **Get** /api/data-services/{id} | Get Data Service



## ApiDataServicesGet

> ModelsPaginatedResultModelsDataService ApiDataServicesGet(ctx).SortBy(sortBy).Limit(limit).Continuation(continuation).Id(id).Name(name).ShortName(shortName).HasIncrementalBackup(hasIncrementalBackup).HasFullBackup(hasFullBackup).ComingSoon(comingSoon).Execute()

List Data Services



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
    sortBy := "sortBy_example" // string | A given Data Service attribute to sort results by (one of: id, name, short_name, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id := "id_example" // string | Filter results by Data Service ID (optional)
    name := "name_example" // string | Filter results by Data Service name (optional)
    shortName := "shortName_example" // string | Filter results by Data Service short name (optional)
    hasIncrementalBackup := true // bool | Filter results based on incremental backup eligibility (optional)
    hasFullBackup := true // bool | Filter results based on vault full backup eligibility (optional)
    comingSoon := true // bool | Filter results based on 'Coming soon' flag (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataServicesApi.ApiDataServicesGet(context.Background()).SortBy(sortBy).Limit(limit).Continuation(continuation).Id(id).Name(name).ShortName(shortName).HasIncrementalBackup(hasIncrementalBackup).HasFullBackup(hasFullBackup).ComingSoon(comingSoon).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataServicesApi.ApiDataServicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDataServicesGet`: ModelsPaginatedResultModelsDataService
    fmt.Fprintf(os.Stdout, "Response from `DataServicesApi.ApiDataServicesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDataServicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **string** | A given Data Service attribute to sort results by (one of: id, name, short_name, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id** | **string** | Filter results by Data Service ID | 
 **name** | **string** | Filter results by Data Service name | 
 **shortName** | **string** | Filter results by Data Service short name | 
 **hasIncrementalBackup** | **bool** | Filter results based on incremental backup eligibility | 
 **hasFullBackup** | **bool** | Filter results based on vault full backup eligibility | 
 **comingSoon** | **bool** | Filter results based on &#39;Coming soon&#39; flag | 

### Return type

[**ModelsPaginatedResultModelsDataService**](ModelsPaginatedResultModelsDataService.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDataServicesIdGet

> ModelsDataService ApiDataServicesIdGet(ctx, id).Execute()

Get Data Service



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
    id := "id_example" // string | Data Service ID (must be a valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataServicesApi.ApiDataServicesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataServicesApi.ApiDataServicesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDataServicesIdGet`: ModelsDataService
    fmt.Fprintf(os.Stdout, "Response from `DataServicesApi.ApiDataServicesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Data Service ID (must be a valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDataServicesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsDataService**](ModelsDataService.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

