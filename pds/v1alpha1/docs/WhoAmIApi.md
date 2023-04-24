# \WhoAmIApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiWhoamiGet**](WhoAmIApi.md#ApiWhoamiGet) | **Get** /api/whoami | Get Current Actor



## ApiWhoamiGet

> ControllersWhoAmIResponse ApiWhoamiGet(ctx).Execute()

Get Current Actor



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
    resp, r, err := apiClient.WhoAmIApi.ApiWhoamiGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhoAmIApi.ApiWhoamiGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiWhoamiGet`: ControllersWhoAmIResponse
    fmt.Fprintf(os.Stdout, "Response from `WhoAmIApi.ApiWhoamiGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiWhoamiGetRequest struct via the builder pattern


### Return type

[**ControllersWhoAmIResponse**](ControllersWhoAmIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

