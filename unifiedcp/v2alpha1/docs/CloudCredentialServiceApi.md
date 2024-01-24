# \CloudCredentialServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudCredentialServiceCreateCloudCredential**](CloudCredentialServiceApi.md#CloudCredentialServiceCreateCloudCredential) | **Post** /v1/tenants/{tenantId}/cloudCredentials | Create API creates a set of cloud credentials for a tenant
[**CloudCredentialServiceDeleteCloudCredential**](CloudCredentialServiceApi.md#CloudCredentialServiceDeleteCloudCredential) | **Delete** /v1/cloudCredentials/{id} | Delete API deletes the cloud credentials
[**CloudCredentialServiceGetCloudCredential**](CloudCredentialServiceApi.md#CloudCredentialServiceGetCloudCredential) | **Get** /v1/cloudcredentials/{id} | Get API returns the cloud credential details sans the actual credentials
[**CloudCredentialServiceListCloudCredentials**](CloudCredentialServiceApi.md#CloudCredentialServiceListCloudCredentials) | **Get** /v1/cloudCredentials | List API lists all the cloud credentials for a tenant or project
[**CloudCredentialServiceUpdateCloudCredential**](CloudCredentialServiceApi.md#CloudCredentialServiceUpdateCloudCredential) | **Put** /v1/cloudCredentials/{id} | Update API updates a cloud credential



## CloudCredentialServiceCreateCloudCredential

> V1CloudCredential CloudCredentialServiceCreateCloudCredential(ctx, tenantId).V1CloudCredential(v1CloudCredential).Execute()

Create API creates a set of cloud credentials for a tenant

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
    tenantId := "tenantId_example" // string | The parent tenant id under which cloud credential will be created (-- api-linter: core::0133::request-unknown-fields=disabled     aip.dev/not-precedent: We need this field for to support creation of     the resource in the tenant context. --)
    v1CloudCredential := *openapiclient.NewV1CloudCredential() // V1CloudCredential | Cloud credential configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudCredentialServiceApi.CloudCredentialServiceCreateCloudCredential(context.Background(), tenantId).V1CloudCredential(v1CloudCredential).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialServiceApi.CloudCredentialServiceCreateCloudCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudCredentialServiceCreateCloudCredential`: V1CloudCredential
    fmt.Fprintf(os.Stdout, "Response from `CloudCredentialServiceApi.CloudCredentialServiceCreateCloudCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | The parent tenant id under which cloud credential will be created (-- api-linter: core::0133::request-unknown-fields&#x3D;disabled     aip.dev/not-precedent: We need this field for to support creation of     the resource in the tenant context. --) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudCredentialServiceCreateCloudCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CloudCredential** | [**V1CloudCredential**](V1CloudCredential.md) | Cloud credential configuration | 

### Return type

[**V1CloudCredential**](V1CloudCredential.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudCredentialServiceDeleteCloudCredential

> map[string]interface{} CloudCredentialServiceDeleteCloudCredential(ctx, id).Execute()

Delete API deletes the cloud credentials

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
    id := "id_example" // string | ID of the cloud credential

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudCredentialServiceApi.CloudCredentialServiceDeleteCloudCredential(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialServiceApi.CloudCredentialServiceDeleteCloudCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudCredentialServiceDeleteCloudCredential`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudCredentialServiceApi.CloudCredentialServiceDeleteCloudCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the cloud credential | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudCredentialServiceDeleteCloudCredentialRequest struct via the builder pattern


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


## CloudCredentialServiceGetCloudCredential

> V1CloudCredential CloudCredentialServiceGetCloudCredential(ctx, id).IncludeConfig(includeConfig).Execute()

Get API returns the cloud credential details sans the actual credentials

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
    id := "id_example" // string | ID of the cloud credential
    includeConfig := true // bool | Specifies option to include configuration details excluding sensitive information (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudCredentialServiceApi.CloudCredentialServiceGetCloudCredential(context.Background(), id).IncludeConfig(includeConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialServiceApi.CloudCredentialServiceGetCloudCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudCredentialServiceGetCloudCredential`: V1CloudCredential
    fmt.Fprintf(os.Stdout, "Response from `CloudCredentialServiceApi.CloudCredentialServiceGetCloudCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the cloud credential | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudCredentialServiceGetCloudCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeConfig** | **bool** | Specifies option to include configuration details excluding sensitive information | 

### Return type

[**V1CloudCredential**](V1CloudCredential.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudCredentialServiceListCloudCredentials

> V1ListCloudCredentialsResponse CloudCredentialServiceListCloudCredentials(ctx).TenantId(tenantId).ProjectId(projectId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()

List API lists all the cloud credentials for a tenant or project

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
    tenantId := "tenantId_example" // string | Tenant ID for which the credentials will be listed (optional)
    projectId := "projectId_example" // string | Project ID for which the credentials will be listed (optional)
    paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
    paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudCredentialServiceApi.CloudCredentialServiceListCloudCredentials(context.Background()).TenantId(tenantId).ProjectId(projectId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialServiceApi.CloudCredentialServiceListCloudCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudCredentialServiceListCloudCredentials`: V1ListCloudCredentialsResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudCredentialServiceApi.CloudCredentialServiceListCloudCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudCredentialServiceListCloudCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant ID for which the credentials will be listed | 
 **projectId** | **string** | Project ID for which the credentials will be listed | 
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 

### Return type

[**V1ListCloudCredentialsResponse**](V1ListCloudCredentialsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudCredentialServiceUpdateCloudCredential

> V1CloudCredential CloudCredentialServiceUpdateCloudCredential(ctx, id).V1CloudCredential(v1CloudCredential).UpdateMask(updateMask).Execute()

Update API updates a cloud credential

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
    id := "id_example" // string | id of the cloud credential to be updated
    v1CloudCredential := *openapiclient.NewV1CloudCredential() // V1CloudCredential | Desired cloud credential configuration
    updateMask := "updateMask_example" // string | Specifies the field that should be updated to the value specified in cloud_credentials (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudCredentialServiceApi.CloudCredentialServiceUpdateCloudCredential(context.Background(), id).V1CloudCredential(v1CloudCredential).UpdateMask(updateMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialServiceApi.CloudCredentialServiceUpdateCloudCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudCredentialServiceUpdateCloudCredential`: V1CloudCredential
    fmt.Fprintf(os.Stdout, "Response from `CloudCredentialServiceApi.CloudCredentialServiceUpdateCloudCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the cloud credential to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudCredentialServiceUpdateCloudCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CloudCredential** | [**V1CloudCredential**](V1CloudCredential.md) | Desired cloud credential configuration | 
 **updateMask** | **string** | Specifies the field that should be updated to the value specified in cloud_credentials | 

### Return type

[**V1CloudCredential**](V1CloudCredential.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

