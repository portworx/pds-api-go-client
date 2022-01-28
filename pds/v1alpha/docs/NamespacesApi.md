# \NamespacesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDeploymentTargetsIdNamespacesPost**](NamespacesApi.md#ApiDeploymentTargetsIdNamespacesPost) | **Post** /api/deployment-targets/{id}/namespaces | Create Namespace
[**ApiNamespacesIdGet**](NamespacesApi.md#ApiNamespacesIdGet) | **Get** /api/namespaces/{id} | Get Namespace
[**ApiNamespacesPost**](NamespacesApi.md#ApiNamespacesPost) | **Post** /api/namespaces | Create Namespace



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.ApiDeploymentTargetsIdNamespacesPost(context.Background(), id).Body(body).Execute()
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.ApiNamespacesIdGet(context.Background(), id).Execute()
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


## ApiNamespacesPost

> ModelsNamespace ApiNamespacesPost(ctx).Body(body).Execute()

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
    body := *openapiclient.NewControllersCreateNamespaceRequest() // ControllersCreateNamespaceRequest | Request body containing the new namespace

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.ApiNamespacesPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.ApiNamespacesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiNamespacesPost`: ModelsNamespace
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.ApiNamespacesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiNamespacesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ControllersCreateNamespaceRequest**](ControllersCreateNamespaceRequest.md) | Request body containing the new namespace | 

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

