# \ApplicationServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationServiceListAvailableApplications**](ApplicationServiceApi.md#ApplicationServiceListAvailableApplications) | **Get** /v1/tenants/{tenantId}:listAvailableApplications | ListAvailableApplications API lists all the applications visible to a caller



## ApplicationServiceListAvailableApplications

> V1ListAvailableApplicationsResponse ApplicationServiceListAvailableApplications(ctx, tenantId).Execute()

ListAvailableApplications API lists all the applications visible to a caller

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
    tenantId := "tenantId_example" // string | tenant_id for which list of applications needs to be fetched

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceListAvailableApplications(context.Background(), tenantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceListAvailableApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceListAvailableApplications`: V1ListAvailableApplicationsResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceListAvailableApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | tenant_id for which list of applications needs to be fetched | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceListAvailableApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1ListAvailableApplicationsResponse**](V1ListAvailableApplicationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

