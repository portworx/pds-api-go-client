# \BackupTargetsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBackupTargetsIdDelete**](BackupTargetsApi.md#ApiBackupTargetsIdDelete) | **Delete** /api/backup-targets/{id} | Delete BackupTargets
[**ApiBackupTargetsIdGet**](BackupTargetsApi.md#ApiBackupTargetsIdGet) | **Get** /api/backup-targets/{id} | Get BackupTarget
[**ApiBackupTargetsPost**](BackupTargetsApi.md#ApiBackupTargetsPost) | **Post** /api/backup-targets | Create BackupTarget
[**ApiTenantsIdBackupTargetsGet**](BackupTargetsApi.md#ApiTenantsIdBackupTargetsGet) | **Get** /api/tenants/{id}/backup-targets | List Tenant&#39;s BackupTargets



## ApiBackupTargetsIdDelete

> ApiBackupTargetsIdDelete(ctx, id).Execute()

Delete BackupTargets



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupTargetsApi.ApiBackupTargetsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupTargetsApi.ApiBackupTargetsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | BackupTarget ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupTargetsIdDeleteRequest struct via the builder pattern


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


## ApiBackupTargetsIdGet

> ModelsBackupTarget ApiBackupTargetsIdGet(ctx, id).Execute()

Get BackupTarget



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupTargetsApi.ApiBackupTargetsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupTargetsApi.ApiBackupTargetsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiBackupTargetsIdGet`: ModelsBackupTarget
    fmt.Fprintf(os.Stdout, "Response from `BackupTargetsApi.ApiBackupTargetsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | BackupTarget ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupTargetsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsBackupTarget**](ModelsBackupTarget.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiBackupTargetsPost

> ModelsBackupTarget ApiBackupTargetsPost(ctx).Body(body).Execute()

Create BackupTarget



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
    body := *openapiclient.NewControllersCreateBackupTargetRequest() // ControllersCreateBackupTargetRequest | Request body containing the backup target config

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupTargetsApi.ApiBackupTargetsPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupTargetsApi.ApiBackupTargetsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiBackupTargetsPost`: ModelsBackupTarget
    fmt.Fprintf(os.Stdout, "Response from `BackupTargetsApi.ApiBackupTargetsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupTargetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ControllersCreateBackupTargetRequest**](ControllersCreateBackupTargetRequest.md) | Request body containing the backup target config | 

### Return type

[**ModelsBackupTarget**](ModelsBackupTarget.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdBackupTargetsGet

> ControllersPaginatedBackupTargets ApiTenantsIdBackupTargetsGet(ctx, id).SortBy(sortBy).Id2(id2).Name(name).Execute()

List Tenant's BackupTargets



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
    id := "id_example" // string | Tenant ID (must be valid UUID)
    sortBy := "sortBy_example" // string | A given BackupTarget attribute to sort results by (one of: id, name) (optional)
    id2 := "id_example" // string | Filter results by BackupTarget ID (optional)
    name := "name_example" // string | Filter results by BackupTarget name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupTargetsApi.ApiTenantsIdBackupTargetsGet(context.Background(), id).SortBy(sortBy).Id2(id2).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupTargetsApi.ApiTenantsIdBackupTargetsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdBackupTargetsGet`: ControllersPaginatedBackupTargets
    fmt.Fprintf(os.Stdout, "Response from `BackupTargetsApi.ApiTenantsIdBackupTargetsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdBackupTargetsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given BackupTarget attribute to sort results by (one of: id, name) | 
 **id2** | **string** | Filter results by BackupTarget ID | 
 **name** | **string** | Filter results by BackupTarget name | 

### Return type

[**ControllersPaginatedBackupTargets**](ControllersPaginatedBackupTargets.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

