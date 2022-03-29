# \BackupCredentialsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBackupCredentialsIdCredentialsGet**](BackupCredentialsApi.md#ApiBackupCredentialsIdCredentialsGet) | **Get** /api/backup-credentials/{id}/credentials | Get cloud configuration for BackupCredentials
[**ApiBackupCredentialsIdDelete**](BackupCredentialsApi.md#ApiBackupCredentialsIdDelete) | **Delete** /api/backup-credentials/{id} | Delete BackupCredentials
[**ApiBackupCredentialsIdGet**](BackupCredentialsApi.md#ApiBackupCredentialsIdGet) | **Get** /api/backup-credentials/{id} | Get BackupCredentials
[**ApiBackupCredentialsIdPut**](BackupCredentialsApi.md#ApiBackupCredentialsIdPut) | **Put** /api/backup-credentials/{id} | Update BackupCredentials
[**ApiTenantsIdBackupCredentialsGet**](BackupCredentialsApi.md#ApiTenantsIdBackupCredentialsGet) | **Get** /api/tenants/{id}/backup-credentials | List BackupCredentials
[**ApiTenantsIdBackupCredentialsPost**](BackupCredentialsApi.md#ApiTenantsIdBackupCredentialsPost) | **Post** /api/tenants/{id}/backup-credentials | Create BackupCredentials



## ApiBackupCredentialsIdCredentialsGet

> ControllersPartialCredentials ApiBackupCredentialsIdCredentialsGet(ctx, id).Execute()

Get cloud configuration for BackupCredentials



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
    id := "id_example" // string | BackupCredentials ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupCredentialsApi.ApiBackupCredentialsIdCredentialsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupCredentialsApi.ApiBackupCredentialsIdCredentialsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiBackupCredentialsIdCredentialsGet`: ControllersPartialCredentials
    fmt.Fprintf(os.Stdout, "Response from `BackupCredentialsApi.ApiBackupCredentialsIdCredentialsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | BackupCredentials ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupCredentialsIdCredentialsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ControllersPartialCredentials**](ControllersPartialCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiBackupCredentialsIdDelete

> ApiBackupCredentialsIdDelete(ctx, id).Execute()

Delete BackupCredentials



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
    id := "id_example" // string | BackupCredentials ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupCredentialsApi.ApiBackupCredentialsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupCredentialsApi.ApiBackupCredentialsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | BackupCredentials ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupCredentialsIdDeleteRequest struct via the builder pattern


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


## ApiBackupCredentialsIdGet

> ModelsBackupCredentials ApiBackupCredentialsIdGet(ctx, id).Execute()

Get BackupCredentials



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
    id := "id_example" // string | BackupCredentials ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupCredentialsApi.ApiBackupCredentialsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupCredentialsApi.ApiBackupCredentialsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiBackupCredentialsIdGet`: ModelsBackupCredentials
    fmt.Fprintf(os.Stdout, "Response from `BackupCredentialsApi.ApiBackupCredentialsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | BackupCredentials ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupCredentialsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsBackupCredentials**](ModelsBackupCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiBackupCredentialsIdPut

> ModelsBackupCredentials ApiBackupCredentialsIdPut(ctx, id).Body(body).Execute()

Update BackupCredentials



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
    id := "id_example" // string | BackupCredentials ID (must be UUID)
    body := *openapiclient.NewControllersUpdateBackupCredentialsRequest() // ControllersUpdateBackupCredentialsRequest | Request body containing the backup credentials config

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupCredentialsApi.ApiBackupCredentialsIdPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupCredentialsApi.ApiBackupCredentialsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiBackupCredentialsIdPut`: ModelsBackupCredentials
    fmt.Fprintf(os.Stdout, "Response from `BackupCredentialsApi.ApiBackupCredentialsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | BackupCredentials ID (must be UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupCredentialsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersUpdateBackupCredentialsRequest**](ControllersUpdateBackupCredentialsRequest.md) | Request body containing the backup credentials config | 

### Return type

[**ModelsBackupCredentials**](ModelsBackupCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdBackupCredentialsGet

> ControllersPaginatedBackupCredentials ApiTenantsIdBackupCredentialsGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Type_(type_).Execute()

List BackupCredentials



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
    sortBy := "sortBy_example" // string | A given BackupCredentials attribute to sort results by (one of: id, name, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by BackupCredentials id (optional)
    name := "name_example" // string | Filter results by BackupCredentials name (optional)
    type_ := "type__example" // string | Filter results by BackupCredentials type (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupCredentialsApi.ApiTenantsIdBackupCredentialsGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupCredentialsApi.ApiTenantsIdBackupCredentialsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdBackupCredentialsGet`: ControllersPaginatedBackupCredentials
    fmt.Fprintf(os.Stdout, "Response from `BackupCredentialsApi.ApiTenantsIdBackupCredentialsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdBackupCredentialsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given BackupCredentials attribute to sort results by (one of: id, name, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by BackupCredentials id | 
 **name** | **string** | Filter results by BackupCredentials name | 
 **type_** | **string** | Filter results by BackupCredentials type | 

### Return type

[**ControllersPaginatedBackupCredentials**](ControllersPaginatedBackupCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdBackupCredentialsPost

> ModelsBackupCredentials ApiTenantsIdBackupCredentialsPost(ctx, id).Body(body).Execute()

Create BackupCredentials



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
    body := *openapiclient.NewControllersCreateBackupCredentialsRequest() // ControllersCreateBackupCredentialsRequest | Request body containing the backup credentials config

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupCredentialsApi.ApiTenantsIdBackupCredentialsPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupCredentialsApi.ApiTenantsIdBackupCredentialsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdBackupCredentialsPost`: ModelsBackupCredentials
    fmt.Fprintf(os.Stdout, "Response from `BackupCredentialsApi.ApiTenantsIdBackupCredentialsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdBackupCredentialsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersCreateBackupCredentialsRequest**](ControllersCreateBackupCredentialsRequest.md) | Request body containing the backup credentials config | 

### Return type

[**ModelsBackupCredentials**](ModelsBackupCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

