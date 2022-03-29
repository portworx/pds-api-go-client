# \AuthorizerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAuthorizerPost**](AuthorizerApi.md#ApiAuthorizerPost) | **Post** /api/authorizer | Authorize operation



## ApiAuthorizerPost

> ModelsAuthorizerResponse ApiAuthorizerPost(ctx).Body(body).Execute()

Authorize operation



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
    body := *openapiclient.NewModelsAuthorizerRequest() // ModelsAuthorizerRequest | Request body containing the authorizer request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthorizerApi.ApiAuthorizerPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizerApi.ApiAuthorizerPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAuthorizerPost`: ModelsAuthorizerResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizerApi.ApiAuthorizerPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAuthorizerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ModelsAuthorizerRequest**](ModelsAuthorizerRequest.md) | Request body containing the authorizer request | 

### Return type

[**ModelsAuthorizerResponse**](ModelsAuthorizerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

