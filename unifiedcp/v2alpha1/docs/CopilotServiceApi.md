# \CopilotServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopilotServiceSearchQuery**](CopilotServiceApi.md#CopilotServiceSearchQuery) | **Post** /pds/v1/copilot/search | Search dataservice queries.



## CopilotServiceSearchQuery

> V1SearchQueryResponse CopilotServiceSearchQuery(ctx).V1SearchQueryRequest(v1SearchQueryRequest).Execute()

Search dataservice queries.

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
    v1SearchQueryRequest := *openapiclient.NewV1SearchQueryRequest() // V1SearchQueryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CopilotServiceApi.CopilotServiceSearchQuery(context.Background()).V1SearchQueryRequest(v1SearchQueryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CopilotServiceApi.CopilotServiceSearchQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopilotServiceSearchQuery`: V1SearchQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `CopilotServiceApi.CopilotServiceSearchQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCopilotServiceSearchQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1SearchQueryRequest** | [**V1SearchQueryRequest**](V1SearchQueryRequest.md) |  | 

### Return type

[**V1SearchQueryResponse**](V1SearchQueryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

