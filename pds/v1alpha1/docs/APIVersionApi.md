# \APIVersionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiVersionGet**](APIVersionApi.md#ApiVersionGet) | **Get** /api/version | Get version information



## ApiVersionGet

> ControllersAPIVersionResponse ApiVersionGet(ctx).Execute()

Get version information



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIVersionApi.ApiVersionGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIVersionApi.ApiVersionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiVersionGet`: ControllersAPIVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `APIVersionApi.ApiVersionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiVersionGetRequest struct via the builder pattern


### Return type

[**ControllersAPIVersionResponse**](ControllersAPIVersionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

