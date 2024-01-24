# \TargetClusterRegistrationManifestServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifest**](TargetClusterRegistrationManifestServiceApi.md#TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifest) | **Post** /v1/tenants/{tenantId}:registrationManifests | GetTargetClusterRegistrationManifest fetches registration manifest for the given request



## TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifest

> V1TargetClusterRegistrationManifest TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifest(ctx, tenantId).TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody(targetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody).Execute()

GetTargetClusterRegistrationManifest fetches registration manifest for the given request

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
    tenantId := "tenantId_example" // string | tenanat_id is the id of the tenant for which manifest is requested
    targetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody := *openapiclient.NewTargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody() // TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TargetClusterRegistrationManifestServiceApi.TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifest(context.Background(), tenantId).TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody(targetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetClusterRegistrationManifestServiceApi.TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifest`: V1TargetClusterRegistrationManifest
    fmt.Fprintf(os.Stdout, "Response from `TargetClusterRegistrationManifestServiceApi.TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | tenanat_id is the id of the tenant for which manifest is requested | 

### Other Parameters

Other parameters are passed through a pointer to a apiTargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody** | [**TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody**](TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody.md) |  | 

### Return type

[**V1TargetClusterRegistrationManifest**](V1TargetClusterRegistrationManifest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

