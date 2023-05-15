# \RestoresApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBackupJobsIdRestorePost**](RestoresApi.md#ApiBackupJobsIdRestorePost) | **Post** /api/backup-jobs/{id}/restore | Start Restore
[**ApiRestoresIdGet**](RestoresApi.md#ApiRestoresIdGet) | **Get** /api/restores/{id} | Get Restore
[**ApiRestoresIdUpdateManifestPost**](RestoresApi.md#ApiRestoresIdUpdateManifestPost) | **Post** /api/restores/{id}/update-manifest | Update Restore Manifest



## ApiBackupJobsIdRestorePost

> ModelsRestore ApiBackupJobsIdRestorePost(ctx, id).Body(body).Execute()

Start Restore



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
    body := *openapiclient.NewRequestsCreateRestoreRequest() // RequestsCreateRestoreRequest | Request body containing information about required restore

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoresApi.ApiBackupJobsIdRestorePost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoresApi.ApiBackupJobsIdRestorePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiBackupJobsIdRestorePost`: ModelsRestore
    fmt.Fprintf(os.Stdout, "Response from `RestoresApi.ApiBackupJobsIdRestorePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | BackupJob ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupJobsIdRestorePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestsCreateRestoreRequest**](RequestsCreateRestoreRequest.md) | Request body containing information about required restore | 

### Return type

[**ModelsRestore**](ModelsRestore.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRestoresIdGet

> ModelsRestore ApiRestoresIdGet(ctx, id).Execute()

Get Restore



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
    id := "id_example" // string | Restore ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoresApi.ApiRestoresIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoresApi.ApiRestoresIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRestoresIdGet`: ModelsRestore
    fmt.Fprintf(os.Stdout, "Response from `RestoresApi.ApiRestoresIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Restore ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRestoresIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsRestore**](ModelsRestore.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRestoresIdUpdateManifestPost

> ModelsRestore ApiRestoresIdUpdateManifestPost(ctx, id).Body(body).Execute()

Update Restore Manifest



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
    id := "id_example" // string | Restore ID (must be valid UUID)
    body := *openapiclient.NewRequestsUpdateRestoreManifestRequest() // RequestsUpdateRestoreManifestRequest | Request body containing the manifest update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestoresApi.ApiRestoresIdUpdateManifestPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoresApi.ApiRestoresIdUpdateManifestPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRestoresIdUpdateManifestPost`: ModelsRestore
    fmt.Fprintf(os.Stdout, "Response from `RestoresApi.ApiRestoresIdUpdateManifestPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Restore ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRestoresIdUpdateManifestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestsUpdateRestoreManifestRequest**](RequestsUpdateRestoreManifestRequest.md) | Request body containing the manifest update | 

### Return type

[**ModelsRestore**](ModelsRestore.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

