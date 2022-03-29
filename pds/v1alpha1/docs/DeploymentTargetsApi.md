# \DeploymentTargetsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDeploymentTargetsIdCredentialsGet**](DeploymentTargetsApi.md#ApiDeploymentTargetsIdCredentialsGet) | **Get** /api/deployment-targets/{id}/credentials | Get join credentials of a DeploymentTarget
[**ApiDeploymentTargetsIdDelete**](DeploymentTargetsApi.md#ApiDeploymentTargetsIdDelete) | **Delete** /api/deployment-targets/{id} | Delete DeploymentTarget
[**ApiDeploymentTargetsIdGet**](DeploymentTargetsApi.md#ApiDeploymentTargetsIdGet) | **Get** /api/deployment-targets/{id} | Get DeploymentTarget
[**ApiDeploymentTargetsIdHeartbeatPost**](DeploymentTargetsApi.md#ApiDeploymentTargetsIdHeartbeatPost) | **Post** /api/deployment-targets/{id}/heartbeat | Make DeploymentTarget heart beat request
[**ApiDeploymentTargetsIdPut**](DeploymentTargetsApi.md#ApiDeploymentTargetsIdPut) | **Put** /api/deployment-targets/{id} | Update DeploymentTarget
[**ApiProjectsIdDeploymentTargetsGet**](DeploymentTargetsApi.md#ApiProjectsIdDeploymentTargetsGet) | **Get** /api/projects/{id}/deployment-targets | List Project&#39;s DeploymentTargets
[**ApiTenantsIdDeploymentTargetsGet**](DeploymentTargetsApi.md#ApiTenantsIdDeploymentTargetsGet) | **Get** /api/tenants/{id}/deployment-targets | List Tenant&#39;s DeploymentTargets
[**ApiTenantsIdDeploymentTargetsPost**](DeploymentTargetsApi.md#ApiTenantsIdDeploymentTargetsPost) | **Post** /api/tenants/{id}/deployment-targets | Create DeploymentTarget



## ApiDeploymentTargetsIdCredentialsGet

> ControllersDeploymentTargetCredentialsResponse ApiDeploymentTargetsIdCredentialsGet(ctx, id).Execute()

Get join credentials of a DeploymentTarget



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
    resp, r, err := api_client.DeploymentTargetsApi.ApiDeploymentTargetsIdCredentialsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTargetsApi.ApiDeploymentTargetsIdCredentialsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentTargetsIdCredentialsGet`: ControllersDeploymentTargetCredentialsResponse
    fmt.Fprintf(os.Stdout, "Response from `DeploymentTargetsApi.ApiDeploymentTargetsIdCredentialsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DeploymentTarget ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentTargetsIdCredentialsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ControllersDeploymentTargetCredentialsResponse**](ControllersDeploymentTargetCredentialsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeploymentTargetsIdDelete

> ApiDeploymentTargetsIdDelete(ctx, id).Execute()

Delete DeploymentTarget



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
    resp, r, err := api_client.DeploymentTargetsApi.ApiDeploymentTargetsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTargetsApi.ApiDeploymentTargetsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DeploymentTarget ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentTargetsIdDeleteRequest struct via the builder pattern


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


## ApiDeploymentTargetsIdHeartbeatPost

> ControllersDeploymentTargetHeartbeatResponse ApiDeploymentTargetsIdHeartbeatPost(ctx, id).Body(body).Execute()

Make DeploymentTarget heart beat request



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
    body := *openapiclient.NewControllersDeploymentTargetHeartbeatRequest() // ControllersDeploymentTargetHeartbeatRequest | Object with target cluster ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentTargetsApi.ApiDeploymentTargetsIdHeartbeatPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTargetsApi.ApiDeploymentTargetsIdHeartbeatPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentTargetsIdHeartbeatPost`: ControllersDeploymentTargetHeartbeatResponse
    fmt.Fprintf(os.Stdout, "Response from `DeploymentTargetsApi.ApiDeploymentTargetsIdHeartbeatPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DeploymentTarget ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentTargetsIdHeartbeatPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersDeploymentTargetHeartbeatRequest**](ControllersDeploymentTargetHeartbeatRequest.md) | Object with target cluster ID | 

### Return type

[**ControllersDeploymentTargetHeartbeatResponse**](ControllersDeploymentTargetHeartbeatResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
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
    body := *openapiclient.NewControllersUpdateDeploymentTargetRequest() // ControllersUpdateDeploymentTargetRequest | Object with partial update fields (name)

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

 **body** | [**ControllersUpdateDeploymentTargetRequest**](ControllersUpdateDeploymentTargetRequest.md) | Object with partial update fields (name) | 

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


## ApiProjectsIdDeploymentTargetsGet

> ControllersPaginatedDeploymentTargets ApiProjectsIdDeploymentTargetsGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).ClusterId(clusterId).Name(name).Execute()

List Project's DeploymentTargets



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
    id := "id_example" // string | Project ID (must be valid UUID)
    sortBy := "sortBy_example" // string | A given DeploymentTarget attribute to sort results by (one of: id, name) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by DeploymentTarget ID (optional)
    clusterId := "clusterId_example" // string | Filter results by Cluster ID (optional)
    name := "name_example" // string | Filter results by DeploymentTarget name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentTargetsApi.ApiProjectsIdDeploymentTargetsGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).ClusterId(clusterId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTargetsApi.ApiProjectsIdDeploymentTargetsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectsIdDeploymentTargetsGet`: ControllersPaginatedDeploymentTargets
    fmt.Fprintf(os.Stdout, "Response from `DeploymentTargetsApi.ApiProjectsIdDeploymentTargetsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsIdDeploymentTargetsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given DeploymentTarget attribute to sort results by (one of: id, name) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by DeploymentTarget ID | 
 **clusterId** | **string** | Filter results by Cluster ID | 
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


## ApiTenantsIdDeploymentTargetsGet

> ControllersPaginatedDeploymentTargets ApiTenantsIdDeploymentTargetsGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).ClusterId(clusterId).Name(name).Execute()

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
    sortBy := "sortBy_example" // string | A given DeploymentTarget attribute to sort results by (one of: id, name, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by DeploymentTarget ID (optional)
    clusterId := "clusterId_example" // string | Filter results by Cluster ID (optional)
    name := "name_example" // string | Filter results by DeploymentTarget name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentTargetsApi.ApiTenantsIdDeploymentTargetsGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).ClusterId(clusterId).Name(name).Execute()
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

 **sortBy** | **string** | A given DeploymentTarget attribute to sort results by (one of: id, name, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by DeploymentTarget ID | 
 **clusterId** | **string** | Filter results by Cluster ID | 
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


## ApiTenantsIdDeploymentTargetsPost

> ModelsDeploymentTarget ApiTenantsIdDeploymentTargetsPost(ctx, id).Body(body).Execute()

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
    id := "id_example" // string | Tenant ID (must be valid UUID)
    body := *openapiclient.NewControllersCreateTenantDeploymentTarget() // ControllersCreateTenantDeploymentTarget | Request body containing the deployment target config

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentTargetsApi.ApiTenantsIdDeploymentTargetsPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTargetsApi.ApiTenantsIdDeploymentTargetsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdDeploymentTargetsPost`: ModelsDeploymentTarget
    fmt.Fprintf(os.Stdout, "Response from `DeploymentTargetsApi.ApiTenantsIdDeploymentTargetsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdDeploymentTargetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersCreateTenantDeploymentTarget**](ControllersCreateTenantDeploymentTarget.md) | Request body containing the deployment target config | 

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

