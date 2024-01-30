# \DataServicesServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataServicesServiceGetDataService**](DataServicesServiceAPI.md#DataServicesServiceGetDataService) | **Get** /pds/v1/catalog/dataServices/{id} | GetDataService API returns the the data service resource.
[**DataServicesServiceListDataServices**](DataServicesServiceAPI.md#DataServicesServiceListDataServices) | **Get** /pds/v1/catalog/dataServices | ListDataServices API lists the the data service resources.



## DataServicesServiceGetDataService

> V1DataService DataServicesServiceGetDataService(ctx, id).Execute()

GetDataService API returns the the data service resource.

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
	id := "id_example" // string | UID of the Data Service.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataServicesServiceAPI.DataServicesServiceGetDataService(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataServicesServiceAPI.DataServicesServiceGetDataService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataServicesServiceGetDataService`: V1DataService
	fmt.Fprintf(os.Stdout, "Response from `DataServicesServiceAPI.DataServicesServiceGetDataService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UID of the Data Service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataServicesServiceGetDataServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1DataService**](V1DataService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataServicesServiceListDataServices

> V1ListDataServicesResponse DataServicesServiceListDataServices(ctx).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()

ListDataServices API lists the the data service resources.

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
	sortSortBy := "sortSortBy_example" // string | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. (optional) (default to "FIELD_UNSPECIFIED")
	sortSortOrder := "sortSortOrder_example" // string | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. (optional) (default to "VALUE_UNSPECIFIED")
	paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
	paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataServicesServiceAPI.DataServicesServiceListDataServices(context.Background()).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataServicesServiceAPI.DataServicesServiceListDataServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataServicesServiceListDataServices`: V1ListDataServicesResponse
	fmt.Fprintf(os.Stdout, "Response from `DataServicesServiceAPI.DataServicesServiceListDataServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataServicesServiceListDataServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortSortBy** | **string** | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. | [default to &quot;FIELD_UNSPECIFIED&quot;]
 **sortSortOrder** | **string** | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. | [default to &quot;VALUE_UNSPECIFIED&quot;]
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 

### Return type

[**V1ListDataServicesResponse**](V1ListDataServicesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

