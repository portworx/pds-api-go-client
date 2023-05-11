# \DeploymentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDeploymentsIdConnectionInfoGet**](DeploymentsApi.md#ApiDeploymentsIdConnectionInfoGet) | **Get** /api/deployments/{id}/connection-info | Deployment Connection Information
[**ApiDeploymentsIdCredentialsGet**](DeploymentsApi.md#ApiDeploymentsIdCredentialsGet) | **Get** /api/deployments/{id}/credentials | Deployment Credentials
[**ApiDeploymentsIdDelete**](DeploymentsApi.md#ApiDeploymentsIdDelete) | **Delete** /api/deployments/{id} | Delete Deployment
[**ApiDeploymentsIdEventsGet**](DeploymentsApi.md#ApiDeploymentsIdEventsGet) | **Get** /api/deployments/{id}/events | Get Deployment Events
[**ApiDeploymentsIdGet**](DeploymentsApi.md#ApiDeploymentsIdGet) | **Get** /api/deployments/{id} | Get Deployment
[**ApiDeploymentsIdPut**](DeploymentsApi.md#ApiDeploymentsIdPut) | **Put** /api/deployments/{id} | Update Deployment
[**ApiDeploymentsIdStatusGet**](DeploymentsApi.md#ApiDeploymentsIdStatusGet) | **Get** /api/deployments/{id}/status | Get Deployment Status
[**ApiProjectsIdDeploymentsGet**](DeploymentsApi.md#ApiProjectsIdDeploymentsGet) | **Get** /api/projects/{id}/deployments | List Project&#39;s Deployments
[**ApiProjectsIdDeploymentsPost**](DeploymentsApi.md#ApiProjectsIdDeploymentsPost) | **Post** /api/projects/{id}/deployments | Create Deployment



## ApiDeploymentsIdConnectionInfoGet

> DeploymentsConnectionInfo ApiDeploymentsIdConnectionInfoGet(ctx, id).Execute()

Deployment Connection Information



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentsApi.ApiDeploymentsIdConnectionInfoGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ApiDeploymentsIdConnectionInfoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentsIdConnectionInfoGet`: DeploymentsConnectionInfo
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ApiDeploymentsIdConnectionInfoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Deployment ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentsIdConnectionInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeploymentsConnectionInfo**](DeploymentsConnectionInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeploymentsIdCredentialsGet

> DeploymentsCredentials ApiDeploymentsIdCredentialsGet(ctx, id).Execute()

Deployment Credentials



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentsApi.ApiDeploymentsIdCredentialsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ApiDeploymentsIdCredentialsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentsIdCredentialsGet`: DeploymentsCredentials
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ApiDeploymentsIdCredentialsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Deployment ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentsIdCredentialsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeploymentsCredentials**](DeploymentsCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeploymentsIdDelete

> ApiDeploymentsIdDelete(ctx, id).Force(force).Execute()

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
    force := "force_example" // string | Delete deployment even if the deletion job fails on any deployment targets (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentsApi.ApiDeploymentsIdDelete(context.Background(), id).Force(force).Execute()
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

 **force** | **string** | Delete deployment even if the deletion job fails on any deployment targets | 

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


## ApiDeploymentsIdEventsGet

> ServiceDeploymentResourceEvents ApiDeploymentsIdEventsGet(ctx, id).Execute()

Get Deployment Events



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentsApi.ApiDeploymentsIdEventsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ApiDeploymentsIdEventsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentsIdEventsGet`: ServiceDeploymentResourceEvents
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ApiDeploymentsIdEventsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Deployment ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentsIdEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDeploymentResourceEvents**](ServiceDeploymentResourceEvents.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeploymentsIdGet

> ModelsDeployment ApiDeploymentsIdGet(ctx, id).Expand(expand).Execute()

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
    expand := "expand_example" // string | Expand the result with related entities (allowed values: deployment_target, namespace, deployment_manifest) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentsApi.ApiDeploymentsIdGet(context.Background(), id).Expand(expand).Execute()
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

 **expand** | **string** | Expand the result with related entities (allowed values: deployment_target, namespace, deployment_manifest) | 

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentsApi.ApiDeploymentsIdPut(context.Background(), id).Body(body).LocalOnly(localOnly).Execute()
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

> ServiceDeploymentStatus ApiDeploymentsIdStatusGet(ctx, id).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentsApi.ApiDeploymentsIdStatusGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ApiDeploymentsIdStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentsIdStatusGet`: ServiceDeploymentStatus
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

[**ServiceDeploymentStatus**](ServiceDeploymentStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectsIdDeploymentsGet

> ModelsPaginatedResultModelsDeployment ApiProjectsIdDeploymentsGet(ctx, id).Expand(expand).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).ClusterResourceName(clusterResourceName).DataServiceId(dataServiceId).DeploymentTargetId(deploymentTargetId).ImageId(imageId).Name(name).NamespaceId(namespaceId).State(state).Execute()

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
    expand := "expand_example" // string | Expand the result with related entities (allowed values: deployment_target, namespace) (optional)
    sortBy := "sortBy_example" // string | A given Deployment attribute to sort results by (one of: id, name, cluster_resource_name, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by Deployment id (optional)
    clusterResourceName := "clusterResourceName_example" // string | Filter results by Deployment cluster_resource_name (optional)
    dataServiceId := "dataServiceId_example" // string | Filter results by Deployment data_service_id (optional)
    deploymentTargetId := "deploymentTargetId_example" // string | Filter results by Deployment deployment_target_id (optional)
    imageId := "imageId_example" // string | Filter results by Deployment image_id (optional)
    name := "name_example" // string | Filter results by Deployment name (optional)
    namespaceId := "namespaceId_example" // string | Filter results by Deployment namespace_id (optional)
    state := "state_example" // string | Filter results by Deployment state (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentsApi.ApiProjectsIdDeploymentsGet(context.Background(), id).Expand(expand).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).ClusterResourceName(clusterResourceName).DataServiceId(dataServiceId).DeploymentTargetId(deploymentTargetId).ImageId(imageId).Name(name).NamespaceId(namespaceId).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ApiProjectsIdDeploymentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectsIdDeploymentsGet`: ModelsPaginatedResultModelsDeployment
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

 **expand** | **string** | Expand the result with related entities (allowed values: deployment_target, namespace) | 
 **sortBy** | **string** | A given Deployment attribute to sort results by (one of: id, name, cluster_resource_name, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by Deployment id | 
 **clusterResourceName** | **string** | Filter results by Deployment cluster_resource_name | 
 **dataServiceId** | **string** | Filter results by Deployment data_service_id | 
 **deploymentTargetId** | **string** | Filter results by Deployment deployment_target_id | 
 **imageId** | **string** | Filter results by Deployment image_id | 
 **name** | **string** | Filter results by Deployment name | 
 **namespaceId** | **string** | Filter results by Deployment namespace_id | 
 **state** | **string** | Filter results by Deployment state | 

### Return type

[**ModelsPaginatedResultModelsDeployment**](ModelsPaginatedResultModelsDeployment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectsIdDeploymentsPost

> ModelsDeployment ApiProjectsIdDeploymentsPost(ctx, id).Body(body).LocalOnly(localOnly).Execute()

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
    id := "id_example" // string | Project ID (must be valid UUID)
    body := *openapiclient.NewRequestsCreateProjectDeploymentRequest() // RequestsCreateProjectDeploymentRequest | Request body containing the deployment config
    localOnly := true // bool | Set to true to only create the Deployment object in the database (does not create any actual resources) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentsApi.ApiProjectsIdDeploymentsPost(context.Background(), id).Body(body).LocalOnly(localOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ApiProjectsIdDeploymentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectsIdDeploymentsPost`: ModelsDeployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ApiProjectsIdDeploymentsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsIdDeploymentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestsCreateProjectDeploymentRequest**](RequestsCreateProjectDeploymentRequest.md) | Request body containing the deployment config | 
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

