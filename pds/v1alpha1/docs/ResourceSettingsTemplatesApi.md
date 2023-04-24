# \ResourceSettingsTemplatesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiResourceSettingsTemplatesIdDelete**](ResourceSettingsTemplatesApi.md#ApiResourceSettingsTemplatesIdDelete) | **Delete** /api/resource-settings-templates/{id} | Delete ResourceSettingsTemplates
[**ApiResourceSettingsTemplatesIdGet**](ResourceSettingsTemplatesApi.md#ApiResourceSettingsTemplatesIdGet) | **Get** /api/resource-settings-templates/{id} | Get ResourceSettingsTemplate
[**ApiResourceSettingsTemplatesIdPut**](ResourceSettingsTemplatesApi.md#ApiResourceSettingsTemplatesIdPut) | **Put** /api/resource-settings-templates/{id} | Update ResourceSettingsTemplate
[**ApiTenantsIdResourceSettingsTemplatesGet**](ResourceSettingsTemplatesApi.md#ApiTenantsIdResourceSettingsTemplatesGet) | **Get** /api/tenants/{id}/resource-settings-templates | List ResourceSettingsTemplates
[**ApiTenantsIdResourceSettingsTemplatesPost**](ResourceSettingsTemplatesApi.md#ApiTenantsIdResourceSettingsTemplatesPost) | **Post** /api/tenants/{id}/resource-settings-templates | Create ResourceSettingsTemplate



## ApiResourceSettingsTemplatesIdDelete

> ApiResourceSettingsTemplatesIdDelete(ctx, id).Execute()

Delete ResourceSettingsTemplates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | ResourceSettingsTemplate ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ResourceSettingsTemplatesApi.ApiResourceSettingsTemplatesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSettingsTemplatesApi.ApiResourceSettingsTemplatesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ResourceSettingsTemplate ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiResourceSettingsTemplatesIdDeleteRequest struct via the builder pattern


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


## ApiResourceSettingsTemplatesIdGet

> ModelsResourceSettingsTemplate ApiResourceSettingsTemplatesIdGet(ctx, id).Execute()

Get ResourceSettingsTemplate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | ResourceSettingsTemplate ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceSettingsTemplatesApi.ApiResourceSettingsTemplatesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSettingsTemplatesApi.ApiResourceSettingsTemplatesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiResourceSettingsTemplatesIdGet`: ModelsResourceSettingsTemplate
    fmt.Fprintf(os.Stdout, "Response from `ResourceSettingsTemplatesApi.ApiResourceSettingsTemplatesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ResourceSettingsTemplate ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiResourceSettingsTemplatesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsResourceSettingsTemplate**](ModelsResourceSettingsTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiResourceSettingsTemplatesIdPut

> ModelsResourceSettingsTemplate ApiResourceSettingsTemplatesIdPut(ctx, id).Body(body).Execute()

Update ResourceSettingsTemplate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | ResourceSettingsTemplate ID (must be valid UUID)
    body := *openapiclient.NewControllersUpdateResourceSettingsTemplateRequest() // ControllersUpdateResourceSettingsTemplateRequest | Request body containing updated template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceSettingsTemplatesApi.ApiResourceSettingsTemplatesIdPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSettingsTemplatesApi.ApiResourceSettingsTemplatesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiResourceSettingsTemplatesIdPut`: ModelsResourceSettingsTemplate
    fmt.Fprintf(os.Stdout, "Response from `ResourceSettingsTemplatesApi.ApiResourceSettingsTemplatesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ResourceSettingsTemplate ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiResourceSettingsTemplatesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersUpdateResourceSettingsTemplateRequest**](ControllersUpdateResourceSettingsTemplateRequest.md) | Request body containing updated template | 

### Return type

[**ModelsResourceSettingsTemplate**](ModelsResourceSettingsTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdResourceSettingsTemplatesGet

> ModelsPaginatedResultModelsResourceSettingsTemplate ApiTenantsIdResourceSettingsTemplatesGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).DataServiceId(dataServiceId).Execute()

List ResourceSettingsTemplates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Tenant ID (must be valid UUID)
    sortBy := "sortBy_example" // string | A given ResourceSettingsTemplates attribute to sort results by (one of: id, name, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by ResourceSettingsTemplates id (optional)
    name := "name_example" // string | Filter results by ResourceSettingsTemplates name (optional)
    dataServiceId := "dataServiceId_example" // string | Filter results by DataService ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceSettingsTemplatesApi.ApiTenantsIdResourceSettingsTemplatesGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).DataServiceId(dataServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSettingsTemplatesApi.ApiTenantsIdResourceSettingsTemplatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdResourceSettingsTemplatesGet`: ModelsPaginatedResultModelsResourceSettingsTemplate
    fmt.Fprintf(os.Stdout, "Response from `ResourceSettingsTemplatesApi.ApiTenantsIdResourceSettingsTemplatesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdResourceSettingsTemplatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given ResourceSettingsTemplates attribute to sort results by (one of: id, name, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by ResourceSettingsTemplates id | 
 **name** | **string** | Filter results by ResourceSettingsTemplates name | 
 **dataServiceId** | **string** | Filter results by DataService ID | 

### Return type

[**ModelsPaginatedResultModelsResourceSettingsTemplate**](ModelsPaginatedResultModelsResourceSettingsTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTenantsIdResourceSettingsTemplatesPost

> ModelsResourceSettingsTemplate ApiTenantsIdResourceSettingsTemplatesPost(ctx, id).Body(body).Execute()

Create ResourceSettingsTemplate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Tenant ID (must be valid UUID)
    body := *openapiclient.NewControllersCreateResourceSettingsTemplateRequest() // ControllersCreateResourceSettingsTemplateRequest | Request body containing the resource settings template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceSettingsTemplatesApi.ApiTenantsIdResourceSettingsTemplatesPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSettingsTemplatesApi.ApiTenantsIdResourceSettingsTemplatesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTenantsIdResourceSettingsTemplatesPost`: ModelsResourceSettingsTemplate
    fmt.Fprintf(os.Stdout, "Response from `ResourceSettingsTemplatesApi.ApiTenantsIdResourceSettingsTemplatesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdResourceSettingsTemplatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersCreateResourceSettingsTemplateRequest**](ControllersCreateResourceSettingsTemplateRequest.md) | Request body containing the resource settings template | 

### Return type

[**ModelsResourceSettingsTemplate**](ModelsResourceSettingsTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

