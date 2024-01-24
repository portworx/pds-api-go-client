# \BackupConfigServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupConfigServiceCreateBackupConfig**](BackupConfigServiceApi.md#BackupConfigServiceCreateBackupConfig) | **Post** /pds/v1/projects/{projectId}/backupConfigs | CreateBackupConfig API creates the backup configuration of the backup.
[**BackupConfigServiceDeleteBackupConfig**](BackupConfigServiceApi.md#BackupConfigServiceDeleteBackupConfig) | **Delete** /pds/v1/backupConfigs/{id} | DeleteBackupConfig API deletes the backup configuration.
[**BackupConfigServiceGetBackupConfig**](BackupConfigServiceApi.md#BackupConfigServiceGetBackupConfig) | **Get** /pds/v1/backupConfigs/{id} | GetBackupConfig API returns the the backup configuration resource.
[**BackupConfigServiceListBackupConfigs**](BackupConfigServiceApi.md#BackupConfigServiceListBackupConfigs) | **Get** /pds/v1/backupConfigs | ListBackupConfigs API lists all the backup configuration for a deployment.
[**BackupConfigServiceUpdateBackupConfig**](BackupConfigServiceApi.md#BackupConfigServiceUpdateBackupConfig) | **Put** /pds/v1/backupConfigs/{backupConfig.meta.uid} | UpdateBackupConfig API updates the backup configuration of the backup.



## BackupConfigServiceCreateBackupConfig

> V1BackupConfig BackupConfigServiceCreateBackupConfig(ctx, projectId).DeploymentId(deploymentId).V1BackupConfig(v1BackupConfig).Execute()

CreateBackupConfig API creates the backup configuration of the backup.

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
    projectId := "projectId_example" // string | (-- api-linter: core::0133::request-required-fields=disabled     aip.dev/not-precedent: We really need this field to be required to support creation of     the resource in the project context. --) (-- api-linter: core::0133::request-unknown-fields=disabled     aip.dev/not-precedent: We really need this field to be required to support creation of     the resource in the project context. --) The parent project id under which backup configuration will be created.
    deploymentId := "deploymentId_example" // string | (-- api-linter: core::0133::request-unknown-fields=disabled     aip.dev/not-precedent: We really need this field to be required     to support creation of the resource in the deployment context.. --) (-- api-linter: core::0133::request-required-fields=disabled     aip.dev/not-precedent: We really need this field to be required     to support creation of the resource in the deployment context.. --) Deployment id associated with the backup configuration.
    v1BackupConfig := *openapiclient.NewV1BackupConfig() // V1BackupConfig | Backup configuration for the backup.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupConfigServiceApi.BackupConfigServiceCreateBackupConfig(context.Background(), projectId).DeploymentId(deploymentId).V1BackupConfig(v1BackupConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupConfigServiceApi.BackupConfigServiceCreateBackupConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupConfigServiceCreateBackupConfig`: V1BackupConfig
    fmt.Fprintf(os.Stdout, "Response from `BackupConfigServiceApi.BackupConfigServiceCreateBackupConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | (-- api-linter: core::0133::request-required-fields&#x3D;disabled     aip.dev/not-precedent: We really need this field to be required to support creation of     the resource in the project context. --) (-- api-linter: core::0133::request-unknown-fields&#x3D;disabled     aip.dev/not-precedent: We really need this field to be required to support creation of     the resource in the project context. --) The parent project id under which backup configuration will be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupConfigServiceCreateBackupConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentId** | **string** | (-- api-linter: core::0133::request-unknown-fields&#x3D;disabled     aip.dev/not-precedent: We really need this field to be required     to support creation of the resource in the deployment context.. --) (-- api-linter: core::0133::request-required-fields&#x3D;disabled     aip.dev/not-precedent: We really need this field to be required     to support creation of the resource in the deployment context.. --) Deployment id associated with the backup configuration. | 
 **v1BackupConfig** | [**V1BackupConfig**](V1BackupConfig.md) | Backup configuration for the backup. | 

### Return type

[**V1BackupConfig**](V1BackupConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupConfigServiceDeleteBackupConfig

> map[string]interface{} BackupConfigServiceDeleteBackupConfig(ctx, id).Execute()

DeleteBackupConfig API deletes the backup configuration.

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
    id := "id_example" // string | ID of the backup configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupConfigServiceApi.BackupConfigServiceDeleteBackupConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupConfigServiceApi.BackupConfigServiceDeleteBackupConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupConfigServiceDeleteBackupConfig`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BackupConfigServiceApi.BackupConfigServiceDeleteBackupConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the backup configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupConfigServiceDeleteBackupConfigRequest struct via the builder pattern


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


## BackupConfigServiceGetBackupConfig

> V1BackupConfig BackupConfigServiceGetBackupConfig(ctx, id).Execute()

GetBackupConfig API returns the the backup configuration resource.

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
    id := "id_example" // string | ID of the backup configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupConfigServiceApi.BackupConfigServiceGetBackupConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupConfigServiceApi.BackupConfigServiceGetBackupConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupConfigServiceGetBackupConfig`: V1BackupConfig
    fmt.Fprintf(os.Stdout, "Response from `BackupConfigServiceApi.BackupConfigServiceGetBackupConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the backup configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupConfigServiceGetBackupConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1BackupConfig**](V1BackupConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupConfigServiceListBackupConfigs

> V1ListBackupConfigsResponse BackupConfigServiceListBackupConfigs(ctx).AccountId(accountId).TenantId(tenantId).ProjectId(projectId).TargetClusterId(targetClusterId).NamespaceId(namespaceId).DeploymentId(deploymentId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).Execute()

ListBackupConfigs API lists all the backup configuration for a deployment.

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
    accountId := "accountId_example" // string | Account ID for which the backup configurations will be listed. (optional)
    tenantId := "tenantId_example" // string | Tenant ID for which the backup configurations will be listed. (optional)
    projectId := "projectId_example" // string | Project ID for which the backup configurations will be listed. (optional)
    targetClusterId := "targetClusterId_example" // string | Cluster ID for which the backup configurations will be listed. (optional)
    namespaceId := "namespaceId_example" // string | Namespace ID for which the backup configurations will be listed. (optional)
    deploymentId := "deploymentId_example" // string | Deployment ID for which the backup configurations will be listed. (optional)
    paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
    paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)
    sortSortBy := "sortSortBy_example" // string | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. (optional) (default to "FIELD_UNSPECIFIED")
    sortSortOrder := "sortSortOrder_example" // string | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. (optional) (default to "VALUE_UNSPECIFIED")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupConfigServiceApi.BackupConfigServiceListBackupConfigs(context.Background()).AccountId(accountId).TenantId(tenantId).ProjectId(projectId).TargetClusterId(targetClusterId).NamespaceId(namespaceId).DeploymentId(deploymentId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupConfigServiceApi.BackupConfigServiceListBackupConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupConfigServiceListBackupConfigs`: V1ListBackupConfigsResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupConfigServiceApi.BackupConfigServiceListBackupConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupConfigServiceListBackupConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Account ID for which the backup configurations will be listed. | 
 **tenantId** | **string** | Tenant ID for which the backup configurations will be listed. | 
 **projectId** | **string** | Project ID for which the backup configurations will be listed. | 
 **targetClusterId** | **string** | Cluster ID for which the backup configurations will be listed. | 
 **namespaceId** | **string** | Namespace ID for which the backup configurations will be listed. | 
 **deploymentId** | **string** | Deployment ID for which the backup configurations will be listed. | 
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 
 **sortSortBy** | **string** | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. | [default to &quot;FIELD_UNSPECIFIED&quot;]
 **sortSortOrder** | **string** | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. | [default to &quot;VALUE_UNSPECIFIED&quot;]

### Return type

[**V1ListBackupConfigsResponse**](V1ListBackupConfigsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupConfigServiceUpdateBackupConfig

> V1BackupConfig BackupConfigServiceUpdateBackupConfig(ctx, backupConfigMetaUid).DesiredBackupConfiguration(desiredBackupConfiguration).Execute()

UpdateBackupConfig API updates the backup configuration of the backup.

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
    backupConfigMetaUid := "backupConfigMetaUid_example" // string | UID of the resource of the format <resource prefix>-<uuid>.
    desiredBackupConfiguration := *openapiclient.NewDesiredBackupConfiguration() // DesiredBackupConfiguration | Desired backup configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupConfigServiceApi.BackupConfigServiceUpdateBackupConfig(context.Background(), backupConfigMetaUid).DesiredBackupConfiguration(desiredBackupConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupConfigServiceApi.BackupConfigServiceUpdateBackupConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupConfigServiceUpdateBackupConfig`: V1BackupConfig
    fmt.Fprintf(os.Stdout, "Response from `BackupConfigServiceApi.BackupConfigServiceUpdateBackupConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupConfigMetaUid** | **string** | UID of the resource of the format &lt;resource prefix&gt;-&lt;uuid&gt;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupConfigServiceUpdateBackupConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **desiredBackupConfiguration** | [**DesiredBackupConfiguration**](DesiredBackupConfiguration.md) | Desired backup configuration. | 

### Return type

[**V1BackupConfig**](V1BackupConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

