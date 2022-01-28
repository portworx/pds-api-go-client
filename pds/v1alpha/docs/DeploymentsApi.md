# \DeploymentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDeploymentsIdBackupsGet**](DeploymentsApi.md#ApiDeploymentsIdBackupsGet) | **Get** /api/deployments/{id}/backups | Get Deployment Backups
[**ApiDeploymentsIdConnectionStringsGet**](DeploymentsApi.md#ApiDeploymentsIdConnectionStringsGet) | **Get** /api/deployments/{id}/connection-strings | Deployment Connection String
[**ApiDeploymentsIdDelete**](DeploymentsApi.md#ApiDeploymentsIdDelete) | **Delete** /api/deployments/{id} | Delete Deployment
[**ApiDeploymentsIdGet**](DeploymentsApi.md#ApiDeploymentsIdGet) | **Get** /api/deployments/{id} | Get Deployment
[**ApiDeploymentsIdHealthGet**](DeploymentsApi.md#ApiDeploymentsIdHealthGet) | **Get** /api/deployments/{id}/health | Get Deployment Health
[**ApiDeploymentsIdPut**](DeploymentsApi.md#ApiDeploymentsIdPut) | **Put** /api/deployments/{id} | Update Deployment
[**ApiDeploymentsIdStatusGet**](DeploymentsApi.md#ApiDeploymentsIdStatusGet) | **Get** /api/deployments/{id}/status | Get Deployment Status
[**ApiDeploymentsPost**](DeploymentsApi.md#ApiDeploymentsPost) | **Post** /api/deployments | Create Deployment
[**ApiProjectsIdDeploymentsGet**](DeploymentsApi.md#ApiProjectsIdDeploymentsGet) | **Get** /api/projects/{id}/deployments | List Project&#39;s Deployments



## ApiDeploymentsIdBackupsGet

> PaginationPaginatedResponse ApiDeploymentsIdBackupsGet(ctx, id).Execute()

Get Deployment Backups



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
    id := "id_example" // string | Deployment ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.ApiDeploymentsIdBackupsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ApiDeploymentsIdBackupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentsIdBackupsGet`: PaginationPaginatedResponse
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ApiDeploymentsIdBackupsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Deployment ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentsIdBackupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaginationPaginatedResponse**](PaginationPaginatedResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeploymentsIdConnectionStringsGet

> ControllersDeploymentConnectionStringsResponse ApiDeploymentsIdConnectionStringsGet(ctx, id).Execute()

Deployment Connection String



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
    id := "id_example" // string | Deployment ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.ApiDeploymentsIdConnectionStringsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ApiDeploymentsIdConnectionStringsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentsIdConnectionStringsGet`: ControllersDeploymentConnectionStringsResponse
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ApiDeploymentsIdConnectionStringsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Deployment ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentsIdConnectionStringsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ControllersDeploymentConnectionStringsResponse**](ControllersDeploymentConnectionStringsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeploymentsIdDelete

> ApiDeploymentsIdDelete(ctx, id).LocalOnly(localOnly).Execute()

Delete Deployment



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
    id := "id_example" // string | Deployment ID (must be a valid UUID)
    localOnly := true // bool | Set to true to only delete the Deployment object in the database (does not delete any actual resources) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.ApiDeploymentsIdDelete(context.Background(), id).LocalOnly(localOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ApiDeploymentsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Deployment ID (must be a valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **localOnly** | **bool** | Set to true to only delete the Deployment object in the database (does not delete any actual resources) | 

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


## ApiDeploymentsIdGet

> ModelsDeployment ApiDeploymentsIdGet(ctx, id).Execute()

Get Deployment



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
    id := "id_example" // string | Deployment ID (either id or deployment_id field)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.ApiDeploymentsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ApiDeploymentsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentsIdGet`: ModelsDeployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ApiDeploymentsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Deployment ID (either id or deployment_id field) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsDeployment**](ModelsDeployment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeploymentsIdHealthGet

> StatusDeploymentHealth ApiDeploymentsIdHealthGet(ctx, id).Execute()

Get Deployment Health



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
    id := "id_example" // string | Deployment ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.ApiDeploymentsIdHealthGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ApiDeploymentsIdHealthGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentsIdHealthGet`: StatusDeploymentHealth
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ApiDeploymentsIdHealthGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Deployment ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentsIdHealthGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatusDeploymentHealth**](StatusDeploymentHealth.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeploymentsIdPut

> ModelsDeployment ApiDeploymentsIdPut(ctx, id).Body(body).LocalOnly(localOnly).Execute()

Update Deployment



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
    id := "id_example" // string | Deployment ID (must be valid UUID)
    body := *openapiclient.NewControllersUpdateDeploymentRequest() // ControllersUpdateDeploymentRequest | Request body containing the updated deployment config
    localOnly := true // bool | Set to true to only update the Deployment object in the database (does not create any actual resources) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.ApiDeploymentsIdPut(context.Background(), id).Body(body).LocalOnly(localOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ApiDeploymentsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentsIdPut`: ModelsDeployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ApiDeploymentsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Deployment ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersUpdateDeploymentRequest**](ControllersUpdateDeploymentRequest.md) | Request body containing the updated deployment config | 
 **localOnly** | **bool** | Set to true to only update the Deployment object in the database (does not create any actual resources) | 

### Return type

[**ModelsDeployment**](ModelsDeployment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeploymentsIdStatusGet

> StatusDeploymentStatus ApiDeploymentsIdStatusGet(ctx, id).Execute()

Get Deployment Status



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
    id := "id_example" // string | Deployment ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.ApiDeploymentsIdStatusGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ApiDeploymentsIdStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentsIdStatusGet`: StatusDeploymentStatus
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ApiDeploymentsIdStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Deployment ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentsIdStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatusDeploymentStatus**](StatusDeploymentStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeploymentsPost

> ModelsDeployment ApiDeploymentsPost(ctx).Body(body).LocalOnly(localOnly).Execute()

Create Deployment



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
    body := *openapiclient.NewControllersCreateDeploymentRequest() // ControllersCreateDeploymentRequest | Request body containing the deployment config
    localOnly := true // bool | Set to true to only create the Deployment object in the database (does not create any actual resources) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.ApiDeploymentsPost(context.Background()).Body(body).LocalOnly(localOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ApiDeploymentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentsPost`: ModelsDeployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ApiDeploymentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ControllersCreateDeploymentRequest**](ControllersCreateDeploymentRequest.md) | Request body containing the deployment config | 
 **localOnly** | **bool** | Set to true to only create the Deployment object in the database (does not create any actual resources) | 

### Return type

[**ModelsDeployment**](ModelsDeployment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectsIdDeploymentsGet

> ControllersPaginatedDeployments ApiProjectsIdDeploymentsGet(ctx, id).SortBy(sortBy).Id2(id2).ClusterName(clusterName).DataServiceId(dataServiceId).DeploymentTargetId(deploymentTargetId).ImageId(imageId).NamespaceId(namespaceId).Execute()

List Project's Deployments



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
    sortBy := "sortBy_example" // string | A given Deployment attribute to sort results by (one of: id, cluster_name) (optional)
    id2 := "id_example" // string | Filter results by Deployment id (optional)
    clusterName := "clusterName_example" // string | Filter results by Deployment cluster_name (optional)
    dataServiceId := "dataServiceId_example" // string | Filter results by Deployment data_service_id (optional)
    deploymentTargetId := "deploymentTargetId_example" // string | Filter results by Deployment deployment_target_id (optional)
    imageId := "imageId_example" // string | Filter results by Deployment image_id (optional)
    namespaceId := "namespaceId_example" // string | Filter results by Deployment namespace_id (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.ApiProjectsIdDeploymentsGet(context.Background(), id).SortBy(sortBy).Id2(id2).ClusterName(clusterName).DataServiceId(dataServiceId).DeploymentTargetId(deploymentTargetId).ImageId(imageId).NamespaceId(namespaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ApiProjectsIdDeploymentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectsIdDeploymentsGet`: ControllersPaginatedDeployments
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ApiProjectsIdDeploymentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsIdDeploymentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given Deployment attribute to sort results by (one of: id, cluster_name) | 
 **id2** | **string** | Filter results by Deployment id | 
 **clusterName** | **string** | Filter results by Deployment cluster_name | 
 **dataServiceId** | **string** | Filter results by Deployment data_service_id | 
 **deploymentTargetId** | **string** | Filter results by Deployment deployment_target_id | 
 **imageId** | **string** | Filter results by Deployment image_id | 
 **namespaceId** | **string** | Filter results by Deployment namespace_id | 

### Return type

[**ControllersPaginatedDeployments**](ControllersPaginatedDeployments.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

