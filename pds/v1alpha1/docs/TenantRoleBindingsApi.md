# \TenantRoleBindingsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTenantsIdRoleBindingsDelete**](TenantRoleBindingsApi.md#ApiTenantsIdRoleBindingsDelete) | **Delete** /api/tenants/{id}/role-bindings | Delete TenantRoleBinding
[**ApiTenantsIdRoleBindingsGet**](TenantRoleBindingsApi.md#ApiTenantsIdRoleBindingsGet) | **Get** /api/tenants/{id}/role-bindings | List TenantRoleBindings
[**ApiTenantsIdRoleBindingsPut**](TenantRoleBindingsApi.md#ApiTenantsIdRoleBindingsPut) | **Put** /api/tenants/{id}/role-bindings | Create TenantRoleBinding



## ApiTenantsIdRoleBindingsDelete

> ApiTenantsIdRoleBindingsDelete(ctx, id).Body(body).Execute()

Delete TenantRoleBinding



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Tenant ID (must be valid UUID)
    body := *openapiclient.NewRequestsDeleteRoleBindingRequest() // RequestsDeleteRoleBindingRequest | Request body containing the tenant role binding

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TenantRoleBindingsApi.ApiTenantsIdRoleBindingsDelete(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantRoleBindingsApi.ApiTenantsIdRoleBindingsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdRoleBindingsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestsDeleteRoleBindingRequest**](RequestsDeleteRoleBindingRequest.md) | Request body containing the tenant role binding | 

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


## ApiTenantsIdRoleBindingsGet

> ControllersPaginatedTenantRoleBindings ApiTenantsIdRoleBindingsGet(ctx, id).SortBy(sortBy).RoleName(roleName).ActorId(actorId).ActorType(actorType).Execute()

List TenantRoleBindings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Tenant ID (must be valid UUID)
    sortBy := "sortBy_example" // string | A given TenantRoleBinding attribute to sort results by (one of: role_name, actor_id) (optional)
    roleName := "roleName_example" // string | Filter results by TenantRoleBinding assigned role name (optional)
    actorId := "actorId_example" // string | Filter results by TenantRoleBinding actor id (optional)
    actorType := "actorType_example" // string | Filter results by TenantRoleBinding actor type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantRoleBindingsApi.ApiTenantsIdRoleBindingsGet(context.Background(), id).SortBy(sortBy).RoleName(roleName).ActorId(actorId).ActorType(actorType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantRoleBindingsApi.ApiTenantsIdRoleBindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdRoleBindingsGet`: ControllersPaginatedTenantRoleBindings
    fmt.Fprintf(os.Stdout, "Response from `TenantRoleBindingsApi.ApiTenantsIdRoleBindingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdRoleBindingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given TenantRoleBinding attribute to sort results by (one of: role_name, actor_id) | 
 **roleName** | **string** | Filter results by TenantRoleBinding assigned role name | 
 **actorId** | **string** | Filter results by TenantRoleBinding actor id | 
 **actorType** | **string** | Filter results by TenantRoleBinding actor type | 

### Return type

[**ControllersPaginatedTenantRoleBindings**](ControllersPaginatedTenantRoleBindings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdRoleBindingsPut

> ModelsTenantRoleBinding ApiTenantsIdRoleBindingsPut(ctx, id).Body(body).Execute()

Create TenantRoleBinding



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Tenant ID (must be valid UUID)
    body := *openapiclient.NewControllersUpsertTenantRoleBindingRequest() // ControllersUpsertTenantRoleBindingRequest | Request body containing the tenant role binding

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantRoleBindingsApi.ApiTenantsIdRoleBindingsPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantRoleBindingsApi.ApiTenantsIdRoleBindingsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdRoleBindingsPut`: ModelsTenantRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `TenantRoleBindingsApi.ApiTenantsIdRoleBindingsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdRoleBindingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersUpsertTenantRoleBindingRequest**](ControllersUpsertTenantRoleBindingRequest.md) | Request body containing the tenant role binding | 

### Return type

[**ModelsTenantRoleBinding**](ModelsTenantRoleBinding.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

