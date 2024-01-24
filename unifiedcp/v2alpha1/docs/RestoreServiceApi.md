# \RestoreServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestoreServiceCreateRestore**](RestoreServiceApi.md#RestoreServiceCreateRestore) | **Post** /pds/v1/namespaces/{namespaceId}/restores | CreateRestore API creates the Restore resource.
[**RestoreServiceDeleteRestore**](RestoreServiceApi.md#RestoreServiceDeleteRestore) | **Delete** /pds/v1/restores/{id} | DeleteRestore API deletes the restore.
[**RestoreServiceGetRestore**](RestoreServiceApi.md#RestoreServiceGetRestore) | **Get** /pds/v1/restores/{id} | GetRestore API returns the Restore resource.
[**RestoreServiceListRestores**](RestoreServiceApi.md#RestoreServiceListRestores) | **Get** /pds/v1/restores | ListRestore API lists the Restore resources.
[**RestoreServiceRecreateRestore**](RestoreServiceApi.md#RestoreServiceRecreateRestore) | **Post** /pds/v1/restores/{id}:recreate | RecreateRestore API recreates a already failed restore.



## RestoreServiceCreateRestore

> V1Restore RestoreServiceCreateRestore(ctx, namespaceId).V1Restore(v1Restore).Execute()

CreateRestore API creates the Restore resource.

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
    namespaceId := "namespaceId_example" // string | UID of the namespace where restore will be created (-- api-linter: core::0133::request-unknown-fields=disabled     aip.dev/not-precedent: We really need this field because namespace is the parent of the restore. --)
    v1Restore := *openapiclient.NewV1Restore() // V1Restore | Restore resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoreServiceApi.RestoreServiceCreateRestore(context.Background(), namespaceId).V1Restore(v1Restore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreServiceApi.RestoreServiceCreateRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreServiceCreateRestore`: V1Restore
    fmt.Fprintf(os.Stdout, "Response from `RestoreServiceApi.RestoreServiceCreateRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceId** | **string** | UID of the namespace where restore will be created (-- api-linter: core::0133::request-unknown-fields&#x3D;disabled     aip.dev/not-precedent: We really need this field because namespace is the parent of the restore. --) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreServiceCreateRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1Restore** | [**V1Restore**](V1Restore.md) | Restore resource. | 

### Return type

[**V1Restore**](V1Restore.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreServiceDeleteRestore

> map[string]interface{} RestoreServiceDeleteRestore(ctx, id).Execute()

DeleteRestore API deletes the restore.

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
    id := "id_example" // string | UID of the Restore.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoreServiceApi.RestoreServiceDeleteRestore(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreServiceApi.RestoreServiceDeleteRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreServiceDeleteRestore`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RestoreServiceApi.RestoreServiceDeleteRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UID of the Restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreServiceDeleteRestoreRequest struct via the builder pattern


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


## RestoreServiceGetRestore

> V1Restore RestoreServiceGetRestore(ctx, id).Execute()

GetRestore API returns the Restore resource.

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
    id := "id_example" // string | UID of the restore.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoreServiceApi.RestoreServiceGetRestore(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreServiceApi.RestoreServiceGetRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreServiceGetRestore`: V1Restore
    fmt.Fprintf(os.Stdout, "Response from `RestoreServiceApi.RestoreServiceGetRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UID of the restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreServiceGetRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1Restore**](V1Restore.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreServiceListRestores

> V1ListRestoresResponse RestoreServiceListRestores(ctx).AccountId(accountId).TenantId(tenantId).ProjectId(projectId).DeploymentId(deploymentId).BackupId(backupId).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()

ListRestore API lists the Restore resources.

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
    accountId := "accountId_example" // string | Account ID for which the restore will be listed. (optional)
    tenantId := "tenantId_example" // string | Tenant ID for which the restore will be listed. (optional)
    projectId := "projectId_example" // string | Project ID for which the restore will be listed. (optional)
    deploymentId := "deploymentId_example" // string | Deployment ID for which the restore will be listed. (optional)
    backupId := "backupId_example" // string | Backup ID for which the restore will be listed. (optional)
    sortSortBy := "sortSortBy_example" // string | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. (optional) (default to "FIELD_UNSPECIFIED")
    sortSortOrder := "sortSortOrder_example" // string | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. (optional) (default to "VALUE_UNSPECIFIED")
    paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
    paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoreServiceApi.RestoreServiceListRestores(context.Background()).AccountId(accountId).TenantId(tenantId).ProjectId(projectId).DeploymentId(deploymentId).BackupId(backupId).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreServiceApi.RestoreServiceListRestores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreServiceListRestores`: V1ListRestoresResponse
    fmt.Fprintf(os.Stdout, "Response from `RestoreServiceApi.RestoreServiceListRestores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestoreServiceListRestoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Account ID for which the restore will be listed. | 
 **tenantId** | **string** | Tenant ID for which the restore will be listed. | 
 **projectId** | **string** | Project ID for which the restore will be listed. | 
 **deploymentId** | **string** | Deployment ID for which the restore will be listed. | 
 **backupId** | **string** | Backup ID for which the restore will be listed. | 
 **sortSortBy** | **string** | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. | [default to &quot;FIELD_UNSPECIFIED&quot;]
 **sortSortOrder** | **string** | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. | [default to &quot;VALUE_UNSPECIFIED&quot;]
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 

### Return type

[**V1ListRestoresResponse**](V1ListRestoresResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreServiceRecreateRestore

> V1Restore RestoreServiceRecreateRestore(ctx, id).RestoreServiceRecreateRestoreBody(restoreServiceRecreateRestoreBody).Execute()

RecreateRestore API recreates a already failed restore.

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
    id := "id_example" // string | UID of the existing restore
    restoreServiceRecreateRestoreBody := *openapiclient.NewRestoreServiceRecreateRestoreBody() // RestoreServiceRecreateRestoreBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoreServiceApi.RestoreServiceRecreateRestore(context.Background(), id).RestoreServiceRecreateRestoreBody(restoreServiceRecreateRestoreBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreServiceApi.RestoreServiceRecreateRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreServiceRecreateRestore`: V1Restore
    fmt.Fprintf(os.Stdout, "Response from `RestoreServiceApi.RestoreServiceRecreateRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UID of the existing restore | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreServiceRecreateRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restoreServiceRecreateRestoreBody** | [**RestoreServiceRecreateRestoreBody**](RestoreServiceRecreateRestoreBody.md) |  | 

### Return type

[**V1Restore**](V1Restore.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

