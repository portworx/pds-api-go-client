# \CopilotApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCopilotSearchPost**](CopilotApi.md#ApiCopilotSearchPost) | **Post** /api/copilot/search | Search Database query



## ApiCopilotSearchPost

> ModelsCopilotSearchResponse ApiCopilotSearchPost(ctx).Body(body).Execute()

Search Database query



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
    body := *openapiclient.NewRequestsCreateCopilotSearchRequest() // RequestsCreateCopilotSearchRequest | Request body containing the search config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CopilotApi.ApiCopilotSearchPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CopilotApi.ApiCopilotSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiCopilotSearchPost`: ModelsCopilotSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `CopilotApi.ApiCopilotSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiCopilotSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RequestsCreateCopilotSearchRequest**](RequestsCreateCopilotSearchRequest.md) | Request body containing the search config | 

### Return type

[**ModelsCopilotSearchResponse**](ModelsCopilotSearchResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

