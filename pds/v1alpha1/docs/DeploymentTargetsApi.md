# \DeploymentTargetsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDeploymentTargetsIdAgentMetadataPatch**](DeploymentTargetsApi.md#ApiDeploymentTargetsIdAgentMetadataPatch) | **Patch** /api/deployment-targets/{id}/agent-metadata | Patch DeploymentTarget agent-metadata
[**ApiDeploymentTargetsIdConfigGet**](DeploymentTargetsApi.md#ApiDeploymentTargetsIdConfigGet) | **Get** /api/deployment-targets/{id}/config | Get configuration values for a DeploymentTarget
[**ApiDeploymentTargetsIdCredentialsGet**](DeploymentTargetsApi.md#ApiDeploymentTargetsIdCredentialsGet) | **Get** /api/deployment-targets/{id}/credentials | Get join credentials of a DeploymentTarget
[**ApiDeploymentTargetsIdDelete**](DeploymentTargetsApi.md#ApiDeploymentTargetsIdDelete) | **Delete** /api/deployment-targets/{id} | Delete DeploymentTarget
[**ApiDeploymentTargetsIdGet**](DeploymentTargetsApi.md#ApiDeploymentTargetsIdGet) | **Get** /api/deployment-targets/{id} | Get DeploymentTarget
[**ApiDeploymentTargetsIdHeartbeatPost**](DeploymentTargetsApi.md#ApiDeploymentTargetsIdHeartbeatPost) | **Post** /api/deployment-targets/{id}/heartbeat | Make DeploymentTarget heart beat request
[**ApiDeploymentTargetsIdMetadataPost**](DeploymentTargetsApi.md#ApiDeploymentTargetsIdMetadataPost) | **Post** /api/deployment-targets/{id}/metadata | Update DeploymentTarget metadata
[**ApiDeploymentTargetsIdOperatorHeartbeatPost**](DeploymentTargetsApi.md#ApiDeploymentTargetsIdOperatorHeartbeatPost) | **Post** /api/deployment-targets/{id}/operator-heartbeat | Make DeploymentTarget operator heart beat request
[**ApiDeploymentTargetsIdOperatorMetadataPatch**](DeploymentTargetsApi.md#ApiDeploymentTargetsIdOperatorMetadataPatch) | **Patch** /api/deployment-targets/{id}/operator-metadata | Update target operator metadata
[**ApiDeploymentTargetsIdPut**](DeploymentTargetsApi.md#ApiDeploymentTargetsIdPut) | **Put** /api/deployment-targets/{id} | Update DeploymentTarget
[**ApiDeploymentTargetsIdUpdateCapabilitiesPost**](DeploymentTargetsApi.md#ApiDeploymentTargetsIdUpdateCapabilitiesPost) | **Post** /api/deployment-targets/{id}/update-capabilities | Update target capabilities
[**ApiProjectsIdDeploymentTargetsGet**](DeploymentTargetsApi.md#ApiProjectsIdDeploymentTargetsGet) | **Get** /api/projects/{id}/deployment-targets | List Project&#39;s DeploymentTargets
[**ApiTenantsIdDeploymentTargetsGet**](DeploymentTargetsApi.md#ApiTenantsIdDeploymentTargetsGet) | **Get** /api/tenants/{id}/deployment-targets | List Tenant&#39;s DeploymentTargets
[**ApiTenantsIdDeploymentTargetsPost**](DeploymentTargetsApi.md#ApiTenantsIdDeploymentTargetsPost) | **Post** /api/tenants/{id}/deployment-targets | Create DeploymentTarget



## ApiDeploymentTargetsIdAgentMetadataPatch

> ApiDeploymentTargetsIdAgentMetadataPatch(ctx, id).Body(body).Execute()

Patch DeploymentTarget agent-metadata



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
    body := *openapiclient.NewRequestsPatchDeploymentTargetsAgentMetadataRequest() // RequestsPatchDeploymentTargetsAgentMetadataRequest | Object with target cluster ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentTargetsApi.ApiDeploymentTargetsIdAgentMetadataPatch(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTargetsApi.ApiDeploymentTargetsIdAgentMetadataPatch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiDeploymentTargetsIdAgentMetadataPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestsPatchDeploymentTargetsAgentMetadataRequest**](RequestsPatchDeploymentTargetsAgentMetadataRequest.md) | Object with target cluster ID | 

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


## ApiDeploymentTargetsIdConfigGet

> ControllersTargetClusterConfigResponse ApiDeploymentTargetsIdConfigGet(ctx, id).Execute()

Get configuration values for a DeploymentTarget



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentTargetsApi.ApiDeploymentTargetsIdConfigGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTargetsApi.ApiDeploymentTargetsIdConfigGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentTargetsIdConfigGet`: ControllersTargetClusterConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `DeploymentTargetsApi.ApiDeploymentTargetsIdConfigGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DeploymentTarget ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentTargetsIdConfigGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ControllersTargetClusterConfigResponse**](ControllersTargetClusterConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentTargetsApi.ApiDeploymentTargetsIdCredentialsGet(context.Background(), id).Execute()
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentTargetsApi.ApiDeploymentTargetsIdDelete(context.Background(), id).Execute()
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentTargetsApi.ApiDeploymentTargetsIdGet(context.Background(), id).Execute()
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentTargetsApi.ApiDeploymentTargetsIdHeartbeatPost(context.Background(), id).Body(body).Execute()
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


## ApiDeploymentTargetsIdMetadataPost

> ApiDeploymentTargetsIdMetadataPost(ctx, id).Body(body).Execute()

Update DeploymentTarget metadata



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
    body := *openapiclient.NewControllersDeploymentTargetMetadataRequest() // ControllersDeploymentTargetMetadataRequest | Object with target cluster ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentTargetsApi.ApiDeploymentTargetsIdMetadataPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTargetsApi.ApiDeploymentTargetsIdMetadataPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiDeploymentTargetsIdMetadataPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersDeploymentTargetMetadataRequest**](ControllersDeploymentTargetMetadataRequest.md) | Object with target cluster ID | 

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


## ApiDeploymentTargetsIdOperatorHeartbeatPost

> ControllersDeploymentTargetHeartbeatResponse ApiDeploymentTargetsIdOperatorHeartbeatPost(ctx, id).Body(body).Execute()

Make DeploymentTarget operator heart beat request



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentTargetsApi.ApiDeploymentTargetsIdOperatorHeartbeatPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTargetsApi.ApiDeploymentTargetsIdOperatorHeartbeatPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentTargetsIdOperatorHeartbeatPost`: ControllersDeploymentTargetHeartbeatResponse
    fmt.Fprintf(os.Stdout, "Response from `DeploymentTargetsApi.ApiDeploymentTargetsIdOperatorHeartbeatPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DeploymentTarget ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentTargetsIdOperatorHeartbeatPostRequest struct via the builder pattern


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


## ApiDeploymentTargetsIdOperatorMetadataPatch

> ApiDeploymentTargetsIdOperatorMetadataPatch(ctx, id).Body(body).Execute()

Update target operator metadata



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
    body := *openapiclient.NewRequestsUpdateOperatorMetadataRequest() // RequestsUpdateOperatorMetadataRequest | Body with the new metadata.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentTargetsApi.ApiDeploymentTargetsIdOperatorMetadataPatch(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTargetsApi.ApiDeploymentTargetsIdOperatorMetadataPatch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiDeploymentTargetsIdOperatorMetadataPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestsUpdateOperatorMetadataRequest**](RequestsUpdateOperatorMetadataRequest.md) | Body with the new metadata. | 

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentTargetsApi.ApiDeploymentTargetsIdPut(context.Background(), id).Body(body).Execute()
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


## ApiDeploymentTargetsIdUpdateCapabilitiesPost

> ApiDeploymentTargetsIdUpdateCapabilitiesPost(ctx, id).Body(body).Execute()

Update target capabilities



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
    body := *openapiclient.NewModelsDeploymentTargetCapabilities() // ModelsDeploymentTargetCapabilities | Body containing supported capabilities versions (must be valid semver).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentTargetsApi.ApiDeploymentTargetsIdUpdateCapabilitiesPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTargetsApi.ApiDeploymentTargetsIdUpdateCapabilitiesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiDeploymentTargetsIdUpdateCapabilitiesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ModelsDeploymentTargetCapabilities**](ModelsDeploymentTargetCapabilities.md) | Body containing supported capabilities versions (must be valid semver). | 

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


## ApiProjectsIdDeploymentTargetsGet

> ModelsPaginatedResultModelsDeploymentTarget ApiProjectsIdDeploymentTargetsGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).ClusterId(clusterId).Name(name).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentTargetsApi.ApiProjectsIdDeploymentTargetsGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).ClusterId(clusterId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTargetsApi.ApiProjectsIdDeploymentTargetsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectsIdDeploymentTargetsGet`: ModelsPaginatedResultModelsDeploymentTarget
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

[**ModelsPaginatedResultModelsDeploymentTarget**](ModelsPaginatedResultModelsDeploymentTarget.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdDeploymentTargetsGet

> ModelsPaginatedResultModelsDeploymentTarget ApiTenantsIdDeploymentTargetsGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).ClusterId(clusterId).Name(name).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentTargetsApi.ApiTenantsIdDeploymentTargetsGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).ClusterId(clusterId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTargetsApi.ApiTenantsIdDeploymentTargetsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdDeploymentTargetsGet`: ModelsPaginatedResultModelsDeploymentTarget
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

[**ModelsPaginatedResultModelsDeploymentTarget**](ModelsPaginatedResultModelsDeploymentTarget.md)

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
    body := *openapiclient.NewRequestsCreateDeploymentTargetRequest() // RequestsCreateDeploymentTargetRequest | Request body containing the deployment target config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentTargetsApi.ApiTenantsIdDeploymentTargetsPost(context.Background(), id).Body(body).Execute()
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

 **body** | [**RequestsCreateDeploymentTargetRequest**](RequestsCreateDeploymentTargetRequest.md) | Request body containing the deployment target config | 

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

