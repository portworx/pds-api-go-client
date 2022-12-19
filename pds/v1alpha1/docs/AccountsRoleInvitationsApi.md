# \AccountsRoleInvitationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountRoleInvitationsIdDelete**](AccountsRoleInvitationsApi.md#ApiAccountRoleInvitationsIdDelete) | **Delete** /api/account-role-invitations/{id} | Delete AccountRoleInvitation
[**ApiAccountRoleInvitationsIdPatch**](AccountsRoleInvitationsApi.md#ApiAccountRoleInvitationsIdPatch) | **Patch** /api/account-role-invitations/{id} | Patch AccountRoleInvitation role_name



## ApiAccountRoleInvitationsIdDelete

> ApiAccountRoleInvitationsIdDelete(ctx, id).Execute()

Delete AccountRoleInvitation



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
    id := "id_example" // string | AccountRoleInvitation ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsRoleInvitationsApi.ApiAccountRoleInvitationsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsRoleInvitationsApi.ApiAccountRoleInvitationsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | AccountRoleInvitation ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountRoleInvitationsIdDeleteRequest struct via the builder pattern


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


## ApiAccountRoleInvitationsIdPatch

> ApiAccountRoleInvitationsIdPatch(ctx, id).Body(body).Execute()

Patch AccountRoleInvitation role_name



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
    id := "id_example" // string | AccountRoleInvitation ID (must be valid UUID)
    body := *openapiclient.NewRequestsPatchAccountRoleInvitationRequest() // RequestsPatchAccountRoleInvitationRequest | Object with patched role name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsRoleInvitationsApi.ApiAccountRoleInvitationsIdPatch(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsRoleInvitationsApi.ApiAccountRoleInvitationsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | AccountRoleInvitation ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountRoleInvitationsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestsPatchAccountRoleInvitationRequest**](RequestsPatchAccountRoleInvitationRequest.md) | Object with patched role name | 

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

