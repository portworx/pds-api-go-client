# \TaskServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TaskServiceGetTask**](TaskServiceApi.md#TaskServiceGetTask) | **Get** /pds/v1/tasks/{id} | GetTask API returns the information about the specified task.
[**TaskServiceListTasks**](TaskServiceApi.md#TaskServiceListTasks) | **Get** /pds/v1/tasks | ListTasks returns the list of tasks.



## TaskServiceGetTask

> V1Task TaskServiceGetTask(ctx, id).Execute()

GetTask API returns the information about the specified task.

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
    id := "id_example" // string | UID of the task.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskServiceApi.TaskServiceGetTask(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskServiceApi.TaskServiceGetTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskServiceGetTask`: V1Task
    fmt.Fprintf(os.Stdout, "Response from `TaskServiceApi.TaskServiceGetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UID of the task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskServiceGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1Task**](V1Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskServiceListTasks

> V1ListTasksResponse TaskServiceListTasks(ctx).ProjectId(projectId).DeploymentId(deploymentId).BackupId(backupId).RestoreId(restoreId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).Execute()

ListTasks returns the list of tasks.

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
    projectId := "projectId_example" // string | UID of the project. (optional)
    deploymentId := "deploymentId_example" // string | UID of the deployment. (optional)
    backupId := "backupId_example" // string | UID of the backup. (optional)
    restoreId := "restoreId_example" // string | UID of the restore. (optional)
    paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
    paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)
    sortSortBy := "sortSortBy_example" // string | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. (optional) (default to "FIELD_UNSPECIFIED")
    sortSortOrder := "sortSortOrder_example" // string | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. (optional) (default to "VALUE_UNSPECIFIED")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskServiceApi.TaskServiceListTasks(context.Background()).ProjectId(projectId).DeploymentId(deploymentId).BackupId(backupId).RestoreId(restoreId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskServiceApi.TaskServiceListTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskServiceListTasks`: V1ListTasksResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskServiceApi.TaskServiceListTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskServiceListTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | UID of the project. | 
 **deploymentId** | **string** | UID of the deployment. | 
 **backupId** | **string** | UID of the backup. | 
 **restoreId** | **string** | UID of the restore. | 
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 
 **sortSortBy** | **string** | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. | [default to &quot;FIELD_UNSPECIFIED&quot;]
 **sortSortOrder** | **string** | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. | [default to &quot;VALUE_UNSPECIFIED&quot;]

### Return type

[**V1ListTasksResponse**](V1ListTasksResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

