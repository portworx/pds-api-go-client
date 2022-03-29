# \ProjectRoleBindingsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProjectsIdRoleBindingsDelete**](ProjectRoleBindingsApi.md#ApiProjectsIdRoleBindingsDelete) | **Delete** /api/projects/{id}/role-bindings | Delete ProjectRoleBinding
[**ApiProjectsIdRoleBindingsGet**](ProjectRoleBindingsApi.md#ApiProjectsIdRoleBindingsGet) | **Get** /api/projects/{id}/role-bindings | List ProjectRoleBindings
[**ApiProjectsIdRoleBindingsPut**](ProjectRoleBindingsApi.md#ApiProjectsIdRoleBindingsPut) | **Put** /api/projects/{id}/role-bindings | Create ProjectRoleBinding



## ApiProjectsIdRoleBindingsDelete

> ApiProjectsIdRoleBindingsDelete(ctx, id).ActorType(actorType).Execute()

Delete ProjectRoleBinding



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
    actorType := "actorType_example" // string | ProjectRoleBinding actor type (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectRoleBindingsApi.ApiProjectsIdRoleBindingsDelete(context.Background(), id).ActorType(actorType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectRoleBindingsApi.ApiProjectsIdRoleBindingsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiProjectsIdRoleBindingsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actorType** | **string** | ProjectRoleBinding actor type | 

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


## ApiProjectsIdRoleBindingsGet

> ControllersPaginatedProjectRoleBindings ApiProjectsIdRoleBindingsGet(ctx, id).SortBy(sortBy).RoleName(roleName).ActorId(actorId).ActorType(actorType).Execute()

List ProjectRoleBindings



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
    sortBy := "sortBy_example" // string | A given ProjectRoleBinding attribute to sort results by (one of: role_name, actor_id) (optional)
    roleName := "roleName_example" // string | Filter results by ProjectRoleBinding assigned role name (optional)
    actorId := "actorId_example" // string | Filter results by ProjectRoleBinding actor id (optional)
    actorType := "actorType_example" // string | Filter results by ProjectRoleBinding actor type (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectRoleBindingsApi.ApiProjectsIdRoleBindingsGet(context.Background(), id).SortBy(sortBy).RoleName(roleName).ActorId(actorId).ActorType(actorType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectRoleBindingsApi.ApiProjectsIdRoleBindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectsIdRoleBindingsGet`: ControllersPaginatedProjectRoleBindings
    fmt.Fprintf(os.Stdout, "Response from `ProjectRoleBindingsApi.ApiProjectsIdRoleBindingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsIdRoleBindingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given ProjectRoleBinding attribute to sort results by (one of: role_name, actor_id) | 
 **roleName** | **string** | Filter results by ProjectRoleBinding assigned role name | 
 **actorId** | **string** | Filter results by ProjectRoleBinding actor id | 
 **actorType** | **string** | Filter results by ProjectRoleBinding actor type | 

### Return type

[**ControllersPaginatedProjectRoleBindings**](ControllersPaginatedProjectRoleBindings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectsIdRoleBindingsPut

> ModelsProjectRoleBinding ApiProjectsIdRoleBindingsPut(ctx, id).Body(body).Execute()

Create ProjectRoleBinding



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
    body := *openapiclient.NewControllersUpsertProjectRoleBindingRequest() // ControllersUpsertProjectRoleBindingRequest | Request body containing the project role binding

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectRoleBindingsApi.ApiProjectsIdRoleBindingsPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectRoleBindingsApi.ApiProjectsIdRoleBindingsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectsIdRoleBindingsPut`: ModelsProjectRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `ProjectRoleBindingsApi.ApiProjectsIdRoleBindingsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsIdRoleBindingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersUpsertProjectRoleBindingRequest**](ControllersUpsertProjectRoleBindingRequest.md) | Request body containing the project role binding | 

### Return type

[**ModelsProjectRoleBinding**](ModelsProjectRoleBinding.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

