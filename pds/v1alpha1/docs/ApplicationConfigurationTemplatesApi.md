# \ApplicationConfigurationTemplatesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiApplicationConfigurationTemplatesIdDelete**](ApplicationConfigurationTemplatesApi.md#ApiApplicationConfigurationTemplatesIdDelete) | **Delete** /api/application-configuration-templates/{id} | Delete ApplicationConfigurationTemplates
[**ApiApplicationConfigurationTemplatesIdGet**](ApplicationConfigurationTemplatesApi.md#ApiApplicationConfigurationTemplatesIdGet) | **Get** /api/application-configuration-templates/{id} | Get ApplicationConfigurationTemplate
[**ApiApplicationConfigurationTemplatesIdPut**](ApplicationConfigurationTemplatesApi.md#ApiApplicationConfigurationTemplatesIdPut) | **Put** /api/application-configuration-templates/{id} | Update ApplicationConfigurationTemplate
[**ApiTenantsIdApplicationConfigurationTemplatesGet**](ApplicationConfigurationTemplatesApi.md#ApiTenantsIdApplicationConfigurationTemplatesGet) | **Get** /api/tenants/{id}/application-configuration-templates | List ApplicationConfigurationTemplates
[**ApiTenantsIdApplicationConfigurationTemplatesPost**](ApplicationConfigurationTemplatesApi.md#ApiTenantsIdApplicationConfigurationTemplatesPost) | **Post** /api/tenants/{id}/application-configuration-templates | Create ApplicationConfigurationTemplates



## ApiApplicationConfigurationTemplatesIdDelete

> ApiApplicationConfigurationTemplatesIdDelete(ctx, id).Execute()

Delete ApplicationConfigurationTemplates



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
    id := "id_example" // string | ApplicationConfigurationTemplate ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationConfigurationTemplatesApi.ApiApplicationConfigurationTemplatesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConfigurationTemplatesApi.ApiApplicationConfigurationTemplatesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ApplicationConfigurationTemplate ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiApplicationConfigurationTemplatesIdDeleteRequest struct via the builder pattern


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


## ApiApplicationConfigurationTemplatesIdGet

> ModelsApplicationConfigurationTemplate ApiApplicationConfigurationTemplatesIdGet(ctx, id).Execute()

Get ApplicationConfigurationTemplate



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
    id := "id_example" // string | ApplicationConfigurationTemplate ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationConfigurationTemplatesApi.ApiApplicationConfigurationTemplatesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConfigurationTemplatesApi.ApiApplicationConfigurationTemplatesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiApplicationConfigurationTemplatesIdGet`: ModelsApplicationConfigurationTemplate
    fmt.Fprintf(os.Stdout, "Response from `ApplicationConfigurationTemplatesApi.ApiApplicationConfigurationTemplatesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ApplicationConfigurationTemplate ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiApplicationConfigurationTemplatesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsApplicationConfigurationTemplate**](ModelsApplicationConfigurationTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiApplicationConfigurationTemplatesIdPut

> ModelsApplicationConfigurationTemplate ApiApplicationConfigurationTemplatesIdPut(ctx, id).Body(body).Execute()

Update ApplicationConfigurationTemplate



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
    id := "id_example" // string | ApplicationConfigurationTemplate ID (must be valid UUID)
    body := *openapiclient.NewControllersUpdateApplicationConfigurationTemplateRequest() // ControllersUpdateApplicationConfigurationTemplateRequest | Request body containing updated template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationConfigurationTemplatesApi.ApiApplicationConfigurationTemplatesIdPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConfigurationTemplatesApi.ApiApplicationConfigurationTemplatesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiApplicationConfigurationTemplatesIdPut`: ModelsApplicationConfigurationTemplate
    fmt.Fprintf(os.Stdout, "Response from `ApplicationConfigurationTemplatesApi.ApiApplicationConfigurationTemplatesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ApplicationConfigurationTemplate ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiApplicationConfigurationTemplatesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersUpdateApplicationConfigurationTemplateRequest**](ControllersUpdateApplicationConfigurationTemplateRequest.md) | Request body containing updated template | 

### Return type

[**ModelsApplicationConfigurationTemplate**](ModelsApplicationConfigurationTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdApplicationConfigurationTemplatesGet

> ControllersPaginatedApplicationConfigurationTemplates ApiTenantsIdApplicationConfigurationTemplatesGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).DataServiceId(dataServiceId).Execute()

List ApplicationConfigurationTemplates



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
    sortBy := "sortBy_example" // string | A given ApplicationConfigurationTemplates attribute to sort results by (one of: id, name, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by ApplicationConfigurationTemplates id (optional)
    name := "name_example" // string | Filter results by ApplicationConfigurationTemplates name (optional)
    dataServiceId := "dataServiceId_example" // string | Filter results by DataService ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationConfigurationTemplatesApi.ApiTenantsIdApplicationConfigurationTemplatesGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).DataServiceId(dataServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConfigurationTemplatesApi.ApiTenantsIdApplicationConfigurationTemplatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdApplicationConfigurationTemplatesGet`: ControllersPaginatedApplicationConfigurationTemplates
    fmt.Fprintf(os.Stdout, "Response from `ApplicationConfigurationTemplatesApi.ApiTenantsIdApplicationConfigurationTemplatesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdApplicationConfigurationTemplatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given ApplicationConfigurationTemplates attribute to sort results by (one of: id, name, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by ApplicationConfigurationTemplates id | 
 **name** | **string** | Filter results by ApplicationConfigurationTemplates name | 
 **dataServiceId** | **string** | Filter results by DataService ID | 

### Return type

[**ControllersPaginatedApplicationConfigurationTemplates**](ControllersPaginatedApplicationConfigurationTemplates.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdApplicationConfigurationTemplatesPost

> ModelsApplicationConfigurationTemplate ApiTenantsIdApplicationConfigurationTemplatesPost(ctx, id).Body(body).Execute()

Create ApplicationConfigurationTemplates



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
    body := *openapiclient.NewControllersCreateApplicationConfigurationTemplatesRequest() // ControllersCreateApplicationConfigurationTemplatesRequest | Request body containing the application configuration template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationConfigurationTemplatesApi.ApiTenantsIdApplicationConfigurationTemplatesPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConfigurationTemplatesApi.ApiTenantsIdApplicationConfigurationTemplatesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdApplicationConfigurationTemplatesPost`: ModelsApplicationConfigurationTemplate
    fmt.Fprintf(os.Stdout, "Response from `ApplicationConfigurationTemplatesApi.ApiTenantsIdApplicationConfigurationTemplatesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdApplicationConfigurationTemplatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersCreateApplicationConfigurationTemplatesRequest**](ControllersCreateApplicationConfigurationTemplatesRequest.md) | Request body containing the application configuration template | 

### Return type

[**ModelsApplicationConfigurationTemplate**](ModelsApplicationConfigurationTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

