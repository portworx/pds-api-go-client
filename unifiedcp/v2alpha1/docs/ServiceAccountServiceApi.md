# \ServiceAccountServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceAccountServiceCreateServiceAccount**](ServiceAccountServiceApi.md#ServiceAccountServiceCreateServiceAccount) | **Post** /v1/tenants/{tenantId}/serviceAccounts | Create a requested service account.
[**ServiceAccountServiceDeleteServiceAccount**](ServiceAccountServiceApi.md#ServiceAccountServiceDeleteServiceAccount) | **Delete** /v1/serviceAccounts/{id} | Inititates deletion of a service account.
[**ServiceAccountServiceGetAccessToken**](ServiceAccountServiceApi.md#ServiceAccountServiceGetAccessToken) | **Post** /v1/tenants/{tenantId}:getToken | Get access token for a service account. (-- api-linter: core::0131::http-body&#x3D;disabled     api-linter: core::0131::http-method&#x3D;disabled     aip.dev/not-precedent: We need to do this because reasons API needs input from request body. --)
[**ServiceAccountServiceGetServiceAccount**](ServiceAccountServiceApi.md#ServiceAccountServiceGetServiceAccount) | **Get** /v1/serviceAccounts/{id} | Returns a requested service account.
[**ServiceAccountServiceListServiceAccount**](ServiceAccountServiceApi.md#ServiceAccountServiceListServiceAccount) | **Get** /v1/serviceAccounts | Returns a requested list of service accounts.
[**ServiceAccountServiceRegenerateServiceAccountSecret**](ServiceAccountServiceApi.md#ServiceAccountServiceRegenerateServiceAccountSecret) | **Get** /v1/serviceAccounts/{id}:regenerate | Regenerate access token for a service account.
[**ServiceAccountServiceUpdateServiceAccount**](ServiceAccountServiceApi.md#ServiceAccountServiceUpdateServiceAccount) | **Put** /v1/serviceAccounts/{id} | Updates a service account.



## ServiceAccountServiceCreateServiceAccount

> V1ServiceAccount ServiceAccountServiceCreateServiceAccount(ctx, tenantId).V1ServiceAccount(v1ServiceAccount).Execute()

Create a requested service account.

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
    tenantId := "tenantId_example" // string | The parent tenant under which the service account will be created (-- api-linter: core::0133::request-unknown-fields=disabled     aip.dev/not-precedent: We need this field for to support creation of     the resource in the tenant context. --)
    v1ServiceAccount := *openapiclient.NewV1ServiceAccount() // V1ServiceAccount | Service account details.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountServiceApi.ServiceAccountServiceCreateServiceAccount(context.Background(), tenantId).V1ServiceAccount(v1ServiceAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountServiceApi.ServiceAccountServiceCreateServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceAccountServiceCreateServiceAccount`: V1ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountServiceApi.ServiceAccountServiceCreateServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | The parent tenant under which the service account will be created (-- api-linter: core::0133::request-unknown-fields&#x3D;disabled     aip.dev/not-precedent: We need this field for to support creation of     the resource in the tenant context. --) | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceAccountServiceCreateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1ServiceAccount** | [**V1ServiceAccount**](V1ServiceAccount.md) | Service account details. | 

### Return type

[**V1ServiceAccount**](V1ServiceAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceAccountServiceDeleteServiceAccount

> map[string]interface{} ServiceAccountServiceDeleteServiceAccount(ctx, id).Execute()

Inititates deletion of a service account.

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
    id := "id_example" // string | Unique identifier for the service account to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountServiceApi.ServiceAccountServiceDeleteServiceAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountServiceApi.ServiceAccountServiceDeleteServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceAccountServiceDeleteServiceAccount`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountServiceApi.ServiceAccountServiceDeleteServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the service account to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceAccountServiceDeleteServiceAccountRequest struct via the builder pattern


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


## ServiceAccountServiceGetAccessToken

> V1AccessToken ServiceAccountServiceGetAccessToken(ctx, tenantId).ServiceAccountServiceGetAccessTokenBody(serviceAccountServiceGetAccessTokenBody).Execute()

Get access token for a service account. (-- api-linter: core::0131::http-body=disabled     api-linter: core::0131::http-method=disabled     aip.dev/not-precedent: We need to do this because reasons API needs input from request body. --)

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
    tenantId := "tenantId_example" // string | ID of the tenant under which the service account was created.
    serviceAccountServiceGetAccessTokenBody := *openapiclient.NewServiceAccountServiceGetAccessTokenBody() // ServiceAccountServiceGetAccessTokenBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountServiceApi.ServiceAccountServiceGetAccessToken(context.Background(), tenantId).ServiceAccountServiceGetAccessTokenBody(serviceAccountServiceGetAccessTokenBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountServiceApi.ServiceAccountServiceGetAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceAccountServiceGetAccessToken`: V1AccessToken
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountServiceApi.ServiceAccountServiceGetAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant under which the service account was created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceAccountServiceGetAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceAccountServiceGetAccessTokenBody** | [**ServiceAccountServiceGetAccessTokenBody**](ServiceAccountServiceGetAccessTokenBody.md) |  | 

### Return type

[**V1AccessToken**](V1AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceAccountServiceGetServiceAccount

> V1ServiceAccount ServiceAccountServiceGetServiceAccount(ctx, id).TenantId(tenantId).Execute()

Returns a requested service account.

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
    id := "id_example" // string | Unique identifier for the service account to be fetched.
    tenantId := "tenantId_example" // string | Tenant id to which a service account is associated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountServiceApi.ServiceAccountServiceGetServiceAccount(context.Background(), id).TenantId(tenantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountServiceApi.ServiceAccountServiceGetServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceAccountServiceGetServiceAccount`: V1ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountServiceApi.ServiceAccountServiceGetServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the service account to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceAccountServiceGetServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant id to which a service account is associated. | 

### Return type

[**V1ServiceAccount**](V1ServiceAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceAccountServiceListServiceAccount

> V1ListServiceAccountResponse ServiceAccountServiceListServiceAccount(ctx).TenantId(tenantId).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()

Returns a requested list of service accounts.

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
    tenantId := "tenantId_example" // string | id of tenant on which service account should be listed. If not provided, then list will filtered on account id present in the context. (optional)
    sortSortBy := "sortSortBy_example" // string | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. (optional) (default to "FIELD_UNSPECIFIED")
    sortSortOrder := "sortSortOrder_example" // string | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. (optional) (default to "VALUE_UNSPECIFIED")
    paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
    paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountServiceApi.ServiceAccountServiceListServiceAccount(context.Background()).TenantId(tenantId).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountServiceApi.ServiceAccountServiceListServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceAccountServiceListServiceAccount`: V1ListServiceAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountServiceApi.ServiceAccountServiceListServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceAccountServiceListServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | id of tenant on which service account should be listed. If not provided, then list will filtered on account id present in the context. | 
 **sortSortBy** | **string** | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. | [default to &quot;FIELD_UNSPECIFIED&quot;]
 **sortSortOrder** | **string** | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. | [default to &quot;VALUE_UNSPECIFIED&quot;]
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 

### Return type

[**V1ListServiceAccountResponse**](V1ListServiceAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceAccountServiceRegenerateServiceAccountSecret

> V1ServiceAccount ServiceAccountServiceRegenerateServiceAccountSecret(ctx, id).Execute()

Regenerate access token for a service account.

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
    id := "id_example" // string | Unique identifier for the service account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountServiceApi.ServiceAccountServiceRegenerateServiceAccountSecret(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountServiceApi.ServiceAccountServiceRegenerateServiceAccountSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceAccountServiceRegenerateServiceAccountSecret`: V1ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountServiceApi.ServiceAccountServiceRegenerateServiceAccountSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceAccountServiceRegenerateServiceAccountSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1ServiceAccount**](V1ServiceAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceAccountServiceUpdateServiceAccount

> V1ServiceAccount ServiceAccountServiceUpdateServiceAccount(ctx, id).V1ServiceAccount(v1ServiceAccount).Execute()

Updates a service account.

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
    id := "id_example" // string | Unique identifier for the service account.
    v1ServiceAccount := *openapiclient.NewV1ServiceAccount() // V1ServiceAccount | Service account to be updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountServiceApi.ServiceAccountServiceUpdateServiceAccount(context.Background(), id).V1ServiceAccount(v1ServiceAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountServiceApi.ServiceAccountServiceUpdateServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceAccountServiceUpdateServiceAccount`: V1ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountServiceApi.ServiceAccountServiceUpdateServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceAccountServiceUpdateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1ServiceAccount** | [**V1ServiceAccount**](V1ServiceAccount.md) | Service account to be updated. | 

### Return type

[**V1ServiceAccount**](V1ServiceAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

