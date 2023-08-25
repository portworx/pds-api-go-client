# \EventsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDeploymentTargetsIdDeploymentEventsPost**](EventsApi.md#ApiDeploymentTargetsIdDeploymentEventsPost) | **Post** /api/deployment-targets/{id}/deployment/events | Adds k8s generated events for a deployment
[**ApiDeploymentTargetsIdEventsLastSyncedTimeGet**](EventsApi.md#ApiDeploymentTargetsIdEventsLastSyncedTimeGet) | **Get** /api/deployment-targets/{id}/events/last-synced-time | Get the timestamp of the latest event
[**ApiDeploymentsIdEventsGet**](EventsApi.md#ApiDeploymentsIdEventsGet) | **Get** /api/deployments/{id}/events | Get Deployment Events



## ApiDeploymentTargetsIdDeploymentEventsPost

> ApiDeploymentTargetsIdDeploymentEventsPost(ctx, id).Body(body).Execute()

Adds k8s generated events for a deployment



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
    body := *openapiclient.NewRequestsCreateDeploymentEventsRequest([]openapiclient.RequestsDeploymentEvent{*openapiclient.NewRequestsDeploymentEvent("DeploymentId_example", []openapiclient.RequestsK8sEvent{*openapiclient.NewRequestsK8sEvent(int64(123), "0/3 nodes are available", "cas-events-0h35cp-0.16f4c1e217362a3e", "FailedScheduling", "ResourceKind_example", "ResourceName_example", int64(123), "Type_example")}, "Namespace_example")}) // RequestsCreateDeploymentEventsRequest | Request body containing list of events

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.ApiDeploymentTargetsIdDeploymentEventsPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ApiDeploymentTargetsIdDeploymentEventsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DeploymentTarget ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentTargetsIdDeploymentEventsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestsCreateDeploymentEventsRequest**](RequestsCreateDeploymentEventsRequest.md) | Request body containing list of events | 

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


## ApiDeploymentsIdEventsGet

> []ModelsDeploymentTargetDeploymentEvent ApiDeploymentsIdEventsGet(ctx, id).Execute()

Get Deployment Events



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
    id := "id_example" // string | Deployment ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.ApiDeploymentsIdEventsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ApiDeploymentsIdEventsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeploymentsIdEventsGet`: []ModelsDeploymentTargetDeploymentEvent
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ApiDeploymentsIdEventsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Deployment ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentsIdEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ModelsDeploymentTargetDeploymentEvent**](ModelsDeploymentTargetDeploymentEvent.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

