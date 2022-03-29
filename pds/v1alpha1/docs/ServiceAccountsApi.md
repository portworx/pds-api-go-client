# \ServiceAccountsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServiceAccountsIdDelete**](ServiceAccountsApi.md#ApiServiceAccountsIdDelete) | **Delete** /api/service-accounts/{id} | Delete ServiceAccounts
[**ApiServiceAccountsIdGet**](ServiceAccountsApi.md#ApiServiceAccountsIdGet) | **Get** /api/service-accounts/{id} | Get ServiceAccounts
[**ApiServiceAccountsIdTokenGet**](ServiceAccountsApi.md#ApiServiceAccountsIdTokenGet) | **Get** /api/service-accounts/{id}/token | Get token of a ServiceAccount
[**ApiTenantsIdServiceAccountsGet**](ServiceAccountsApi.md#ApiTenantsIdServiceAccountsGet) | **Get** /api/tenants/{id}/service-accounts | List Tenant&#39;s ServiceAccounts
[**ApiTenantsIdServiceAccountsPost**](ServiceAccountsApi.md#ApiTenantsIdServiceAccountsPost) | **Post** /api/tenants/{id}/service-accounts | Create ServiceAccounts



## ApiServiceAccountsIdDelete

> ApiServiceAccountsIdDelete(ctx, id).Execute()

Delete ServiceAccounts



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
    id := "id_example" // string | ServiceAccount ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccountsApi.ApiServiceAccountsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.ApiServiceAccountsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ServiceAccount ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountsIdDeleteRequest struct via the builder pattern


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


## ApiServiceAccountsIdGet

> ControllersServiceAccountResponse ApiServiceAccountsIdGet(ctx, id).Execute()

Get ServiceAccounts



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
    id := "id_example" // string | ServiceAccount ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccountsApi.ApiServiceAccountsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.ApiServiceAccountsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceAccountsIdGet`: ControllersServiceAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.ApiServiceAccountsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ServiceAccount ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ControllersServiceAccountResponse**](ControllersServiceAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiServiceAccountsIdTokenGet

> ControllersServiceAccountTokenResponse ApiServiceAccountsIdTokenGet(ctx, id).Execute()

Get token of a ServiceAccount



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
    id := "id_example" // string | ServiceAccount ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccountsApi.ApiServiceAccountsIdTokenGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.ApiServiceAccountsIdTokenGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceAccountsIdTokenGet`: ControllersServiceAccountTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.ApiServiceAccountsIdTokenGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ServiceAccount ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountsIdTokenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ControllersServiceAccountTokenResponse**](ControllersServiceAccountTokenResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdServiceAccountsGet

> ControllersPaginatedServiceAccounts ApiTenantsIdServiceAccountsGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Name(name).Id2(id2).Token(token).Execute()

List Tenant's ServiceAccounts



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
    sortBy := "sortBy_example" // string | A given ServiceAccount attribute to sort results by (one of: name, id, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    name := "name_example" // string | Filter results by ServiceAccount name (optional)
    id2 := "id_example" // string | Filter results by ServiceAccount id (optional)
    token := "token_example" // string | Filter results by ServiceAccount token (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccountsApi.ApiTenantsIdServiceAccountsGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Name(name).Id2(id2).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.ApiTenantsIdServiceAccountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdServiceAccountsGet`: ControllersPaginatedServiceAccounts
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.ApiTenantsIdServiceAccountsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdServiceAccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given ServiceAccount attribute to sort results by (one of: name, id, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **name** | **string** | Filter results by ServiceAccount name | 
 **id2** | **string** | Filter results by ServiceAccount id | 
 **token** | **string** | Filter results by ServiceAccount token | 

### Return type

[**ControllersPaginatedServiceAccounts**](ControllersPaginatedServiceAccounts.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdServiceAccountsPost

> ModelsServiceAccount ApiTenantsIdServiceAccountsPost(ctx, id).Body(body).Execute()

Create ServiceAccounts



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
    body := *openapiclient.NewControllersCreateServiceAccountRequest() // ControllersCreateServiceAccountRequest | Request body containing name of the service account

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccountsApi.ApiTenantsIdServiceAccountsPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.ApiTenantsIdServiceAccountsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdServiceAccountsPost`: ModelsServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.ApiTenantsIdServiceAccountsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdServiceAccountsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersCreateServiceAccountRequest**](ControllersCreateServiceAccountRequest.md) | Request body containing name of the service account | 

### Return type

[**ModelsServiceAccount**](ModelsServiceAccount.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

