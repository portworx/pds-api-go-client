# \AccountsDNSApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountsIdDnsDetailsPut**](AccountsDNSApi.md#ApiAccountsIdDnsDetailsPut) | **Put** /api/accounts/{id}/dns-details | Update DNS Details



## ApiAccountsIdDnsDetailsPut

> ApiAccountsIdDnsDetailsPut(ctx, id).Body(body).Execute()

Update DNS Details



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
    id := "id_example" // string | Account ID (must be valid UUID)
    body := *openapiclient.NewModelsDNSDetails() // ModelsDNSDetails | Request body containing a new DNS details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccountsDNSApi.ApiAccountsIdDnsDetailsPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsDNSApi.ApiAccountsIdDnsDetailsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsIdDnsDetailsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ModelsDNSDetails**](ModelsDNSDetails.md) | Request body containing a new DNS details | 

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

