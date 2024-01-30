# \DeploymentConfigUpdateServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeploymentConfigUpdateServiceCreateDeploymentConfigUpdate**](DeploymentConfigUpdateServiceAPI.md#DeploymentConfigUpdateServiceCreateDeploymentConfigUpdate) | **Put** /pds/v1/deployments/{deploymentConfigUpdate.config.deploymentMeta.uid}/deploymentConfigUpdates | CreateDeploymentConfigUpdate API creates a deployment config update (-- api-linter: core::0133::http-method&#x3D;disabled     aip.dev/not-precedent: We need to do this because here we are actually updating deployment config. --)
[**DeploymentConfigUpdateServiceGetDeploymentConfigUpdate**](DeploymentConfigUpdateServiceAPI.md#DeploymentConfigUpdateServiceGetDeploymentConfigUpdate) | **Get** /pds/v1/deploymentConfigUpdates/{id} | GetDeploymentConfigUpdate API returns a deployment config update by id
[**DeploymentConfigUpdateServiceListDeploymentConfigUpdates**](DeploymentConfigUpdateServiceAPI.md#DeploymentConfigUpdateServiceListDeploymentConfigUpdates) | **Get** /pds/v1/deploymentConfigUpdates | ListDeploymentConfigUpdates API returns a list of deployment config updates
[**DeploymentConfigUpdateServiceRetryDeploymentConfigUpdate**](DeploymentConfigUpdateServiceAPI.md#DeploymentConfigUpdateServiceRetryDeploymentConfigUpdate) | **Post** /pds/v1/deploymentConfigUpdates/{id}:retry | RetryDeploymentConfigUpdate API retries a deployment config update



## DeploymentConfigUpdateServiceCreateDeploymentConfigUpdate

> V1DeploymentConfigUpdate DeploymentConfigUpdateServiceCreateDeploymentConfigUpdate(ctx, deploymentConfigUpdateConfigDeploymentMetaUid).ConfigOfTheDeploymentForWhichConfigUpdateIsRequested(configOfTheDeploymentForWhichConfigUpdateIsRequested).Execute()

CreateDeploymentConfigUpdate API creates a deployment config update (-- api-linter: core::0133::http-method=disabled     aip.dev/not-precedent: We need to do this because here we are actually updating deployment config. --)

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
	deploymentConfigUpdateConfigDeploymentMetaUid := "deploymentConfigUpdateConfigDeploymentMetaUid_example" // string | UID of the resource of the format <resource prefix>-<uuid>.
	configOfTheDeploymentForWhichConfigUpdateIsRequested := *openapiclient.NewConfigOfTheDeploymentForWhichConfigUpdateIsRequested() // ConfigOfTheDeploymentForWhichConfigUpdateIsRequested | Config of the deployment for which config update is requested

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentConfigUpdateServiceAPI.DeploymentConfigUpdateServiceCreateDeploymentConfigUpdate(context.Background(), deploymentConfigUpdateConfigDeploymentMetaUid).ConfigOfTheDeploymentForWhichConfigUpdateIsRequested(configOfTheDeploymentForWhichConfigUpdateIsRequested).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentConfigUpdateServiceAPI.DeploymentConfigUpdateServiceCreateDeploymentConfigUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeploymentConfigUpdateServiceCreateDeploymentConfigUpdate`: V1DeploymentConfigUpdate
	fmt.Fprintf(os.Stdout, "Response from `DeploymentConfigUpdateServiceAPI.DeploymentConfigUpdateServiceCreateDeploymentConfigUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentConfigUpdateConfigDeploymentMetaUid** | **string** | UID of the resource of the format &lt;resource prefix&gt;-&lt;uuid&gt;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentConfigUpdateServiceCreateDeploymentConfigUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configOfTheDeploymentForWhichConfigUpdateIsRequested** | [**ConfigOfTheDeploymentForWhichConfigUpdateIsRequested**](ConfigOfTheDeploymentForWhichConfigUpdateIsRequested.md) | Config of the deployment for which config update is requested | 

### Return type

[**V1DeploymentConfigUpdate**](V1DeploymentConfigUpdate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeploymentConfigUpdateServiceGetDeploymentConfigUpdate

> V1DeploymentConfigUpdate DeploymentConfigUpdateServiceGetDeploymentConfigUpdate(ctx, id).Execute()

GetDeploymentConfigUpdate API returns a deployment config update by id

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
	id := "id_example" // string | UID of the deployment config update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentConfigUpdateServiceAPI.DeploymentConfigUpdateServiceGetDeploymentConfigUpdate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentConfigUpdateServiceAPI.DeploymentConfigUpdateServiceGetDeploymentConfigUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeploymentConfigUpdateServiceGetDeploymentConfigUpdate`: V1DeploymentConfigUpdate
	fmt.Fprintf(os.Stdout, "Response from `DeploymentConfigUpdateServiceAPI.DeploymentConfigUpdateServiceGetDeploymentConfigUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UID of the deployment config update | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentConfigUpdateServiceGetDeploymentConfigUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1DeploymentConfigUpdate**](V1DeploymentConfigUpdate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeploymentConfigUpdateServiceListDeploymentConfigUpdates

> V1ListDeploymentConfigUpdatesResponse DeploymentConfigUpdateServiceListDeploymentConfigUpdates(ctx).DeploymentId(deploymentId).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()

ListDeploymentConfigUpdates API returns a list of deployment config updates

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
	deploymentId := "deploymentId_example" // string | UID of the deployment (optional)
	sortSortBy := "sortSortBy_example" // string | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. (optional) (default to "FIELD_UNSPECIFIED")
	sortSortOrder := "sortSortOrder_example" // string | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. (optional) (default to "VALUE_UNSPECIFIED")
	paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
	paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentConfigUpdateServiceAPI.DeploymentConfigUpdateServiceListDeploymentConfigUpdates(context.Background()).DeploymentId(deploymentId).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentConfigUpdateServiceAPI.DeploymentConfigUpdateServiceListDeploymentConfigUpdates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeploymentConfigUpdateServiceListDeploymentConfigUpdates`: V1ListDeploymentConfigUpdatesResponse
	fmt.Fprintf(os.Stdout, "Response from `DeploymentConfigUpdateServiceAPI.DeploymentConfigUpdateServiceListDeploymentConfigUpdates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentConfigUpdateServiceListDeploymentConfigUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentId** | **string** | UID of the deployment | 
 **sortSortBy** | **string** | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. | [default to &quot;FIELD_UNSPECIFIED&quot;]
 **sortSortOrder** | **string** | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. | [default to &quot;VALUE_UNSPECIFIED&quot;]
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 

### Return type

[**V1ListDeploymentConfigUpdatesResponse**](V1ListDeploymentConfigUpdatesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeploymentConfigUpdateServiceRetryDeploymentConfigUpdate

> V1DeploymentConfigUpdate DeploymentConfigUpdateServiceRetryDeploymentConfigUpdate(ctx, id).Body(body).Execute()

RetryDeploymentConfigUpdate API retries a deployment config update

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
	id := "id_example" // string | UID of the deployment config update
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentConfigUpdateServiceAPI.DeploymentConfigUpdateServiceRetryDeploymentConfigUpdate(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentConfigUpdateServiceAPI.DeploymentConfigUpdateServiceRetryDeploymentConfigUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeploymentConfigUpdateServiceRetryDeploymentConfigUpdate`: V1DeploymentConfigUpdate
	fmt.Fprintf(os.Stdout, "Response from `DeploymentConfigUpdateServiceAPI.DeploymentConfigUpdateServiceRetryDeploymentConfigUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UID of the deployment config update | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentConfigUpdateServiceRetryDeploymentConfigUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**V1DeploymentConfigUpdate**](V1DeploymentConfigUpdate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

