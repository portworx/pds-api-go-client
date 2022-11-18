# \BackupPoliciesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBackupPoliciesIdDelete**](BackupPoliciesApi.md#ApiBackupPoliciesIdDelete) | **Delete** /api/backup-policies/{id} | Delete BackupPolicies
[**ApiBackupPoliciesIdGet**](BackupPoliciesApi.md#ApiBackupPoliciesIdGet) | **Get** /api/backup-policies/{id} | Get BackupPolicy
[**ApiBackupPoliciesIdPut**](BackupPoliciesApi.md#ApiBackupPoliciesIdPut) | **Put** /api/backup-policies/{id} | Update BackupPolicy
[**ApiTenantsIdBackupPoliciesGet**](BackupPoliciesApi.md#ApiTenantsIdBackupPoliciesGet) | **Get** /api/tenants/{id}/backup-policies | List BackupPolicies
[**ApiTenantsIdBackupPoliciesPost**](BackupPoliciesApi.md#ApiTenantsIdBackupPoliciesPost) | **Post** /api/tenants/{id}/backup-policies | Create BackupPolicy



## ApiBackupPoliciesIdDelete

> ApiBackupPoliciesIdDelete(ctx, id).Execute()

Delete BackupPolicies



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
    id := "id_example" // string | BackupPolicy ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPoliciesApi.ApiBackupPoliciesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPoliciesApi.ApiBackupPoliciesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | BackupPolicy ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupPoliciesIdDeleteRequest struct via the builder pattern


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


## ApiBackupPoliciesIdGet

> ModelsBackupPolicy ApiBackupPoliciesIdGet(ctx, id).Execute()

Get BackupPolicy



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
    id := "id_example" // string | BackupPolicy ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPoliciesApi.ApiBackupPoliciesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPoliciesApi.ApiBackupPoliciesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiBackupPoliciesIdGet`: ModelsBackupPolicy
    fmt.Fprintf(os.Stdout, "Response from `BackupPoliciesApi.ApiBackupPoliciesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | BackupPolicy ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupPoliciesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsBackupPolicy**](ModelsBackupPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiBackupPoliciesIdPut

> ModelsBackupPolicy ApiBackupPoliciesIdPut(ctx, id).Body(body).Execute()

Update BackupPolicy



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
    id := "id_example" // string | BackupPolicy ID (must be valid UUID)
    body := *openapiclient.NewControllersUpdateBackupPolicyRequest() // ControllersUpdateBackupPolicyRequest | Request body containing updated backup policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPoliciesApi.ApiBackupPoliciesIdPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPoliciesApi.ApiBackupPoliciesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiBackupPoliciesIdPut`: ModelsBackupPolicy
    fmt.Fprintf(os.Stdout, "Response from `BackupPoliciesApi.ApiBackupPoliciesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | BackupPolicy ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBackupPoliciesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersUpdateBackupPolicyRequest**](ControllersUpdateBackupPolicyRequest.md) | Request body containing updated backup policy | 

### Return type

[**ModelsBackupPolicy**](ModelsBackupPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdBackupPoliciesGet

> ModelsPaginatedResultModelsBackupPolicy ApiTenantsIdBackupPoliciesGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Execute()

List BackupPolicies



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
    sortBy := "sortBy_example" // string | A given BackupPolicy attribute to sort results by (one of: id, name, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by BackupPolicy id (optional)
    name := "name_example" // string | Filter results by BackupPolicy name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPoliciesApi.ApiTenantsIdBackupPoliciesGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPoliciesApi.ApiTenantsIdBackupPoliciesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdBackupPoliciesGet`: ModelsPaginatedResultModelsBackupPolicy
    fmt.Fprintf(os.Stdout, "Response from `BackupPoliciesApi.ApiTenantsIdBackupPoliciesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdBackupPoliciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given BackupPolicy attribute to sort results by (one of: id, name, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by BackupPolicy id | 
 **name** | **string** | Filter results by BackupPolicy name | 

### Return type

[**ModelsPaginatedResultModelsBackupPolicy**](ModelsPaginatedResultModelsBackupPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdBackupPoliciesPost

> ModelsBackupPolicy ApiTenantsIdBackupPoliciesPost(ctx, id).Body(body).Execute()

Create BackupPolicy



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
    body := *openapiclient.NewControllersCreateBackupPolicyRequest() // ControllersCreateBackupPolicyRequest | Request body containing the backup policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPoliciesApi.ApiTenantsIdBackupPoliciesPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPoliciesApi.ApiTenantsIdBackupPoliciesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdBackupPoliciesPost`: ModelsBackupPolicy
    fmt.Fprintf(os.Stdout, "Response from `BackupPoliciesApi.ApiTenantsIdBackupPoliciesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdBackupPoliciesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersCreateBackupPolicyRequest**](ControllersCreateBackupPolicyRequest.md) | Request body containing the backup policy | 

### Return type

[**ModelsBackupPolicy**](ModelsBackupPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

