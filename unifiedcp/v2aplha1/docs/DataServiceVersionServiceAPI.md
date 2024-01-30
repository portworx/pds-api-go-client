# \DataServiceVersionServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataServiceVersionServiceGetDataServiceVersion**](DataServiceVersionServiceAPI.md#DataServiceVersionServiceGetDataServiceVersion) | **Get** /pds/v1/catalog/dataServiceVersions/{id} | GetDataServiceVersion returns a data service version
[**DataServiceVersionServiceListCompatibleDataServiceVersions**](DataServiceVersionServiceAPI.md#DataServiceVersionServiceListCompatibleDataServiceVersions) | **Get** /pds/v1/catalog/dataServiceVersions:listCompatibleVersions | ListCompatibleDataServiceVersions lists all the data service versions compatible with other version of a data service
[**DataServiceVersionServiceListDataServiceVersions**](DataServiceVersionServiceAPI.md#DataServiceVersionServiceListDataServiceVersions) | **Get** /pds/v1/catalog/dataServices/{dataServiceId}/dataServiceVersions | ListDataServiceVersions lists all the versions of a data service



## DataServiceVersionServiceGetDataServiceVersion

> V1DataServiceVersion DataServiceVersionServiceGetDataServiceVersion(ctx, id).Execute()

GetDataServiceVersion returns a data service version

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
	id := "id_example" // string | UID of the version

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataServiceVersionServiceAPI.DataServiceVersionServiceGetDataServiceVersion(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataServiceVersionServiceAPI.DataServiceVersionServiceGetDataServiceVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataServiceVersionServiceGetDataServiceVersion`: V1DataServiceVersion
	fmt.Fprintf(os.Stdout, "Response from `DataServiceVersionServiceAPI.DataServiceVersionServiceGetDataServiceVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UID of the version | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataServiceVersionServiceGetDataServiceVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1DataServiceVersion**](V1DataServiceVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataServiceVersionServiceListCompatibleDataServiceVersions

> V1ListCompatibleDataServiceVersionsResponse DataServiceVersionServiceListCompatibleDataServiceVersions(ctx).DataServiceId(dataServiceId).VersionId(versionId).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()

ListCompatibleDataServiceVersions lists all the data service versions compatible with other version of a data service

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
	dataServiceId := "dataServiceId_example" // string | UID of the data service for which compatible data service versions are requested (optional)
	versionId := "versionId_example" // string | UID of the data service version for which compatible versions are requested (optional)
	sortSortBy := "sortSortBy_example" // string | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. (optional) (default to "FIELD_UNSPECIFIED")
	sortSortOrder := "sortSortOrder_example" // string | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. (optional) (default to "VALUE_UNSPECIFIED")
	paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
	paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataServiceVersionServiceAPI.DataServiceVersionServiceListCompatibleDataServiceVersions(context.Background()).DataServiceId(dataServiceId).VersionId(versionId).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataServiceVersionServiceAPI.DataServiceVersionServiceListCompatibleDataServiceVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataServiceVersionServiceListCompatibleDataServiceVersions`: V1ListCompatibleDataServiceVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DataServiceVersionServiceAPI.DataServiceVersionServiceListCompatibleDataServiceVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataServiceId** | **string** | UID of the data service for which compatible data service versions are requested | 
 **versionId** | **string** | UID of the data service version for which compatible versions are requested | 
 **sortSortBy** | **string** | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. | [default to &quot;FIELD_UNSPECIFIED&quot;]
 **sortSortOrder** | **string** | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. | [default to &quot;VALUE_UNSPECIFIED&quot;]
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 

### Return type

[**V1ListCompatibleDataServiceVersionsResponse**](V1ListCompatibleDataServiceVersionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataServiceVersionServiceListDataServiceVersions

> V1ListDataServiceVersionsResponse DataServiceVersionServiceListDataServiceVersions(ctx, dataServiceId).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()

ListDataServiceVersions lists all the versions of a data service

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
	dataServiceId := "dataServiceId_example" // string | UID of the data service
	sortSortBy := "sortSortBy_example" // string | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. (optional) (default to "FIELD_UNSPECIFIED")
	sortSortOrder := "sortSortOrder_example" // string | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. (optional) (default to "VALUE_UNSPECIFIED")
	paginationPageNumber := "paginationPageNumber_example" // string | Page number is the page number to return based on the size (optional)
	paginationPageSize := "paginationPageSize_example" // string | Page size is the maximum number of records to include per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataServiceVersionServiceAPI.DataServiceVersionServiceListDataServiceVersions(context.Background(), dataServiceId).SortSortBy(sortSortBy).SortSortOrder(sortSortOrder).PaginationPageNumber(paginationPageNumber).PaginationPageSize(paginationPageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataServiceVersionServiceAPI.DataServiceVersionServiceListDataServiceVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataServiceVersionServiceListDataServiceVersions`: V1ListDataServiceVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DataServiceVersionServiceAPI.DataServiceVersionServiceListDataServiceVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataServiceId** | **string** | UID of the data service | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataServiceVersionServiceListDataServiceVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortSortBy** | **string** | Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource. | [default to &quot;FIELD_UNSPECIFIED&quot;]
 **sortSortOrder** | **string** | Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending. | [default to &quot;VALUE_UNSPECIFIED&quot;]
 **paginationPageNumber** | **string** | Page number is the page number to return based on the size | 
 **paginationPageSize** | **string** | Page size is the maximum number of records to include per page | 

### Return type

[**V1ListDataServiceVersionsResponse**](V1ListDataServiceVersionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

