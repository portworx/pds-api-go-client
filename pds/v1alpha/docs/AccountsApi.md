# \AccountsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountsGet**](AccountsApi.md#ApiAccountsGet) | **Get** /api/accounts | List Accounts
[**ApiAccountsIdGet**](AccountsApi.md#ApiAccountsIdGet) | **Get** /api/accounts/{id} | Get Account
[**ApiAccountsIdUsersGet**](AccountsApi.md#ApiAccountsIdUsersGet) | **Get** /api/accounts/{id}/users | List Account Users
[**ApiAccountsPost**](AccountsApi.md#ApiAccountsPost) | **Post** /api/accounts | Create Account



## ApiAccountsGet

> ControllersPaginatedAccounts ApiAccountsGet(ctx).SortBy(sortBy).Id(id).ActorId(actorId).Execute()

List Accounts



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
    sortBy := "sortBy_example" // string | A given Accounts attribute to sort results by (one of: id, name) (optional)
    id := "id_example" // string | Filter results by Accounts id (optional)
    actorId := "actorId_example" // string | Filter results by Accounts name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ApiAccountsGet(context.Background()).SortBy(sortBy).Id(id).ActorId(actorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ApiAccountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsGet`: ControllersPaginatedAccounts
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ApiAccountsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **string** | A given Accounts attribute to sort results by (one of: id, name) | 
 **id** | **string** | Filter results by Accounts id | 
 **actorId** | **string** | Filter results by Accounts name | 

### Return type

[**ControllersPaginatedAccounts**](ControllersPaginatedAccounts.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Account ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ApiAccountsIdGet(context.Background(), id).Execute()
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


## ApiAccountsIdUsersGet

> ControllersPaginatedUsers ApiAccountsIdUsersGet(ctx, id).SortBy(sortBy).Id2(id2).Email(email).Execute()

List Account Users



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
    id := TODO // interface{} | Account ID (must be valid UUID)
    sortBy := "sortBy_example" // string | A given User attribute to sort results by (one of: id, email) (optional)
    id2 := "id_example" // string | Filter results by User id (optional)
    email := "email_example" // string | Filter results by User email (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ApiAccountsIdUsersGet(context.Background(), id).SortBy(sortBy).Id2(id2).Email(email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ApiAccountsIdUsersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsIdUsersGet`: ControllersPaginatedUsers
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ApiAccountsIdUsersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | Account ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsIdUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given User attribute to sort results by (one of: id, email) | 
 **id2** | **string** | Filter results by User id | 
 **email** | **string** | Filter results by User email | 

### Return type

[**ControllersPaginatedUsers**](ControllersPaginatedUsers.md)

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
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewControllersCreateAccountRequest() // ControllersCreateAccountRequest | Request body containing name of the account

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ApiAccountsPost(context.Background()).Body(body).Execute()
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

