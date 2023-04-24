# \MetadataApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiMetadataGet**](MetadataApi.md#ApiMetadataGet) | **Get** /api/metadata | Get metadata information



## ApiMetadataGet

> ControllersAPIMetadataResponse ApiMetadataGet(ctx).Execute()

Get metadata information



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
    resp, r, err := apiClient.MetadataApi.ApiMetadataGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.ApiMetadataGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiMetadataGet`: ControllersAPIMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.ApiMetadataGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiMetadataGetRequest struct via the builder pattern


### Return type

[**ControllersAPIMetadataResponse**](ControllersAPIMetadataResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

