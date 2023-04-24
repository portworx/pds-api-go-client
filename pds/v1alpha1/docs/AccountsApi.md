# \AccountsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountsGet**](AccountsApi.md#ApiAccountsGet) | **Get** /api/accounts | List Accounts
[**ApiAccountsIdAccountRoleInvitationsGet**](AccountsApi.md#ApiAccountsIdAccountRoleInvitationsGet) | **Get** /api/accounts/{id}/account-role-invitations | List Account Role Invitations
[**ApiAccountsIdEulaPut**](AccountsApi.md#ApiAccountsIdEulaPut) | **Put** /api/accounts/{id}/eula | Accept EULA
[**ApiAccountsIdGet**](AccountsApi.md#ApiAccountsIdGet) | **Get** /api/accounts/{id} | Get Account
[**ApiAccountsIdGlobalConfigPut**](AccountsApi.md#ApiAccountsIdGlobalConfigPut) | **Put** /api/accounts/{id}/global-config | Update AccountGlobalConfig
[**ApiAccountsIdUsersGet**](AccountsApi.md#ApiAccountsIdUsersGet) | **Get** /api/accounts/{id}/users | List Account Users
[**ApiAccountsPost**](AccountsApi.md#ApiAccountsPost) | **Post** /api/accounts | Create Account



## ApiAccountsGet

> ModelsPaginatedResultModelsAccount ApiAccountsGet(ctx).SortBy(sortBy).Limit(limit).Continuation(continuation).Id(id).Name(name).Execute()

List Accounts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    sortBy := "sortBy_example" // string | A given Accounts attribute to sort results by (one of: id, name, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id := "id_example" // string | Filter results by Accounts id (optional)
    name := "name_example" // string | Filter results by Accounts name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.ApiAccountsGet(context.Background()).SortBy(sortBy).Limit(limit).Continuation(continuation).Id(id).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ApiAccountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsGet`: ModelsPaginatedResultModelsAccount
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ApiAccountsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **string** | A given Accounts attribute to sort results by (one of: id, name, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id** | **string** | Filter results by Accounts id | 
 **name** | **string** | Filter results by Accounts name | 

### Return type

[**ModelsPaginatedResultModelsAccount**](ModelsPaginatedResultModelsAccount.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsIdAccountRoleInvitationsGet

> ModelsPaginatedResultModelsAccountRoleInvitation ApiAccountsIdAccountRoleInvitationsGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Email(email).RoleName(roleName).Execute()

List Account Role Invitations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Account ID (must be valid UUID)
    sortBy := "sortBy_example" // string | A given User attribute to sort results by (one of: id, email, role_name, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by invitation id (optional)
    email := "email_example" // string | Filter results by User email (optional)
    roleName := "roleName_example" // string | Filter results by assigned role name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.ApiAccountsIdAccountRoleInvitationsGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Email(email).RoleName(roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ApiAccountsIdAccountRoleInvitationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsIdAccountRoleInvitationsGet`: ModelsPaginatedResultModelsAccountRoleInvitation
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ApiAccountsIdAccountRoleInvitationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsIdAccountRoleInvitationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given User attribute to sort results by (one of: id, email, role_name, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by invitation id | 
 **email** | **string** | Filter results by User email | 
 **roleName** | **string** | Filter results by assigned role name | 

### Return type

[**ModelsPaginatedResultModelsAccountRoleInvitation**](ModelsPaginatedResultModelsAccountRoleInvitation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsIdEulaPut

> ApiAccountsIdEulaPut(ctx, id).Body(body).Execute()

Accept EULA



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Account ID (must be valid UUID)
    body := *openapiclient.NewControllersAcceptEULARequest() // ControllersAcceptEULARequest | Request body containing the version of the EULA.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccountsApi.ApiAccountsIdEulaPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ApiAccountsIdEulaPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsIdEulaPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersAcceptEULARequest**](ControllersAcceptEULARequest.md) | Request body containing the version of the EULA. | 

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


## ApiAccountsIdGet

> ModelsAccount ApiAccountsIdGet(ctx, id).Execute()

Get Account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Account ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.ApiAccountsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ApiAccountsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsIdGet`: ModelsAccount
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ApiAccountsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsAccount**](ModelsAccount.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsIdGlobalConfigPut

> ApiAccountsIdGlobalConfigPut(ctx, id).Body(body).Execute()

Update AccountGlobalConfig



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Account ID (must be valid UUID)
    body := *openapiclient.NewControllersUpdateGlobalConfigRequest() // ControllersUpdateGlobalConfigRequest | Request body containing the global config values. Empty values are ignored.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccountsApi.ApiAccountsIdGlobalConfigPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ApiAccountsIdGlobalConfigPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsIdGlobalConfigPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ControllersUpdateGlobalConfigRequest**](ControllersUpdateGlobalConfigRequest.md) | Request body containing the global config values. Empty values are ignored. | 

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


## ApiAccountsIdUsersGet

> ModelsPaginatedResultModelsUser ApiAccountsIdUsersGet(ctx, id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Email(email).Execute()

List Account Users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Account ID (must be valid UUID)
    sortBy := "sortBy_example" // string | A given User attribute to sort results by (one of: id, email, created_at) (optional)
    limit := "limit_example" // string | Maximum number of rows to return (could be less) (optional)
    continuation := "continuation_example" // string | Use a token returned by a previous query to continue listing with the next batch of rows (optional)
    id2 := "id_example" // string | Filter results by User id (optional)
    email := "email_example" // string | Filter results by User email (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.ApiAccountsIdUsersGet(context.Background(), id).SortBy(sortBy).Limit(limit).Continuation(continuation).Id2(id2).Email(email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ApiAccountsIdUsersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsIdUsersGet`: ModelsPaginatedResultModelsUser
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ApiAccountsIdUsersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsIdUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given User attribute to sort results by (one of: id, email, created_at) | 
 **limit** | **string** | Maximum number of rows to return (could be less) | 
 **continuation** | **string** | Use a token returned by a previous query to continue listing with the next batch of rows | 
 **id2** | **string** | Filter results by User id | 
 **email** | **string** | Filter results by User email | 

### Return type

[**ModelsPaginatedResultModelsUser**](ModelsPaginatedResultModelsUser.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsPost

> ModelsAccount ApiAccountsPost(ctx).Body(body).Execute()

Create Account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewControllersCreateAccountRequest() // ControllersCreateAccountRequest | Request body containing name of the account

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.ApiAccountsPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ApiAccountsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsPost`: ModelsAccount
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ApiAccountsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ControllersCreateAccountRequest**](ControllersCreateAccountRequest.md) | Request body containing name of the account | 

### Return type

[**ModelsAccount**](ModelsAccount.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

