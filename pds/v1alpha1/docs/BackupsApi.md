# \BackupsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBackupsIdDelete**](BackupsApi.md#ApiBackupsIdDelete) | **Delete** /api/backups/{id} | Delete Backup
[**ApiBackupsIdGet**](BackupsApi.md#ApiBackupsIdGet) | **Get** /api/backups/{id} | Get Backup
[**ApiBackupsIdPut**](BackupsApi.md#ApiBackupsIdPut) | **Put** /api/backups/{id} | Update Backup
[**ApiDeploymentsIdBackupsPost**](BackupsApi.md#ApiDeploymentsIdBackupsPost) | **Post** /api/deployments/{id}/backups | Create Backup



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

