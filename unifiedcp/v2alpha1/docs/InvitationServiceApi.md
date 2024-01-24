# \InvitationServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvitationServiceAcceptInvitation**](InvitationServiceApi.md#InvitationServiceAcceptInvitation) | **Post** /v1/invite:accept | AcceptInvitation API accepts the invitation from the system
[**InvitationServiceCreateInvitation**](InvitationServiceApi.md#InvitationServiceCreateInvitation) | **Post** /v1/invite | CreateInvitation API creates an invitation
[**InvitationServiceDeleteInvitation**](InvitationServiceApi.md#InvitationServiceDeleteInvitation) | **Delete** /v1/invite/{uid} | DeleteInvitation deletes the specified invitation from the system
[**InvitationServiceGetInvitation**](InvitationServiceApi.md#InvitationServiceGetInvitation) | **Get** /v1/invite/{uid} | GetInvitation deletes the specified invitation from the system
[**InvitationServiceListInvitations**](InvitationServiceApi.md#InvitationServiceListInvitations) | **Get** /v1/invite | ListInvitation API lists invitation in an account/tenant/project



## InvitationServiceAcceptInvitation

> V1AcceptInvitationResponse InvitationServiceAcceptInvitation(ctx).V1AcceptInvitationRequest(v1AcceptInvitationRequest).Execute()

AcceptInvitation API accepts the invitation from the system

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
    v1AcceptInvitationRequest := *openapiclient.NewV1AcceptInvitationRequest("AccountId_example") // V1AcceptInvitationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitationServiceApi.InvitationServiceAcceptInvitation(context.Background()).V1AcceptInvitationRequest(v1AcceptInvitationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationServiceApi.InvitationServiceAcceptInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvitationServiceAcceptInvitation`: V1AcceptInvitationResponse
    fmt.Fprintf(os.Stdout, "Response from `InvitationServiceApi.InvitationServiceAcceptInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvitationServiceAcceptInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1AcceptInvitationRequest** | [**V1AcceptInvitationRequest**](V1AcceptInvitationRequest.md) |  | 

### Return type

[**V1AcceptInvitationResponse**](V1AcceptInvitationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationServiceCreateInvitation

> V1Invitation InvitationServiceCreateInvitation(ctx).V1CreateInvitationRequest(v1CreateInvitationRequest).Execute()

CreateInvitation API creates an invitation

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
    v1CreateInvitationRequest := *openapiclient.NewV1CreateInvitationRequest() // V1CreateInvitationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitationServiceApi.InvitationServiceCreateInvitation(context.Background()).V1CreateInvitationRequest(v1CreateInvitationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationServiceApi.InvitationServiceCreateInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvitationServiceCreateInvitation`: V1Invitation
    fmt.Fprintf(os.Stdout, "Response from `InvitationServiceApi.InvitationServiceCreateInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvitationServiceCreateInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1CreateInvitationRequest** | [**V1CreateInvitationRequest**](V1CreateInvitationRequest.md) |  | 

### Return type

[**V1Invitation**](V1Invitation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationServiceDeleteInvitation

> map[string]interface{} InvitationServiceDeleteInvitation(ctx, uid).Execute()

DeleteInvitation deletes the specified invitation from the system

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
    uid := "uid_example" // string | UID is the unique ID of the invitation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitationServiceApi.InvitationServiceDeleteInvitation(context.Background(), uid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationServiceApi.InvitationServiceDeleteInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvitationServiceDeleteInvitation`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InvitationServiceApi.InvitationServiceDeleteInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID is the unique ID of the invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvitationServiceDeleteInvitationRequest struct via the builder pattern


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


## InvitationServiceGetInvitation

> V1Invitation InvitationServiceGetInvitation(ctx, uid).Execute()

GetInvitation deletes the specified invitation from the system

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
    uid := "uid_example" // string | UID of the invitation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitationServiceApi.InvitationServiceGetInvitation(context.Background(), uid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationServiceApi.InvitationServiceGetInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvitationServiceGetInvitation`: V1Invitation
    fmt.Fprintf(os.Stdout, "Response from `InvitationServiceApi.InvitationServiceGetInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the invitation | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvitationServiceGetInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1Invitation**](V1Invitation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationServiceListInvitations

> V1ListInvitationsResponse InvitationServiceListInvitations(ctx).AccountId(accountId).TenantId(tenantId).ProjectId(projectId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()

ListInvitation API lists invitation in an account/tenant/project

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
    accountId := "accountId_example" // string | Account UID to which the user has been invited. User can be invited to multiple accounts (optional)
    tenantId := "tenantId_example" // string | Tenant UID to which the user has been invited. User can be invited to multiple tenants (optional)
    projectId := "projectId_example" // string | Project UID to which the user has been invited. User can be invited to multiple projects (optional)
    paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
    paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitationServiceApi.InvitationServiceListInvitations(context.Background()).AccountId(accountId).TenantId(tenantId).ProjectId(projectId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationServiceApi.InvitationServiceListInvitations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvitationServiceListInvitations`: V1ListInvitationsResponse
    fmt.Fprintf(os.Stdout, "Response from `InvitationServiceApi.InvitationServiceListInvitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvitationServiceListInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Account UID to which the user has been invited. User can be invited to multiple accounts | 
 **tenantId** | **string** | Tenant UID to which the user has been invited. User can be invited to multiple tenants | 
 **projectId** | **string** | Project UID to which the user has been invited. User can be invited to multiple projects | 
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 

### Return type

[**V1ListInvitationsResponse**](V1ListInvitationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

