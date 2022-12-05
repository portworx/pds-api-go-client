# \StorageOptionsTemplatesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiStorageOptionsTemplatesIdDelete**](StorageOptionsTemplatesApi.md#ApiStorageOptionsTemplatesIdDelete) | **Delete** /api/storage-options-templates/{id} | Delete StorageOptionsTemplates
[**ApiStorageOptionsTemplatesIdGet**](StorageOptionsTemplatesApi.md#ApiStorageOptionsTemplatesIdGet) | **Get** /api/storage-options-templates/{id} | Get StorageOptionsTemplate
[**ApiStorageOptionsTemplatesIdPut**](StorageOptionsTemplatesApi.md#ApiStorageOptionsTemplatesIdPut) | **Put** /api/storage-options-templates/{id} | Update StorageOptionsTemplate
[**ApiTenantsIdStorageOptionsTemplatesGet**](StorageOptionsTemplatesApi.md#ApiTenantsIdStorageOptionsTemplatesGet) | **Get** /api/tenants/{id}/storage-options-templates | List StorageOptionsTemplates
[**ApiTenantsIdStorageOptionsTemplatesPost**](StorageOptionsTemplatesApi.md#ApiTenantsIdStorageOptionsTemplatesPost) | **Post** /api/tenants/{id}/storage-options-templates | Create StorageOptionsTemplate



## ApiStorageOptionsTemplatesIdDelete

> ApiStorageOptionsTemplatesIdDelete(ctx, id).Execute()

Delete StorageOptionsTemplates



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
    id := "id_example" // string | StorageOptionsTemplate ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageOptionsTemplatesApi.ApiStorageOptionsTemplatesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageOptionsTemplatesApi.ApiStorageOptionsTemplatesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | StorageOptionsTemplate ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiStorageOptionsTemplatesIdDeleteRequest struct via the builder pattern


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


## ApiStorageOptionsTemplatesIdGet

> ModelsStorageOptionsTemplate ApiStorageOptionsTemplatesIdGet(ctx, id).Execute()

Get StorageOptionsTemplate



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
    id := "id_example" // string | StorageOptionsTemplate ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageOptionsTemplatesApi.ApiStorageOptionsTemplatesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageOptionsTemplatesApi.ApiStorageOptionsTemplatesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiStorageOptionsTemplatesIdGet`: ModelsStorageOptionsTemplate
    fmt.Fprintf(os.Stdout, "Response from `StorageOptionsTemplatesApi.ApiStorageOptionsTemplatesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | StorageOptionsTemplate ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiStorageOptionsTemplatesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsStorageOptionsTemplate**](ModelsStorageOptionsTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiStorageOptionsTemplatesIdPut

> ModelsStorageOptionsTemplate ApiStorageOptionsTemplatesIdPut(ctx, id).Body(body).Execute()

Update StorageOptionsTemplate



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
    id := "id_example" // string | StorageOptionsTemplate ID (must be valid UUID)
    body := *openapiclient.NewControllersUpdateStorageOptionsTemplateRequest() // ControllersUpdateStorageOptionsTemplateRequest | Request body containing updated template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageOptionsTemplatesApi.ApiStorageOptionsTemplatesIdPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageOptionsTemplatesApi.ApiStorageOptionsTemplatesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiStorageOptionsTemplatesIdPut`: ModelsStorageOptionsTemplate
    fmt.Fprintf(os.Stdout, "Response from `StorageOptionsTemplatesApi.ApiStorageOptionsTemplatesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | StorageOptionsTemplate ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiStorageOptionsTemplatesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersUpdateStorageOptionsTemplateRequest**](ControllersUpdateStorageOptionsTemplateRequest.md) | Request body containing updated template | 

### Return type

[**ModelsStorageOptionsTemplate**](ModelsStorageOptionsTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdStorageOptionsTemplatesGet

> ModelsPaginatedResultModelsStorageOptionsTemplate ApiTenantsIdStorageOptionsTemplatesGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Provisioner(provisioner).Execute()

List StorageOptionsTemplates



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
    sortBy := "sortBy_example" // string | A given StorageOptionsTemplates attribute to sort results by (one of: id, name, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by StorageOptionsTemplates id (optional)
    name := "name_example" // string | Filter results by StorageOptionsTemplates name (optional)
    provisioner := []string{"Provisioner_example"} // []string | Filter results by StorageOptionsTemplates provisioner (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageOptionsTemplatesApi.ApiTenantsIdStorageOptionsTemplatesGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).Provisioner(provisioner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageOptionsTemplatesApi.ApiTenantsIdStorageOptionsTemplatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdStorageOptionsTemplatesGet`: ModelsPaginatedResultModelsStorageOptionsTemplate
    fmt.Fprintf(os.Stdout, "Response from `StorageOptionsTemplatesApi.ApiTenantsIdStorageOptionsTemplatesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdStorageOptionsTemplatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given StorageOptionsTemplates attribute to sort results by (one of: id, name, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by StorageOptionsTemplates id | 
 **name** | **string** | Filter results by StorageOptionsTemplates name | 
 **provisioner** | **[]string** | Filter results by StorageOptionsTemplates provisioner | 

### Return type

[**ModelsPaginatedResultModelsStorageOptionsTemplate**](ModelsPaginatedResultModelsStorageOptionsTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdStorageOptionsTemplatesPost

> ModelsStorageOptionsTemplate ApiTenantsIdStorageOptionsTemplatesPost(ctx, id).Body(body).Execute()

Create StorageOptionsTemplate



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
    body := *openapiclient.NewControllersCreateStorageOptionsTemplateRequest() // ControllersCreateStorageOptionsTemplateRequest | Request body containing the storage options template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageOptionsTemplatesApi.ApiTenantsIdStorageOptionsTemplatesPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageOptionsTemplatesApi.ApiTenantsIdStorageOptionsTemplatesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdStorageOptionsTemplatesPost`: ModelsStorageOptionsTemplate
    fmt.Fprintf(os.Stdout, "Response from `StorageOptionsTemplatesApi.ApiTenantsIdStorageOptionsTemplatesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdStorageOptionsTemplatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersCreateStorageOptionsTemplateRequest**](ControllersCreateStorageOptionsTemplateRequest.md) | Request body containing the storage options template | 

### Return type

[**ModelsStorageOptionsTemplate**](ModelsStorageOptionsTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

