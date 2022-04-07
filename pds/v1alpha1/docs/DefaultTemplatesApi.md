# \DefaultTemplatesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTenantsIdSeedDefaultTemplatesPost**](DefaultTemplatesApi.md#ApiTenantsIdSeedDefaultTemplatesPost) | **Post** /api/tenants/{id}/seed-default-templates | Seed default templates



## ApiTenantsIdSeedDefaultTemplatesPost

> ApiTenantsIdSeedDefaultTemplatesPost(ctx, id).Execute()

Seed default templates



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
    id := "id_example" // string | Tenant ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultTemplatesApi.ApiTenantsIdSeedDefaultTemplatesPost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultTemplatesApi.ApiTenantsIdSeedDefaultTemplatesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTenantsIdSeedDefaultTemplatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

