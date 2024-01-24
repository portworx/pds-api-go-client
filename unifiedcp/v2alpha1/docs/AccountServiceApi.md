# \AccountServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountServiceCreateAccount**](AccountServiceApi.md#AccountServiceCreateAccount) | **Post** /v1/accounts | CreateAccount API creates a new Account
[**AccountServiceDeleteAccount**](AccountServiceApi.md#AccountServiceDeleteAccount) | **Delete** /v1/accounts/{accountId} | DeleteAccount API deletes the specified account
[**AccountServiceGetAccount**](AccountServiceApi.md#AccountServiceGetAccount) | **Get** /v1/accounts/{accountId} | GetAccount API returns the info about account for given account id
[**AccountServiceListAccounts**](AccountServiceApi.md#AccountServiceListAccounts) | **Get** /v1/accounts | ListAccounts API lists the accounts visible to the caller
[**AccountServiceUpdateAccount**](AccountServiceApi.md#AccountServiceUpdateAccount) | **Put** /v1/accounts/{account.meta.uid} | UpdateAccount API updates the specified account.This API requires the caller to first get an Account from GetAccount().



## AccountServiceCreateAccount

> V1Account AccountServiceCreateAccount(ctx).V1Account(v1Account).Execute()

CreateAccount API creates a new Account

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
    v1Account := *openapiclient.NewV1Account() // V1Account | account to be created

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountServiceApi.AccountServiceCreateAccount(context.Background()).V1Account(v1Account).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceApi.AccountServiceCreateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountServiceCreateAccount`: V1Account
    fmt.Fprintf(os.Stdout, "Response from `AccountServiceApi.AccountServiceCreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountServiceCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1Account** | [**V1Account**](V1Account.md) | account to be created | 

### Return type

[**V1Account**](V1Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountServiceDeleteAccount

> map[string]interface{} AccountServiceDeleteAccount(ctx, accountId).Execute()

DeleteAccount API deletes the specified account

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
    accountId := "accountId_example" // string | account id for which account details need to be fetched

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountServiceApi.AccountServiceDeleteAccount(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceApi.AccountServiceDeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountServiceDeleteAccount`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountServiceApi.AccountServiceDeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id for which account details need to be fetched | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountServiceDeleteAccountRequest struct via the builder pattern


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


## AccountServiceGetAccount

> V1Account AccountServiceGetAccount(ctx, accountId).Execute()

GetAccount API returns the info about account for given account id

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
    accountId := "accountId_example" // string | account id for which account details need to be fetched

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountServiceApi.AccountServiceGetAccount(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceApi.AccountServiceGetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountServiceGetAccount`: V1Account
    fmt.Fprintf(os.Stdout, "Response from `AccountServiceApi.AccountServiceGetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id for which account details need to be fetched | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountServiceGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1Account**](V1Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountServiceListAccounts

> V1ListAccountsResponse AccountServiceListAccounts(ctx).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()

ListAccounts API lists the accounts visible to the caller

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
    paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
    paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountServiceApi.AccountServiceListAccounts(context.Background()).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceApi.AccountServiceListAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountServiceListAccounts`: V1ListAccountsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountServiceApi.AccountServiceListAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountServiceListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 

### Return type

[**V1ListAccountsResponse**](V1ListAccountsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountServiceUpdateAccount

> V1Account AccountServiceUpdateAccount(ctx, accountMetaUid).AccountForWhichNameNeedsToBeUpdated(accountForWhichNameNeedsToBeUpdated).Execute()

UpdateAccount API updates the specified account.This API requires the caller to first get an Account from GetAccount().

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
    accountMetaUid := "accountMetaUid_example" // string | UID of the resource of the format <resource prefix>-<uuid>.
    accountForWhichNameNeedsToBeUpdated := *openapiclient.NewAccountForWhichNameNeedsToBeUpdated() // AccountForWhichNameNeedsToBeUpdated | account  for which name needs to be updated

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountServiceApi.AccountServiceUpdateAccount(context.Background(), accountMetaUid).AccountForWhichNameNeedsToBeUpdated(accountForWhichNameNeedsToBeUpdated).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceApi.AccountServiceUpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountServiceUpdateAccount`: V1Account
    fmt.Fprintf(os.Stdout, "Response from `AccountServiceApi.AccountServiceUpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountMetaUid** | **string** | UID of the resource of the format &lt;resource prefix&gt;-&lt;uuid&gt;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountServiceUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountForWhichNameNeedsToBeUpdated** | [**AccountForWhichNameNeedsToBeUpdated**](AccountForWhichNameNeedsToBeUpdated.md) | account  for which name needs to be updated | 

### Return type

[**V1Account**](V1Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

