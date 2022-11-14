# \AccountRoleBindingsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountsIdInvitationsPost**](AccountRoleBindingsApi.md#ApiAccountsIdInvitationsPost) | **Post** /api/accounts/{id}/invitations | Create Invitation
[**ApiAccountsIdRoleBindingsDelete**](AccountRoleBindingsApi.md#ApiAccountsIdRoleBindingsDelete) | **Delete** /api/accounts/{id}/role-bindings | Delete AccountRoleBinding
[**ApiAccountsIdRoleBindingsGet**](AccountRoleBindingsApi.md#ApiAccountsIdRoleBindingsGet) | **Get** /api/accounts/{id}/role-bindings | List AccountRoleBinding
[**ApiAccountsIdRoleBindingsPut**](AccountRoleBindingsApi.md#ApiAccountsIdRoleBindingsPut) | **Put** /api/accounts/{id}/role-bindings | Create/Update AccountRoleBinding
[**ApiUsersIdAccountRoleBindingsGet**](AccountRoleBindingsApi.md#ApiUsersIdAccountRoleBindingsGet) | **Get** /api/users/{id}/account-role-bindings | List User&#39;s AccountRoleBindings



## ApiAccountsIdInvitationsPost

> ApiAccountsIdInvitationsPost(ctx, id).Body(body).Execute()

Create Invitation



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
    id := "id_example" // string | Account ID (must be valid UUID)
    body := *openapiclient.NewControllersInvitationRequest() // ControllersInvitationRequest | Request body containing the invitation details.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountRoleBindingsApi.ApiAccountsIdInvitationsPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountRoleBindingsApi.ApiAccountsIdInvitationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsIdInvitationsPostRequest struct via the builder pattern


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


## ApiAccountsIdRoleBindingsDelete

> ApiAccountsIdRoleBindingsDelete(ctx, id).Body(body).Execute()

Delete AccountRoleBinding



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
    id := "id_example" // string | Account ID (must be valid UUID)
    body := *openapiclient.NewRequestsDeleteRoleBindingRequest() // RequestsDeleteRoleBindingRequest | Request body containing the account role binding

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountRoleBindingsApi.ApiAccountsIdRoleBindingsDelete(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountRoleBindingsApi.ApiAccountsIdRoleBindingsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsIdRoleBindingsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestsDeleteRoleBindingRequest**](RequestsDeleteRoleBindingRequest.md) | Request body containing the account role binding | 

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


## ApiAccountsIdRoleBindingsGet

> ControllersPaginatedAccountRoleBindings ApiAccountsIdRoleBindingsGet(ctx, id).SortBy(sortBy).RoleName(roleName).ActorId(actorId).ActorType(actorType).Execute()

List AccountRoleBinding



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
    id := "id_example" // string | Account ID (must be valid UUID)
    sortBy := "sortBy_example" // string | A given AccountRoleBinding attribute to sort results by (one of: role_name, actor_id) (optional)
    roleName := "roleName_example" // string | Filter results by AccountRoleBinding assigned role name (optional)
    actorId := "actorId_example" // string | Filter results by AccountRoleBinding actor id (optional)
    actorType := "actorType_example" // string | Filter results by AccountRoleBinding actor type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountRoleBindingsApi.ApiAccountsIdRoleBindingsGet(context.Background(), id).SortBy(sortBy).RoleName(roleName).ActorId(actorId).ActorType(actorType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountRoleBindingsApi.ApiAccountsIdRoleBindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsIdRoleBindingsGet`: ControllersPaginatedAccountRoleBindings
    fmt.Fprintf(os.Stdout, "Response from `AccountRoleBindingsApi.ApiAccountsIdRoleBindingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsIdRoleBindingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given AccountRoleBinding attribute to sort results by (one of: role_name, actor_id) | 
 **roleName** | **string** | Filter results by AccountRoleBinding assigned role name | 
 **actorId** | **string** | Filter results by AccountRoleBinding actor id | 
 **actorType** | **string** | Filter results by AccountRoleBinding actor type | 

### Return type

[**ControllersPaginatedAccountRoleBindings**](ControllersPaginatedAccountRoleBindings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsIdRoleBindingsPut

> ModelsAccountRoleBinding ApiAccountsIdRoleBindingsPut(ctx, id).Body(body).Execute()

Create/Update AccountRoleBinding



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
    id := "id_example" // string | Account ID (must be valid UUID)
    body := *openapiclient.NewControllersUpsertAccountRoleBindingRequest() // ControllersUpsertAccountRoleBindingRequest | Request body containing the account role binding

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountRoleBindingsApi.ApiAccountsIdRoleBindingsPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountRoleBindingsApi.ApiAccountsIdRoleBindingsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsIdRoleBindingsPut`: ModelsAccountRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `AccountRoleBindingsApi.ApiAccountsIdRoleBindingsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsIdRoleBindingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersUpsertAccountRoleBindingRequest**](ControllersUpsertAccountRoleBindingRequest.md) | Request body containing the account role binding | 

### Return type

[**ModelsAccountRoleBinding**](ModelsAccountRoleBinding.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersIdAccountRoleBindingsGet

> ControllersPaginatedAccountRoleBindings ApiUsersIdAccountRoleBindingsGet(ctx, id).SortBy(sortBy).RoleName(roleName).Execute()

List User's AccountRoleBindings



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
    sortBy := "sortBy_example" // string | A given AccountRoleBinding attribute to sort results by (allowed: role_name) (optional)
    roleName := "roleName_example" // string | Filter results by role_name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountRoleBindingsApi.ApiUsersIdAccountRoleBindingsGet(context.Background(), id).SortBy(sortBy).RoleName(roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountRoleBindingsApi.ApiUsersIdAccountRoleBindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsersIdAccountRoleBindingsGet`: ControllersPaginatedAccountRoleBindings
    fmt.Fprintf(os.Stdout, "Response from `AccountRoleBindingsApi.ApiUsersIdAccountRoleBindingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersIdAccountRoleBindingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given AccountRoleBinding attribute to sort results by (allowed: role_name) | 
 **roleName** | **string** | Filter results by role_name | 

### Return type

[**ControllersPaginatedAccountRoleBindings**](ControllersPaginatedAccountRoleBindings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

