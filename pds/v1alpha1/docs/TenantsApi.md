# \TenantsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountsIdTenantsGet**](TenantsApi.md#ApiAccountsIdTenantsGet) | **Get** /api/accounts/{id}/tenants | List Account&#39;s Tenants
[**ApiTenantsIdDnsDetailsGet**](TenantsApi.md#ApiTenantsIdDnsDetailsGet) | **Get** /api/tenants/{id}/dns-details | Get DNS details for Tenant
[**ApiTenantsIdGet**](TenantsApi.md#ApiTenantsIdGet) | **Get** /api/tenants/{id} | Get Tenant



## ApiAccountsIdTenantsGet

> ControllersPaginatedAccountTenants ApiAccountsIdTenantsGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Execute()

List Account's Tenants



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
    sortBy := "sortBy_example" // string | A given Tenant attribute to sort results by (one of: id, name, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by Tenant id (optional)
    name := "name_example" // string | Filter results by Tenant name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.ApiAccountsIdTenantsGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.ApiAccountsIdTenantsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsIdTenantsGet`: ControllersPaginatedAccountTenants
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.ApiAccountsIdTenantsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsIdTenantsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given Tenant attribute to sort results by (one of: id, name, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by Tenant id | 
 **name** | **string** | Filter results by Tenant name | 

### Return type

[**ControllersPaginatedAccountTenants**](ControllersPaginatedAccountTenants.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdDnsDetailsGet

> ModelsDNSDetails ApiTenantsIdDnsDetailsGet(ctx, id).Execute()

Get DNS details for Tenant



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.ApiTenantsIdDnsDetailsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.ApiTenantsIdDnsDetailsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdDnsDetailsGet`: ModelsDNSDetails
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.ApiTenantsIdDnsDetailsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdDnsDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsDNSDetails**](ModelsDNSDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdGet

> ModelsTenant ApiTenantsIdGet(ctx, id).Execute()

Get Tenant



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.ApiTenantsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.ApiTenantsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdGet`: ModelsTenant
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.ApiTenantsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsTenant**](ModelsTenant.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

