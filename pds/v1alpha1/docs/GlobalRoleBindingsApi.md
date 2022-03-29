# \GlobalRoleBindingsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiGlobalRoleBindingsDelete**](GlobalRoleBindingsApi.md#ApiGlobalRoleBindingsDelete) | **Delete** /api/global-role-bindings | Delete GlobalRoleBinding
[**ApiGlobalRoleBindingsGet**](GlobalRoleBindingsApi.md#ApiGlobalRoleBindingsGet) | **Get** /api/global-role-bindings | List GlobalRoleBindings
[**ApiGlobalRoleBindingsPut**](GlobalRoleBindingsApi.md#ApiGlobalRoleBindingsPut) | **Put** /api/global-role-bindings | Set GlobalRoleBinding



## ApiGlobalRoleBindingsDelete

> ApiGlobalRoleBindingsDelete(ctx).ActorType(actorType).Execute()

Delete GlobalRoleBinding



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
    actorType := "actorType_example" // string | GlobalRoleBinding actor type (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GlobalRoleBindingsApi.ApiGlobalRoleBindingsDelete(context.Background()).ActorType(actorType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalRoleBindingsApi.ApiGlobalRoleBindingsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiGlobalRoleBindingsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actorType** | **string** | GlobalRoleBinding actor type | 

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


## ApiGlobalRoleBindingsGet

> ControllersPaginatedGlobalRoleBindings ApiGlobalRoleBindingsGet(ctx).SortBy(sortBy).RoleName(roleName).ActorId(actorId).ActorType(actorType).Execute()

List GlobalRoleBindings



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
    sortBy := "sortBy_example" // string | A given GlobalRoleBinding attribute to sort results by (one of: role_name, actor_id) (optional)
    roleName := "roleName_example" // string | Filter results by GlobalRoleBinding assigned role name (optional)
    actorId := "actorId_example" // string | Filter results by GlobalRoleBinding actor id (optional)
    actorType := "actorType_example" // string | Filter results by GlobalRoleBinding actor type (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GlobalRoleBindingsApi.ApiGlobalRoleBindingsGet(context.Background()).SortBy(sortBy).RoleName(roleName).ActorId(actorId).ActorType(actorType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalRoleBindingsApi.ApiGlobalRoleBindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiGlobalRoleBindingsGet`: ControllersPaginatedGlobalRoleBindings
    fmt.Fprintf(os.Stdout, "Response from `GlobalRoleBindingsApi.ApiGlobalRoleBindingsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiGlobalRoleBindingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **string** | A given GlobalRoleBinding attribute to sort results by (one of: role_name, actor_id) | 
 **roleName** | **string** | Filter results by GlobalRoleBinding assigned role name | 
 **actorId** | **string** | Filter results by GlobalRoleBinding actor id | 
 **actorType** | **string** | Filter results by GlobalRoleBinding actor type | 

### Return type

[**ControllersPaginatedGlobalRoleBindings**](ControllersPaginatedGlobalRoleBindings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiGlobalRoleBindingsPut

> ModelsGlobalRoleBinding ApiGlobalRoleBindingsPut(ctx).Body(body).Execute()

Set GlobalRoleBinding



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
    body := *openapiclient.NewModelsGlobalRoleBinding() // ModelsGlobalRoleBinding | Request body containing the global role binding

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GlobalRoleBindingsApi.ApiGlobalRoleBindingsPut(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalRoleBindingsApi.ApiGlobalRoleBindingsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiGlobalRoleBindingsPut`: ModelsGlobalRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `GlobalRoleBindingsApi.ApiGlobalRoleBindingsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiGlobalRoleBindingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ModelsGlobalRoleBinding**](ModelsGlobalRoleBinding.md) | Request body containing the global role binding | 

### Return type

[**ModelsGlobalRoleBinding**](ModelsGlobalRoleBinding.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

