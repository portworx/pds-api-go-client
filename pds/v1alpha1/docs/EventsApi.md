# \EventsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDeploymentTargetsIdEventsLastSyncedTimeGet**](EventsApi.md#ApiDeploymentTargetsIdEventsLastSyncedTimeGet) | **Get** /api/deployment-targets/{id}/events/last-synced-time | Get the timestamp of the latest event



## ApiDeploymentTargetsIdEventsLastSyncedTimeGet

> ModelsDeploymentTargetLastSyncedEvent ApiDeploymentTargetsIdEventsLastSyncedTimeGet(ctx, id).Execute()

Get the timestamp of the latest event



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
    id := "id_example" // string | DeploymentTarget ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.ApiDeploymentTargetsIdEventsLastSyncedTimeGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ApiDeploymentTargetsIdEventsLastSyncedTimeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentTargetsIdEventsLastSyncedTimeGet`: ModelsDeploymentTargetLastSyncedEvent
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ApiDeploymentTargetsIdEventsLastSyncedTimeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DeploymentTarget ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentTargetsIdEventsLastSyncedTimeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsDeploymentTargetLastSyncedEvent**](ModelsDeploymentTargetLastSyncedEvent.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

