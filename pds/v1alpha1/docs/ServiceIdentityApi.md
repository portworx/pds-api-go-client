# \ServiceIdentityApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountsIdServiceIdentityPost**](ServiceIdentityApi.md#ApiAccountsIdServiceIdentityPost) | **Post** /api/accounts/{id}/service-identity | Create a Service Identity
[**ApiAccountsServiceIdentityregenerateIdGet**](ServiceIdentityApi.md#ApiAccountsServiceIdentityregenerateIdGet) | **Get** /api/accounts/service-identity:regenerate/{id} | Regenerate service identity credentials
[**ServiceIdentityGenerateTokenPost**](ServiceIdentityApi.md#ServiceIdentityGenerateTokenPost) | **Post** /service-identity/generate-token | Generate JWT token for service identity



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
    body := *openapiclient.NewRequestsServiceIdentityRequest("Name_example") // RequestsServiceIdentityRequest | Request body containing the service identity details.

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


## ApiAccountsServiceIdentityregenerateIdGet

> ModelsServiceIdentityWithToken ApiAccountsServiceIdentityregenerateIdGet(ctx, id).Execute()

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
    resp, r, err := apiClient.ServiceIdentityApi.ApiAccountsServiceIdentityregenerateIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIdentityApi.ApiAccountsServiceIdentityregenerateIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsServiceIdentityregenerateIdGet`: ModelsServiceIdentityWithToken
    fmt.Fprintf(os.Stdout, "Response from `ServiceIdentityApi.ApiAccountsServiceIdentityregenerateIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Service Identity ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsServiceIdentityregenerateIdGetRequest struct via the builder pattern


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

