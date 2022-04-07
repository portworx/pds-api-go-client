# \BackupTargetsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBackupTargetsIdDelete**](BackupTargetsApi.md#ApiBackupTargetsIdDelete) | **Delete** /api/backup-targets/{id} | Delete BackupTargets
[**ApiBackupTargetsIdGet**](BackupTargetsApi.md#ApiBackupTargetsIdGet) | **Get** /api/backup-targets/{id} | Get BackupTarget
[**ApiBackupTargetsIdPut**](BackupTargetsApi.md#ApiBackupTargetsIdPut) | **Put** /api/backup-targets/{id} | Update BackupTarget
[**ApiBackupTargetsIdRetryPost**](BackupTargetsApi.md#ApiBackupTargetsIdRetryPost) | **Post** /api/backup-targets/{id}/retry | Retry sync of a BackupTarget
[**ApiBackupTargetsIdStatesGet**](BackupTargetsApi.md#ApiBackupTargetsIdStatesGet) | **Get** /api/backup-targets/{id}/states | List BackupTarget&#39;s BackupTargetStates
[**ApiTenantsIdBackupTargetsGet**](BackupTargetsApi.md#ApiTenantsIdBackupTargetsGet) | **Get** /api/tenants/{id}/backup-targets | List Tenant&#39;s BackupTargets
[**ApiTenantsIdBackupTargetsPost**](BackupTargetsApi.md#ApiTenantsIdBackupTargetsPost) | **Post** /api/tenants/{id}/backup-targets | Create BackupTarget



## ApiBackupTargetsIdDelete

> ApiBackupTargetsIdDelete(ctx, id).Force(force).Execute()

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
    force := "force_example" // string | Delete backup target even if the deletion job fails on any deployment targets (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupTargetsApi.ApiBackupTargetsIdDelete(context.Background(), id).Force(force).Execute()
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

 **force** | **string** | Delete backup target even if the deletion job fails on any deployment targets | 

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupTargetsApi.ApiBackupTargetsIdGet(context.Background(), id).Execute()
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


## ApiBackupTargetsIdPut

> ModelsBackupTarget ApiBackupTargetsIdPut(ctx, id).Body(body).Execute()

Update BackupTarget



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
    body := *openapiclient.NewControllersUpdateBackupTargetRequest() // ControllersUpdateBackupTargetRequest | Object with partial update fields

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupTargetsApi.ApiBackupTargetsIdPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupTargetsApi.ApiBackupTargetsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiBackupTargetsIdPut`: ModelsBackupTarget
    fmt.Fprintf(os.Stdout, "Response from `BackupTargetsApi.ApiBackupTargetsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | BackupTarget ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupTargetsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersUpdateBackupTargetRequest**](ControllersUpdateBackupTargetRequest.md) | Object with partial update fields | 

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


## ApiBackupTargetsIdRetryPost

> ApiBackupTargetsIdRetryPost(ctx, id).Execute()

Retry sync of a BackupTarget



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
    id := "id_example" // string | BackupTargetState ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupTargetsApi.ApiBackupTargetsIdRetryPost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupTargetsApi.ApiBackupTargetsIdRetryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | BackupTargetState ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupTargetsIdRetryPostRequest struct via the builder pattern


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


## ApiBackupTargetsIdStatesGet

> ControllersPaginatedBackupTargetStates ApiBackupTargetsIdStatesGet(ctx, id).Limit(limit).Continuation(continuation).SortBy(sortBy).BackupTargetId(backupTargetId).DeploymentTargetId(deploymentTargetId).State(state).Execute()

List BackupTarget's BackupTargetStates



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
    id := "id_example" // string | Backup Target ID (must be valid UUID)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    sortBy := "sortBy_example" // string | A given BackupTargetState attribute to sort results by (one of: state, deployment_target_id, backup_target_id) (optional)
    backupTargetId := "backupTargetId_example" // string | Filter results by BackupTarget ID (optional)
    deploymentTargetId := "deploymentTargetId_example" // string | Filter results by DeploymentTarget ID (optional)
    state := "state_example" // string | Filter results by state (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupTargetsApi.ApiBackupTargetsIdStatesGet(context.Background(), id).Limit(limit).Continuation(continuation).SortBy(sortBy).BackupTargetId(backupTargetId).DeploymentTargetId(deploymentTargetId).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupTargetsApi.ApiBackupTargetsIdStatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiBackupTargetsIdStatesGet`: ControllersPaginatedBackupTargetStates
    fmt.Fprintf(os.Stdout, "Response from `BackupTargetsApi.ApiBackupTargetsIdStatesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Backup Target ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupTargetsIdStatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **sortBy** | **string** | A given BackupTargetState attribute to sort results by (one of: state, deployment_target_id, backup_target_id) | 
 **backupTargetId** | **string** | Filter results by BackupTarget ID | 
 **deploymentTargetId** | **string** | Filter results by DeploymentTarget ID | 
 **state** | **string** | Filter results by state | 

### Return type

[**ControllersPaginatedBackupTargetStates**](ControllersPaginatedBackupTargetStates.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdBackupTargetsGet

> ControllersPaginatedBackupTargets ApiTenantsIdBackupTargetsGet(ctx, id).Limit(limit).Continuation(continuation).SortBy(sortBy).Id2(id2).Name(name).Type_(type_).Bucket(bucket).Region(region).BackupCredentialsId(backupCredentialsId).Execute()

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
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    sortBy := "sortBy_example" // string | A given BackupTarget attribute to sort results by (one of: id, name, type, created_at) (optional)
    id2 := "id_example" // string | Filter results by BackupTarget ID (optional)
    name := "name_example" // string | Filter results by BackupTarget name (optional)
    type_ := "type__example" // string | Filter results by BackupTarget type (optional)
    bucket := "bucket_example" // string | Filter results by bucket (optional)
    region := "region_example" // string | Filter results by region (optional)
    backupCredentialsId := "backupCredentialsId_example" // string | Filter results by BackupCredentials ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupTargetsApi.ApiTenantsIdBackupTargetsGet(context.Background(), id).Limit(limit).Continuation(continuation).SortBy(sortBy).Id2(id2).Name(name).Type_(type_).Bucket(bucket).Region(region).BackupCredentialsId(backupCredentialsId).Execute()
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

 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **sortBy** | **string** | A given BackupTarget attribute to sort results by (one of: id, name, type, created_at) | 
 **id2** | **string** | Filter results by BackupTarget ID | 
 **name** | **string** | Filter results by BackupTarget name | 
 **type_** | **string** | Filter results by BackupTarget type | 
 **bucket** | **string** | Filter results by bucket | 
 **region** | **string** | Filter results by region | 
 **backupCredentialsId** | **string** | Filter results by BackupCredentials ID | 

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


## ApiTenantsIdBackupTargetsPost

> ModelsBackupTarget ApiTenantsIdBackupTargetsPost(ctx, id).Body(body).Execute()

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
    id := "id_example" // string | Tenant ID (must be valid UUID)
    body := *openapiclient.NewControllersCreateTenantBackupTarget() // ControllersCreateTenantBackupTarget | Request body containing the backup target config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupTargetsApi.ApiTenantsIdBackupTargetsPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupTargetsApi.ApiTenantsIdBackupTargetsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdBackupTargetsPost`: ModelsBackupTarget
    fmt.Fprintf(os.Stdout, "Response from `BackupTargetsApi.ApiTenantsIdBackupTargetsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdBackupTargetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersCreateTenantBackupTarget**](ControllersCreateTenantBackupTarget.md) | Request body containing the backup target config | 

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

