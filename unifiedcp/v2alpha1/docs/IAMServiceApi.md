# \IAMServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IAMServiceCreateIAM**](IAMServiceApi.md#IAMServiceCreateIAM) | **Post** /v1/iam | CreateIAM API creates a new IAM rolebinding
[**IAMServiceDeleteIAM**](IAMServiceApi.md#IAMServiceDeleteIAM) | **Delete** /v1/iam/{actorId} | DeleteIAM API delete IAM, currently required only for name
[**IAMServiceGetIAM**](IAMServiceApi.md#IAMServiceGetIAM) | **Get** /v1/iam/{actorId} | GetIAM API returns the info about IAM for given IAM id
[**IAMServiceGrantIAM**](IAMServiceApi.md#IAMServiceGrantIAM) | **Post** /v1/iam/{iam.config.actorId}:grant | GrantIAM API creates new IAM rolebinding at tenant, project and account level
[**IAMServiceListIAM**](IAMServiceApi.md#IAMServiceListIAM) | **Get** /v1/iam | ListIAM API lists the role bindings
[**IAMServiceRevokeIAM**](IAMServiceApi.md#IAMServiceRevokeIAM) | **Post** /v1/iam/{iam.config.actorId}:revoke | RevokeIAM API delete IAM rolebinding at tenant, project and account level
[**IAMServiceUpdateIAM**](IAMServiceApi.md#IAMServiceUpdateIAM) | **Put** /v1/iam/{iam.meta.uid} | UpdateIAM API updates IAM with the new set of role bindings. The request replaces the existing set of bindings.



## IAMServiceCreateIAM

> V1IAM IAMServiceCreateIAM(ctx).V1IAM(v1IAM).Execute()

CreateIAM API creates a new IAM rolebinding

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
    v1IAM := *openapiclient.NewV1IAM() // V1IAM | IAM to be created

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAMServiceApi.IAMServiceCreateIAM(context.Background()).V1IAM(v1IAM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMServiceApi.IAMServiceCreateIAM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IAMServiceCreateIAM`: V1IAM
    fmt.Fprintf(os.Stdout, "Response from `IAMServiceApi.IAMServiceCreateIAM`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIAMServiceCreateIAMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1IAM** | [**V1IAM**](V1IAM.md) | IAM to be created | 

### Return type

[**V1IAM**](V1IAM.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IAMServiceDeleteIAM

> map[string]interface{} IAMServiceDeleteIAM(ctx, actorId).Execute()

DeleteIAM API delete IAM, currently required only for name

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
    actorId := "actorId_example" // string | Actor ID for which the IAM should be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAMServiceApi.IAMServiceDeleteIAM(context.Background(), actorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMServiceApi.IAMServiceDeleteIAM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IAMServiceDeleteIAM`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IAMServiceApi.IAMServiceDeleteIAM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID for which the IAM should be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIAMServiceDeleteIAMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IAMServiceGetIAM

> V1IAM IAMServiceGetIAM(ctx, actorId).Execute()

GetIAM API returns the info about IAM for given IAM id

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
    actorId := "actorId_example" // string | Actor ID for which the details need to be fetched

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAMServiceApi.IAMServiceGetIAM(context.Background(), actorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMServiceApi.IAMServiceGetIAM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IAMServiceGetIAM`: V1IAM
    fmt.Fprintf(os.Stdout, "Response from `IAMServiceApi.IAMServiceGetIAM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID for which the details need to be fetched | 

### Other Parameters

Other parameters are passed through a pointer to a apiIAMServiceGetIAMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1IAM**](V1IAM.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IAMServiceGrantIAM

> V1GrantIAMResponse IAMServiceGrantIAM(ctx, iamConfigActorId).IAMServiceGrantIAMBody(iAMServiceGrantIAMBody).Execute()

GrantIAM API creates new IAM rolebinding at tenant, project and account level

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
    iamConfigActorId := "iamConfigActorId_example" // string | Actor ID for the associated actor
    iAMServiceGrantIAMBody := *openapiclient.NewIAMServiceGrantIAMBody() // IAMServiceGrantIAMBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAMServiceApi.IAMServiceGrantIAM(context.Background(), iamConfigActorId).IAMServiceGrantIAMBody(iAMServiceGrantIAMBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMServiceApi.IAMServiceGrantIAM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IAMServiceGrantIAM`: V1GrantIAMResponse
    fmt.Fprintf(os.Stdout, "Response from `IAMServiceApi.IAMServiceGrantIAM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iamConfigActorId** | **string** | Actor ID for the associated actor | 

### Other Parameters

Other parameters are passed through a pointer to a apiIAMServiceGrantIAMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iAMServiceGrantIAMBody** | [**IAMServiceGrantIAMBody**](IAMServiceGrantIAMBody.md) |  | 

### Return type

[**V1GrantIAMResponse**](V1GrantIAMResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IAMServiceListIAM

> V1ListIAMResponse IAMServiceListIAM(ctx).ActorId(actorId).AccountId(accountId).TenantId(tenantId).ProjectId(projectId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()

ListIAM API lists the role bindings

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
    actorId := "actorId_example" // string | Actor ID for which the IAM should be listed (optional)
    accountId := "accountId_example" // string | Account UID for which the IAM should be listed (optional)
    tenantId := "tenantId_example" // string | Tenant UID for which the IAM should be listed (optional)
    projectId := "projectId_example" // string | Project UID for which the IAM should be listed (optional)
    paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
    paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAMServiceApi.IAMServiceListIAM(context.Background()).ActorId(actorId).AccountId(accountId).TenantId(tenantId).ProjectId(projectId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMServiceApi.IAMServiceListIAM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IAMServiceListIAM`: V1ListIAMResponse
    fmt.Fprintf(os.Stdout, "Response from `IAMServiceApi.IAMServiceListIAM`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIAMServiceListIAMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actorId** | **string** | Actor ID for which the IAM should be listed | 
 **accountId** | **string** | Account UID for which the IAM should be listed | 
 **tenantId** | **string** | Tenant UID for which the IAM should be listed | 
 **projectId** | **string** | Project UID for which the IAM should be listed | 
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 

### Return type

[**V1ListIAMResponse**](V1ListIAMResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IAMServiceRevokeIAM

> V1RevokeIAMResponse IAMServiceRevokeIAM(ctx, iamConfigActorId).IAMServiceRevokeIAMBody(iAMServiceRevokeIAMBody).Execute()

RevokeIAM API delete IAM rolebinding at tenant, project and account level

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
    iamConfigActorId := "iamConfigActorId_example" // string | Actor ID for the associated actor
    iAMServiceRevokeIAMBody := *openapiclient.NewIAMServiceRevokeIAMBody() // IAMServiceRevokeIAMBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAMServiceApi.IAMServiceRevokeIAM(context.Background(), iamConfigActorId).IAMServiceRevokeIAMBody(iAMServiceRevokeIAMBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMServiceApi.IAMServiceRevokeIAM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IAMServiceRevokeIAM`: V1RevokeIAMResponse
    fmt.Fprintf(os.Stdout, "Response from `IAMServiceApi.IAMServiceRevokeIAM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iamConfigActorId** | **string** | Actor ID for the associated actor | 

### Other Parameters

Other parameters are passed through a pointer to a apiIAMServiceRevokeIAMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iAMServiceRevokeIAMBody** | [**IAMServiceRevokeIAMBody**](IAMServiceRevokeIAMBody.md) |  | 

### Return type

[**V1RevokeIAMResponse**](V1RevokeIAMResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IAMServiceUpdateIAM

> V1IAM IAMServiceUpdateIAM(ctx, iamMetaUid).IAMToBeUpdated(iAMToBeUpdated).Execute()

UpdateIAM API updates IAM with the new set of role bindings. The request replaces the existing set of bindings.

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
    iamMetaUid := "iamMetaUid_example" // string | UID of the resource of the format <resource prefix>-<uuid>.
    iAMToBeUpdated := *openapiclient.NewIAMToBeUpdated() // IAMToBeUpdated | IAM to be updated

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAMServiceApi.IAMServiceUpdateIAM(context.Background(), iamMetaUid).IAMToBeUpdated(iAMToBeUpdated).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMServiceApi.IAMServiceUpdateIAM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IAMServiceUpdateIAM`: V1IAM
    fmt.Fprintf(os.Stdout, "Response from `IAMServiceApi.IAMServiceUpdateIAM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iamMetaUid** | **string** | UID of the resource of the format &lt;resource prefix&gt;-&lt;uuid&gt;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIAMServiceUpdateIAMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iAMToBeUpdated** | [**IAMToBeUpdated**](IAMToBeUpdated.md) | IAM to be updated | 

### Return type

[**V1IAM**](V1IAM.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

