# \SampleTemplatesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDefaultTemplatesApplicationConfigurationGet**](SampleTemplatesApi.md#ApiDefaultTemplatesApplicationConfigurationGet) | **Get** /api/default-templates/application-configuration | List ApplicationConfigurationSamples
[**ApiDefaultTemplatesBackupPoliciesGet**](SampleTemplatesApi.md#ApiDefaultTemplatesBackupPoliciesGet) | **Get** /api/default-templates/backup-policies | List BackupPolicySamples
[**ApiDefaultTemplatesResourceSettingsGet**](SampleTemplatesApi.md#ApiDefaultTemplatesResourceSettingsGet) | **Get** /api/default-templates/resource-settings | List ResourceSettingsSamples
[**ApiDefaultTemplatesStorageOptionsGet**](SampleTemplatesApi.md#ApiDefaultTemplatesStorageOptionsGet) | **Get** /api/default-templates/storage-options | List StorageOptionsSamples



## ApiDefaultTemplatesApplicationConfigurationGet

> ControllersApplicationConfigurationSamples ApiDefaultTemplatesApplicationConfigurationGet(ctx).DataServiceId(dataServiceId).Execute()

List ApplicationConfigurationSamples



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
    dataServiceId := "dataServiceId_example" // string | Filter results by DataService ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SampleTemplatesApi.ApiDefaultTemplatesApplicationConfigurationGet(context.Background()).DataServiceId(dataServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SampleTemplatesApi.ApiDefaultTemplatesApplicationConfigurationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDefaultTemplatesApplicationConfigurationGet`: ControllersApplicationConfigurationSamples
    fmt.Fprintf(os.Stdout, "Response from `SampleTemplatesApi.ApiDefaultTemplatesApplicationConfigurationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDefaultTemplatesApplicationConfigurationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataServiceId** | **string** | Filter results by DataService ID | 

### Return type

[**ControllersApplicationConfigurationSamples**](ControllersApplicationConfigurationSamples.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDefaultTemplatesBackupPoliciesGet

> ControllersBackupPolicySamples ApiDefaultTemplatesBackupPoliciesGet(ctx).Execute()

List BackupPolicySamples



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SampleTemplatesApi.ApiDefaultTemplatesBackupPoliciesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SampleTemplatesApi.ApiDefaultTemplatesBackupPoliciesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDefaultTemplatesBackupPoliciesGet`: ControllersBackupPolicySamples
    fmt.Fprintf(os.Stdout, "Response from `SampleTemplatesApi.ApiDefaultTemplatesBackupPoliciesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiDefaultTemplatesBackupPoliciesGetRequest struct via the builder pattern


### Return type

[**ControllersBackupPolicySamples**](ControllersBackupPolicySamples.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDefaultTemplatesResourceSettingsGet

> ControllersResourceSettingsSamples ApiDefaultTemplatesResourceSettingsGet(ctx).DataServiceId(dataServiceId).Execute()

List ResourceSettingsSamples



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
    dataServiceId := "dataServiceId_example" // string | Filter results by DataService ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SampleTemplatesApi.ApiDefaultTemplatesResourceSettingsGet(context.Background()).DataServiceId(dataServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SampleTemplatesApi.ApiDefaultTemplatesResourceSettingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDefaultTemplatesResourceSettingsGet`: ControllersResourceSettingsSamples
    fmt.Fprintf(os.Stdout, "Response from `SampleTemplatesApi.ApiDefaultTemplatesResourceSettingsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDefaultTemplatesResourceSettingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataServiceId** | **string** | Filter results by DataService ID | 

### Return type

[**ControllersResourceSettingsSamples**](ControllersResourceSettingsSamples.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDefaultTemplatesStorageOptionsGet

> ControllersStorageOptionsSamples ApiDefaultTemplatesStorageOptionsGet(ctx).Execute()

List StorageOptionsSamples



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SampleTemplatesApi.ApiDefaultTemplatesStorageOptionsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SampleTemplatesApi.ApiDefaultTemplatesStorageOptionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDefaultTemplatesStorageOptionsGet`: ControllersStorageOptionsSamples
    fmt.Fprintf(os.Stdout, "Response from `SampleTemplatesApi.ApiDefaultTemplatesStorageOptionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiDefaultTemplatesStorageOptionsGetRequest struct via the builder pattern


### Return type

[**ControllersStorageOptionsSamples**](ControllersStorageOptionsSamples.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

