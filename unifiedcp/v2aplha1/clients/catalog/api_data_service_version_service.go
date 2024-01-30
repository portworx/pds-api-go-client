/*
public/portworx/pds/catalog/dataservices/apiv1/dataservices.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package catalog

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// DataServiceVersionServiceAPIService DataServiceVersionServiceAPI service
type DataServiceVersionServiceAPIService service

type ApiDataServiceVersionServiceGetDataServiceVersionRequest struct {
	ctx context.Context
	ApiService *DataServiceVersionServiceAPIService
	id string
}

func (r ApiDataServiceVersionServiceGetDataServiceVersionRequest) Execute() (*V1DataServiceVersion, *http.Response, error) {
	return r.ApiService.DataServiceVersionServiceGetDataServiceVersionExecute(r)
}

/*
DataServiceVersionServiceGetDataServiceVersion GetDataServiceVersion returns a data service version

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id UID of the version
 @return ApiDataServiceVersionServiceGetDataServiceVersionRequest
*/
func (a *DataServiceVersionServiceAPIService) DataServiceVersionServiceGetDataServiceVersion(ctx context.Context, id string) ApiDataServiceVersionServiceGetDataServiceVersionRequest {
	return ApiDataServiceVersionServiceGetDataServiceVersionRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return V1DataServiceVersion
func (a *DataServiceVersionServiceAPIService) DataServiceVersionServiceGetDataServiceVersionExecute(r ApiDataServiceVersionServiceGetDataServiceVersionRequest) (*V1DataServiceVersion, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1DataServiceVersion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataServiceVersionServiceAPIService.DataServiceVersionServiceGetDataServiceVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pds/v1/catalog/dataServiceVersions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest struct {
	ctx context.Context
	ApiService *DataServiceVersionServiceAPIService
	dataServiceId *string
	versionId *string
	sortSortBy *string
	sortSortOrder *string
	paginationPageNumber *string
	paginationPageSize *string
}

// UID of the data service for which compatible data service versions are requested
func (r ApiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest) DataServiceId(dataServiceId string) ApiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest {
	r.dataServiceId = &dataServiceId
	return r
}

// UID of the data service version for which compatible versions are requested
func (r ApiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest) VersionId(versionId string) ApiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest {
	r.versionId = &versionId
	return r
}

// Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource.
func (r ApiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest) SortSortBy(sortSortBy string) ApiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest {
	r.sortSortBy = &sortSortBy
	return r
}

// Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending.
func (r ApiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest) SortSortOrder(sortSortOrder string) ApiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest {
	r.sortSortOrder = &sortSortOrder
	return r
}

// Page number is the page number to return based on the size
func (r ApiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest) PaginationPageNumber(paginationPageNumber string) ApiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest {
	r.paginationPageNumber = &paginationPageNumber
	return r
}

// Page size is the maximum number of records to include per page
func (r ApiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest) PaginationPageSize(paginationPageSize string) ApiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest {
	r.paginationPageSize = &paginationPageSize
	return r
}

func (r ApiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest) Execute() (*V1ListCompatibleDataServiceVersionsResponse, *http.Response, error) {
	return r.ApiService.DataServiceVersionServiceListCompatibleDataServiceVersionsExecute(r)
}

/*
DataServiceVersionServiceListCompatibleDataServiceVersions ListCompatibleDataServiceVersions lists all the data service versions compatible with other version of a data service

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest
*/
func (a *DataServiceVersionServiceAPIService) DataServiceVersionServiceListCompatibleDataServiceVersions(ctx context.Context) ApiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest {
	return ApiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V1ListCompatibleDataServiceVersionsResponse
func (a *DataServiceVersionServiceAPIService) DataServiceVersionServiceListCompatibleDataServiceVersionsExecute(r ApiDataServiceVersionServiceListCompatibleDataServiceVersionsRequest) (*V1ListCompatibleDataServiceVersionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1ListCompatibleDataServiceVersionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataServiceVersionServiceAPIService.DataServiceVersionServiceListCompatibleDataServiceVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pds/v1/catalog/dataServiceVersions:listCompatibleVersions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.dataServiceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dataServiceId", r.dataServiceId, "")
	}
	if r.versionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "versionId", r.versionId, "")
	}
	if r.sortSortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort.sortBy", r.sortSortBy, "")
	} else {
		var defaultValue string = "FIELD_UNSPECIFIED"
		r.sortSortBy = &defaultValue
	}
	if r.sortSortOrder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort.sortOrder", r.sortSortOrder, "")
	} else {
		var defaultValue string = "VALUE_UNSPECIFIED"
		r.sortSortOrder = &defaultValue
	}
	if r.paginationPageNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pagination.pageNumber", r.paginationPageNumber, "")
	}
	if r.paginationPageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pagination.pageSize", r.paginationPageSize, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDataServiceVersionServiceListDataServiceVersionsRequest struct {
	ctx context.Context
	ApiService *DataServiceVersionServiceAPIService
	dataServiceId string
	sortSortBy *string
	sortSortOrder *string
	paginationPageNumber *string
	paginationPageSize *string
}

// Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource.
func (r ApiDataServiceVersionServiceListDataServiceVersionsRequest) SortSortBy(sortSortBy string) ApiDataServiceVersionServiceListDataServiceVersionsRequest {
	r.sortSortBy = &sortSortBy
	return r
}

// Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending.
func (r ApiDataServiceVersionServiceListDataServiceVersionsRequest) SortSortOrder(sortSortOrder string) ApiDataServiceVersionServiceListDataServiceVersionsRequest {
	r.sortSortOrder = &sortSortOrder
	return r
}

// Page number is the page number to return based on the size
func (r ApiDataServiceVersionServiceListDataServiceVersionsRequest) PaginationPageNumber(paginationPageNumber string) ApiDataServiceVersionServiceListDataServiceVersionsRequest {
	r.paginationPageNumber = &paginationPageNumber
	return r
}

// Page size is the maximum number of records to include per page
func (r ApiDataServiceVersionServiceListDataServiceVersionsRequest) PaginationPageSize(paginationPageSize string) ApiDataServiceVersionServiceListDataServiceVersionsRequest {
	r.paginationPageSize = &paginationPageSize
	return r
}

func (r ApiDataServiceVersionServiceListDataServiceVersionsRequest) Execute() (*V1ListDataServiceVersionsResponse, *http.Response, error) {
	return r.ApiService.DataServiceVersionServiceListDataServiceVersionsExecute(r)
}

/*
DataServiceVersionServiceListDataServiceVersions ListDataServiceVersions lists all the versions of a data service

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dataServiceId UID of the data service
 @return ApiDataServiceVersionServiceListDataServiceVersionsRequest
*/
func (a *DataServiceVersionServiceAPIService) DataServiceVersionServiceListDataServiceVersions(ctx context.Context, dataServiceId string) ApiDataServiceVersionServiceListDataServiceVersionsRequest {
	return ApiDataServiceVersionServiceListDataServiceVersionsRequest{
		ApiService: a,
		ctx: ctx,
		dataServiceId: dataServiceId,
	}
}

// Execute executes the request
//  @return V1ListDataServiceVersionsResponse
func (a *DataServiceVersionServiceAPIService) DataServiceVersionServiceListDataServiceVersionsExecute(r ApiDataServiceVersionServiceListDataServiceVersionsRequest) (*V1ListDataServiceVersionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1ListDataServiceVersionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataServiceVersionServiceAPIService.DataServiceVersionServiceListDataServiceVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pds/v1/catalog/dataServices/{dataServiceId}/dataServiceVersions"
	localVarPath = strings.Replace(localVarPath, "{"+"dataServiceId"+"}", url.PathEscape(parameterValueToString(r.dataServiceId, "dataServiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sortSortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort.sortBy", r.sortSortBy, "")
	} else {
		var defaultValue string = "FIELD_UNSPECIFIED"
		r.sortSortBy = &defaultValue
	}
	if r.sortSortOrder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort.sortOrder", r.sortSortOrder, "")
	} else {
		var defaultValue string = "VALUE_UNSPECIFIED"
		r.sortSortOrder = &defaultValue
	}
	if r.paginationPageNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pagination.pageNumber", r.paginationPageNumber, "")
	}
	if r.paginationPageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pagination.pageSize", r.paginationPageSize, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
