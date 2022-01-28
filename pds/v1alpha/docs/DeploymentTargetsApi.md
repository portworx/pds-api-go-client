# \DeploymentTargetsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDeploymentTargetsIdGet**](DeploymentTargetsApi.md#ApiDeploymentTargetsIdGet) | **Get** /api/deployment-targets/{id} | Get DeploymentTarget
[**ApiDeploymentTargetsIdPut**](DeploymentTargetsApi.md#ApiDeploymentTargetsIdPut) | **Put** /api/deployment-targets/{id} | Update DeploymentTarget
[**ApiDeploymentTargetsPost**](DeploymentTargetsApi.md#ApiDeploymentTargetsPost) | **Post** /api/deployment-targets | Create DeploymentTarget
[**ApiTenantsIdDeploymentTargetsGet**](DeploymentTargetsApi.md#ApiTenantsIdDeploymentTargetsGet) | **Get** /api/tenants/{id}/deployment-targets | List Tenant&#39;s DeploymentTargets



## ApiDeploymentTargetsIdGet

> ModelsDeploymentTarget ApiDeploymentTargetsIdGet(ctx, id).Execute()

Get DeploymentTarget



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentTargetsApi.ApiDeploymentTargetsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTargetsApi.ApiDeploymentTargetsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentTargetsIdGet`: ModelsDeploymentTarget
    fmt.Fprintf(os.Stdout, "Response from `DeploymentTargetsApi.ApiDeploymentTargetsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DeploymentTarget ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentTargetsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsDeploymentTarget**](ModelsDeploymentTarget.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeploymentTargetsIdPut

> ModelsDeploymentTarget ApiDeploymentTargetsIdPut(ctx, id).Body(body).Execute()

Update DeploymentTarget



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
    body := *openapiclient.NewControllersUpdateDeploymentTargetRequest() // ControllersUpdateDeploymentTargetRequest | Object with partial update fields (name, tenant_id and cluster_id are mandatory)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentTargetsApi.ApiDeploymentTargetsIdPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTargetsApi.ApiDeploymentTargetsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentTargetsIdPut`: ModelsDeploymentTarget
    fmt.Fprintf(os.Stdout, "Response from `DeploymentTargetsApi.ApiDeploymentTargetsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DeploymentTarget ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentTargetsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersUpdateDeploymentTargetRequest**](ControllersUpdateDeploymentTargetRequest.md) | Object with partial update fields (name, tenant_id and cluster_id are mandatory) | 

### Return type

[**ModelsDeploymentTarget**](ModelsDeploymentTarget.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeploymentTargetsPost

> ModelsDeploymentTarget ApiDeploymentTargetsPost(ctx).Body(body).Execute()

Create DeploymentTarget



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
    body := *openapiclient.NewControllersCreateDeploymentTargetRequest() // ControllersCreateDeploymentTargetRequest | Request body containing the deployment target config

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentTargetsApi.ApiDeploymentTargetsPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTargetsApi.ApiDeploymentTargetsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentTargetsPost`: ModelsDeploymentTarget
    fmt.Fprintf(os.Stdout, "Response from `DeploymentTargetsApi.ApiDeploymentTargetsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentTargetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ControllersCreateDeploymentTargetRequest**](ControllersCreateDeploymentTargetRequest.md) | Request body containing the deployment target config | 

### Return type

[**ModelsDeploymentTarget**](ModelsDeploymentTarget.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdDeploymentTargetsGet

> ControllersPaginatedDeploymentTargets ApiTenantsIdDeploymentTargetsGet(ctx, id).SortBy(sortBy).Id2(id2).Name(name).Execute()

List Tenant's DeploymentTargets



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
    id := "id_example" // string | Tenant ID (must be valid UUID)
    sortBy := "sortBy_example" // string | A given DeploymentTarget attribute to sort results by (one of: id, name) (optional)
    id2 := "id_example" // string | Filter results by DeploymentTarget ID (optional)
    name := "name_example" // string | Filter results by DeploymentTarget name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentTargetsApi.ApiTenantsIdDeploymentTargetsGet(context.Background(), id).SortBy(sortBy).Id2(id2).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTargetsApi.ApiTenantsIdDeploymentTargetsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdDeploymentTargetsGet`: ControllersPaginatedDeploymentTargets
    fmt.Fprintf(os.Stdout, "Response from `DeploymentTargetsApi.ApiTenantsIdDeploymentTargetsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdDeploymentTargetsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given DeploymentTarget attribute to sort results by (one of: id, name) | 
 **id2** | **string** | Filter results by DeploymentTarget ID | 
 **name** | **string** | Filter results by DeploymentTarget name | 

### Return type

[**ControllersPaginatedDeploymentTargets**](ControllersPaginatedDeploymentTargets.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

