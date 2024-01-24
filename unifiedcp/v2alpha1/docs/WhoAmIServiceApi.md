# \WhoAmIServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WhoAmIServiceWhoAmI**](WhoAmIServiceApi.md#WhoAmIServiceWhoAmI) | **Get** /v1/whoami | WhoAmI API returns the current actor accounts.



## WhoAmIServiceWhoAmI

> V1WhoAmIResponse WhoAmIServiceWhoAmI(ctx).Execute()

WhoAmI API returns the current actor accounts.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WhoAmIServiceApi.WhoAmIServiceWhoAmI(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhoAmIServiceApi.WhoAmIServiceWhoAmI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WhoAmIServiceWhoAmI`: V1WhoAmIResponse
    fmt.Fprintf(os.Stdout, "Response from `WhoAmIServiceApi.WhoAmIServiceWhoAmI`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWhoAmIServiceWhoAmIRequest struct via the builder pattern


### Return type

[**V1WhoAmIResponse**](V1WhoAmIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

