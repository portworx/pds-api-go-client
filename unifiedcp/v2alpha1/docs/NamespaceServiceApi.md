# \NamespaceServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NamespaceServiceListNamespaces**](NamespaceServiceApi.md#NamespaceServiceListNamespaces) | **Get** /v1/clusters/{clusterId}/namespaces | ListNamespaces API will return all the namespaces on the control plane



## NamespaceServiceListNamespaces

> V1ListNamespacesResponse NamespaceServiceListNamespaces(ctx, clusterId).TenantId(tenantId).ProjectId(projectId).Label(label).Execute()

ListNamespaces API will return all the namespaces on the control plane

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
    clusterId := "clusterId_example" // string | cluster ID for which namespaces need to be fetched
    tenantId := "tenantId_example" // string | tenant ID for which namespace need to be fetched (optional)
    projectId := "projectId_example" // string | project ID for which namespaces need to be fetched (optional)
    label := "label_example" // string | for label based filtering of the namespaces (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceServiceApi.NamespaceServiceListNamespaces(context.Background(), clusterId).TenantId(tenantId).ProjectId(projectId).Label(label).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceServiceApi.NamespaceServiceListNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamespaceServiceListNamespaces`: V1ListNamespacesResponse
    fmt.Fprintf(os.Stdout, "Response from `NamespaceServiceApi.NamespaceServiceListNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster ID for which namespaces need to be fetched | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamespaceServiceListNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | tenant ID for which namespace need to be fetched | 
 **projectId** | **string** | project ID for which namespaces need to be fetched | 
 **label** | **string** | for label based filtering of the namespaces | 

### Return type

[**V1ListNamespacesResponse**](V1ListNamespacesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

