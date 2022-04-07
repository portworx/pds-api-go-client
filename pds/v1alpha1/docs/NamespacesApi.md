# \NamespacesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDeploymentTargetsIdNamespacesGet**](NamespacesApi.md#ApiDeploymentTargetsIdNamespacesGet) | **Get** /api/deployment-targets/{id}/namespaces | List DeploymentTarget&#39;s Namespaces
[**ApiDeploymentTargetsIdNamespacesPost**](NamespacesApi.md#ApiDeploymentTargetsIdNamespacesPost) | **Post** /api/deployment-targets/{id}/namespaces | Create Namespace
[**ApiNamespacesIdDelete**](NamespacesApi.md#ApiNamespacesIdDelete) | **Delete** /api/namespaces/{id} | Delete Namespace
[**ApiNamespacesIdGet**](NamespacesApi.md#ApiNamespacesIdGet) | **Get** /api/namespaces/{id} | Get Namespace
[**ApiNamespacesIdPut**](NamespacesApi.md#ApiNamespacesIdPut) | **Put** /api/namespaces/{id} | Update namespace



## ApiDeploymentTargetsIdNamespacesGet

> ControllersPaginatedNamespaces ApiDeploymentTargetsIdNamespacesGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Status(status).Execute()

List DeploymentTarget's Namespaces



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
    id := "id_example" // string | DeploymentTarget ID (must be valid UUID)
    sortBy := "sortBy_example" // string | A given Namespace attribute to sort results by (one of: id, name, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by Namespace ID (optional)
    name := "name_example" // string | Filter results by Namespace name (optional)
    status := "status_example" // string | Filter results by Namespace status (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespacesApi.ApiDeploymentTargetsIdNamespacesGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.ApiDeploymentTargetsIdNamespacesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentTargetsIdNamespacesGet`: ControllersPaginatedNamespaces
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.ApiDeploymentTargetsIdNamespacesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DeploymentTarget ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentTargetsIdNamespacesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given Namespace attribute to sort results by (one of: id, name, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by Namespace ID | 
 **name** | **string** | Filter results by Namespace name | 
 **status** | **string** | Filter results by Namespace status | 

### Return type

[**ControllersPaginatedNamespaces**](ControllersPaginatedNamespaces.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeploymentTargetsIdNamespacesPost

> ModelsNamespace ApiDeploymentTargetsIdNamespacesPost(ctx, id).Body(body).Execute()

Create Namespace



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
    id := "id_example" // string | DeploymentTarget ID (must be valid UUID)
    body := *openapiclient.NewControllersCreateNamespace() // ControllersCreateNamespace | Request body containing the new namespace

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespacesApi.ApiDeploymentTargetsIdNamespacesPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.ApiDeploymentTargetsIdNamespacesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentTargetsIdNamespacesPost`: ModelsNamespace
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.ApiDeploymentTargetsIdNamespacesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DeploymentTarget ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentTargetsIdNamespacesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersCreateNamespace**](ControllersCreateNamespace.md) | Request body containing the new namespace | 

### Return type

[**ModelsNamespace**](ModelsNamespace.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiNamespacesIdDelete

> ApiNamespacesIdDelete(ctx, id).Execute()

Delete Namespace



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
    id := "id_example" // string | Namespace ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespacesApi.ApiNamespacesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.ApiNamespacesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Namespace ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiNamespacesIdDeleteRequest struct via the builder pattern


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


## ApiNamespacesIdGet

> ModelsNamespace ApiNamespacesIdGet(ctx, id).Execute()

Get Namespace



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
    id := "id_example" // string | Namespace ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespacesApi.ApiNamespacesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.ApiNamespacesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiNamespacesIdGet`: ModelsNamespace
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.ApiNamespacesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Namespace ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiNamespacesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsNamespace**](ModelsNamespace.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiNamespacesIdPut

> ModelsNamespace ApiNamespacesIdPut(ctx, id).Body(body).Execute()

Update namespace



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
    id := "id_example" // string | Namespace ID (must be valid UUID)
    body := *openapiclient.NewControllersUpdateNamespaceRequest() // ControllersUpdateNamespaceRequest | Object with partial update of fields (status)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespacesApi.ApiNamespacesIdPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.ApiNamespacesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiNamespacesIdPut`: ModelsNamespace
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.ApiNamespacesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Namespace ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiNamespacesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersUpdateNamespaceRequest**](ControllersUpdateNamespaceRequest.md) | Object with partial update of fields (status) | 

### Return type

[**ModelsNamespace**](ModelsNamespace.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

