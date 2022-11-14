# \ProjectRoleBindingsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProjectsIdInvitationsPost**](ProjectRoleBindingsApi.md#ApiProjectsIdInvitationsPost) | **Post** /api/projects/{id}/invitations | Create Project Invitation
[**ApiProjectsIdRoleBindingsDelete**](ProjectRoleBindingsApi.md#ApiProjectsIdRoleBindingsDelete) | **Delete** /api/projects/{id}/role-bindings | Delete ProjectRoleBinding
[**ApiProjectsIdRoleBindingsGet**](ProjectRoleBindingsApi.md#ApiProjectsIdRoleBindingsGet) | **Get** /api/projects/{id}/role-bindings | List ProjectRoleBindings
[**ApiProjectsIdRoleBindingsPut**](ProjectRoleBindingsApi.md#ApiProjectsIdRoleBindingsPut) | **Put** /api/projects/{id}/role-bindings | Create ProjectRoleBinding
[**ApiUsersIdTenantsTenantidProjectRoleBindingsGet**](ProjectRoleBindingsApi.md#ApiUsersIdTenantsTenantidProjectRoleBindingsGet) | **Get** /api/users/{id}/tenants/{tenantid}/project-role-bindings | List User&#39;s ProjectRoleBindings



## ApiProjectsIdInvitationsPost

> ApiProjectsIdInvitationsPost(ctx, id).Body(body).Execute()

Create Project Invitation



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
    body := *openapiclient.NewControllersInvitationRequest() // ControllersInvitationRequest | Request body containing the invitation details.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectRoleBindingsApi.ApiProjectsIdInvitationsPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectRoleBindingsApi.ApiProjectsIdInvitationsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiProjectsIdInvitationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersInvitationRequest**](ControllersInvitationRequest.md) | Request body containing the invitation details. | 

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


## ApiProjectsIdRoleBindingsDelete

> ApiProjectsIdRoleBindingsDelete(ctx, id).Body(body).Execute()

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
    body := *openapiclient.NewRequestsDeleteRoleBindingRequest() // RequestsDeleteRoleBindingRequest | Request body containing the project role binding

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectRoleBindingsApi.ApiProjectsIdRoleBindingsDelete(context.Background(), id).Body(body).Execute()
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

 **body** | [**RequestsDeleteRoleBindingRequest**](RequestsDeleteRoleBindingRequest.md) | Request body containing the project role binding | 

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectRoleBindingsApi.ApiProjectsIdRoleBindingsGet(context.Background(), id).SortBy(sortBy).RoleName(roleName).ActorId(actorId).ActorType(actorType).Execute()
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectRoleBindingsApi.ApiProjectsIdRoleBindingsPut(context.Background(), id).Body(body).Execute()
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


## ApiUsersIdTenantsTenantidProjectRoleBindingsGet

> ControllersPaginatedProjectRoleBindings ApiUsersIdTenantsTenantidProjectRoleBindingsGet(ctx, id, tenantid).SortBy(sortBy).RoleName(roleName).Execute()

List User's ProjectRoleBindings



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
    sortBy := "sortBy_example" // string | A given ProjectRoleBinding attribute to sort results by (allowed: role_name) (optional)
    roleName := "roleName_example" // string | Filter results by role_name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectRoleBindingsApi.ApiUsersIdTenantsTenantidProjectRoleBindingsGet(context.Background(), id, tenantid).SortBy(sortBy).RoleName(roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectRoleBindingsApi.ApiUsersIdTenantsTenantidProjectRoleBindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsersIdTenantsTenantidProjectRoleBindingsGet`: ControllersPaginatedProjectRoleBindings
    fmt.Fprintf(os.Stdout, "Response from `ProjectRoleBindingsApi.ApiUsersIdTenantsTenantidProjectRoleBindingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User ID (must be valid UUID) | 
**tenantid** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersIdTenantsTenantidProjectRoleBindingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortBy** | **string** | A given ProjectRoleBinding attribute to sort results by (allowed: role_name) | 
 **roleName** | **string** | Filter results by role_name | 

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

