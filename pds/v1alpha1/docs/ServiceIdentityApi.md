# \ServiceIdentityApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountsIdServiceIdentityPost**](ServiceIdentityApi.md#ApiAccountsIdServiceIdentityPost) | **Post** /api/accounts/{id}/service-identity | Create a Service Identity



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

