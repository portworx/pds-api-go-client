# \TargetClusterServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TargetClusterServiceDeleteTargetCluster**](TargetClusterServiceApi.md#TargetClusterServiceDeleteTargetCluster) | **Delete** /v1/clusters/{id} | DeleteTargetCluster API deletes the specified TargetCluster
[**TargetClusterServiceGetTargetCluster**](TargetClusterServiceApi.md#TargetClusterServiceGetTargetCluster) | **Get** /v1/clusters/{id} | GetTargetCluster API returns the info about the TargetCluster for given name
[**TargetClusterServiceListTargetClusters**](TargetClusterServiceApi.md#TargetClusterServiceListTargetClusters) | **Get** /v1/clusters | ListTargetCluster API lists the TargetClusters visible to the caller
[**TargetClusterServiceUpdateTargetCluster**](TargetClusterServiceApi.md#TargetClusterServiceUpdateTargetCluster) | **Patch** /v1/clusters/{targetCluster.meta.uid} | UpdateTargetCluster API updates the metadata(e.g name or labels) of the specified TargetCluster



## TargetClusterServiceDeleteTargetCluster

> map[string]interface{} TargetClusterServiceDeleteTargetCluster(ctx, id).Force(force).Execute()

DeleteTargetCluster API deletes the specified TargetCluster

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
    id := "id_example" // string | Unique identifier of the cluster whose details needs to be fetched
    force := true // bool | Flag to indicate force delete (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TargetClusterServiceApi.TargetClusterServiceDeleteTargetCluster(context.Background(), id).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetClusterServiceApi.TargetClusterServiceDeleteTargetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TargetClusterServiceDeleteTargetCluster`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TargetClusterServiceApi.TargetClusterServiceDeleteTargetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the cluster whose details needs to be fetched | 

### Other Parameters

Other parameters are passed through a pointer to a apiTargetClusterServiceDeleteTargetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Flag to indicate force delete | 

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


## TargetClusterServiceGetTargetCluster

> V1TargetCluster TargetClusterServiceGetTargetCluster(ctx, id).Execute()

GetTargetCluster API returns the info about the TargetCluster for given name

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
    id := "id_example" // string | Unique identifier of the cluster whose details needs to be fetched

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TargetClusterServiceApi.TargetClusterServiceGetTargetCluster(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetClusterServiceApi.TargetClusterServiceGetTargetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TargetClusterServiceGetTargetCluster`: V1TargetCluster
    fmt.Fprintf(os.Stdout, "Response from `TargetClusterServiceApi.TargetClusterServiceGetTargetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the cluster whose details needs to be fetched | 

### Other Parameters

Other parameters are passed through a pointer to a apiTargetClusterServiceGetTargetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1TargetCluster**](V1TargetCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TargetClusterServiceListTargetClusters

> V1ListTargetClustersResponse TargetClusterServiceListTargetClusters(ctx).AccountId(accountId).ProjectId(projectId).TenantId(tenantId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()

ListTargetCluster API lists the TargetClusters visible to the caller

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
    accountId := "accountId_example" // string | list target clusters associated with an account (optional)
    projectId := "projectId_example" // string | list target clusters associated with a project (optional)
    tenantId := "tenantId_example" // string | list target clusters associated with a tenant (optional)
    paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
    paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TargetClusterServiceApi.TargetClusterServiceListTargetClusters(context.Background()).AccountId(accountId).ProjectId(projectId).TenantId(tenantId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetClusterServiceApi.TargetClusterServiceListTargetClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TargetClusterServiceListTargetClusters`: V1ListTargetClustersResponse
    fmt.Fprintf(os.Stdout, "Response from `TargetClusterServiceApi.TargetClusterServiceListTargetClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTargetClusterServiceListTargetClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | list target clusters associated with an account | 
 **projectId** | **string** | list target clusters associated with a project | 
 **tenantId** | **string** | list target clusters associated with a tenant | 
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 

### Return type

[**V1ListTargetClustersResponse**](V1ListTargetClustersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TargetClusterServiceUpdateTargetCluster

> V1TargetCluster TargetClusterServiceUpdateTargetCluster(ctx, targetClusterMetaUid).TargetClusterDetails(targetClusterDetails).Execute()

UpdateTargetCluster API updates the metadata(e.g name or labels) of the specified TargetCluster

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
    targetClusterMetaUid := "targetClusterMetaUid_example" // string | UID of the resource of the format <resource prefix>-<uuid>.
    targetClusterDetails := *openapiclient.NewTargetClusterDetails() // TargetClusterDetails | Target cluster details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TargetClusterServiceApi.TargetClusterServiceUpdateTargetCluster(context.Background(), targetClusterMetaUid).TargetClusterDetails(targetClusterDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetClusterServiceApi.TargetClusterServiceUpdateTargetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TargetClusterServiceUpdateTargetCluster`: V1TargetCluster
    fmt.Fprintf(os.Stdout, "Response from `TargetClusterServiceApi.TargetClusterServiceUpdateTargetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetClusterMetaUid** | **string** | UID of the resource of the format &lt;resource prefix&gt;-&lt;uuid&gt;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTargetClusterServiceUpdateTargetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetClusterDetails** | [**TargetClusterDetails**](TargetClusterDetails.md) | Target cluster details | 

### Return type

[**V1TargetCluster**](V1TargetCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

