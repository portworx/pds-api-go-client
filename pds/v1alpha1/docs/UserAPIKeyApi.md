# \UserAPIKeyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiUserApiKeyGet**](UserAPIKeyApi.md#ApiUserApiKeyGet) | **Get** /api/user-api-key | List UserAPIKeys
[**ApiUserApiKeyIdDelete**](UserAPIKeyApi.md#ApiUserApiKeyIdDelete) | **Delete** /api/user-api-key/{id} | Delete UserAPIKey
[**ApiUserApiKeyIdGet**](UserAPIKeyApi.md#ApiUserApiKeyIdGet) | **Get** /api/user-api-key/{id} | Get UserAPIKey
[**ApiUserApiKeyIdPatch**](UserAPIKeyApi.md#ApiUserApiKeyIdPatch) | **Patch** /api/user-api-key/{id} | Patch UserAPIKey
[**ApiUserApiKeyPost**](UserAPIKeyApi.md#ApiUserApiKeyPost) | **Post** /api/user-api-key | Create UserAPIKey



## ApiUserApiKeyGet

> ModelsPaginatedResultModelsUserAPIKey ApiUserApiKeyGet(ctx).Id(id).Name(name).Enabled(enabled).SortBy(sortBy).Limit(limit).Continuation(continuation).Execute()

List UserAPIKeys



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
    id := "id_example" // string | Filter results by id (optional)
    name := "name_example" // string | Filter results by name (optional)
    enabled := true // bool | Filter results by enabled (optional)
    sortBy := "sortBy_example" // string | A given UserAPIKey attribute to sort results by (one of: id, name, created_at, enabled, expires_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIKeyApi.ApiUserApiKeyGet(context.Background()).Id(id).Name(name).Enabled(enabled).SortBy(sortBy).Limit(limit).Continuation(continuation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIKeyApi.ApiUserApiKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUserApiKeyGet`: ModelsPaginatedResultModelsUserAPIKey
    fmt.Fprintf(os.Stdout, "Response from `UserAPIKeyApi.ApiUserApiKeyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiUserApiKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Filter results by id | 
 **name** | **string** | Filter results by name | 
 **enabled** | **bool** | Filter results by enabled | 
 **sortBy** | **string** | A given UserAPIKey attribute to sort results by (one of: id, name, created_at, enabled, expires_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 

### Return type

[**ModelsPaginatedResultModelsUserAPIKey**](ModelsPaginatedResultModelsUserAPIKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUserApiKeyIdDelete

> ApiUserApiKeyIdDelete(ctx, id).Execute()

Delete UserAPIKey



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
    id := "id_example" // string | UserAPIKey ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPIKeyApi.ApiUserApiKeyIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIKeyApi.ApiUserApiKeyIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UserAPIKey ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserApiKeyIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUserApiKeyIdGet

> ModelsUserAPIKey ApiUserApiKeyIdGet(ctx, id).Execute()

Get UserAPIKey



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
    id := "id_example" // string | UserAPIKey ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIKeyApi.ApiUserApiKeyIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIKeyApi.ApiUserApiKeyIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUserApiKeyIdGet`: ModelsUserAPIKey
    fmt.Fprintf(os.Stdout, "Response from `UserAPIKeyApi.ApiUserApiKeyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UserAPIKey ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserApiKeyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsUserAPIKey**](ModelsUserAPIKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUserApiKeyIdPatch

> ApiUserApiKeyIdPatch(ctx, id).Body(body).Execute()

Patch UserAPIKey



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
    id := "id_example" // string | UserAPIKey ID (must be valid UUID)
    body := *openapiclient.NewRequestsPatchUserAPIKeyRequest() // RequestsPatchUserAPIKeyRequest | Request body containing the new values.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPIKeyApi.ApiUserApiKeyIdPatch(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIKeyApi.ApiUserApiKeyIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UserAPIKey ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserApiKeyIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestsPatchUserAPIKeyRequest**](RequestsPatchUserAPIKeyRequest.md) | Request body containing the new values. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUserApiKeyPost

> ModelsUserAPIKey ApiUserApiKeyPost(ctx).Body(body).Execute()

Create UserAPIKey



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
    body := *openapiclient.NewRequestsCreateUserAPIKeyRequest() // RequestsCreateUserAPIKeyRequest | Request body containing the necessary information to create a key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIKeyApi.ApiUserApiKeyPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIKeyApi.ApiUserApiKeyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUserApiKeyPost`: ModelsUserAPIKey
    fmt.Fprintf(os.Stdout, "Response from `UserAPIKeyApi.ApiUserApiKeyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiUserApiKeyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RequestsCreateUserAPIKeyRequest**](RequestsCreateUserAPIKeyRequest.md) | Request body containing the necessary information to create a key. | 

### Return type

[**ModelsUserAPIKey**](ModelsUserAPIKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

