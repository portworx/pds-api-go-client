# \TenantServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenantServiceCreateTenant**](TenantServiceApi.md#TenantServiceCreateTenant) | **Post** /v1/tenants | CreateTenant API creates a new Tenant
[**TenantServiceDeleteTenant**](TenantServiceApi.md#TenantServiceDeleteTenant) | **Delete** /v1/tenants/{tenantId} | Delete tenant removes a tenant record.
[**TenantServiceGetTenant**](TenantServiceApi.md#TenantServiceGetTenant) | **Get** /v1/tenants/{tenantId} | GetTenant API returns the info about  for given tenant id
[**TenantServiceListTenants**](TenantServiceApi.md#TenantServiceListTenants) | **Get** /v1/tenants | ListTenants API lists the tenants visible to the caller for the provided account.
[**TenantServiceListTenants2**](TenantServiceApi.md#TenantServiceListTenants2) | **Get** /v1/accounts/{accountId}/tenants | ListTenants API lists the tenants visible to the caller for the provided account.
[**TenantServiceUpdateTenant**](TenantServiceApi.md#TenantServiceUpdateTenant) | **Put** /v1/tenants/{tenant.meta.uid} | UpdateTenant API updates tenant.



## TenantServiceCreateTenant

> V1Tenant TenantServiceCreateTenant(ctx).V1Tenant(v1Tenant).Execute()

CreateTenant API creates a new Tenant

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
    v1Tenant := *openapiclient.NewV1Tenant() // V1Tenant | tenant to be created

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantServiceApi.TenantServiceCreateTenant(context.Background()).V1Tenant(v1Tenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantServiceApi.TenantServiceCreateTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantServiceCreateTenant`: V1Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantServiceApi.TenantServiceCreateTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantServiceCreateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1Tenant** | [**V1Tenant**](V1Tenant.md) | tenant to be created | 

### Return type

[**V1Tenant**](V1Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantServiceDeleteTenant

> map[string]interface{} TenantServiceDeleteTenant(ctx, tenantId).Execute()

Delete tenant removes a tenant record.

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
    tenantId := "tenantId_example" // string | ID of the tenant which needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantServiceApi.TenantServiceDeleteTenant(context.Background(), tenantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantServiceApi.TenantServiceDeleteTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantServiceDeleteTenant`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TenantServiceApi.TenantServiceDeleteTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant which needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantServiceDeleteTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantServiceGetTenant

> V1Tenant TenantServiceGetTenant(ctx, tenantId).Execute()

GetTenant API returns the info about  for given tenant id

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
    tenantId := "tenantId_example" // string | ID of the tenant which needs to get info.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantServiceApi.TenantServiceGetTenant(context.Background(), tenantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantServiceApi.TenantServiceGetTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantServiceGetTenant`: V1Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantServiceApi.TenantServiceGetTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant which needs to get info. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantServiceGetTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1Tenant**](V1Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantServiceListTenants

> V1ListTenantsResponse TenantServiceListTenants(ctx).AccountId(accountId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()

ListTenants API lists the tenants visible to the caller for the provided account.

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
    accountId := "accountId_example" // string | ID of the account which needs to get list tenant. (optional)
    paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
    paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantServiceApi.TenantServiceListTenants(context.Background()).AccountId(accountId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantServiceApi.TenantServiceListTenants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantServiceListTenants`: V1ListTenantsResponse
    fmt.Fprintf(os.Stdout, "Response from `TenantServiceApi.TenantServiceListTenants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantServiceListTenantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | ID of the account which needs to get list tenant. | 
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 

### Return type

[**V1ListTenantsResponse**](V1ListTenantsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantServiceListTenants2

> V1ListTenantsResponse TenantServiceListTenants2(ctx, accountId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()

ListTenants API lists the tenants visible to the caller for the provided account.

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
    accountId := "accountId_example" // string | ID of the account which needs to get list tenant.
    paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
    paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantServiceApi.TenantServiceListTenants2(context.Background(), accountId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantServiceApi.TenantServiceListTenants2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantServiceListTenants2`: V1ListTenantsResponse
    fmt.Fprintf(os.Stdout, "Response from `TenantServiceApi.TenantServiceListTenants2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | ID of the account which needs to get list tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantServiceListTenants2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 

### Return type

[**V1ListTenantsResponse**](V1ListTenantsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantServiceUpdateTenant

> V1Tenant TenantServiceUpdateTenant(ctx, tenantMetaUid).TenantWhichNeedsToBeUpdated(tenantWhichNeedsToBeUpdated).Execute()

UpdateTenant API updates tenant.

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
    tenantMetaUid := "tenantMetaUid_example" // string | UID of the resource of the format <resource prefix>-<uuid>.
    tenantWhichNeedsToBeUpdated := *openapiclient.NewTenantWhichNeedsToBeUpdated() // TenantWhichNeedsToBeUpdated | tenant which  needs to be updated

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantServiceApi.TenantServiceUpdateTenant(context.Background(), tenantMetaUid).TenantWhichNeedsToBeUpdated(tenantWhichNeedsToBeUpdated).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantServiceApi.TenantServiceUpdateTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantServiceUpdateTenant`: V1Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantServiceApi.TenantServiceUpdateTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantMetaUid** | **string** | UID of the resource of the format &lt;resource prefix&gt;-&lt;uuid&gt;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantServiceUpdateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantWhichNeedsToBeUpdated** | [**TenantWhichNeedsToBeUpdated**](TenantWhichNeedsToBeUpdated.md) | tenant which  needs to be updated | 

### Return type

[**V1Tenant**](V1Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

