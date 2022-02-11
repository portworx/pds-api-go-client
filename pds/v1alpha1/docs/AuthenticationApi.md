# \AuthenticationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthAuthorizationCodePost**](AuthenticationApi.md#AuthAuthorizationCodePost) | **Post** /auth/authorization-code | Authorization code exchange
[**AuthOidcInfoGet**](AuthenticationApi.md#AuthOidcInfoGet) | **Get** /auth/oidc-info | OIDC info
[**AuthRefreshTokenPost**](AuthenticationApi.md#AuthRefreshTokenPost) | **Post** /auth/refresh-token | Use refresh token to generate new tokens.



## AuthAuthorizationCodePost

> ControllersOIDCTokenResponse AuthAuthorizationCodePost(ctx).Body(body).Execute()

Authorization code exchange



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
    body := *openapiclient.NewControllersAuthorizationCodeRequest() // ControllersAuthorizationCodeRequest | Request body containing the received authorization code.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.AuthAuthorizationCodePost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.AuthAuthorizationCodePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthAuthorizationCodePost`: ControllersOIDCTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.AuthAuthorizationCodePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthAuthorizationCodePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ControllersAuthorizationCodeRequest**](ControllersAuthorizationCodeRequest.md) | Request body containing the received authorization code. | 

### Return type

[**ControllersOIDCTokenResponse**](ControllersOIDCTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthOidcInfoGet

> ControllersOIDCInfoResponse AuthOidcInfoGet(ctx).Execute()

OIDC info



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.AuthOidcInfoGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.AuthOidcInfoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthOidcInfoGet`: ControllersOIDCInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.AuthOidcInfoGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthOidcInfoGetRequest struct via the builder pattern


### Return type

[**ControllersOIDCInfoResponse**](ControllersOIDCInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthRefreshTokenPost

> ControllersOIDCTokenResponse AuthRefreshTokenPost(ctx).Body(body).Execute()

Use refresh token to generate new tokens.

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
    body := *openapiclient.NewControllersRefreshTokenRequest() // ControllersRefreshTokenRequest | Request body containing the refresh token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.AuthRefreshTokenPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.AuthRefreshTokenPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthRefreshTokenPost`: ControllersOIDCTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.AuthRefreshTokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthRefreshTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ControllersRefreshTokenRequest**](ControllersRefreshTokenRequest.md) | Request body containing the refresh token. | 

### Return type

[**ControllersOIDCTokenResponse**](ControllersOIDCTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

