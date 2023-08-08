# \IAMApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountsIdIamActorIdGet**](IAMApi.md#ApiAccountsIdIamActorIdGet) | **Get** /api/accounts/{id}/iam/{actor-id} | Get IAM
[**ApiAccountsIdIamGet**](IAMApi.md#ApiAccountsIdIamGet) | **Get** /api/accounts/{id}/iam | List IAM



## ApiAccountsIdIamActorIdGet

> ModelsIAM ApiAccountsIdIamActorIdGet(ctx, id, actorId).Execute()

Get IAM



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
    actorId := "actorId_example" // string | Actor ID (must be valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAMApi.ApiAccountsIdIamActorIdGet(context.Background(), id, actorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMApi.ApiAccountsIdIamActorIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsIdIamActorIdGet`: ModelsIAM
    fmt.Fprintf(os.Stdout, "Response from `IAMApi.ApiAccountsIdIamActorIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID (must be valid UUID) | 
**actorId** | **string** | Actor ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsIdIamActorIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelsIAM**](ModelsIAM.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsIdIamGet

> []ModelsIAM ApiAccountsIdIamGet(ctx, id).SortBy(sortBy).ActorId(actorId).ActorType(actorType).Execute()

List IAM



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
    sortBy := "sortBy_example" // string | A given IAM attribute to sort results by (one of: account_id, actor_id) (optional)
    actorId := "actorId_example" // string | Filter results by IAM actor id (optional)
    actorType := "actorType_example" // string | Filter results by IAM actor type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAMApi.ApiAccountsIdIamGet(context.Background(), id).SortBy(sortBy).ActorId(actorId).ActorType(actorType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMApi.ApiAccountsIdIamGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsIdIamGet`: []ModelsIAM
    fmt.Fprintf(os.Stdout, "Response from `IAMApi.ApiAccountsIdIamGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID (must be valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsIdIamGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | A given IAM attribute to sort results by (one of: account_id, actor_id) | 
 **actorId** | **string** | Filter results by IAM actor id | 
 **actorType** | **string** | Filter results by IAM actor type | 

### Return type

[**[]ModelsIAM**](ModelsIAM.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

