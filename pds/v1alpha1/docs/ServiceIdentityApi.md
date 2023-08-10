# \ServiceIdentityApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountsIdServiceIdentityGet**](ServiceIdentityApi.md#ApiAccountsIdServiceIdentityGet) | **Get** /api/accounts/{id}/service-identity | List Service Identity
[**ApiAccountsIdServiceIdentityPost**](ServiceIdentityApi.md#ApiAccountsIdServiceIdentityPost) | **Post** /api/accounts/{id}/service-identity | Create a Service Identity
[**ApiServiceIdentityIdDelete**](ServiceIdentityApi.md#ApiServiceIdentityIdDelete) | **Delete** /api/service-identity/{id} | Delete service identity
[**ApiServiceIdentityIdGet**](ServiceIdentityApi.md#ApiServiceIdentityIdGet) | **Get** /api/service-identity/{id} | Get service identity by ID
[**ApiServiceIdentityIdPut**](ServiceIdentityApi.md#ApiServiceIdentityIdPut) | **Put** /api/service-identity/{id} | Update service identity
[**ApiServiceIdentityregenerateIdGet**](ServiceIdentityApi.md#ApiServiceIdentityregenerateIdGet) | **Get** /api/service-identity:regenerate/{id} | Regenerate service identity credentials
[**ServiceIdentityGenerateTokenPost**](ServiceIdentityApi.md#ServiceIdentityGenerateTokenPost) | **Post** /service-identity/generate-token | Generate JWT token for service identity



## ApiAccountsIdServiceIdentityGet

> ModelsPaginatedResultModelsServiceIdentity ApiAccountsIdServiceIdentityGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).CreatedBy(createdBy).ClientId(clientId).Enabled(enabled).Execute()

List Service Identity



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
    id := "id_example" // string | Account ID (must be valid UUID)
    sortBy := "sortBy_example" // string | A given Service Identity attribute to sort results by (one of: id, name, created_at, created_by) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by service identity id (optional)
    name := "name_example" // string | Filter results by name (optional)
    createdBy := "createdBy_example" // string | Filter results by created_by (optional)
    clientId := "clientId_example" // string | Filter results by client_id (optional)
    enabled := true // bool | Filter results by enabled (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceIdentityApi.ApiAccountsIdServiceIdentityGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Name(name).CreatedBy(createdBy).ClientId(clientId).Enabled(enabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIdentityApi.ApiAccountsIdServiceIdentityGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsIdServiceIdentityGet`: ModelsPaginatedResultModelsServiceIdentity
    fmt.Fprintf(os.Stdout, "Response from `ServiceIdentityApi.ApiAccountsIdServiceIdentityGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsIdServiceIdentityGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given Service Identity attribute to sort results by (one of: id, name, created_at, created_by) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by service identity id | 
 **name** | **string** | Filter results by name | 
 **createdBy** | **string** | Filter results by created_by | 
 **clientId** | **string** | Filter results by client_id | 
 **enabled** | **bool** | Filter results by enabled | 

### Return type

[**ModelsPaginatedResultModelsServiceIdentity**](ModelsPaginatedResultModelsServiceIdentity.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsIdServiceIdentityPost

> ModelsServiceIdentityWithToken ApiAccountsIdServiceIdentityPost(ctx, id).Body(body).Execute()

Create a Service Identity



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
    id := "id_example" // string | Account ID (must be valid UUID)
    body := *openapiclient.NewRequestsServiceIdentityRequest(false, "Name_example") // RequestsServiceIdentityRequest | Request body containing the service identity details.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceIdentityApi.ApiAccountsIdServiceIdentityPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIdentityApi.ApiAccountsIdServiceIdentityPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsIdServiceIdentityPost`: ModelsServiceIdentityWithToken
    fmt.Fprintf(os.Stdout, "Response from `ServiceIdentityApi.ApiAccountsIdServiceIdentityPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsIdServiceIdentityPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestsServiceIdentityRequest**](RequestsServiceIdentityRequest.md) | Request body containing the service identity details. | 

### Return type

[**ModelsServiceIdentityWithToken**](ModelsServiceIdentityWithToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiServiceIdentityIdDelete

> ApiServiceIdentityIdDelete(ctx, id).Execute()

Delete service identity



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
    id := "id_example" // string | Service Identity ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceIdentityApi.ApiServiceIdentityIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIdentityApi.ApiServiceIdentityIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Service Identity ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceIdentityIdDeleteRequest struct via the builder pattern


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


## ApiServiceIdentityIdGet

> ModelsServiceIdentity ApiServiceIdentityIdGet(ctx, id).Execute()

Get service identity by ID



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
    id := "id_example" // string | Service Identity ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceIdentityApi.ApiServiceIdentityIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIdentityApi.ApiServiceIdentityIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceIdentityIdGet`: ModelsServiceIdentity
    fmt.Fprintf(os.Stdout, "Response from `ServiceIdentityApi.ApiServiceIdentityIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Service Identity ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceIdentityIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsServiceIdentity**](ModelsServiceIdentity.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiServiceIdentityIdPut

> ApiServiceIdentityIdPut(ctx, id).Body(body).Execute()

Update service identity



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
    id := "id_example" // string | Service Identity ID (must be valid UUID)
    body := *openapiclient.NewRequestsServiceIdentityRequest(false, "Name_example") // RequestsServiceIdentityRequest | Request body containing the service identity details.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceIdentityApi.ApiServiceIdentityIdPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIdentityApi.ApiServiceIdentityIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Service Identity ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceIdentityIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestsServiceIdentityRequest**](RequestsServiceIdentityRequest.md) | Request body containing the service identity details. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiServiceIdentityregenerateIdGet

> ModelsServiceIdentityWithToken ApiServiceIdentityregenerateIdGet(ctx, id).Execute()

Regenerate service identity credentials



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
    id := "id_example" // string | Service Identity ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceIdentityApi.ApiServiceIdentityregenerateIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIdentityApi.ApiServiceIdentityregenerateIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceIdentityregenerateIdGet`: ModelsServiceIdentityWithToken
    fmt.Fprintf(os.Stdout, "Response from `ServiceIdentityApi.ApiServiceIdentityregenerateIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Service Identity ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceIdentityregenerateIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsServiceIdentityWithToken**](ModelsServiceIdentityWithToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceIdentityGenerateTokenPost

> ControllersGenerateTokenResponse ServiceIdentityGenerateTokenPost(ctx).Body(body).Execute()

Generate JWT token for service identity



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
    body := *openapiclient.NewControllersGenerateTokenRequest() // ControllersGenerateTokenRequest | Request body containing the client id and client token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceIdentityApi.ServiceIdentityGenerateTokenPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIdentityApi.ServiceIdentityGenerateTokenPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceIdentityGenerateTokenPost`: ControllersGenerateTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceIdentityApi.ServiceIdentityGenerateTokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceIdentityGenerateTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ControllersGenerateTokenRequest**](ControllersGenerateTokenRequest.md) | Request body containing the client id and client token. | 

### Return type

[**ControllersGenerateTokenResponse**](ControllersGenerateTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

