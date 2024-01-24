# \DeploymentServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeploymentServiceCreateDeployment**](DeploymentServiceApi.md#DeploymentServiceCreateDeployment) | **Post** /pds/v1/namespace/{namespaceId}/deployment | CreateDeployment API creates the Deployment resource.
[**DeploymentServiceDeleteDeployment**](DeploymentServiceApi.md#DeploymentServiceDeleteDeployment) | **Delete** /pds/v1/deployments/{id} | DeleteDeployment API deletes the Deployment resource.
[**DeploymentServiceGetDeployment**](DeploymentServiceApi.md#DeploymentServiceGetDeployment) | **Get** /pds/v1/deployments/{id} | GetDeployment API returns the Deployment resource.
[**DeploymentServiceGetDeploymentCredentials**](DeploymentServiceApi.md#DeploymentServiceGetDeploymentCredentials) | **Get** /pds/v1/deployments/{id}:credentials | GetDeploymentCredentials API returns the Credentials to be used to access the Deployment .
[**DeploymentServiceListDeployments**](DeploymentServiceApi.md#DeploymentServiceListDeployments) | **Get** /pds/v1/deployments | ListDeployments API lists the Deployment resources.
[**DeploymentServiceUpdateDeployment**](DeploymentServiceApi.md#DeploymentServiceUpdateDeployment) | **Put** /pds/v1/deployments | UpdateDeployment API updates the Deployment resource.



## DeploymentServiceCreateDeployment

> V1Deployment DeploymentServiceCreateDeployment(ctx, namespaceId).V1Deployment(v1Deployment).Execute()

CreateDeployment API creates the Deployment resource.

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
    namespaceId := "namespaceId_example" // string | UID of the namespace resource where this deployment will be created. (-- api-linter: core::0133::request-unknown-fields=disabled     aip.dev/not-precedent: We need this field for to support creation of     the resource in the namespace context. --)
    v1Deployment := *openapiclient.NewV1Deployment() // V1Deployment | Deployment resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentServiceApi.DeploymentServiceCreateDeployment(context.Background(), namespaceId).V1Deployment(v1Deployment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentServiceApi.DeploymentServiceCreateDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeploymentServiceCreateDeployment`: V1Deployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentServiceApi.DeploymentServiceCreateDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceId** | **string** | UID of the namespace resource where this deployment will be created. (-- api-linter: core::0133::request-unknown-fields&#x3D;disabled     aip.dev/not-precedent: We need this field for to support creation of     the resource in the namespace context. --) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentServiceCreateDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1Deployment** | [**V1Deployment**](V1Deployment.md) | Deployment resource. | 

### Return type

[**V1Deployment**](V1Deployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeploymentServiceDeleteDeployment

> map[string]interface{} DeploymentServiceDeleteDeployment(ctx, id).Execute()

DeleteDeployment API deletes the Deployment resource.

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
    id := "id_example" // string | UID of the Deployment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentServiceApi.DeploymentServiceDeleteDeployment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentServiceApi.DeploymentServiceDeleteDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeploymentServiceDeleteDeployment`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeploymentServiceApi.DeploymentServiceDeleteDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UID of the Deployment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentServiceDeleteDeploymentRequest struct via the builder pattern


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


## DeploymentServiceGetDeployment

> V1Deployment DeploymentServiceGetDeployment(ctx, id).Execute()

GetDeployment API returns the Deployment resource.

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
    id := "id_example" // string | UID of the Deployment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentServiceApi.DeploymentServiceGetDeployment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentServiceApi.DeploymentServiceGetDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeploymentServiceGetDeployment`: V1Deployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentServiceApi.DeploymentServiceGetDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UID of the Deployment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentServiceGetDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1Deployment**](V1Deployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeploymentServiceGetDeploymentCredentials

> V1DeploymentCredentials DeploymentServiceGetDeploymentCredentials(ctx, id).Execute()

GetDeploymentCredentials API returns the Credentials to be used to access the Deployment .

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
    id := "id_example" // string | UID of the Deployment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentServiceApi.DeploymentServiceGetDeploymentCredentials(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentServiceApi.DeploymentServiceGetDeploymentCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeploymentServiceGetDeploymentCredentials`: V1DeploymentCredentials
    fmt.Fprintf(os.Stdout, "Response from `DeploymentServiceApi.DeploymentServiceGetDeploymentCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UID of the Deployment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentServiceGetDeploymentCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1DeploymentCredentials**](V1DeploymentCredentials.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeploymentServiceListDeployments

> V1ListDeploymentsResponse DeploymentServiceListDeployments(ctx).AccountId(accountId).TenantId(tenantId).ClusterId(clusterId).NamespaceId(namespaceId).ProjectId(projectId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).Execute()

ListDeployments API lists the Deployment resources.

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
    accountId := "accountId_example" // string | UID of the account. (optional)
    tenantId := "tenantId_example" // string | UID of the tenant. (optional)
    clusterId := "clusterId_example" // string | UID of the target cluster. (optional)
    namespaceId := "namespaceId_example" // string | UID of the namespace. (optional)
    projectId := "projectId_example" // string | UID of the project. (optional)
    paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
    paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)
    sortSortBy := "sortSortBy_example" // string | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. (optional) (default to "FIELD_UNSPECIFIED")
    sortSortOrder := "sortSortOrder_example" // string | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. (optional) (default to "VALUE_UNSPECIFIED")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentServiceApi.DeploymentServiceListDeployments(context.Background()).AccountId(accountId).TenantId(tenantId).ClusterId(clusterId).NamespaceId(namespaceId).ProjectId(projectId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentServiceApi.DeploymentServiceListDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeploymentServiceListDeployments`: V1ListDeploymentsResponse
    fmt.Fprintf(os.Stdout, "Response from `DeploymentServiceApi.DeploymentServiceListDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentServiceListDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | UID of the account. | 
 **tenantId** | **string** | UID of the tenant. | 
 **clusterId** | **string** | UID of the target cluster. | 
 **namespaceId** | **string** | UID of the namespace. | 
 **projectId** | **string** | UID of the project. | 
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 
 **sortSortBy** | **string** | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. | [default to &quot;FIELD_UNSPECIFIED&quot;]
 **sortSortOrder** | **string** | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. | [default to &quot;VALUE_UNSPECIFIED&quot;]

### Return type

[**V1ListDeploymentsResponse**](V1ListDeploymentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeploymentServiceUpdateDeployment

> V1Deployment DeploymentServiceUpdateDeployment(ctx).V1Deployment(v1Deployment).Execute()

UpdateDeployment API updates the Deployment resource.

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
    v1Deployment := *openapiclient.NewV1Deployment() // V1Deployment | Deployment resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentServiceApi.DeploymentServiceUpdateDeployment(context.Background()).V1Deployment(v1Deployment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentServiceApi.DeploymentServiceUpdateDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeploymentServiceUpdateDeployment`: V1Deployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentServiceApi.DeploymentServiceUpdateDeployment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentServiceUpdateDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1Deployment** | [**V1Deployment**](V1Deployment.md) | Deployment resource. | 

### Return type

[**V1Deployment**](V1Deployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

