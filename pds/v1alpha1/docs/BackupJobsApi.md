# \BackupJobsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBackupJobsIdGet**](BackupJobsApi.md#ApiBackupJobsIdGet) | **Get** /api/backup-jobs/{id} | Get BackupJob
[**ApiBackupsIdJobsGet**](BackupJobsApi.md#ApiBackupsIdJobsGet) | **Get** /api/backups/{id}/jobs | List Backup&#39;s Jobs



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


## ApiBackupsIdJobsGet

> ControllersBackupJobsResponse ApiBackupsIdJobsGet(ctx, id).Execute()

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
    // response from `ApiBackupsIdJobsGet`: ControllersBackupJobsResponse
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

[**ControllersBackupJobsResponse**](ControllersBackupJobsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

