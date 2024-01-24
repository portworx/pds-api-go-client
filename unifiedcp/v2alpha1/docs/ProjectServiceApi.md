# \ProjectServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectServiceAssociateResources**](ProjectServiceApi.md#ProjectServiceAssociateResources) | **Put** /v1/projects/{projectId}:associate | AssociateResources will append the supplied resources to existing infra resources of a given project config. (-- api-linter: core::0136::http-method&#x3D;disabled     aip.dev/not-precedent: We need to do this because renaming creating issues with other apis)
[**ProjectServiceCreateProject**](ProjectServiceApi.md#ProjectServiceCreateProject) | **Post** /v1/tenants/{tenantId}/projects | Create project api creates a project.
[**ProjectServiceDeleteProject**](ProjectServiceApi.md#ProjectServiceDeleteProject) | **Delete** /v1/projects/{projectId} | Deletes a project and its associated resources.
[**ProjectServiceDisassociateResources**](ProjectServiceApi.md#ProjectServiceDisassociateResources) | **Put** /v1/projects/{projectId}:disassociate | DisassociateResource will remove the infra resources supplied in the request from the project config. (-- api-linter: core::0136::http-method&#x3D;disabled     aip.dev/not-precedent: We need to do this because renaming creating issues with other apis)
[**ProjectServiceGetProject**](ProjectServiceApi.md#ProjectServiceGetProject) | **Get** /v1/projects/{projectId} | Get project apis returns a requested project.
[**ProjectServiceListProjects**](ProjectServiceApi.md#ProjectServiceListProjects) | **Get** /v1/projects | ListProjects API lists the projects visible to the caller for the provided tenant.



## ProjectServiceAssociateResources

> V1Project ProjectServiceAssociateResources(ctx, projectId).ProjectServiceAssociateResourcesBody(projectServiceAssociateResourcesBody).Execute()

AssociateResources will append the supplied resources to existing infra resources of a given project config. (-- api-linter: core::0136::http-method=disabled     aip.dev/not-precedent: We need to do this because renaming creating issues with other apis)

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
    projectId := "projectId_example" // string | ID of the project.
    projectServiceAssociateResourcesBody := *openapiclient.NewProjectServiceAssociateResourcesBody() // ProjectServiceAssociateResourcesBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceAssociateResources(context.Background(), projectId).ProjectServiceAssociateResourcesBody(projectServiceAssociateResourcesBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceAssociateResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceAssociateResources`: V1Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceAssociateResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceAssociateResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectServiceAssociateResourcesBody** | [**ProjectServiceAssociateResourcesBody**](ProjectServiceAssociateResourcesBody.md) |  | 

### Return type

[**V1Project**](V1Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectServiceCreateProject

> V1Project ProjectServiceCreateProject(ctx, tenantId).ProjectServiceCreateProjectBody(projectServiceCreateProjectBody).Execute()

Create project api creates a project.

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
    tenantId := "tenantId_example" // string | The parent tenant under which project will be created (-- api-linter: core::0133::request-unknown-fields=disabled     aip.dev/not-precedent: We need this field for to support creation of     the resource in the tenant context. --)
    projectServiceCreateProjectBody := *openapiclient.NewProjectServiceCreateProjectBody() // ProjectServiceCreateProjectBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceCreateProject(context.Background(), tenantId).ProjectServiceCreateProjectBody(projectServiceCreateProjectBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceCreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceCreateProject`: V1Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceCreateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | The parent tenant under which project will be created (-- api-linter: core::0133::request-unknown-fields&#x3D;disabled     aip.dev/not-precedent: We need this field for to support creation of     the resource in the tenant context. --) | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectServiceCreateProjectBody** | [**ProjectServiceCreateProjectBody**](ProjectServiceCreateProjectBody.md) |  | 

### Return type

[**V1Project**](V1Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectServiceDeleteProject

> map[string]interface{} ProjectServiceDeleteProject(ctx, projectId).Execute()

Deletes a project and its associated resources.

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
    projectId := "projectId_example" // string | ID of the project which needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceDeleteProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceDeleteProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceDeleteProject`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceDeleteProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project which needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceDeleteProjectRequest struct via the builder pattern


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


## ProjectServiceDisassociateResources

> V1Project ProjectServiceDisassociateResources(ctx, projectId).ProjectServiceDisassociateResourcesBody(projectServiceDisassociateResourcesBody).Execute()

DisassociateResource will remove the infra resources supplied in the request from the project config. (-- api-linter: core::0136::http-method=disabled     aip.dev/not-precedent: We need to do this because renaming creating issues with other apis)

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
    projectId := "projectId_example" // string | ID of the project from which infra resources to be disassociated.
    projectServiceDisassociateResourcesBody := *openapiclient.NewProjectServiceDisassociateResourcesBody() // ProjectServiceDisassociateResourcesBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceDisassociateResources(context.Background(), projectId).ProjectServiceDisassociateResourcesBody(projectServiceDisassociateResourcesBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceDisassociateResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceDisassociateResources`: V1Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceDisassociateResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project from which infra resources to be disassociated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceDisassociateResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectServiceDisassociateResourcesBody** | [**ProjectServiceDisassociateResourcesBody**](ProjectServiceDisassociateResourcesBody.md) |  | 

### Return type

[**V1Project**](V1Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectServiceGetProject

> V1Project ProjectServiceGetProject(ctx, projectId).Execute()

Get project apis returns a requested project.

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
    projectId := "projectId_example" // string | Project id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceGetProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceGetProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceGetProject`: V1Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceGetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1Project**](V1Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectServiceListProjects

> V1ListProjectsResponse ProjectServiceListProjects(ctx).QueryAppResource(queryAppResource).QueryInfraResource(queryInfraResource).QueryResourceId(queryResourceId).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()

ListProjects API lists the projects visible to the caller for the provided tenant.

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
    queryAppResource := "queryAppResource_example" // string | Type of the application managed resource for which projects to be listed.   - TYPE_UNSPECIFIED: Unspecified, do not use.  - PDS_DEPLOYMENT: List of supported PDS application resources. PDS application resource of type deployment.  - PDS_BACKUP: PDS application resource of type backup.  - PDS_RESTORE: PDS application resource of type restore.  - BAAS_BACKUP: List of supported BAAS application resources. BAAS application resource of type backup. (optional) (default to "TYPE_UNSPECIFIED")
    queryInfraResource := "queryInfraResource_example" // string | Type of the infra resource for which projects to be listed.   - TYPE_UNSPECIFIED: Unspecified, do not use.  - ACCOUNT: Currently supported infra resources. Infra resource of type account.  - TENANT: Infra resource of type tenant.  - PROJECT: Infra resource of type projects.  - TARGET_CLUSTER: Infra resource target cluster.  - NAMESPACE: Infra resource of type namespace.  - CREDENTIAL: Infra resource of type credential.  - BACKUP_LOCATION: Infra resource of type backup location. (optional) (default to "TYPE_UNSPECIFIED")
    queryResourceId := "queryResourceId_example" // string | ID of the resource for which projects to be listed. (optional)
    sortSortBy := "sortSortBy_example" // string | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. (optional) (default to "FIELD_UNSPECIFIED")
    sortSortOrder := "sortSortOrder_example" // string | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. (optional) (default to "VALUE_UNSPECIFIED")
    paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
    paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceListProjects(context.Background()).QueryAppResource(queryAppResource).QueryInfraResource(queryInfraResource).QueryResourceId(queryResourceId).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceListProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceListProjects`: V1ListProjectsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceListProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceListProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryAppResource** | **string** | Type of the application managed resource for which projects to be listed.   - TYPE_UNSPECIFIED: Unspecified, do not use.  - PDS_DEPLOYMENT: List of supported PDS application resources. PDS application resource of type deployment.  - PDS_BACKUP: PDS application resource of type backup.  - PDS_RESTORE: PDS application resource of type restore.  - BAAS_BACKUP: List of supported BAAS application resources. BAAS application resource of type backup. | [default to &quot;TYPE_UNSPECIFIED&quot;]
 **queryInfraResource** | **string** | Type of the infra resource for which projects to be listed.   - TYPE_UNSPECIFIED: Unspecified, do not use.  - ACCOUNT: Currently supported infra resources. Infra resource of type account.  - TENANT: Infra resource of type tenant.  - PROJECT: Infra resource of type projects.  - TARGET_CLUSTER: Infra resource target cluster.  - NAMESPACE: Infra resource of type namespace.  - CREDENTIAL: Infra resource of type credential.  - BACKUP_LOCATION: Infra resource of type backup location. | [default to &quot;TYPE_UNSPECIFIED&quot;]
 **queryResourceId** | **string** | ID of the resource for which projects to be listed. | 
 **sortSortBy** | **string** | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. | [default to &quot;FIELD_UNSPECIFIED&quot;]
 **sortSortOrder** | **string** | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. | [default to &quot;VALUE_UNSPECIFIED&quot;]
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 

### Return type

[**V1ListProjectsResponse**](V1ListProjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

