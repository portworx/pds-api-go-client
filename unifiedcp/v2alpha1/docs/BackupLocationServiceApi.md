# \BackupLocationServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupLocationServiceCreateBackupLocation**](BackupLocationServiceApi.md#BackupLocationServiceCreateBackupLocation) | **Post** /v1/tenants/{tenantId}/backupLocations | Create API creates a backup location for a tenant
[**BackupLocationServiceDeleteBackupLocation**](BackupLocationServiceApi.md#BackupLocationServiceDeleteBackupLocation) | **Delete** /v1/backupLocations/{id} | Delete API deletes a backup location
[**BackupLocationServiceGetBackupLocation**](BackupLocationServiceApi.md#BackupLocationServiceGetBackupLocation) | **Get** /v1/backupLocations/{id} | Get API returns the backup location
[**BackupLocationServiceListBackupLocations**](BackupLocationServiceApi.md#BackupLocationServiceListBackupLocations) | **Get** /v1/backupLocations | List API lists all the backup locations for a tenant or project
[**BackupLocationServiceUpdateBackupLocation**](BackupLocationServiceApi.md#BackupLocationServiceUpdateBackupLocation) | **Put** /v1/backupLocations/{id} | Update API updates a backup location



## BackupLocationServiceCreateBackupLocation

> V1BackupLocation BackupLocationServiceCreateBackupLocation(ctx, tenantId).V1BackupLocation(v1BackupLocation).Execute()

Create API creates a backup location for a tenant

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
    tenantId := "tenantId_example" // string | The parent tenant id under which backup location will be created (-- api-linter: core::0133::request-unknown-fields=disabled     aip.dev/not-precedent: We need this field for to support creation of     the resource in the tenant context. --)
    v1BackupLocation := *openapiclient.NewV1BackupLocation() // V1BackupLocation | Backup location configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupLocationServiceApi.BackupLocationServiceCreateBackupLocation(context.Background(), tenantId).V1BackupLocation(v1BackupLocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupLocationServiceApi.BackupLocationServiceCreateBackupLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupLocationServiceCreateBackupLocation`: V1BackupLocation
    fmt.Fprintf(os.Stdout, "Response from `BackupLocationServiceApi.BackupLocationServiceCreateBackupLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | The parent tenant id under which backup location will be created (-- api-linter: core::0133::request-unknown-fields&#x3D;disabled     aip.dev/not-precedent: We need this field for to support creation of     the resource in the tenant context. --) | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupLocationServiceCreateBackupLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1BackupLocation** | [**V1BackupLocation**](V1BackupLocation.md) | Backup location configuration | 

### Return type

[**V1BackupLocation**](V1BackupLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupLocationServiceDeleteBackupLocation

> map[string]interface{} BackupLocationServiceDeleteBackupLocation(ctx, id).Execute()

Delete API deletes a backup location

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
    id := "id_example" // string | ID of the backup location

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupLocationServiceApi.BackupLocationServiceDeleteBackupLocation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupLocationServiceApi.BackupLocationServiceDeleteBackupLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupLocationServiceDeleteBackupLocation`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BackupLocationServiceApi.BackupLocationServiceDeleteBackupLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the backup location | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupLocationServiceDeleteBackupLocationRequest struct via the builder pattern


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


## BackupLocationServiceGetBackupLocation

> V1BackupLocation BackupLocationServiceGetBackupLocation(ctx, id).Execute()

Get API returns the backup location

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
    id := "id_example" // string | ID of the backup location

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupLocationServiceApi.BackupLocationServiceGetBackupLocation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupLocationServiceApi.BackupLocationServiceGetBackupLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupLocationServiceGetBackupLocation`: V1BackupLocation
    fmt.Fprintf(os.Stdout, "Response from `BackupLocationServiceApi.BackupLocationServiceGetBackupLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the backup location | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupLocationServiceGetBackupLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1BackupLocation**](V1BackupLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupLocationServiceListBackupLocations

> V1ListBackupLocationsResponse BackupLocationServiceListBackupLocations(ctx).TenantId(tenantId).ProjectId(projectId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()

List API lists all the backup locations for a tenant or project

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
    tenantId := "tenantId_example" // string | Tenant ID for which the backup locations will be listed (optional)
    projectId := "projectId_example" // string | Project ID for which the backup locations will be listed (optional)
    paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
    paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupLocationServiceApi.BackupLocationServiceListBackupLocations(context.Background()).TenantId(tenantId).ProjectId(projectId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupLocationServiceApi.BackupLocationServiceListBackupLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupLocationServiceListBackupLocations`: V1ListBackupLocationsResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupLocationServiceApi.BackupLocationServiceListBackupLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupLocationServiceListBackupLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant ID for which the backup locations will be listed | 
 **projectId** | **string** | Project ID for which the backup locations will be listed | 
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 

### Return type

[**V1ListBackupLocationsResponse**](V1ListBackupLocationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupLocationServiceUpdateBackupLocation

> V1BackupLocation BackupLocationServiceUpdateBackupLocation(ctx, id).V1BackupLocation(v1BackupLocation).UpdateMask(updateMask).Execute()

Update API updates a backup location

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
    id := "id_example" // string | ID of the backup location
    v1BackupLocation := *openapiclient.NewV1BackupLocation() // V1BackupLocation | Desired backup location configuration
    updateMask := "updateMask_example" // string | Specifies the field that should be updated to the value specified in backup_location (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupLocationServiceApi.BackupLocationServiceUpdateBackupLocation(context.Background(), id).V1BackupLocation(v1BackupLocation).UpdateMask(updateMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupLocationServiceApi.BackupLocationServiceUpdateBackupLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupLocationServiceUpdateBackupLocation`: V1BackupLocation
    fmt.Fprintf(os.Stdout, "Response from `BackupLocationServiceApi.BackupLocationServiceUpdateBackupLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the backup location | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupLocationServiceUpdateBackupLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1BackupLocation** | [**V1BackupLocation**](V1BackupLocation.md) | Desired backup location configuration | 
 **updateMask** | **string** | Specifies the field that should be updated to the value specified in backup_location | 

### Return type

[**V1BackupLocation**](V1BackupLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

