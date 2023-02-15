# \DeploymentManifestsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDeploymentsIdUpdateManifestPost**](DeploymentManifestsApi.md#ApiDeploymentsIdUpdateManifestPost) | **Post** /api/deployments/{id}/update-manifest | Post DeploymentManifest



## ApiDeploymentsIdUpdateManifestPost

> ApiDeploymentsIdUpdateManifestPost(ctx, id).Body(body).Execute()

Post DeploymentManifest



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
    body := *openapiclient.NewRequestsUpsertDeploymentManifestRequest() // RequestsUpsertDeploymentManifestRequest | Body with the Deployment Manifest

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentManifestsApi.ApiDeploymentsIdUpdateManifestPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentManifestsApi.ApiDeploymentsIdUpdateManifestPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Deployment ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeploymentsIdUpdateManifestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestsUpsertDeploymentManifestRequest**](RequestsUpsertDeploymentManifestRequest.md) | Body with the Deployment Manifest | 

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

