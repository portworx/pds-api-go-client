# \BackupsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBackupTargetsIdBackupsGet**](BackupsApi.md#ApiBackupTargetsIdBackupsGet) | **Get** /api/backup-targets/{id}/backups | List BackupTarget&#39;s Backups
[**ApiBackupsIdDelete**](BackupsApi.md#ApiBackupsIdDelete) | **Delete** /api/backups/{id} | Delete Backup
[**ApiBackupsIdGet**](BackupsApi.md#ApiBackupsIdGet) | **Get** /api/backups/{id} | Get Backup
[**ApiBackupsIdJobsNameDelete**](BackupsApi.md#ApiBackupsIdJobsNameDelete) | **Delete** /api/backups/{id}/jobs/{name} | Delete Backup jobs
[**ApiBackupsIdPut**](BackupsApi.md#ApiBackupsIdPut) | **Put** /api/backups/{id} | Update Backup
[**ApiDeploymentsIdBackupsGet**](BackupsApi.md#ApiDeploymentsIdBackupsGet) | **Get** /api/deployments/{id}/backups | List Deployment&#39;s Backups
[**ApiDeploymentsIdBackupsPost**](BackupsApi.md#ApiDeploymentsIdBackupsPost) | **Post** /api/deployments/{id}/backups | Create Backup



## ApiBackupTargetsIdBackupsGet

> ModelsPaginatedResultModelsBackup ApiBackupTargetsIdBackupsGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).ClusterResourceName(clusterResourceName).State(state).Suspend(suspend).BackupType(backupType).BackupLevel(backupLevel).Execute()

List BackupTarget's Backups



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
    id := "id_example" // string | BackupTarget ID (must be valid UUID)
    sortBy := "sortBy_example" // string | A given Backup attribute to sort results by (one of: id, cluster_resource_name, created_at, backup_type, backup_level) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by Backup id (optional)
    clusterResourceName := "clusterResourceName_example" // string | Filter results by Backup cluster_resource_name (optional)
    state := "state_example" // string | Filter results by Backup state (optional)
    suspend := "suspend_example" // string | Filter results by Backup suspend flag (optional)
    backupType := "backupType_example" // string | Filter results by Backup type (one of: adhoc,scheduled) (optional)
    backupLevel := "backupLevel_example" // string | Filter results by Backup type (one of: snapshot,incremental) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupsApi.ApiBackupTargetsIdBackupsGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).ClusterResourceName(clusterResourceName).State(state).Suspend(suspend).BackupType(backupType).BackupLevel(backupLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.ApiBackupTargetsIdBackupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiBackupTargetsIdBackupsGet`: ModelsPaginatedResultModelsBackup
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.ApiBackupTargetsIdBackupsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | BackupTarget ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupTargetsIdBackupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given Backup attribute to sort results by (one of: id, cluster_resource_name, created_at, backup_type, backup_level) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by Backup id | 
 **clusterResourceName** | **string** | Filter results by Backup cluster_resource_name | 
 **state** | **string** | Filter results by Backup state | 
 **suspend** | **string** | Filter results by Backup suspend flag | 
 **backupType** | **string** | Filter results by Backup type (one of: adhoc,scheduled) | 
 **backupLevel** | **string** | Filter results by Backup type (one of: snapshot,incremental) | 

### Return type

[**ModelsPaginatedResultModelsBackup**](ModelsPaginatedResultModelsBackup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiBackupsIdDelete

> ApiBackupsIdDelete(ctx, id).LocalOnly(localOnly).Execute()

Delete Backup



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
    id := "id_example" // string | Backup ID (must be a valid UUID)
    localOnly := true // bool | Set to true to only delete the Backup object in the database (does not delete any actual resources) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupsApi.ApiBackupsIdDelete(context.Background(), id).LocalOnly(localOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.ApiBackupsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Backup ID (must be a valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **localOnly** | **bool** | Set to true to only delete the Backup object in the database (does not delete any actual resources) | 

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


## ApiBackupsIdGet

> ModelsBackup ApiBackupsIdGet(ctx, id).Execute()

Get Backup



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
    id := "id_example" // string | Backup ID (must be a valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupsApi.ApiBackupsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.ApiBackupsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiBackupsIdGet`: ModelsBackup
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.ApiBackupsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Backup ID (must be a valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsBackup**](ModelsBackup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiBackupsIdJobsNameDelete

> ApiBackupsIdJobsNameDelete(ctx, id, name).Execute()

Delete Backup jobs



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
    id := "id_example" // string | Backup ID (must be a valid UUID)
    name := "name_example" // string | Backup job name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupsApi.ApiBackupsIdJobsNameDelete(context.Background(), id, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.ApiBackupsIdJobsNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Backup ID (must be a valid UUID) | 
**name** | **string** | Backup job name | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupsIdJobsNameDeleteRequest struct via the builder pattern


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


## ApiBackupsIdPut

> ModelsBackup ApiBackupsIdPut(ctx, id).Body(body).LocalOnly(localOnly).Execute()

Update Backup



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
    id := "id_example" // string | Backup ID (must be UUID)
    body := *openapiclient.NewControllersUpdateBackupRequest() // ControllersUpdateBackupRequest | Request body containing updated backup
    localOnly := true // bool | Set to true to only update the Backup object in the database (does not create any actual resources) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupsApi.ApiBackupsIdPut(context.Background(), id).Body(body).LocalOnly(localOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.ApiBackupsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiBackupsIdPut`: ModelsBackup
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.ApiBackupsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Backup ID (must be UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersUpdateBackupRequest**](ControllersUpdateBackupRequest.md) | Request body containing updated backup | 
 **localOnly** | **bool** | Set to true to only update the Backup object in the database (does not create any actual resources) | 

### Return type

[**ModelsBackup**](ModelsBackup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeploymentsIdBackupsGet

> ModelsPaginatedResultModelsBackup ApiDeploymentsIdBackupsGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).ClusterResourceName(clusterResourceName).State(state).Suspend(suspend).BackupType(backupType).BackupLevel(backupLevel).Execute()

List Deployment's Backups



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
    id := "id_example" // string | Deployment ID (must be valid UUID)
    sortBy := "sortBy_example" // string | A given Backup attribute to sort results by (one of: id, cluster_resource_name, created_at, backup_type, backup_level) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by Backup id (optional)
    clusterResourceName := "clusterResourceName_example" // string | Filter results by Backup cluster_resource_name (optional)
    state := "state_example" // string | Filter results by Backup state (optional)
    suspend := "suspend_example" // string | Filter results by Backup suspend flag (optional)
    backupType := "backupType_example" // string | Filter results by Backup type (one of: adhoc,scheduled) (optional)
    backupLevel := "backupLevel_example" // string | Filter results by Backup type (one of: snapshot,incremental) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupsApi.ApiDeploymentsIdBackupsGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).ClusterResourceName(clusterResourceName).State(state).Suspend(suspend).BackupType(backupType).BackupLevel(backupLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.ApiDeploymentsIdBackupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentsIdBackupsGet`: ModelsPaginatedResultModelsBackup
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.ApiDeploymentsIdBackupsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Deployment ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentsIdBackupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given Backup attribute to sort results by (one of: id, cluster_resource_name, created_at, backup_type, backup_level) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by Backup id | 
 **clusterResourceName** | **string** | Filter results by Backup cluster_resource_name | 
 **state** | **string** | Filter results by Backup state | 
 **suspend** | **string** | Filter results by Backup suspend flag | 
 **backupType** | **string** | Filter results by Backup type (one of: adhoc,scheduled) | 
 **backupLevel** | **string** | Filter results by Backup type (one of: snapshot,incremental) | 

### Return type

[**ModelsPaginatedResultModelsBackup**](ModelsPaginatedResultModelsBackup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeploymentsIdBackupsPost

> ModelsBackup ApiDeploymentsIdBackupsPost(ctx, id).Body(body).LocalOnly(localOnly).Execute()

Create Backup



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
    id := "id_example" // string | Deployment ID (must be valid UUID)
    body := *openapiclient.NewControllersCreateDeploymentBackup() // ControllersCreateDeploymentBackup | Request body containing the backup config
    localOnly := true // bool | Set to true to only create the Backup object in the database (does not create any actual resources) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupsApi.ApiDeploymentsIdBackupsPost(context.Background(), id).Body(body).LocalOnly(localOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.ApiDeploymentsIdBackupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentsIdBackupsPost`: ModelsBackup
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.ApiDeploymentsIdBackupsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Deployment ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentsIdBackupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersCreateDeploymentBackup**](ControllersCreateDeploymentBackup.md) | Request body containing the backup config | 
 **localOnly** | **bool** | Set to true to only create the Backup object in the database (does not create any actual resources) | 

### Return type

[**ModelsBackup**](ModelsBackup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

