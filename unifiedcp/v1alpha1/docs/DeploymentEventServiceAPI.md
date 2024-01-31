# \DeploymentEventServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeploymentEventServiceListDeploymentEvents**](DeploymentEventServiceAPI.md#DeploymentEventServiceListDeploymentEvents) | **Get** /pds/v1/deployments/{deploymentId}/events | ListDeploymentEvents API returns the list of kubernetes events related to a deployment.



## DeploymentEventServiceListDeploymentEvents

> V1ListDeploymentEventsResponse DeploymentEventServiceListDeploymentEvents(ctx, deploymentId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).Execute()

ListDeploymentEvents API returns the list of kubernetes events related to a deployment.

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
	deploymentId := "deploymentId_example" // string | Deployment id for which events need to be listed.
	paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
	paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)
	sortSortBy := "sortSortBy_example" // string | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. (optional) (default to "FIELD_UNSPECIFIED")
	sortSortOrder := "sortSortOrder_example" // string | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. (optional) (default to "VALUE_UNSPECIFIED")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentEventServiceAPI.DeploymentEventServiceListDeploymentEvents(context.Background(), deploymentId).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentEventServiceAPI.DeploymentEventServiceListDeploymentEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeploymentEventServiceListDeploymentEvents`: V1ListDeploymentEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `DeploymentEventServiceAPI.DeploymentEventServiceListDeploymentEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** | Deployment id for which events need to be listed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentEventServiceListDeploymentEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 
 **sortSortBy** | **string** | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. | [default to &quot;FIELD_UNSPECIFIED&quot;]
 **sortSortOrder** | **string** | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. | [default to &quot;VALUE_UNSPECIFIED&quot;]

### Return type

[**V1ListDeploymentEventsResponse**](V1ListDeploymentEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

