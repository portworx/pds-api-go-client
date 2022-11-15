# \ProjectsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProjectsIdAssignDeploymentTargetsPost**](ProjectsApi.md#ApiProjectsIdAssignDeploymentTargetsPost) | **Post** /api/projects/{id}/assign-deployment-targets | Assign Deployment Targets to Project
[**ApiProjectsIdGet**](ProjectsApi.md#ApiProjectsIdGet) | **Get** /api/projects/{id} | Get Project
[**ApiProjectsIdPatch**](ProjectsApi.md#ApiProjectsIdPatch) | **Patch** /api/projects/{id} | Patch Project
[**ApiProjectsIdUnassignDeploymentTargetsPost**](ProjectsApi.md#ApiProjectsIdUnassignDeploymentTargetsPost) | **Post** /api/projects/{id}/unassign-deployment-targets | Unassign Deployment Targets from Project
[**ApiProjectsIdUsersGet**](ProjectsApi.md#ApiProjectsIdUsersGet) | **Get** /api/projects/{id}/users | List Project Users
[**ApiTenantsIdProjectsGet**](ProjectsApi.md#ApiTenantsIdProjectsGet) | **Get** /api/tenants/{id}/projects | List Tenant&#39;s Projects
[**ApiTenantsIdProjectsPost**](ProjectsApi.md#ApiTenantsIdProjectsPost) | **Post** /api/tenants/{id}/projects | Create Project
[**ApiUsersIdTenantsTenantidProjectsGet**](ProjectsApi.md#ApiUsersIdTenantsTenantidProjectsGet) | **Get** /api/users/{id}/tenants/{tenantid}/projects | List User&#39;s Projects



## ApiProjectsIdAssignDeploymentTargetsPost

> ApiProjectsIdAssignDeploymentTargetsPost(ctx, id).Body(body).Execute()

Assign Deployment Targets to Project



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
    body := *openapiclient.NewRequestsDeploymentTargetIDArrayRequest() // RequestsDeploymentTargetIDArrayRequest | Request body containing the deployment target IDs to assign to the project.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiProjectsIdAssignDeploymentTargetsPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiProjectsIdAssignDeploymentTargetsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsIdAssignDeploymentTargetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestsDeploymentTargetIDArrayRequest**](RequestsDeploymentTargetIDArrayRequest.md) | Request body containing the deployment target IDs to assign to the project. | 

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


## ApiProjectsIdGet

> ModelsProject ApiProjectsIdGet(ctx, id).Execute()

Get Project



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiProjectsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiProjectsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectsIdGet`: ModelsProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiProjectsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsProject**](ModelsProject.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectsIdPatch

> ModelsProject ApiProjectsIdPatch(ctx, id).Body(body).Execute()

Patch Project



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
    body := *openapiclient.NewRequestsPatchProjectRequest() // RequestsPatchProjectRequest | Request body containing the new project values

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiProjectsIdPatch(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiProjectsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectsIdPatch`: ModelsProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiProjectsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestsPatchProjectRequest**](RequestsPatchProjectRequest.md) | Request body containing the new project values | 

### Return type

[**ModelsProject**](ModelsProject.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectsIdUnassignDeploymentTargetsPost

> ApiProjectsIdUnassignDeploymentTargetsPost(ctx, id).Body(body).Execute()

Unassign Deployment Targets from Project



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
    body := *openapiclient.NewRequestsDeploymentTargetIDArrayRequest() // RequestsDeploymentTargetIDArrayRequest | Request body containing the deployment target IDs to unassign from the project.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiProjectsIdUnassignDeploymentTargetsPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiProjectsIdUnassignDeploymentTargetsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsIdUnassignDeploymentTargetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestsDeploymentTargetIDArrayRequest**](RequestsDeploymentTargetIDArrayRequest.md) | Request body containing the deployment target IDs to unassign from the project. | 

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


## ApiProjectsIdUsersGet

> ModelsPaginatedResultModelsUser ApiProjectsIdUsersGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Email(email).Execute()

List Project Users



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
    sortBy := "sortBy_example" // string | A given User attribute to sort results by (one of: id, email, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by User id (optional)
    email := "email_example" // string | Filter results by User email (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiProjectsIdUsersGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Email(email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiProjectsIdUsersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectsIdUsersGet`: ModelsPaginatedResultModelsUser
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiProjectsIdUsersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsIdUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given User attribute to sort results by (one of: id, email, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by User id | 
 **email** | **string** | Filter results by User email | 

### Return type

[**ModelsPaginatedResultModelsUser**](ModelsPaginatedResultModelsUser.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdProjectsGet

> ModelsPaginatedResultModelsProject ApiTenantsIdProjectsGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Execute()

List Tenant's Projects



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
    sortBy := "sortBy_example" // string | A given Project attribute to sort results by (one of: id, name, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by Project id (optional)
    name := "name_example" // string | Filter results by Project name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiTenantsIdProjectsGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiTenantsIdProjectsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdProjectsGet`: ModelsPaginatedResultModelsProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiTenantsIdProjectsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdProjectsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given Project attribute to sort results by (one of: id, name, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by Project id | 
 **name** | **string** | Filter results by Project name | 

### Return type

[**ModelsPaginatedResultModelsProject**](ModelsPaginatedResultModelsProject.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdProjectsPost

> ModelsProject ApiTenantsIdProjectsPost(ctx, id).Body(body).Execute()

Create Project



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
    body := *openapiclient.NewRequestsCreateProjectRequest() // RequestsCreateProjectRequest | Request body containing the project values

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiTenantsIdProjectsPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiTenantsIdProjectsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdProjectsPost`: ModelsProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiTenantsIdProjectsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdProjectsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestsCreateProjectRequest**](RequestsCreateProjectRequest.md) | Request body containing the project values | 

### Return type

[**ModelsProject**](ModelsProject.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersIdTenantsTenantidProjectsGet

> ModelsPaginatedResultModelsProject ApiUsersIdTenantsTenantidProjectsGet(ctx, id, tenantid).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Execute()

List User's Projects



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
    id := "id_example" // string | User ID (must be valid UUID)
    tenantid := "tenantid_example" // string | Tenant ID (must be valid UUID)
    sortBy := "sortBy_example" // string | A given Project attribute to sort results by (one of: id, name, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by Project id (optional)
    name := "name_example" // string | Filter results by Project name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiUsersIdTenantsTenantidProjectsGet(context.Background(), id, tenantid).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiUsersIdTenantsTenantidProjectsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsersIdTenantsTenantidProjectsGet`: ModelsPaginatedResultModelsProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiUsersIdTenantsTenantidProjectsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User ID (must be valid UUID) | 
**tenantid** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersIdTenantsTenantidProjectsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortBy** | **string** | A given Project attribute to sort results by (one of: id, name, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by Project id | 
 **name** | **string** | Filter results by Project name | 

### Return type

[**ModelsPaginatedResultModelsProject**](ModelsPaginatedResultModelsProject.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

