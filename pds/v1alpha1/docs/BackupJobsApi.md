# \BackupJobsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBackupJobsIdDelete**](BackupJobsApi.md#ApiBackupJobsIdDelete) | **Delete** /api/backup-jobs/{id} | Delete BackupJob
[**ApiBackupJobsIdGet**](BackupJobsApi.md#ApiBackupJobsIdGet) | **Get** /api/backup-jobs/{id} | Get BackupJob
[**ApiBackupJobsIdPut**](BackupJobsApi.md#ApiBackupJobsIdPut) | **Put** /api/backup-jobs/{id} | Upsert BackupJob
[**ApiBackupsIdJobsGet**](BackupJobsApi.md#ApiBackupsIdJobsGet) | **Get** /api/backups/{id}/jobs | List Backup&#39;s Jobs
[**ApiProjectsIdBackupJobsGet**](BackupJobsApi.md#ApiProjectsIdBackupJobsGet) | **Get** /api/projects/{id}/backup-jobs | List Project&#39;s BackupJobs



## ApiBackupJobsIdDelete

> ApiBackupJobsIdDelete(ctx, id).Execute()

Delete BackupJob



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
    id := "id_example" // string | BackupJob ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupJobsApi.ApiBackupJobsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupJobsApi.ApiBackupJobsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | BackupJob ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupJobsIdDeleteRequest struct via the builder pattern


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


## ApiBackupJobsIdGet

> ModelsBackupJob ApiBackupJobsIdGet(ctx, id).Execute()

Get BackupJob



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
    id := "id_example" // string | BackupJob ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupJobsApi.ApiBackupJobsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupJobsApi.ApiBackupJobsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiBackupJobsIdGet`: ModelsBackupJob
    fmt.Fprintf(os.Stdout, "Response from `BackupJobsApi.ApiBackupJobsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | BackupJob ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupJobsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsBackupJob**](ModelsBackupJob.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiBackupJobsIdPut

> ModelsBackupJob ApiBackupJobsIdPut(ctx, id).Body(body).Execute()

Upsert BackupJob



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
    id := "id_example" // string | BackupJob ID (must be UUID)
    body := *openapiclient.NewRequestsPutBackupJobRequest("ProjectId_example") // RequestsPutBackupJobRequest | Request body containing backup job details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupJobsApi.ApiBackupJobsIdPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupJobsApi.ApiBackupJobsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiBackupJobsIdPut`: ModelsBackupJob
    fmt.Fprintf(os.Stdout, "Response from `BackupJobsApi.ApiBackupJobsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | BackupJob ID (must be UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupJobsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestsPutBackupJobRequest**](RequestsPutBackupJobRequest.md) | Request body containing backup job details | 

### Return type

[**ModelsBackupJob**](ModelsBackupJob.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiBackupsIdJobsGet

> ControllersListBackupJobsStatusResponse ApiBackupsIdJobsGet(ctx, id).Execute()

List Backup's Jobs



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
    id := "id_example" // string | Backup ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupJobsApi.ApiBackupsIdJobsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupJobsApi.ApiBackupsIdJobsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiBackupsIdJobsGet`: ControllersListBackupJobsStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupJobsApi.ApiBackupsIdJobsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Backup ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupsIdJobsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ControllersListBackupJobsStatusResponse**](ControllersListBackupJobsStatusResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectsIdBackupJobsGet

> ModelsPaginatedResultModelsBackupJob ApiProjectsIdBackupJobsGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).DeploymentId(deploymentId).DeploymentTargetId(deploymentTargetId).BackupId(backupId).Name(name).NamespaceId(namespaceId).Execute()

List Project's BackupJobs



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
    id := "id_example" // string | Project ID (must be valid UUID)
    sortBy := "sortBy_example" // string | A given BackupJob attribute to sort results by (one of: id, name, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by BackupJob id (optional)
    deploymentId := "deploymentId_example" // string | Filter results by BackupJob deployment_id (optional)
    deploymentTargetId := "deploymentTargetId_example" // string | Filter results by BackupJob deployment_target_id (optional)
    backupId := "backupId_example" // string | Filter results by BackupJob backup_id (optional)
    name := "name_example" // string | Filter results by BackupJob name (optional)
    namespaceId := "namespaceId_example" // string | Filter results by BackupJob namespace_id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupJobsApi.ApiProjectsIdBackupJobsGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).DeploymentId(deploymentId).DeploymentTargetId(deploymentTargetId).BackupId(backupId).Name(name).NamespaceId(namespaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupJobsApi.ApiProjectsIdBackupJobsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProjectsIdBackupJobsGet`: ModelsPaginatedResultModelsBackupJob
    fmt.Fprintf(os.Stdout, "Response from `BackupJobsApi.ApiProjectsIdBackupJobsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsIdBackupJobsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given BackupJob attribute to sort results by (one of: id, name, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by BackupJob id | 
 **deploymentId** | **string** | Filter results by BackupJob deployment_id | 
 **deploymentTargetId** | **string** | Filter results by BackupJob deployment_target_id | 
 **backupId** | **string** | Filter results by BackupJob backup_id | 
 **name** | **string** | Filter results by BackupJob name | 
 **namespaceId** | **string** | Filter results by BackupJob namespace_id | 

### Return type

[**ModelsPaginatedResultModelsBackupJob**](ModelsPaginatedResultModelsBackupJob.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

