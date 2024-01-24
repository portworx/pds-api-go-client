# \BackupServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupServiceDeleteBackup**](BackupServiceApi.md#BackupServiceDeleteBackup) | **Delete** /pds/v1/backups/{id} | DeleteBackup API deletes the Backup resource.
[**BackupServiceGetBackup**](BackupServiceApi.md#BackupServiceGetBackup) | **Get** /pds/v1/backups/{id} | GetBackup API returns the Backup resource.
[**BackupServiceListBackups**](BackupServiceApi.md#BackupServiceListBackups) | **Get** /pds/v1/backups | ListBackups API lists the Backup resources.



## BackupServiceDeleteBackup

> map[string]interface{} BackupServiceDeleteBackup(ctx, id).Execute()

DeleteBackup API deletes the Backup resource.

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
    id := "id_example" // string | UID of the Backup.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupServiceApi.BackupServiceDeleteBackup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupServiceApi.BackupServiceDeleteBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupServiceDeleteBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BackupServiceApi.BackupServiceDeleteBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UID of the Backup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupServiceDeleteBackupRequest struct via the builder pattern


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


## BackupServiceGetBackup

> V1Backup BackupServiceGetBackup(ctx, id).Execute()

GetBackup API returns the Backup resource.

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
    id := "id_example" // string | UID of the Backup.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupServiceApi.BackupServiceGetBackup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupServiceApi.BackupServiceGetBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupServiceGetBackup`: V1Backup
    fmt.Fprintf(os.Stdout, "Response from `BackupServiceApi.BackupServiceGetBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UID of the Backup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupServiceGetBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1Backup**](V1Backup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupServiceListBackups

> V1ListBackupsResponse BackupServiceListBackups(ctx).AccountId(accountId).TenantId(tenantId).ClusterId(clusterId).NamespaceId(namespaceId).ProjectId(projectId).BackupConfigId(backupConfigId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).Execute()

ListBackups API lists the Backup resources.

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
    accountId := "accountId_example" // string | UID of the account. (optional)
    tenantId := "tenantId_example" // string | UID of the tenant. (optional)
    clusterId := "clusterId_example" // string | UID of the target cluster. (optional)
    namespaceId := "namespaceId_example" // string | UID of the namespace. (optional)
    projectId := "projectId_example" // string | UID of the project. (optional)
    backupConfigId := "backupConfigId_example" // string | UID of the backupConfiguration (optional)
    paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
    paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)
    sortSortBy := "sortSortBy_example" // string | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. (optional) (default to "FIELD_UNSPECIFIED")
    sortSortOrder := "sortSortOrder_example" // string | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. (optional) (default to "VALUE_UNSPECIFIED")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupServiceApi.BackupServiceListBackups(context.Background()).AccountId(accountId).TenantId(tenantId).ClusterId(clusterId).NamespaceId(namespaceId).ProjectId(projectId).BackupConfigId(backupConfigId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupServiceApi.BackupServiceListBackups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupServiceListBackups`: V1ListBackupsResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupServiceApi.BackupServiceListBackups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupServiceListBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | UID of the account. | 
 **tenantId** | **string** | UID of the tenant. | 
 **clusterId** | **string** | UID of the target cluster. | 
 **namespaceId** | **string** | UID of the namespace. | 
 **projectId** | **string** | UID of the project. | 
 **backupConfigId** | **string** | UID of the backupConfiguration | 
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 
 **sortSortBy** | **string** | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. | [default to &quot;FIELD_UNSPECIFIED&quot;]
 **sortSortOrder** | **string** | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. | [default to &quot;VALUE_UNSPECIFIED&quot;]

### Return type

[**V1ListBackupsResponse**](V1ListBackupsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

